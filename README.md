# 2022_MATDAQ

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/671bfd1715484de7b05928a0cc33991b)](https://app.codacy.com/gh/matf-pp/2022_MATDAQ?utm_source=github.com&utm_medium=referral&utm_content=matf-pp/2022_MATDAQ&utm_campaign=Badge_Grade_Settings)

## About

- MATDAQ is a stock exchange system developed in Rust and Go

## Building and Running

- Use `docker-compose` to build and run the project

```bash
  $ docker-compose up -d 
```

- To connect to a specific client run 
  - `request-creator`
    ```bash
      $ docker attach 2022_matdaq-request-creator-1
    ```
  - `price-display`
    ```bash
      $ docker attach 2022_matdaq-price-display-1
    ```

- To trigger a rebuild of images use

```bash
  $ docker-compose up -d --build
```

- To stop containers and tear down the infrastructure

```bash $ 
  docker-compose down
```

## Dev setup

### Prerequisites

- Install `docker` and `docker-compose`

### Setup

- Navigate to the `dev` folder and run the setup script in order to properly configure your environment 

```bash
  $ cd dev 
  $ ./setup.sh
```
- Currently developed only on `Linux`

## Developers

- [Aleksandar Šmigić](https://github.com/smiga287)
- [Dimitrije Marković](https://github.com/dimitrijemarkovic)
- [Ilija Stojanović](https://github.com/ilija-s)
