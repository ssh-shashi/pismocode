Transaction Service
A RESTful transaction service written in Go, designed to manage accounts and transactions with maintainability, simplicity, and testability in mind. This service provides basic endpoints to create and retrieve accounts, and log transactions.

Features
Create Account: Register a new account with a unique document number.
Retrieve Account: Fetch details of an account by its ID.
Log Transactions: Record transactions linked to an account with various operation types (purchase, withdrawal, credit voucher).
Dockerized Deployment: Ready-to-use Docker configuration for streamlined deployment.
Unit Testing: Basic tests included for core functionality.
Repository
Access the source code at: Transaction Service Repository

Prerequisites
Go 1.19+
Docker (optional for containerized deployment)
Git to clone the repository
Setup Instructions
Clone the Repository

git clone https://github.com/ssh-shashi/pismocode.git
cd pismocode
Run Locally (Without Docker)
Initialize the Go module and download dependencies:


go mod tidy
Run the application:


go run *.go
The service will start at http://localhost:8080.

Run with Docker
Build the Docker image:


docker build -t transaction-service .
Run the container:


docker run -p 8080:8080 transaction-service
Alternatively, use Docker Compose:


docker-compose up --build
API Endpoints
1. Create Account
Endpoint: POST /accounts
Registers a new account with a unique document number.

Request Body:

{
  "document_number": "12345678900"
}
Response:

{
  "account_id": 1,
  "document_number": "12345678900"
}
2. Retrieve Account
Endpoint: GET /accounts/{account_id}
Fetches details of an account by its ID.

Response:

{
  "account_id": 1,
  "document_number": "12345678900"
}
3. Create Transaction
Endpoint: POST /transactions
Logs a transaction for an account.

Request Body:


{
  "account_id": 1,
  "operation_type_id": 4,
  "amount": 123.45
}
Response:

json
Copy code
{
  "transaction_id": 1,
  "account_id": 1,
  "operation_type_id": 4,
  "amount": 123.45,
  "event_date": "2024-11-25T14:48:00Z"
}
Operation Types
ID	Description	Amount Type
1	Normal Purchase	Negative
2	Purchase with Installments	Negative
3	Withdrawal	Negative
4	Credit Voucher	Positive
Testing
Unit tests are provided for core functionality.

Run tests:


go test ./...
Sample output:


=== RUN   TestCreateAccount
--- PASS: TestCreateAccount (0.00s)
=== RUN   TestGetAccount
--- PASS: TestGetAccount (0.00s)
PASS
ok      github.com/ssh-shashi/pismocode  0.003s
