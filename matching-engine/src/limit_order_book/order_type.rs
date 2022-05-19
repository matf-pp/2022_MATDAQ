use rand::distributions::{Distribution, Standard};
use rand::Rng;
use std::fmt;

#[derive(Clone, Copy, PartialEq)]
pub enum OrderType {
    Limit,
    Market,
}

/*
When generating random OrderType's 1 in 10 orders should be Market order
 */
impl Distribution<OrderType> for Standard {
    fn sample<R: Rng + ?Sized>(&self, rng: &mut R) -> OrderType {
        match rng.gen_range(1..=10) {
            1 => OrderType::Market,
            _ => OrderType::Limit,
        }
    }
}

impl fmt::Display for OrderType {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match *self {
            OrderType::Limit => write!(f, "Limit"),
            OrderType::Market => write!(f, "Market"),
        }
    }
}
