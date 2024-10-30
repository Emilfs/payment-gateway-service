# Payment Gateway Service

The Payment Gateway Service is a Go-based service designed to handle fund deposits, fund withdrawals, and transaction status updates through various payment gateways.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Docker and Docker Compose are installed on your machine.
- Go (version 1.23 or later) is installed on your machine.

## Building the Service

To build the Payment Gateway Service, follow these steps:

1. Clone the repository to your local machine:

```bash
git clone https://github.com/Emilfs/payment-gateway-service.git
cd payment-gateway-service
```


2. Build the Docker image:

```bash
docker-compose build
```

This command builds a Docker image for the service based on the `Dockerfile` in the project root. It also caches Go modules for faster subsequent builds.

## Running the Service

After building the image, you can run the service using Docker Compose:

```bash
docker-compose up
```

This command starts the service container. By default, the service listens on port 8080. You can access the API endpoints at `http://localhost:8080`.

## Running Tests

To run the Go tests for the service, follow these steps:

1. Ensure you are in the project root directory.

2. Run the tests using the Go toolchain:

```bash
go test ./...
```

This command recursively runs all tests in the project. You should see output indicating the status of each test suite.

## Reference
- [Design document](https://docs.google.com/document/d/1bvihhXCYbx_UjuGyZP7jI3H3c3JJC6B8OA0hUY_iRso/edit?usp=sharing)