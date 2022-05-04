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
    let mut rng = rand::thread_rng();
    let mut lob : LimitOrderBook = LimitOrderBook::new();
    let mut prices = Vec::with_capacity(250000);
    let mut amounts = Vec::with_capacity(250000);
    let normal = Normal::new(200.0, 4.0).unwrap();

    // generate 250k random prices and amounts
    for _i in 0..250000 {
        let price: f64 = normal.sample(&mut rand::thread_rng());
        let amount = rng.gen_range(0..1000);
        prices.push((price * 100.0).round().div(100.0));
        amounts.push(amount)
    }

    // start time
    let start = SystemTime::now();

    // generate 250k orders
    for i in 0..250000 {
        let side: Side = rand::random();
        lob.add_order(rng.gen::<u64>(), OrderType::Limit, side, amounts[i], prices[i], 1);
    }

    // end time
    let end = SystemTime::now();

    let since_the_epoch = end
        .duration_since(start)
        .expect("Time went backwards");

    // println!("-----------------------------------------------------");
    println!("Time taken to execute 250k orders: {}ms", since_the_epoch.as_millis());
    println!("Orders on Buy Side: {}, Orders on Sell Side: {}", lob.buy_side_size(), lob.sell_side_size());
}
