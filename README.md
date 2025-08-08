# back-end-golang

This project is a Go-based backend application designed to provide a robust and scalable API service. It leverages several popular Go libraries and frameworks to ensure high performance and maintainability.

## Features

- **RESTful API**: Built using the Gin framework for fast and efficient HTTP request handling.
- **Database Integration**: Utilizes PostgreSQL for data storage, with connection management handled by pgx.
- **Environment Configuration**: Uses `godotenv` for managing environment variables.
- **JSON Handling**: Employs `go-json` for efficient JSON encoding and decoding.
- **JWT Authentication**: Supports JWT-based authentication using `golang-jwt`.
- **Dockerized**: Fully containerized using Docker, with a multi-stage build for optimized image size.

## Getting Started

### Prerequisites

- [Go 1.23.3](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop)

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/keymanzx/back-end-golang.git
   cd back-end-golang
   ```

2. **Set up environment variables**:
   Create a `.env` file in the root directory and configure your environment variables as needed.

3. **Build and run the application**:
   Use Docker Compose to build and run the application:
   ```bash
   docker compose up --build
   ```

   The application will be available at `http://localhost:8000`.

4. **Run Make**:
   The Makefile in this project provides several commands to automate common tasks. Here are some useful targets:
   - `make build`: Compiles the Go application.
   - `make dev`: Starts the Go application in development mode.
   - `make test`: Runs the test suite.
   - `make clean`: Cleans up build artifacts.
   - `make run`: Runs the application locally without Docker.
   - `make swagger`: Generates the Swagger documentation for the API.
   
   To execute a make command, use:
   ```bash
   make <target>
   ```

   ### Using Swagger

   Swagger is a tool that helps in creating interactive API documentation, making it easier for developers to test and understand the API.

   1. **Generate Swagger Documentation**:
      - Use the command `make swagger` to generate Swagger documentation for your API.
      - This command will create Swagger documentation files in JSON or YAML format that can be used with Swagger UI.

   2. **Open Swagger UI**:
      - After generating the Swagger documentation, you can use Swagger UI to view and test the API.
      - Open a browser and go to `http://localhost:8000/swagger/index.html` to access Swagger UI.

   3. **Test the API**:
      - In Swagger UI, you can view details of each API endpoint, including required parameters and response examples.
      - You can test the API directly from Swagger UI by filling in the necessary information and clicking the "Execute" button.

   Using Swagger makes API development and testing efficient and straightforward.
