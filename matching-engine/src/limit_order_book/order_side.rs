use rand::Rng;
use rand::distributions::{Distribution, Standard};

#[derive(Clone, Copy)]
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
