use crate::limit_order_book::order_side::Side;
use crate::limit_order_book::order_type::OrderType;
use std::fmt;
use std::cmp::Ordering;

pub type OrderId = u64;
pub type SecurityId = i32;
pub type SenderId = [u8; 20];

#[derive(Clone, Copy)]
pub struct Order {
    pub order_id: OrderId,
    pub side: Side,
    pub ord_type: OrderType,
    pub amount: u32,
    // TODO: don't forget to change this after changing limit_price from f64 to i32
    pub limit_price: ordered_float::NotNan<f64>,
    pub time: u128,
    pub security_id: SecurityId,
    pub sender_id: SenderId,
}

impl Order {
    pub fn is_filled(&self) -> bool {
        self.amount == 0
    }
}

impl fmt::Display for Order {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "Order id: {}\t Amount: {}\t Limit Price: {}\t Type: {}\t Side: {}", self.order_id, self.amount, self.limit_price, self.ord_type, self.side)
    }
}

impl Ord for Order {
    fn cmp(&self, other: &Self) -> Ordering {
        let res = self.limit_price.cmp(&other.limit_price);
        if res == Ordering::Equal {
            return self.time.cmp(&other.time)
        }
        res
    }
}

impl PartialOrd for Order {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.limit_price.cmp(&other.limit_price))
    }
}

impl PartialEq for Order {
    fn eq(&self, other: &Self) -> bool {
        self.limit_price == other.limit_price && self.time == other.time
    }
}

impl Eq for Order { }
