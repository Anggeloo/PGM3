## PGM3 - Basic HTTP Server in Go

## 📖 Description

PGM3 is a simple HTTP server implemented in Go. On startup, the server listens on port 8080 and responds with the message:
“Welcome to the GO language!”.

## Note: This project currently does not include a Dockerfile. However, it can be easily adapted to run in a Docker container if needed.

## 🚀 Features.

- Implements an HTTP server listening on port 8080.
- Responds with a simple message when accessing the main / path.
- Basic error handling for server startup.

## 🛠️ Prerequisites
Before you start, make sure you have installed:

- Go (version 1.19 or higher).

## ⚙️ Installation and execution

1. Clone this repository on your local machine:

```bash
git clone https://github.com/Anggeloo/PGM3.git
```

2. Navigate to the project directory:

```bash
cd PGM3
```

3. Run the program:

```bash
go run main.go
```

4. Open your web browser and visit the URL:

```bash
http://localhost:8080
```

5. You will see the message:
“Welcome to the GO language!”

## 🌐 Port customization.
If you want to change the port, simply update the port variable in the PGM3.go file:

```bash
port := “:8080” // Change the 8080 to the desired port.
```

## 🌐 Online project.

The server is deployed and accessible at: https://pgm3.onrender.com
