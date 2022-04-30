# User service

## TODO - Dimitrije

- Role of the user service is to manage money for users that are participating 

- Each user should have:
  - Unique id 
  - Money amount
  - Something else?
- User service should support:
  - Basic authentication
  - Getting money for some user
  - Decrement money amount for some user
- For now implement each feature as an HTTP endpoint (e.g. REST)

## Redis setup

### Installation

- Version: `6.2.6`
- For local development you'll want to install `redis` 
  - Arch
    ```bash
    $ yay -S redis
    ```
  - Ubuntu
    ```bash
    $ curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

    $ echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

    $ sudo apt-get update
    $ sudo apt-get install redis
    ```
- Start the `systemd` service
  ```bash
  $ sudo systemctl start redis.service
  ```
- Test whether installation is working
  ```bash
  $ redis-cli ping
  ```
  - You should receive `PONG` written to `stdout`

### Redis Insight

- Optionally install [Redis Insight](https://redis.com/redis-enterprise/redis-insight/) which will give you a GUI to the `Redis` database

## Build and run commands

- Use the following commands for building and running 
  - `user-server`
    ```bash
      $ make us
    ```
- In case you need more granular control
  - Use the following command for clearing old build files and building new ones
    ```bash
      $ make clean
      $ make build
    ```
  - Use the following command for running the binaries without worrying about whether they need to be rebuilt 
  - `user-server`
    ```bash
      $ make run-us
    ```
## Project structure

```
user-service
├── bin/
│   └── user-server
├── cmd/
│   └── user-server/
│       └── main.go
├── pkg/
│   └── README.md
├── go.mod
├── Makefile
└── README.md
```


    

