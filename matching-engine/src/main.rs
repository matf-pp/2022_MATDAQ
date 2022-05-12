extern crate ordered_float;
extern crate rand;
extern crate rand_distr;

mod limit_order_book;

use crate::limit_order_book::{
    limit_order_book::LimitOrderBook,
    limit_order_book_manager::LimitOrderBookManager,
    order::{Order, SecurityId, SenderId},
    order_side::Side,
    order_type::OrderType,
};
use rand::Rng;
use rand_distr::{Distribution, Normal};
use std::collections::HashMap;
use std::time::SystemTime;
use std::time::UNIX_EPOCH;

fn main() {
    let num_of_orders = 300;
    let mut rng = rand::thread_rng();
    let mut prices = Vec::with_capacity(num_of_orders);
    let mut amounts = Vec::with_capacity(num_of_orders);

    // generate random prices and amounts
    for _i in 0..num_of_orders {
        let price: i32 = rng.gen_range(180..=220);
        let amount: i32 = rng.gen_range(1..=1000);
        prices.push(price);
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
    //     limit_price: 100,
    //     time: SystemTime::now().duration_since(UNIX_EPOCH).expect("Time went backwards").as_nanos(),
    // };

    /*
    we get new order from Request-Service-Server containing information about new order
    inside there is SecurityId field, based on that field we get the LOB on which we add new order
    if current SecurityId from request doesnt exist, we create new LOB for that request
     */
    let mut lobs: HashMap<SecurityId, LimitOrderBook> = HashMap::new();

    // start time
    let start = SystemTime::now();

    // generate orders
    for i in 0..num_of_orders {
        let side: Side = rand::random();
        let sender_id: [u8; 20] = [0; 20];
        let ord_type: OrderType = rand::random();
        let price = prices[i];
        let curr_time = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .expect("Time went backwards")
            .as_nanos();
        let security_id = rng.gen_range(1..=20);

        // Check if LOB contains security with given security_id
        // If it doesn't then add it to the LOB
        if !lobs.contains_key(&security_id) {
            lobs.insert(security_id, LimitOrderBook::new());
        }
        /*
        Get LOB from Orders security_id then create new order
         */
        if let Some(limit_ord_book) = lobs.get_mut(&security_id) {
            // LimitOrderBook::add_order(
            //     limit_ord_book,
            //     rng.gen::<u64>(),
            //     security_id,
            //     sender_id,
            //     ord_type,
            //     side,
            //     amounts[i],
            //     price,
            //     curr_time,
            // );
        }
    }

    // for key in lobs.keys() {
    //     println!("Limit Order Book: {}", key);
    //     lobs.get(&key).unwrap().print_book();
    // }

    // end time
    let end = SystemTime::now();

    let since_the_epoch = end.duration_since(start).expect("Time went backwards");

    println!(
        "Time taken to execute {} orders: {}ms",
        num_of_orders,
        since_the_epoch.as_millis()
    );
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn adding_limit_buy_order() {
        let mut limit_order_book = LimitOrderBook::new();
        let order_id = 1;
        let sender_id = [0; 20];
        let time = 0;
        let security_id = 0;

        limit_order_book.add_order(
            order_id,
            security_id,
            sender_id,
            OrderType::Limit,
            Side::Buy,
            45,
            100,
            time,
        );

        let lob = limit_order_book.to_string();

        assert_eq!(
            lob,
            "Order id: 1\t Amount: 45\t Limit Price: 100\t Type: Limit\t Side: Buy\n"
        );
    }
    #[test]
    fn adding_limit_sell_order() {
        let mut limit_order_book = LimitOrderBook::new();
        let order_id = 1;
        let sender_id = [0; 20];
        let time = 0;
        let security_id = 0;

        limit_order_book.add_order(
            order_id,
            security_id,
            sender_id,
            OrderType::Limit,
            Side::Sell,
            45,
            100,
            time,
        );

        let lob = limit_order_book.to_string();

        assert_eq!(
            lob,
            "Order id: 1\t Amount: 45\t Limit Price: 100\t Type: Limit\t Side: Sell\n"
        );
    }

    #[test]
    fn adding_market_buy_and_sell_order() {
        let mut limit_order_book = LimitOrderBook::new();
        let sender_id = [0; 20];
        let time = 0;
        let security_id = 0;

        limit_order_book.add_order(
            1,
            security_id,
            sender_id,
            OrderType::Market,
            Side::Buy,
            45,
            0,
            time,
        );

        limit_order_book.add_order(
            2,
            security_id,
            sender_id,
            OrderType::Market,
            Side::Sell,
            90,
            0,
            time,
        );

        let lob = limit_order_book.to_string();

        assert_eq!(
            lob,
            "Order id: 2\t Amount: 90\t Limit Price: 0\t Type: Market\t Side: Sell\n\
        Order id: 1\t Amount: 45\t Limit Price: 2147483647\t Type: Market\t Side: Buy\n"
        );
    }

    #[test]
    fn executing_limit_sell_order() {
        let mut limit_order_book = LimitOrderBook::new();
        let sender_id = [0; 20];
        let time = 0;
        let security_id = 0;

        limit_order_book.add_order(
            1,
            security_id,
            sender_id,
            OrderType::Market,
            Side::Buy,
            45,
            0,
            time,
        );

        limit_order_book.add_order(
            2,
            security_id,
            sender_id,
            OrderType::Market,
            Side::Buy,
            14,
            0,
            time,
        );

        limit_order_book.add_order(
            3,
            security_id,
            sender_id,
            OrderType::Limit,
            Side::Sell,
            14,
            100,
            time,
        );

        let lob = limit_order_book.to_string();

        assert_eq!(
            lob,
            "Order id: 2\t Amount: 14\t Limit Price: 2147483647\t Type: Market\t Side: Buy\n\
        Order id: 1\t Amount: 31\t Limit Price: 2147483647\t Type: Market\t Side: Buy\n"
        );
    }

    #[test]
    fn random_test() {
        let mut limit_order_book = LimitOrderBook::new();
        let sender_id = [0; 20];
        let time = 0;
        let security_id = 0;

        limit_order_book.add_order(
            1,
            security_id,
            sender_id,
            OrderType::Limit,
            Side::Buy,
            45,
            99,
            time,
        );

        limit_order_book.add_order(
            2,
            security_id,
            sender_id,
            OrderType::Limit,
            Side::Sell,
            10,
            101,
            time,
        );

        limit_order_book.add_order(
            3,
            security_id,
            sender_id,
            OrderType::Limit,
            Side::Sell,
            15,
            100,
            time,
        );

        limit_order_book.add_order(
            4,
            security_id,
            sender_id,
            OrderType::Limit,
            Side::Buy,
            20,
            101,
            time,
        );

        limit_order_book.add_order(
            5,
            security_id,
            sender_id,
            OrderType::Limit,
            Side::Sell,
            46,
            98,
            time,
        );

        limit_order_book.add_order(
            6,
            security_id,
            sender_id,
            OrderType::Limit,
            Side::Buy,
            13,
            95,
            time,
        );

        let lob = limit_order_book.to_string();

        assert_eq!(
            lob,
            "Order id: 2\t Amount: 5\t Limit Price: 101\t Type: Limit\t Side: Sell\n\
        Order id: 5\t Amount: 1\t Limit Price: 98\t Type: Limit\t Side: Sell\n\
        Order id: 6\t Amount: 13\t Limit Price: 95\t Type: Limit\t Side: Buy\n"
        );
    }
}
