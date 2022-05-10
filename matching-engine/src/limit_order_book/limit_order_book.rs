extern crate ordered_float;

use crate::limit_order_book::{
    order::{Order, OrderId, SecurityId, SenderId},
    order_side::Side,
    order_type::OrderType,
    request_handlers::{request_money_update, request_new_order, request_trade},
};
use rand_distr::num_traits::ToPrimitive;
use std::cmp::Reverse;
use std::collections::{BinaryHeap, HashMap};

pub struct LimitOrderBook {
    buy_side: BinaryHeap<Order>,
    sell_side: BinaryHeap<Reverse<Order>>,
    orders: HashMap<OrderId, Order>,
}

impl LimitOrderBook {
    pub fn new() -> LimitOrderBook {
        LimitOrderBook {
            buy_side: BinaryHeap::new(),
            sell_side: BinaryHeap::new(),
            orders: HashMap::new(),
        }
    }

    pub fn print_book(&self) {
        let sell_side = self.sell_side.clone();
        let sorted_sells = sell_side.into_sorted_vec();
        let buy_side = self.buy_side.clone();
        let mut sorted_buys = buy_side.into_sorted_vec();
        sorted_buys.reverse();
        println!("----------------------------------------------------------");
        println!("Sell Side:");
        for ord in sorted_sells {
            println!("\x1b[31m{}\x1b[0m", ord.0);
        }
        println!("----------------------------------------------------------");
        println!("Buy Side:");
        for ord in sorted_buys {
            println!("\x1b[32m{}\x1b[0m", ord);
        }
        println!("----------------------------------------------------------");
    }

    pub fn buy_side_size(&self) -> usize {
        self.buy_side.len()
    }

    pub fn sell_side_size(&self) -> usize {
        self.sell_side.len()
    }

    /*
       Checks if the order is aggressive

       Order is aggressive if it is on the Buy side and it's limit_price is bigger
       than the best ask price, or if it is on the Sell side and it's limit_price is
       less than the best bid price
    */
    fn is_aggressive(&self, order: &Order) -> bool {
        if order.side == Side::Buy {
            order.limit_price >= self.sell_side.peek().unwrap().0.limit_price
        } else {
            order.limit_price <= self.buy_side.peek().unwrap().limit_price
        }
    }

    /*
       Adds new order to the Limit Order Book if the order is passive or executes
       the order if the order is aggressive
    */
    pub fn add_order(
        &mut self,
        order_id: u64,
        security_id: SecurityId,
        sender_id: SenderId,
        ord_type: OrderType,
        side: Side,
        amount: u32,
        limit_price: i32,
        time: u128,
    ) {
        let mut order = Order {
            order_id,
            sender_id,
            ord_type,
            side,
            amount,
            limit_price,
            time,
            security_id,
        };

        // Set limit price of a Market order according to the side it's on
        if order.ord_type == OrderType::Market {
            if order.side == Side::Sell {
                order.limit_price = 0;
            } else {
                order.limit_price = i32::MAX;
            }
        }

        let mut ord = order.clone();

        let mut total_money_exchanged = 0;
        /*
           Here we check the order's type(Market/Limit) and order's side(Buy/Sell)
           and based on that we add or execute the given order
           If the order is executed then total_money_exchanged won't be equal to 0
        */
        match order.side {
            Side::Buy => match order.ord_type {
                OrderType::Limit => {
                    if !self.sell_side.is_empty() && self.is_aggressive(&ord) {
                        /*
                        returns the amount by which we should lower current user's money
                         */
                        total_money_exchanged = self.execute_buy_limit_order(&mut ord);
                    } else {
                        /*
                           after adding new orders we should send update to Price-Display client
                        */
                        self.orders.insert(order.order_id, order);
                        self.buy_side.push(ord);
                        request_new_order(ord.security_id, ord.limit_price, ord.amount, ord.side);
                    }
                }
                OrderType::Market => {
                    if !self.sell_side.is_empty() {
                        total_money_exchanged = self.execute_buy_market_order(&mut ord);
                    } else {
                        self.orders.insert(order.order_id, order);
                        self.buy_side.push(ord);
                        request_new_order(ord.security_id, ord.limit_price, ord.amount, ord.side);
                    }
                }
            },
            Side::Sell => match order.ord_type {
                OrderType::Limit => {
                    if !self.buy_side.is_empty() && self.is_aggressive(&ord) {
                        total_money_exchanged = self.execute_sell_limit_order(&mut ord)
                    } else {
                        self.orders.insert(order.order_id, order);
                        self.sell_side.push(Reverse(ord));
                        request_new_order(ord.security_id, ord.limit_price, ord.amount, ord.side);
                    }
                }
                OrderType::Market => {
                    if !self.buy_side.is_empty() {
                        total_money_exchanged = self.execute_sell_market_order(&mut ord)
                    } else {
                        self.orders.insert(order.order_id, order);
                        self.sell_side.push(Reverse(ord));
                        request_new_order(ord.security_id, ord.limit_price, ord.amount, ord.side);
                    }
                }
            },
        }

        // Send request to update user's money
        request_money_update(order.sender_id, total_money_exchanged);
    }

    /*
       Processes the order that is being executed against curr_best order and calculates
       the amount of money that the user traded

       It also sends request to User-Service to update the money of the user that placed curr_best order
    */
    fn process_curr_order(&mut self, curr_best: &mut Order, ord: &mut Order) -> i32 {
        let money_amount: i32;
        let mut curr_best_price: i32 = 0;

        /*
           Set best price for buying/selling

           If at the top of the book there is a Market Order on the opposite side of our order
           and our order is Limit Order then we execute at the price of our Limit Order
           Else, if our Order is Market Order then we dont execute it, we just add it to the LOB
        */
        if curr_best.ord_type == OrderType::Market {
            if ord.ord_type == OrderType::Market && ord.side == Side::Buy {
                self.buy_side.push(*ord);
                request_new_order(ord.security_id, ord.limit_price, ord.amount, ord.side);
                return 0;
            } else if ord.ord_type == OrderType::Market && ord.side == Side::Sell {
                self.sell_side.push(Reverse(*ord));
                request_new_order(ord.security_id, ord.limit_price, ord.amount, ord.side);
                return 0;
            } else {
                curr_best_price = ord.limit_price;
            }
        } else {
            curr_best_price = curr_best.limit_price;
        }

        /*
           If the amount of the order at the top of the book can't fill our order then we
           update our order's amount and we remove the order it matched against

           Else if the amount of the order at the top of the book can fill our order then we
           update it's amount by decreasing it by our order's amount, our order is now filled
           so we can set its amount to 0, lastly we need to update the order at the top of the
           book that our order matched against
        */
        if curr_best.amount <= ord.amount {
            ord.amount -= curr_best.amount;
            money_amount = curr_best.amount.to_i32().unwrap() * curr_best_price;

            /*
               We check the side on which the curr_best order sits
               If it sits on the Buy side that means that we should decrease the amount of money
               of the user that placed it
               If it sits on the Sell side that means that we should increase the amount of money
               of the user that placed it
            */
            match curr_best.side {
                Side::Buy => {
                    request_money_update(curr_best.sender_id, -money_amount);
                    request_trade(curr_best.security_id, curr_best.amount, Side::Buy);
                    self.buy_side.pop();
                }
                Side::Sell => {
                    request_money_update(curr_best.sender_id, money_amount);
                    request_trade(curr_best.security_id, curr_best.amount, Side::Sell);
                    self.sell_side.pop();
                }
            }
        } else {
            curr_best.amount -= ord.amount;
            money_amount = ord.amount.to_i32().unwrap() * curr_best_price;
            ord.amount = 0;
            self.orders.remove(&curr_best.order_id);
            if ord.side == Side::Buy {
                request_money_update(curr_best.sender_id, money_amount);
                request_trade(curr_best.security_id, curr_best.amount, Side::Buy);
                self.sell_side.pop();
                self.sell_side.push(Reverse(*curr_best));
            } else {
                request_money_update(curr_best.sender_id, -money_amount);
                request_trade(curr_best.security_id, curr_best.amount, Side::Sell);
                self.buy_side.pop();
                self.buy_side.push(*curr_best);
            }
            self.orders.insert(curr_best.order_id, *curr_best);
        }

        return money_amount;
    }

    fn execute_buy_limit_order(&mut self, ord: &mut Order) -> i32 {
        let mut money_amount = 0;

        // order already exists in a LOB
        if self.orders.contains_key(&ord.order_id) {
            println!("Order already exists");
            return money_amount;
        }

        // if we dont have orders to match against then return
        if self.sell_side.is_empty() {
            return money_amount;
        }

        let mut best_ask = self.sell_side.peek().unwrap().0.clone();
        while ord.amount > 0
            && !self.sell_side.is_empty()
            && ord.limit_price >= best_ask.limit_price
        {
            // we are doing -= because we are buying so our money should go down
            money_amount -= self.process_curr_order(&mut best_ask, ord);

            if !self.sell_side.is_empty() {
                best_ask = self.sell_side.peek().unwrap().0.clone();
            }
        }

        // if order is not filled we create new order with remaining amount
        if !ord.is_filled() {
            let ord_clone = ord.clone();
            request_new_order(ord.security_id, ord.limit_price, ord.amount, ord.side);
            self.buy_side.push(ord_clone);
            self.orders.insert(ord.order_id, ord_clone);
        }

        return money_amount;
    }

    fn execute_buy_market_order(&mut self, ord: &mut Order) -> i32 {
        let mut money_amount = 0;

        if self.orders.contains_key(&ord.order_id) {
            println!("Order already exists");
            return 0;
        }

        if self.sell_side.is_empty() {
            return 0;
        }

        let mut best_ask = self.sell_side.peek().unwrap().0.clone();
        while ord.amount > 0 && !self.sell_side.is_empty() {
            money_amount -= self.process_curr_order(&mut best_ask, ord);

            if money_amount == 0 {
                return 0;
            }

            if !self.sell_side.is_empty() {
                best_ask = self.sell_side.peek().unwrap().0.clone();
            }
        }

        if !ord.is_filled() {
            let ord_clone = ord.clone();
            request_new_order(ord.security_id, ord.limit_price, ord.amount, ord.side);
            self.buy_side.push(ord_clone);
            self.orders.insert(ord.order_id, ord_clone);
        }

        return money_amount;
    }

    fn execute_sell_limit_order(&mut self, ord: &mut Order) -> i32 {
        let mut money_amount = 0;

        if self.orders.contains_key(&ord.order_id) {
            println!("Order already exists");
            return 0;
        }

        if self.buy_side.is_empty() {
            return 0;
        }

        let mut best_bid = self.buy_side.peek().unwrap().clone();
        while ord.amount > 0 && !self.buy_side.is_empty() && ord.limit_price <= best_bid.limit_price
        {
            // money_amount is positive because we are selling stock
            money_amount += self.process_curr_order(&mut best_bid, ord);

            if !self.buy_side.is_empty() {
                best_bid = self.buy_side.peek().unwrap().clone();
            }
        }

        if !ord.is_filled() {
            let ord_clone = ord.clone();
            request_new_order(ord.security_id, ord.limit_price, ord.amount, ord.side);
            self.sell_side.push(Reverse(ord_clone));
            self.orders.insert(ord.order_id, ord_clone);
        }

        return money_amount;
    }

    fn execute_sell_market_order(&mut self, ord: &mut Order) -> i32 {
        let mut money_amount = 0;

        if self.orders.contains_key(&ord.order_id) {
            println!("Order already exists");
            return 0;
        }

        if self.buy_side.is_empty() {
            return 0;
        }

        let mut best_bid = self.buy_side.peek().unwrap().clone();
        while ord.amount > 0 && !self.buy_side.is_empty() {
            money_amount += self.process_curr_order(&mut best_bid, ord);

            if money_amount == 0 {
                return 0;
            }

            if !self.buy_side.is_empty() {
                best_bid = self.buy_side.peek().unwrap().clone();
            }
        }

        if !ord.is_filled() {
            let ord_clone = ord.clone();
            request_new_order(ord.security_id, ord.limit_price, ord.amount, ord.side);
            self.sell_side.push(Reverse(ord_clone));
            self.orders.insert(ord.order_id, ord_clone);
        }

        return money_amount;
    }
}
