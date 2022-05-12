use std::collections::HashMap;
use tokio::sync::mpsc::{Receiver, Sender};

use std::convert::TryInto;

// TODO: find a way to import from server.rs
pub mod matching_engine {
    tonic::include_proto!("matching_engine");
}

use matching_engine::{CreateOrderRequest, PublishOrderResponse, PublishTradeResponse};

use tonic::Status;

use crate::limit_order_book::{
    limit_order_book::LimitOrderBook,
    order::{Order, SecurityId},
    order_side::Side,
    order_type::OrderType,
};

use std::time::SystemTime;
use std::time::UNIX_EPOCH;

pub type SenderChannel<T> = Sender<Result<T, Status>>;
pub type ReceiverChannel<T> = Receiver<Result<T, Status>>;

pub struct LimitOrderBookManager<'a, 'b> {
    order_books: HashMap<SecurityId, LimitOrderBook>,
    trade_channel_senders: Vec<&'a SenderChannel<PublishTradeResponse>>,
    order_channel_senders: Vec<&'b SenderChannel<PublishOrderResponse>>,
}

impl<'a, 'b> LimitOrderBookManager<'a, 'b> {
    pub fn new() -> LimitOrderBookManager<'a, 'b> {
        // TODO: setup all limit order books
        LimitOrderBookManager {
            order_books: HashMap::from([
                (1, LimitOrderBook::new("AAPL")),
                (2, LimitOrderBook::new("AMZN")),
                (3, LimitOrderBook::new("MSFT")),
            ]),
            trade_channel_senders: Vec::new(),
            order_channel_senders: Vec::new(),
        }
    }

    pub fn add_trade_channel_sender(&mut self, sender: &'a SenderChannel<PublishTradeResponse>) {
        self.trade_channel_senders.push(sender);
    }

    pub fn add_order_channel_sender(&mut self, sender: &'b SenderChannel<PublishOrderResponse>) {
        self.order_channel_senders.push(sender);
    }

    pub fn create_order(&mut self, order_req: CreateOrderRequest) {
        let security_order = order_req.security_order.unwrap();
        let limit_price = security_order.price;
        let security_id = security_order.security_id;
        let amount = security_order.order_quantity;
        let side = match security_order.order_side {
            0 => Side::Buy,
            _ => Side::Sell,
        };
        let order_type = match order_req.order_type {
            0 => OrderType::Limit,
            _ => OrderType::Market,
        };
        let order_id = order_req.order_id;
        let sender_id = (|| -> [u8; 20] {
            order_req.sender_id.try_into().unwrap_or_else(|v: Vec<u8>| {
                panic!("Expected a Vec of length 20 but it was {}", v.len())
            })
        })();
        let time = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .unwrap()
            .as_nanos();

        let order = Order {
            limit_price,
            security_id,
            amount,
            side,
            order_type,
            order_id,
            sender_id,
            time,
        };

        if !self.order_books.contains_key(&security_id) {
            println!("Missing security_id: {}", security_id);
            return;
        }

        if let Some(order_book) = self.order_books.get_mut(&security_id) {
            println!("Added order with security_id: {}", security_id);
            order_book.add_order(order);
        }
    }
}
