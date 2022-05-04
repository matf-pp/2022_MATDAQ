extern crate ordered_float;

use std::collections::{HashMap, BinaryHeap};
use std::cmp::Reverse;
use ordered_float::NotNan;
use crate::limit_order_book::order_side::Side;
use crate::limit_order_book::order_type::OrderType;
use crate::limit_order_book::order::{Order, OrderId};


pub struct LimitOrderBook {
    buy_side: BinaryHeap<Order>,
    sell_side: BinaryHeap<Reverse<Order>>,
    orders: HashMap<OrderId, Order>
}

impl LimitOrderBook {
    pub fn new() -> LimitOrderBook {
        LimitOrderBook { buy_side: BinaryHeap::new(), sell_side: BinaryHeap::new(), orders: HashMap::new() }
    }

    pub fn print_book(&self) {
        println!("----------------------------------------------------------");
        println!("Sell Side:");
        for ord in &self.sell_side {
            println!("{}", ord.0);
        }
        println!("----------------------------------------------------------");
        println!("Buy Side:");
        for ord in &self.buy_side {
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
        time: u64
    ) {
        let order = Order{ order_id, ord_type, side, amount, limit_price: NotNan::new(limit_price).unwrap(), time };
        let ord = order.clone();

        match order.side {
            Side::Buy => {
                match order.ord_type {
                    OrderType::Limit => {
                        if !self.sell_side.is_empty() && ord.limit_price >= self.sell_side.peek().unwrap().0.limit_price {
                            self.execute_order(ord)
                        } else {
                            self.orders.insert(order.order_id, order);
                            self.buy_side.push(ord);
                        }
                    }
                    OrderType::Market => {
                        if !self.sell_side.is_empty() {
                            self.execute_order(ord)
                        } else {
                            self.orders.insert(order.order_id, order);
                            self.buy_side.push(ord);
                        }
                    }
                }
            },
            Side::Sell => {
                match order.ord_type {
                    OrderType::Limit => {
                        if !self.buy_side.is_empty() && ord.limit_price <= self.buy_side.peek().unwrap().limit_price {
                            self.execute_order(ord)
                        } else {
                            self.orders.insert(order.order_id, order);
                            self.sell_side.push(Reverse(ord));
                        }
                    }
                    OrderType::Market => {
                        if !self.buy_side.is_empty() {
                            self.execute_order(ord)
                        } else {
                            self.orders.insert(order.order_id, order);
                            self.sell_side.push(Reverse(ord));
                        }
                    }
                }
            }
        }
    }

    fn execute_order(&mut self, order: Order) {
        let mut ord = order.clone();

        // order already exists in a LOB
        if self.orders.contains_key(&ord.order_id) {
            println!("Order already exists")
        }

        match ord.side {
            Side::Buy => {
                // sets limit price to MAX value if the order type is market order
                if ord.ord_type == OrderType::Market {
                    ord.limit_price = NotNan::new(f64::MAX).expect("Number is NaN");
                }
                if self.sell_side.is_empty() {
                    return
                }

                let mut best_ask = self.sell_side.peek().unwrap().0.clone();
                while ord.amount > 0 && !self.sell_side.is_empty() && ord.limit_price >= best_ask.limit_price {
                    /*
                    if we can't fill the order with the current best_ask amount we will pop from sell_side
                    and lower the order amount by best_ask amount
                    */
                    if best_ask.amount <= ord.amount {
                        ord.amount -= best_ask.amount;
                        self.sell_side.pop();
                        /*
                        else if the order can be filled from current best_ask we decrease the amount
                        of that order, pop that order from book and put in a new order
                        */
                    } else {
                        best_ask.amount -= ord.amount;
                        ord.amount = 0;
                        self.sell_side.pop();
                        self.orders.remove(&best_ask.order_id);
                        self.sell_side.push(Reverse(best_ask));
                        self.orders.insert(best_ask.order_id, best_ask);
                        return
                    }

                    // get new best_ask
                    if !self.sell_side.is_empty() {
                        best_ask = self.sell_side.peek().unwrap().0.clone();
                    }
                }

                // if order is not filled we create new order with remaining amount
                if ord.amount > 0 {
                    self.buy_side.push(ord);
                    self.orders.insert(ord.order_id, ord);
                }
            },
            Side::Sell => {
                // sets limit price to 0 if the order type is market order
                if ord.ord_type == OrderType::Market {
                    ord.limit_price = NotNan::new(0_f64).expect("Number is NaN");
                }
                if self.buy_side.is_empty() {
                    return
                }
                let mut best_bid = self.buy_side.peek().unwrap().clone();
                while ord.amount > 0 && !self.buy_side.is_empty() && ord.limit_price <= best_bid.limit_price {
                    if best_bid.amount <= ord.amount {
                        ord.amount -= best_bid.amount;
                        self.buy_side.pop();
                    } else {
                        best_bid.amount -= ord.amount;
                        ord.amount = 0;
                        self.buy_side.pop();
                        self.buy_side.push(best_bid);
                        self.orders.insert(best_bid.order_id, best_bid);
                    }

                    if !self.buy_side.is_empty() {
                        best_bid = self.buy_side.peek().unwrap().clone();
                    }
                }

                if ord.amount > 0 {
                    self.sell_side.push(Reverse(ord));
                    self.orders.insert(ord.order_id, ord);
                }
            }
        }
    }
}
