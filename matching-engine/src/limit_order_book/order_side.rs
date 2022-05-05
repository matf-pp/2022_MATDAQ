use rand::Rng;
use rand::distributions::{Distribution, Standard};
use std::fmt;

#[derive(Clone, Copy, PartialEq)]
pub enum Side {
    Buy,
    Sell,
}

impl Distribution<Side> for Standard {
    fn sample<R: Rng + ?Sized>(&self, rng: &mut R) -> Side {
        match rng.gen_range(0..=1) {
            0 => Side::Sell,
            _ => Side::Buy,
        }
    }
}

impl fmt::Display for Side {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match *self {
            Side::Sell => write!(f, "Sell"),
            Side::Buy => write!(f, "Buy"),
        }
    }
}
