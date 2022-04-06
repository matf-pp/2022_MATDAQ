# Client

## Build and run commands

- There are two separate binaries that are being built `request-creator` and `price-display` 
- Use the following commands for building and running 
  - `request-creator`
    ```bash
      $ make rc
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
  - `price-display`
    ```bash
      $ make run-pd
    ```
## Project structure

```
client
├── bin/
│   ├── price-display
│   └── request-creator
├── cmd/
│   ├── price-display/
│   │   └── main.go
│   └── request-creator/
│       └── main.go
├── pkg/
│   └── README.md
├── go.mod
├── Makefile
└── README.md
```

    

