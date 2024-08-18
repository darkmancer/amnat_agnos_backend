# Strong Password Recommendation Service

This is a Go-based service that provides recommendations on the number of steps required to make a given password strong. The service checks the password against various criteria and suggests the minimum number of actions needed.

## Features

- Validates password strength based on length, character variety, and repetition.
- Logs request and response data to a PostgreSQL database.
- Dockerized for easy deployment.
- Includes unit tests with table-driven testing.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/doc/install) (version 1.20 or later)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/darkmancer/amnat_agnos_backend.git
cd strong-password-recommendation
```

### 2. Create a .env File

Create a .env file in the root of the project directory with the following content:

```
DB_HOST=postgres
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name

# Application Configuration
APP_PORT=8081

# Nginx Configuration
NGINX_PORT=8080
```

Replace your_db_user, your_db_password, and your_db_name with your desired PostgreSQL credentials.

### 3. Build and Deploy with Docker Compose

Use Docker Compose to build and deploy the application locally.

```
docker-compose up --build
```

### 4. Access the Application

Once the containers are up and running, you can access the service at:

```
http://localhost:8080
```

### 5. Running Unit Tests

To run the unit tests, you can execute the following command:

```
go test ./...
```

### 6. Stopping the Application

To stop the running containers, use:

```
docker-compose down
```

This will stop and remove all the containers associated with the project.
