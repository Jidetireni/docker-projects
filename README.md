# Dockerized Go API Project

This is a simple Go API project that provides endpoints for serving articles.

## Prerequisites

Before running this project, ensure you have the following installed:

- Go programming language: [Download Go](https://golang.org/dl/)
- Docker: [Download Docker](https://www.docker.com/products/docker-desktop)

## Getting Started

Follow these steps to run the project:

1. Clone this repository to your local machine:

```bash
git clone git@github.com:Jidetireni/docker-projects.git
```

2. Navigate to the project directory:

```bash
cd <project-directory>
```

3. Build the Docker image:

```bash
docker-compose up -d .
```

4. Access the API:

Open your web browser and navigate to http://localhost:8081 to access the API.

## Usage

- The API provides the following endpoints:

  - `/`: Homepage endpoint.
  - `/articles`: Endpoint for retrieving articles.

- The API returns JSON responses for the `/articles` endpoint.

