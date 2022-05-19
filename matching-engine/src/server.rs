use std::error::Error;
use std::sync::{Arc, Mutex};
use tokio::sync::mpsc;

use tokio_stream::wrappers::ReceiverStream;
use tonic::{transport::Server, Request, Response, Status};

mod limit_order_book;
use crate::limit_order_book::limit_order_book_manager::{
    matching_engine::matching_engine_server::{MatchingEngine, MatchingEngineServer},
    matching_engine::{
        CreateOrderRequest, CreateOrderResponse, PublishOrderRequest,
        PublishOrderResponse, PublishTradeRequest, PublishTradeResponse,
    },
    LimitOrderBookManager, ReceiverChannel, SenderChannel,
};

pub mod matching_engine {
    tonic::include_proto!("matching_engine");
}

pub struct MatchingEngineService {
    order_book_manager: Arc<Mutex<LimitOrderBookManager>>,
}

#[tonic::async_trait]
impl MatchingEngine for MatchingEngineService {
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
        self.order_book_manager
            .lock()
            .unwrap()
            .add_trade_channel_sender(Box::new(tx));

        Ok(Response::new(ReceiverStream::new(rx)))
    }

    type PublishOrderCreationStream = ReceiverStream<Result<PublishOrderResponse, Status>>;

    async fn publish_order_creation(
        &self,
        _request: Request<PublishOrderRequest>,
    ) -> Result<Response<Self::PublishOrderCreationStream>, Status> {
        let (tx, rx): (
            SenderChannel<PublishOrderResponse>,
            ReceiverChannel<PublishOrderResponse>,
        ) = mpsc::channel(4);
        self.order_book_manager
            .lock()
            .unwrap()
            .add_order_channel_sender(Box::new(tx));

        Ok(Response::new(ReceiverStream::new(rx)))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let addr = "0.0.0.0:10000".parse().unwrap();

    println!("MatchingEngineServer listening on: {}", addr);

    let matching_engine = MatchingEngineService {
        order_book_manager: Arc::new(Mutex::new(LimitOrderBookManager::new())),
    };

    let svc = MatchingEngineServer::new(matching_engine);

    Server::builder().add_service(svc).serve(addr).await?;

    Ok(())
}
