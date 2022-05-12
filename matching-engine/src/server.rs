use std::error::Error;
use std::sync::{Arc, Mutex};
use tokio::sync::mpsc;

use tokio_stream::wrappers::ReceiverStream;
use tonic::{transport::Server, Request, Response, Status};

mod limit_order_book;
use crate::limit_order_book::limit_order_book_manager::{
    matching_engine::matching_engine_server::{MatchingEngine, MatchingEngineServer},
    matching_engine::{
        security_order, CreateOrderRequest, CreateOrderResponse, PublishOrderRequest,
        PublishOrderResponse, PublishTradeRequest, PublishTradeResponse, SecurityOrder,
    },
    LimitOrderBookManager, ReceiverChannel, SenderChannel,
};

pub mod matching_engine {
    tonic::include_proto!("matching_engine");
}

pub struct MatchingEngineService<'a, 'b> {
    order_book_manager: Arc<Mutex<LimitOrderBookManager<'a, 'b>>>,
}

#[tonic::async_trait]
impl MatchingEngine for MatchingEngineService<'static, 'static> {
    async fn create_order(
        &self,
        request: Request<CreateOrderRequest>,
    ) -> Result<Response<CreateOrderResponse>, Status> {
        self.order_book_manager
            .lock()
            .unwrap()
            .create_order(request.into_inner());
        Ok(Response::new(CreateOrderResponse::default()))
    }

    type PublishTradeStream = ReceiverStream<Result<PublishTradeResponse, Status>>;

    async fn publish_trade(
        &self,
        _request: Request<PublishTradeRequest>,
    ) -> Result<Response<Self::PublishTradeStream>, Status> {
        let (tx, rx): (
            SenderChannel<PublishTradeResponse>,
            ReceiverChannel<PublishTradeResponse>,
        ) = mpsc::channel(4);
        // orderBook.addChannel(tx)
        tokio::spawn(async move {
            for i in 1..4 {
                let order_quantity: u32 = i * 10;
                let order_side: i32 = security_order::OrderSide::Buy as i32;
                let security_id: u32 = i;
                let price: i32 = 1;
                let trade = PublishTradeResponse {
                    security_order: Some(SecurityOrder {
                        order_side,
                        price,
                        security_id,
                        order_quantity,
                    }),
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
                        order_side: 0,
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
async fn main() -> Result<(), Box<dyn Error>> {
    let addr = "matching-engine:10000".parse().unwrap();

    println!("MatchingEngineServer listening on: {}", addr);

    let matching_engine = MatchingEngineService {
        order_book_manager: Arc::new(Mutex::new(LimitOrderBookManager::new())),
    };

    let svc = MatchingEngineServer::new(matching_engine);

    Server::builder().add_service(svc).serve(addr).await?;

    Ok(())
}
