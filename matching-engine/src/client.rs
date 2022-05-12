use tonic::{Request, Response};

use user_service::user_client::UserClient;
use user_service::{DecreaseMoneyRequest, DecreaseMoneyResponse};

pub mod user_service {
    tonic::include_proto!("user_service");
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let username = String::from("Ilija");
    let money_amount: i32 = 8;

    const USER_SERVICE_HOST: &str = "http://user-service:9000";
    let mut client = UserClient::connect(USER_SERVICE_HOST).await?;

    let response: Response<DecreaseMoneyResponse> = client
        .decrease_money(Request::new(DecreaseMoneyRequest {
            username,
            money_amount,
        }))
        .await?;

    println!("RESPONSE = {:?}", response);

    Ok(())
}
