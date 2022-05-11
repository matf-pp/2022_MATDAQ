// use std::error::Error;
// use std::time::Duration;

// use futures::stream;
// use rand::rngs::ThreadRng;
// use rand::Rng;
// use tokio::time;
// use tonic::transport::Channel;
use tonic::{Request, Response};

use user_service::user_client::UserClient;
use user_service::{DecreaseMoneyRequest, DecreaseMoneyResponse};

pub mod user_service {
    tonic::include_proto!("user_service");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let USERNAME = String::from("Ilija");
    const MONEY_AMOUNT: i32 = 8;

    const USER_SERVICE_HOST: &str = "http://127.0.0.1:9000";
    let mut client = UserClient::connect(USER_SERVICE_HOST).await?;

    let response: Response<DecreaseMoneyResponse> = client
        .decrease_money(Request::new(DecreaseMoneyRequest {
            username: USERNAME,
            money_amount: MONEY_AMOUNT,
        }))
        .await?;

    println!("RESPONSE = {:?}", response);

    // println!("\n*** SERVER STREAMING ***");
    // print_features(&mut client).await?;

    Ok(())
}
