# pismocode

Transaction Service
A RESTful transaction service written in Go, designed to manage accounts and transactions with maintainability, simplicity, and testability in mind. This service provides basic endpoints to create and retrieve accounts, and log transactions.

Features
Create Account: Register a new account with a unique document number.
Retrieve Account: Fetch details of an account by its ID.
Log Transactions: Record transactions linked to an account with various operation types (purchase, withdrawal, credit voucher).
Dockerized Deployment: Ready-to-use Docker configuration for streamlined deployment.
Unit Testing: Basic tests included for core functionality.
Repository
Access the source code at: https://github.com/ssh-shashi/pismocode

Prerequisites
Go 1.19+: Install from Go Downloads.
Docker: For containerized deployment (optional).
Git: To clone the repository.
Setup Instructions
1. Clone the Repository
sh
Copy code
git clone https://github.com/ssh-shashi/pismocode.git
cd pismocode
2. Run Locally (Without Docker)
Initialize the Go module and download dependencies:

sh
Copy code
go mod tidy
Run the application:

sh
Copy code
go run *.go
The service will start at http://localhost:8080.

3. Run with Docker
Build the Docker image:

sh
Copy code
docker build -t transaction-service .
Run the container:

sh
Copy code
docker run -p 8080:8080 transaction-service
Alternatively, use docker-compose:

sh
Copy code
docker-compose up --build
API Endpoints
1. Create Account
Endpoint: POST /accounts
Registers a new account with a unique document number.

Request Body:

json
Copy code
{
  "document_number": "12345678900"
}
Response:

json
Copy code
{
  "account_id": 1,
  "document_number": "12345678900"
}
2. Retrieve Account
Endpoint: GET /accounts/{account_id}
Fetches details of an account by its ID.

Response:

json
Copy code
{
  "account_id": 1,
  "document_number": "12345678900"
}
3. Create Transaction
Endpoint: POST /transactions
Logs a transaction for an account.

Request Body:

json
Copy code
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
Operation Types:

1: Normal Purchase (negative amount)
2: Purchase with Installments (negative amount)
3: Withdrawal (negative amount)
4: Credit Voucher (positive amount)
Testing
Unit tests are provided for core functionality.

Run tests:

sh
Copy code
go test ./...
Example output:

plaintext
Copy code
=== RUN   TestCreateAccount
--- PASS: TestCreateAccount (0.00s)
=== RUN   TestGetAccount
--- PASS: TestGetAccount (0.00s)
PASS
ok      github.com/ssh-shashi/pismocode  0.003s
