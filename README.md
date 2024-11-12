# PGM3 - Simple GO Web Server

A basic web server written in Go that displays a welcome message. This project demonstrates how to create a simple HTTP server using Go and containerize it with Docker.

## Features

- Basic HTTP server implementation in Go
- Docker containerization
- Exposed on port 8080
- Simple welcome message endpoint

## Prerequisites

- Go 1.23 or higher
- Docker
- Git

## Installation

Clone the repository:

```bash
git clone https://github.com/Anggeloo/PGM3.git
cd PGM3
```

## Running with Docker

1. Build the Docker image:
```bash
docker build -t pgm3 .
```

2. Run the container:
```bash
docker run -p 8080:8080 pgm3
```

3. Access the application at http://localhost:8080

## Project Structure

```
PGM3/
├── PGM3.go        # Main application file
├── Dockerfile     # Docker configuration
├── go.mod         # Go module file
├── go.sum         # Go module checksum
└── README.md      # This file
```

## Accessing the deployed application on Render

1.Open your web browser
2.Visit the following URL:
https://pgm3.onrender.com

3.You will see the Flask application displaying a counter that you can increment by clicking the button.
