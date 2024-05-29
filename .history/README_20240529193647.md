## Golang Developer Assigment Challenge

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

You can start the container with debugging enabled:
```sh
make docker-build
```

Then, you can attach your debugger to the specified port.

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
- `make docker-up`: Start Docker Compose services.
- `make docker-down`: Stop and remove Docker Compose services.
- `make clean`: Clean binary files.
- `make lint`: Lint the code.

### Directory Structure

```
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
├── docker-compose.yml.orig
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
```