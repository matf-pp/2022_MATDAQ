# Client

## Build and run commands

- There are three separate binaries that are being built `request-creator`, `request-creator-server` and `price-display` 
- Use the following commands for building and running 
  - `request-creator`
    ```bash
      $ make rc
    ```
  - `request-creator-server`
    ```bash
      $ make rs
    ```
  - `price-display`
    ```bash
      $ make pd
    ```
- In case you need more granular control
  - Use the following command for clearing old build files and building new ones
    ```bash
      $ make clean
      $ make build
    ```
  - Use the following command for running the binaries without worrying about whether they need to be rebuilt 
  - `request-creator`
    ```bash
      $ make run-rc
    ```
  - `request-creator-server`
    ```bash
      $ make run-pd
    ```
  - `price-display`
    ```bash
      $ make run-pd
    ```
## Project structure

```
client
├── bin/
│   ├── price-display
│   ├── request-creator
│   └── request-creator-server
├── cmd/
│   ├── price-display/
│   │   └── main.go
│   ├── request-creator/
│   │   └── main.go
│   └── request-creator-server/
│       └── main.go
├── pkg/
│   └── README.md
├── go.mod
├── Makefile
└── README.md
```

    

