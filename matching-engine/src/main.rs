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
use limit_order_book::request_handlers::request_money_update;
use rand::Rng;
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
        let amount: u32 = rng.gen_range(1..=1000);
        prices.push(price);
        amounts.push(amount)
    }

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
        let mut price = prices[i];
        let curr_time = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .expect("Time went backwards")
            .as_nanos();
        let security_id = rng.gen_range(1..=2);

        // Set limit price of a Market order according to the side it's on
        if ord_type == OrderType::Market {
            if side == Side::Sell {
                price = 0;
            } else {
                price = i32::MAX;
            }
        }
        // Check if LOB contains security with given security_id
        // If it doesn't then add it to the LOB
        if !lobs.contains_key(&security_id) {
            lobs.insert(security_id, LimitOrderBook::new(&security_id.to_string()));
        }
        let order = Order {
            security_id,
            sender_id,
            order_id: rng.gen::<u64>(),
            side,
            order_type: ord_type,
            amount: amounts[i],
            limit_price: price,
            time: curr_time,
        };
        /*
        Get LOB from Orders security_id then create new order
         */
        if let Some(limit_ord_book) = lobs.get_mut(&security_id) {
            /*
               If an order is aggressive then execute it
               Else add it to the order book
            */
            if limit_ord_book.is_aggressive(&order) {
                let ord = order.clone();
                // Send request to update user's money and return
                // request_money_update(order.sender_id, limit_ord_book.execute_order(ord))
            } else {
                limit_ord_book.add_order(order);
            }
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn adding_limit_buy_order() {
        let mut limit_order_book = LimitOrderBook::new("1");
        let order_id = 1;
        let sender_id = [0; 20];
        let time = 0;
        let security_id = 0;

        let ord1 = Order {
            order_id,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Limit,
            side: Side::Buy,
            amount: 45,
            limit_price: 98,
        };

        limit_order_book.add_order(ord1);

        let lob = limit_order_book.to_string();

        assert_eq!(
            lob,
            "Order id: 1\t Amount: 45\t Limit Price: 98\t Type: Limit\t Side: Buy"
        );
    }
    #[test]
    fn adding_limit_sell_order() {
        let mut limit_order_book = LimitOrderBook::new("1");
        let order_id = 1;
        let sender_id = [0; 20];
        let time = 0;
        let security_id = 0;
        let ord1 = Order {
            order_id,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Limit,
            side: Side::Sell,
            amount: 45,
            limit_price: 100,
        };

        limit_order_book.add_order(ord1);

        let lob = limit_order_book.to_string();

        assert_eq!(
            lob,
            "Order id: 1\t Amount: 45\t Limit Price: 100\t Type: Limit\t Side: Sell"
        );
    }

    #[test]
    fn adding_market_buy_and_sell_order() {
        let mut limit_order_book = LimitOrderBook::new("1");
        let sender_id = [0; 20];
        let time = 0;
        let security_id = 0;

        let ord1 = Order {
            order_id: 1,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Market,
            side: Side::Buy,
            amount: 45,
            limit_price: 0,
        };

        let ord2 = Order {
            order_id: 2,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Market,
            side: Side::Sell,
            amount: 90,
            limit_price: 100,
        };

        limit_order_book.add_order(ord1);
        limit_order_book.add_order(ord2);

        let lob = limit_order_book.to_string();

        assert_eq!(
            lob,
            "Order id: 2\t Amount: 90\t Limit Price: 0\t Type: Market\t Side: Sell\
        Order id: 1\t Amount: 45\t Limit Price: 2147483647\t Type: Market\t Side: Buy"
        );
    }

    #[test]
    fn executing_limit_sell_order() {
        let mut limit_order_book = LimitOrderBook::new("1");
        let sender_id = [0; 20];
        let time = 0;
        let security_id = 0;

        let ord1 = Order {
            order_id: 1,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Market,
            side: Side::Buy,
            amount: 45,
            limit_price: 0,
        };

        let ord2 = Order {
            order_id: 2,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Market,
            side: Side::Buy,
            amount: 14,
            limit_price: 0,
        };

        let ord3 = Order {
            order_id: 3,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Limit,
            side: Side::Sell,
            amount: 14,
            limit_price: 100,
        };

        limit_order_book.add_order(ord1);
        limit_order_book.add_order(ord2);
        limit_order_book.execute_order(ord3);

        let lob = limit_order_book.to_string();

        assert_eq!(
            lob,
            "Order id: 2\t Amount: 14\t Limit Price: 2147483647\t Type: Market\t Side: Buy\
        Order id: 1\t Amount: 31\t Limit Price: 2147483647\t Type: Market\t Side: Buy"
        );
    }

    #[test]
    fn random_test() {
        let mut limit_order_book = LimitOrderBook::new("1");
        let sender_id = [0; 20];
        let time = 0;
        let security_id = 0;

        let ord1 = Order {
            order_id: 1,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Limit,
            side: Side::Buy,
            amount: 45,
            limit_price: 99,
        };

        let ord2 = Order {
            order_id: 2,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Limit,
            side: Side::Sell,
            amount: 10,
            limit_price: 101,
        };

        let ord3 = Order {
            order_id: 3,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Limit,
            side: Side::Sell,
            amount: 15,
            limit_price: 100,
        };

        let ord4 = Order {
            order_id: 4,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Limit,
            side: Side::Buy,
            amount: 20,
            limit_price: 101,
        };

        let ord5 = Order {
            order_id: 5,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Limit,
            side: Side::Sell,
            amount: 46,
            limit_price: 98,
        };

        let ord6 = Order {
            order_id: 6,
            security_id,
            sender_id,
            time,
            order_type: OrderType::Limit,
            side: Side::Buy,
            amount: 13,
            limit_price: 95,
        };

        limit_order_book.add_order(ord1);
        limit_order_book.add_order(ord2);
        limit_order_book.add_order(ord3);
        limit_order_book.execute_order(ord4);
        limit_order_book.execute_order(ord5);
        limit_order_book.execute_order(ord6);

        let lob = limit_order_book.to_string();

        assert_eq!(
            lob,
            "Order id: 2\t Amount: 5\t Limit Price: 101\t Type: Limit\t Side: Sell\
        Order id: 5\t Amount: 1\t Limit Price: 98\t Type: Limit\t Side: Sell\
        Order id: 6\t Amount: 13\t Limit Price: 95\t Type: Limit\t Side: Buyy"
        );
    }
}
