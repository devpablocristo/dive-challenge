# Golang Developer Assignment

This project provides a service that retrieves the Last Traded Price (LTP) of Bitcoin for the following currency pairs:

- BTC/USD
- BTC/CHF
- BTC/EUR

The service is built using Go and provides an API endpoint at `/api/v1/ltp`. The response includes the LTP for the specified currency pairs with time accuracy up to the last minute.

## API Endpoint

### GET /api/v1/ltp

This endpoint retrieves the Last Traded Price of Bitcoin for the specified currency pairs. The request can be made for a single pair or a list of pairs.

#### Request

- **Path:** `/api/v1/ltp`
- **Query Parameters:**
  - `pair`: The currency pair(s) to retrieve the LTP for. Multiple pairs can be specified.

#### Response

The response will be in JSON format with the following structure:

```json
{
  "ltp": [
    {
      "pair": "BTC/CHF",
      "amount": 49000.12
    },
    {
      "pair": "BTC/EUR",
      "amount": 50000.12
    },
    {
      "pair": "BTC/USD",
      "amount": 52000.12
    }
  ]
}
```

## Requirements

- Time accuracy of the data up to the last minute.
- Code hosted in a remote public repository.
- `README.md` includes clear steps to build and run the app.
- Integration tests.
- Dockerized application.
- Documentation.

## Getting Started

### Prerequisites

- Go 1.16+
- Docker
- Docker Compose

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/devpablocristo/dive-challenge.git
   cd dive-challenge
   ```

### Building and Running the Application

#### Using Docker Compose (Recommended)

1. Build and start the Docker containers:
   ```sh
   docker-compose up --build
   ```

2. The application will be available at `http://localhost:8080`.

#### Debugging Inside Docker

To enable debugging inside Docker, ensure that your `Dockerfile` and `docker-compose.yml` are set up to expose the debug port (e.g., 2345) and that Delve is installed in the Docker container.

You can start the container with debugging enabled by using the following command:
```sh
make docker-build
```
This command will build the Docker image and start the container with the necessary configuration for debugging.

If the containers are already built, you can start them without rebuilding by using the `make up` command:
```sh
make up
```
This command will start the Docker Compose services without rebuilding the images, making it quicker if no changes have been made to the Docker setup.

### Running Tests

To run the tests, use the following command:
```sh
make test
```

### Linting the Code

To lint the code, use the following command:
```sh
make lint
```

### Makefile Commands

- `make all`: Build and run the project.
- `make build`: Build the project.
- `make run`: Run the project.
- `make test`: Run tests.
- `make docker-build`: Build and start Docker Compose services.
- `make up`: Start Docker Compose services if they are already built.
- `make docker-down`: Stop and remove Docker Compose services.
- `make clean`: Clean binary files.
- `make lint`: Lint the code.

### Directory Structure

```
.
├── cmd
│   ├── api
│   │   ├── build.go
│   │   ├── config.go
│   │   └── handlers
│   │       ├── presenter
│   │       │   └── ltp.go
│   │       └── rest.go
│   └── main.go
├── docker-compose.yml
├── Dockerfile
├── entrypoint.sh
├── go.mod
├── go.sum
├── internal
│   ├── core
│   │   ├── ltp
│   │   │   ├── api-client_adapter.go
│   │   │   ├── dtos.go
│   │   │   ├── ltp.go
│   │   │   ├── ports.go
│   │   │   └── repository.go
│   │   ├── ports.go
│   │   └── usecase.go
│   └── platform
│       └── stage
│           └── stage.go
├── makefile
└── README.md
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.