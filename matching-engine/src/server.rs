use matching_engine::matching_engine_server::{MatchingEngine, MatchingEngineServer};
use matching_engine::{
    CreateOrderRequest, CreateOrderResponse, PublishOrderRequest, PublishOrderResponse,
    PublishTradeRequest, PublishTradeResponse, SecurityOrder,
};
use tokio::sync::mpsc;

use tokio_stream::wrappers::ReceiverStream;
use tonic::{transport::Server, Request, Response, Status};

pub mod matching_engine {
    tonic::include_proto!("matching_engine");
}

#[derive(Debug)]
pub struct MatchingEngineService {}

#[tonic::async_trait]
impl MatchingEngine for MatchingEngineService {
    async fn create_order(
        &self,
        _request: Request<CreateOrderRequest>,
    ) -> Result<Response<CreateOrderResponse>, Status> {
        // create order
        Ok(Response::new(CreateOrderResponse::default()))
    }

    type PublishTradeStream = ReceiverStream<Result<PublishTradeResponse, Status>>;

    async fn publish_trade(
        &self,
        _request: Request<PublishTradeRequest>,
    ) -> Result<Response<Self::PublishTradeStream>, Status> {
        let (tx, rx) = mpsc::channel(4);
        tokio::spawn(async move {
            for i in 1..4 {
                let trade = PublishTradeResponse {
                    security_id: i,
                    order_quantity: i,
                };
                println!(" -> send {:?}", trade);
                tx.send(Ok(trade)).await.unwrap();
            }

            println!("/// done sending");
        });

        Ok(Response::new(ReceiverStream::new(rx)))
    }

    type PublishOrderCreationStream = ReceiverStream<Result<PublishOrderResponse, Status>>;

    async fn publish_order_creation(
        &self,
        _request: Request<PublishOrderRequest>,
    ) -> Result<Response<Self::PublishOrderCreationStream>, Status> {
        let (tx, rx) = mpsc::channel(4);
        tokio::spawn(async move {
            for i in 1..4 {
                let order = PublishOrderResponse {
                    security_order: Some(SecurityOrder {
                        price: 0,
                        security_id: i,
                        order_quantity: i,
                    }),
                };
                println!(" -> send {:?}", order);
                tx.send(Ok(order)).await.unwrap();
            }

            println!("/// done sending");
        });

        Ok(Response::new(ReceiverStream::new(rx)))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:10000".parse().unwrap();

    println!("MatchingEngineServer listening on: {}", addr);

    let matching_engine = MatchingEngineService {};

    let svc = MatchingEngineServer::new(matching_engine);

    Server::builder().add_service(svc).serve(addr).await?;

    Ok(())
}
