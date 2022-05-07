# Matching Engine

- We're using the latest stable version of `rustc`
- Run the following command to update
    ```bash
    $ rustup update
    ```
- Run the following command to install protoc binary
    ```bash
    $ apt install libprotobuf-dev protobuf-compiler
    ```

## Building and running

- Build & Run
    ```bash
    $ cargo run
    ```
- Build
    ```bash
    $ cargo build
    ```
- ❤️ Rust ❤️

## Implementation details

Limit Order Book is implemented as a struct that holds one Min Heap for sell orders, one Max Heap for buy orders and a HashMap of all the orders.

Currently executes ~250.000 orders per second.

```rust
struct Order {
    order_id: OrderId,
    side: Side,
    ord_type: OrderType,
    amount: u32,
    limit_price: ordered_float::NotNan<f64>,
    time: u64,
}

struct LimitOrderBook {
    buy_side: BinaryHeap<Order>,
    sell_side: BinaryHeap<Reverse<Order>>,
    orders: HashMap<OrderId, Order>
}
```