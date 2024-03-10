# Docker-projects

# Dockerized Go API Project

This is a simple Go API project that provides endpoints for serving articles.

## Prerequisites

Before running this project, ensure you have the following installed:

- Golang
- docker

## Getting Started

Follow these steps to run the project:

1. Clone this repository to your local machine:

git clone git@github.com:Jidetireni/docker-projects.git
```
2. Navigate to the project directory:

cd project_dir
```
3. Build the Docker image:

docker-compose up -d
```
5. Access the API:

Open your web browser and navigate to http://localhost:8081 to access the API.

## Usage

- The API provides the following endpoints:

  - `/`: Homepage endpoint.
  - `/articles`: Endpoint for retrieving articles.

- The API returns JSON responses for the `/articles` endpoint.


