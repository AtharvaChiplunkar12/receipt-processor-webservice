# Receipt Processor
This project is built using **Go** and the **Gin** framework.
It provides an endpoint to process receipts and another to retrieve reward points for a given receipt ID.

## Overview

The application is designed to:
- Accept receipt data in JSON format.
- Validate and process the receipt to assign a unique ID.
- Store receipt information in memory.
- Calculate and return reward points based on the receipt details.

## Tech Stack
- **Go (Golang)** - Main programming language
- **Gin** - HTTP Web Framework
- **Docker** - Containerization and deployment

---

## Installation & Setup

### **Prerequisites**
- Install [Go](https://go.dev/doc/install) (>= 1.18)
- Ensure `GOPATH` and `GOROOT` are set up correctly

### Clone the Repository
```sh
git clone https://github.com/AtharvaChiplunkar12/receipt-processor-webservice.git
cd src
```

### Running the project with Go
#### Install Dependencies
```sh
go mod tidy
```

---

#### Running the Server

To start the API server, run:
```sh
go run main.go
```
The server will start on `http://localhost:8080`.

### Running the project with Docker

#### Build the Docker image
```sh
docker build -t my-receipt-processor .
```

#### Run the container
```sh
docker run -p 8080:8080 my-receipt-processor .
```
Now your app is running inside a Docker container on port 8080.

---

## API Endpoints

### 1. Process a Receipt
- **Endpoint:** `POST /receipts/process`
- **Description:** Accepts a receipt in JSON format, validates it, and returns a unique receipt ID.
- **Request Body:**
```json
{
  "retailer": "Target",
  "purchaseDate": "2023-10-05",
  "purchaseTime": "14:33",
  "total": "35.70",
  "items": [
    { "shortDescription": "Milk", "price": "3.50" },
    { "shortDescription": "Bread", "price": "2.80" }
  ]
}
```
- **Response:**
```json
{
  "id": "a1b2c3d4"
}
```

### 2. Get Receipt Points
- **Endpoint:** `GET /receipts/:id/points`
- **Description:** Retrieves the reward points for a given receipt ID.
- **Response (Success):**
```json
{
  "points": 120
}
```
- **Response (Not Found):**
```json
{
  "error": "Receipt ID not found"
}
```

---

## Project Structure
```
src/
├── controllers/        # API endpoint handlers
├── dtos/               # Data transfer objects (DTOs)
├── models/             # Data models and response structures  
├── processes/          # Business logic for receipt processing
├── routing/            # API route definitions and setup 
├── unit_tests/         # unit tests for testing API for different cases
├── main.go             # Application entry point
├── go.mod              # Go modules file (dependencies and module path)  
├── go.sum              # Dependencies checksum file 
└── README.md           # Project documentation
```

---

## Testing the API
You can test the API using **cURL** or **Postman**.

### Using cURL
#### Process a Receipt:
```sh
curl -X POST "http://localhost:8080/receipts/process" -H "Content-Type: application/json" -d '{
  "retailer": "Target",
  "purchaseDate": "2023-10-05",
  "purchaseTime": "14:33",
  "total": "35.70",
  "items": [
    { "shortDescription": "Milk", "price": "3.50" },
    { "shortDescription": "Bread", "price": "2.80" }
  ]
}'
```
#### Get Points for a Receipt ID:
```sh
curl -X GET "http://localhost:8080/receipts/a1b2c3d4/points"
```

### Using Postman
1. Open Postman.
2. Create a **POST** request to `http://localhost:8080/receipts/process`.
3. Set **Body** -> **Raw** -> **JSON** and paste the request JSON.
4. Send the request and note the returned `id`.
5. Create a **GET** request to `http://localhost:8080/receipts/{id}/points` and send.

**Note** -  I have written a few unit tests to check the API for different cases. To run them, use:
```sh
go test ./unit_tests/...
```

---

## Future Improvements
- Implement a persistent database (e.g., PostgreSQL, MongoDB)
- Add authentication and authorization
- Expand validation rules for receipts


