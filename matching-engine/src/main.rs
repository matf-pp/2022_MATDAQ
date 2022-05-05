extern crate ordered_float;
extern crate rand;
extern crate rand_distr;

mod limit_order_book;

use crate::limit_order_book::{
    limit_order_book::LimitOrderBook,
    order::{Order, SecurityId},
    order_side::Side,
    order_type::OrderType,
};
use ordered_float::NotNan;
use rand::Rng;
use rand_distr::{Distribution, Normal};
use std::borrow::Borrow;
use std::collections::HashMap;
use std::time::UNIX_EPOCH;
use std::{ops::Div, time::SystemTime};

// hashmap iz securityid u LOB

fn main() {
    let num_of_orders = 500;
    let mut rng = rand::thread_rng();
    let mut lob: LimitOrderBook = LimitOrderBook::new();
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

    /*
    load information about order into Order struct
     */
    // let order = Order {
    //     security_id: 1,
    //     sender_id: 1,
    //     order_id: 111,
    //     side: Side::Sell,
    //     ord_type: OrderType::Limit,
    //     amount: 10,
    //     limit_price: NotNan::new(100.0).expect("Value is NaN"),
    //     time: SystemTime::now().duration_since(UNIX_EPOCH).expect("Time went backwards").as_nanos(),
    // };

    /*
    we get new order from Request-Service-Server containing information about new order
    inside there is SecurityId field, based on that field we get the LOB on which we add new order
    if current SecurityId from request doesnt exist, we create new LOB for that request

    LimitOrderBook::add_order() should receive Order type as parameter

    lobs.
     */
    let mut lobs: HashMap<SecurityId, LimitOrderBook> = HashMap::new();

    // start time
    let start = SystemTime::now();

    // generate orders
    for i in 0..num_of_orders {
        let side: Side = rand::random();
        let ord_type: OrderType = rand::random();
        let price = NotNan::new(prices[i]).expect("");
        let curr_time = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .expect("Time went backwards")
            .as_nanos();
        let security_id = rng.gen_range(1..=5);
        if !lobs.contains_key(&security_id) {
            lobs.insert(security_id, LimitOrderBook::new());
        }
        /*
        Get LOB from Orders securityId then create new order and print book after
         */
        if let Some(limit_ord_book) = lobs.get_mut(&security_id) {
            LimitOrderBook::add_order(
                limit_ord_book,
                rng.gen::<u64>(),
                ord_type,
                side,
                amounts[i],
                price,
                curr_time,
            );
        }
    }

    for key in lobs.keys() {
        println!("Limit Order Book: {}", key);
        lobs.get(&key).unwrap().print_book();
    }

    // end time
    let end = SystemTime::now();

    let since_the_epoch = end.duration_since(start).expect("Time went backwards");

    println!(
        "Time taken to execute {} orders: {}ms",
        num_of_orders,
        since_the_epoch.as_millis()
    );
}
