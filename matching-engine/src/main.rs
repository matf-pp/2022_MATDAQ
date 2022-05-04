extern crate ordered_float;
extern crate rand_distr;
extern crate rand;

mod limit_order_book;

use crate::limit_order_book::limit_order_book::LimitOrderBook;
use crate::limit_order_book::order_side::Side;
use crate::limit_order_book::order_type::OrderType;
use std::{
    time::SystemTime,
    ops::Div
};
use rand::Rng;
use rand_distr::{Normal, Distribution};

fn main() {
    let num_of_orders = 250_000;
    let mut rng = rand::thread_rng();
    let mut lob : LimitOrderBook = LimitOrderBook::new();
    let mut prices = Vec::with_capacity(num_of_orders);
    let mut amounts = Vec::with_capacity(num_of_orders);
    let normal = Normal::new(200.0, 5.0).unwrap();

    // generate random prices and amounts
    for _i in 0..num_of_orders {
        let price: f64 = normal.sample(&mut rand::thread_rng());
        let amount = rng.gen_range(0..1000);
        prices.push((price * 100.0).round().div(100.0));
        amounts.push(amount)
    }

    // start time
    let start = SystemTime::now();

    // generate orders
    for i in 0..num_of_orders {
        let side: Side = rand::random();
        lob.add_order(rng.gen::<u64>(), OrderType::Limit, side, amounts[i], prices[i], 1);
    }

    // end time
    let end = SystemTime::now();

    let since_the_epoch = end
        .duration_since(start)
        .expect("Time went backwards");

    println!("Time taken to execute {} orders: {}ms", num_of_orders, since_the_epoch.as_millis());
    println!("Orders on Buy Side: {}, Orders on Sell Side: {}", lob.buy_side_size(), lob.sell_side_size());
}
