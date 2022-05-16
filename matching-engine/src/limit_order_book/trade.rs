use crate::limit_order_book::order::SecurityId;
use crate::limit_order_book::order_side::Side;

pub struct Trade {
    pub side: Side,
    pub amount: u32,
    pub security_id: SecurityId,
}
