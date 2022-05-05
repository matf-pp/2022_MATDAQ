#[derive(Clone, PartialEq, ::prost::Message)]
pub struct SecurityOrder {
    #[prost(sint32, tag="1")]
    pub price: i32,
    #[prost(uint32, tag="2")]
    pub security_id: u32,
    #[prost(uint32, tag="3")]
    pub order_quantity: u32,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CreateOrderRequest {
    #[prost(message, optional, tag="1")]
    pub security_order: ::core::option::Option<SecurityOrder>,
    #[prost(uint64, tag="2")]
    pub order_id: u64,
    #[prost(enumeration="create_order_request::OrderType", tag="3")]
    pub order_type: i32,
    #[prost(bytes="vec", tag="4")]
    pub sender_id: ::prost::alloc::vec::Vec<u8>,
}
/// Nested message and enum types in `CreateOrderRequest`.
pub mod create_order_request {
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
    #[repr(i32)]
    pub enum OrderType {
        /// TODO: check if this is correct
        Limit = 0,
        Market = 1,
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct CreateOrderResponse {
    #[prost(int32, tag="1")]
    pub money: i32,
}
/// in the future the client should probably specify a list of securityIds it's interested in following
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PublishTradeRequest {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PublishTradeResponse {
    #[prost(uint32, tag="1")]
    pub security_id: u32,
    #[prost(uint32, tag="2")]
    pub order_quantity: u32,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PublishOrderRequest {
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PublishOrderResponse {
    #[prost(message, optional, tag="1")]
    pub security_order: ::core::option::Option<SecurityOrder>,
}
/// Generated client implementations.
pub mod matching_engine_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    #[derive(Debug, Clone)]
    pub struct MatchingEngineClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl MatchingEngineClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: std::convert::TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> MatchingEngineClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> MatchingEngineClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            MatchingEngineClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with `gzip`.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_gzip(mut self) -> Self {
            self.inner = self.inner.send_gzip();
            self
        }
        /// Enable decompressing responses with `gzip`.
        #[must_use]
        pub fn accept_gzip(mut self) -> Self {
            self.inner = self.inner.accept_gzip();
            self
        }
        pub async fn create_order(
            &mut self,
            request: impl tonic::IntoRequest<super::CreateOrderRequest>,
        ) -> Result<tonic::Response<super::CreateOrderResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/user_service.MatchingEngine/CreateOrder",
            );
            self.inner.unary(request.into_request(), path, codec).await
        }
        pub async fn publish_trade(
            &mut self,
            request: impl tonic::IntoRequest<super::PublishTradeRequest>,
        ) -> Result<
                tonic::Response<tonic::codec::Streaming<super::PublishTradeResponse>>,
                tonic::Status,
            > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/user_service.MatchingEngine/PublishTrade",
            );
            self.inner.server_streaming(request.into_request(), path, codec).await
        }
        pub async fn publish_order_creation(
            &mut self,
            request: impl tonic::IntoRequest<super::PublishOrderRequest>,
        ) -> Result<
                tonic::Response<tonic::codec::Streaming<super::PublishOrderResponse>>,
                tonic::Status,
            > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/user_service.MatchingEngine/PublishOrderCreation",
            );
            self.inner.server_streaming(request.into_request(), path, codec).await
        }
    }
}
