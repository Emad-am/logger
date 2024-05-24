# Logger

Logger is a logging service with Telegram logging support, written in Golang with Redis and MongoDB.

## Features

- Written in Go for high performance
- Supports logging to Telegram
- Uses Redis for fast, in-memory operations
- Stores logs in MongoDB
- Docker support for easy deployment

## Requirements

- Golang
- Docker
- Redis
- MongoDB

## Getting Started

### Clone the repository

```bash
git clone https://github.com/Emad-am/logger.git
cd logger
```

## Build the project

```bash
go build
```

## Run with Docker

### Ensure you have Docker installed and running

```bash
docker-compose up --build
```

## Usage

### Configuration

Modify the docker-compose.yml file to set up your InfluxDB and Redis configurations.

### Running the application

Run the application using Docker as described above, or build and run it directly using Golang.

### License

This project is licensed under the MIT License. See the LICENSE file for details.

```bash
This README provides an overview, setup instructions, and usage guidelines for the `feeder` project.
```
