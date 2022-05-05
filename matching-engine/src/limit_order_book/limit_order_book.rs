extern crate ordered_float;

use crate::limit_order_book::order::{Order, OrderId};
use crate::limit_order_book::order_side::Side;
use crate::limit_order_book::order_type::OrderType;
use ordered_float::NotNan;
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
            println!("{}", ord.0);
        }
        println!("----------------------------------------------------------");
        println!("Buy Side:");
        for ord in sorted_buys {
            println!("{}", ord);
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
     */
    fn is_aggressive(&self, order: &Order) -> bool {
        if order.side == Side::Buy {
            order.limit_price >= self.sell_side.peek().unwrap().0.limit_price
        } else {
            order.limit_price <= self.buy_side.peek().unwrap().limit_price
        }
    }

    /*
    Adds new order to the Limit Order Book if the order is passive
    if the order is aggressive it executes the order
     */
    pub fn add_order(
        &mut self,
        order_id: u64,
        ord_type: OrderType,
        side: Side,
        amount: u32,
        limit_price: f64,
        time: u64,
    ) {
        let mut order = Order {
            order_id,
            ord_type,
            side,
            amount,
            limit_price: NotNan::new(limit_price).unwrap(),
            time,
        };

        if order.ord_type == OrderType::Market {
            if order.side == Side::Sell {
                order.limit_price = NotNan::new(0.0).expect("");
            } else {
                order.limit_price = NotNan::new(f64::MAX).expect("");
            }
        }

        let mut ord = order.clone();

        match order.side {
            Side::Buy => match order.ord_type {
                OrderType::Limit => {
                    if !self.sell_side.is_empty() && self.is_aggressive(&ord) {
                        self.execute_buy_limit_order(&mut ord)
                    } else {
                        self.orders.insert(order.order_id, order);
                        self.buy_side.push(ord);
                    }
                }
                OrderType::Market => {
                    if !self.sell_side.is_empty() {
                        self.execute_buy_market_order(&mut ord)
                    } else {
                        self.orders.insert(order.order_id, order);
                        self.buy_side.push(ord);
                    }
                }
            },
            Side::Sell => match order.ord_type {
                OrderType::Limit => {
                    if !self.buy_side.is_empty() && self.is_aggressive(&ord) {
                        self.execute_sell_limit_order(&mut ord)
                    } else {
                        self.orders.insert(order.order_id, order);
                        self.sell_side.push(Reverse(ord));
                    }
                }
                OrderType::Market => {
                    if !self.buy_side.is_empty() {
                        self.execute_sell_market_order(&mut ord)
                    } else {
                        self.orders.insert(order.order_id, order);
                        self.sell_side.push(Reverse(ord));
                    }
                }
            },
        }
    }

    fn process_curr_order(&mut self, curr_best: &mut Order, ord: &mut Order) {
        if curr_best.amount <= ord.amount {
            ord.amount -= curr_best.amount;
            if ord.side == Side::Buy {
                self.sell_side.pop();
            } else {
                self.buy_side.pop();
            }
        } else {
            curr_best.amount -= ord.amount;
            ord.amount = 0;
            self.orders.remove(&curr_best.order_id);
            if ord.side == Side::Buy {
                self.sell_side.pop();
                self.sell_side.push(Reverse(*curr_best));
            } else {
                self.buy_side.pop();
                self.buy_side.push(*curr_best);
            }
            self.orders.insert(curr_best.order_id, *curr_best);
        }
    }

    fn execute_buy_limit_order(&mut self, ord: &mut Order) {
        // order already exists in a LOB
        if self.orders.contains_key(&ord.order_id) {
            println!("Order already exists")
        }

        // if we dont have orders to match against then return
        if self.sell_side.is_empty() {
            return;
        }

        let mut best_ask = self.sell_side.peek().unwrap().0.clone();
        while ord.amount > 0
            && !self.sell_side.is_empty()
            && ord.limit_price >= best_ask.limit_price
        {
            self.process_curr_order(&mut best_ask, ord);

            if !self.sell_side.is_empty() {
                best_ask = self.sell_side.peek().unwrap().0.clone();
            }
        }

        // if order is not filled we create new order with remaining amount
        if !ord.is_filled() {
            let ord_clone = ord.clone();
            self.buy_side.push(ord_clone);
            self.orders.insert(ord.order_id, ord_clone);
        }
    }

    fn execute_buy_market_order(&mut self, ord: &mut Order) {
        if self.orders.contains_key(&ord.order_id) {
            println!("Order already exists")
        }

        if self.sell_side.is_empty() {
            return;
        }

        let mut best_ask = self.sell_side.peek().unwrap().0.clone();
        while ord.amount > 0 && !self.sell_side.is_empty() {
            self.process_curr_order(&mut best_ask, ord);

            if !self.sell_side.is_empty() {
                best_ask = self.sell_side.peek().unwrap().0.clone();
            }
        }

        if !ord.is_filled() {
            let ord_clone = ord.clone();
            self.buy_side.push(ord_clone);
            self.orders.insert(ord.order_id, ord_clone);
        }
    }

    fn execute_sell_limit_order(&mut self, ord: &mut Order) {
        if self.orders.contains_key(&ord.order_id) {
            println!("Order already exists")
        }

        if self.buy_side.is_empty() {
            return;
        }

        let mut best_bid = self.buy_side.peek().unwrap().clone();
        while ord.amount > 0 && !self.buy_side.is_empty() && ord.limit_price <= best_bid.limit_price
        {
            self.process_curr_order(&mut best_bid, ord);

            if !self.buy_side.is_empty() {
                best_bid = self.buy_side.peek().unwrap().clone();
            }
        }

        if !ord.is_filled() {
            let ord_clone = ord.clone();
            self.sell_side.push(Reverse(ord_clone));
            self.orders.insert(ord.order_id, ord_clone);
        }
    }

    fn execute_sell_market_order(&mut self, ord: &mut Order) {
        if self.orders.contains_key(&ord.order_id) {
            println!("Order already exists")
        }

        if self.buy_side.is_empty() {
            return;
        }

        let mut best_bid = self.buy_side.peek().unwrap().clone();
        while ord.amount > 0 && !self.buy_side.is_empty() {
            self.process_curr_order(&mut best_bid, ord);

            if !self.buy_side.is_empty() {
                best_bid = self.buy_side.peek().unwrap().clone();
            }
        }

        if !ord.is_filled() {
            let ord_clone = ord.clone();
            self.sell_side.push(Reverse(ord_clone));
            self.orders.insert(ord.order_id, ord_clone);
        }
    }
}
