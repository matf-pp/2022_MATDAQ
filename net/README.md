# Net

## About

TODO

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


    

