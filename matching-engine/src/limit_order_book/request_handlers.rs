use crate::limit_order_book::{order::SecurityId, order::SenderId, order_side::Side};

/*
   Sends request to user-service microservice telling it to update the users
   money amount by given amount.

   Amount can be both positive or negative depending on the side that user traded
   If the user traded on Buy side => money_amount < 0
   If the user traded on Sell side => money_amount > 0

   Note: this function should be called in two cases:
   1) user whose order is being executed at the end calls this function
   2) all the users against which the trade happened should also get their money updated
*/
pub fn request_money_update(user: SenderId, money_amount: i32) {
    if money_amount < 0 {
        println!("\x1b[31m{}\x1b[0m", money_amount);
    } else {
        println!("\x1b[32m{}\x1b[0m", money_amount);
    }
}

/*
   Sends request to price-display microservice telling it that an order in
   the Limit Order Book happened for a given stock, on a given side and the
   amount that was traded
*/
pub fn request_trade(stock: SecurityId, amount_traded: u32, side: Side) {
    if side == Side::Buy {
        println!(
            "Trade happened: Stock id: {}\tAmount traded: {}\tSide: \x1b[32m{}\x1b[0m",
            stock, amount_traded, side
        )
    } else {
        println!(
            "Trade happened: Stock id: {}\tAmount traded: {}\tSide: \x1b[31m{}\x1b[0m",
            stock, amount_traded, side
        )
    }
}

/*
   Sends request to price-display microservice telling it to add new order
   with given parameters
*/
pub fn request_new_order(stock: SecurityId, price: i32, amount: u32, side: Side) {
    if side == Side::Buy {
        println!(
            "New order happened: Stock id: {}\tPrice: {}\tAmount: {}\tSide: \x1b[32m{}\x1b[0m",
            stock, price, amount, side
        )
    } else {
        println!(
            "New order happened: Stock id: {}\tPrice: {}\tAmount: {}\tSide: \x1b[31m{}\x1b[0m",
            stock, price, amount, side
        )
    }
}
