# Receipt Processor

## Project Structure
```
receipt-processor/
|-- Dockerfile
|-- README.md
|-- api.yml
|-- examples/
|   |-- receipt-109.json
|   |-- receipt-18.json
|   |-- receipt-28.json
|   |-- receipt-32.json
|   |-- receipt-15.json
|   |-- receipt-19.json
|   |-- receipt-19.json
|   |-- receipt-31.json
|-- go.mod
|-- go.sum
|-- handlers/
|   |-- handler.go
|-- main.go
|-- models/
|   |-- receipt_model.go
|-- receipt-processor
|-- utils/
|   |-- calculator.go
```
- **Dockerfile**: Contains all the commands a user could call to assemble an image.
- **examples**: Contains JSON files used as test cases with their names indicating the expected output. For example, the output for `receipt-63.json` should be `63`.
- **handlers**: Contains `handler.go` which manages the HTTP request handlers.
- **main.go**: The entry point of the application that ties everything together.
- **models**: Holds data models, for instance, `receipt_model.go` which defines the structure for receipt data.
- **receipt-processor**: The compiled executable for the application.
- **utils**: Includes utility functions like `calculator.go` which contains logic for calculating points.

### Prerequisites
- Go (at least version 1.18) or Docker

### Installation and Running the Service

#### Running Locally with Go
1. Clone the repository:
   ```bash
   git clone https://github.com/MichaelWong1024/receipt-processor-challenge
   cd receipt-processor-challenge
   ```

2. Build the application:
   ```bash
   go build -o receipt-processor
   ```

3. Run the application:
   ```bash
   ./receipt-processor
   ```
   The service will start on port 8080.

#### (Optional) Running with Docker
1. Ensure Docker is installed and running on your machine.

2. Build the Docker image:
   ```bash
   docker build -t receipt-processor-app .
   ```

3. Run the application in a Docker container:
   ```bash
   docker run -d -p 8080:8080 receipt-processor-app
   ```
   This will start the service in the background at http://localhost:8080.

### API Usage

#### Process Receipts
- **Endpoint**: `POST /receipts/process`
- **Payload**: JSON object representing the receipt.
- **Response**: JSON containing a unique ID for the processed receipt.

   Example Request (receipt-28.json):
   ```bash
   curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
     "retailer": "Target",
     "purchaseDate": "2022-01-01",
     "purchaseTime": "13:01",
     "items": [
       {"shortDescription": "Mountain Dew 12PK", "price": 6.49},
       {"shortDescription": "Emils Cheese Pizza", "price": 12.25},
       {"shortDescription": "Knorr Creamy Chicken", "price": 1.26},
       {"shortDescription": "Doritos Nacho Cheese", "price": 3.35},
       {"shortDescription": "Klarbrunn 12-PK 12 FL OZ", "price": 12.00}
     ],
     "total": 35.35
   }'
   ```

   Example Response:
   ```json
   { "id": "7429e645-4b17-47f3-bdec-bd861aba9ad6" }
   ```

#### Get Points
- **Endpoint**: `GET /receipts/{id}/points`
- **Response**: JSON object containing the number of points awarded.

   Example Request and Response for Obtaining Points:
   ```bash
   curl http://localhost:8080/receipts/{id}/points
   {"points":28}
   ```

   Another Test Case (receipt-109.json):
   ```bash
   curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
     "retailer": "M&M Corner Market",
     "purchaseDate": "2022-03-20",
     "purchaseTime": "14:33",
     "items": [{"shortDescription": "Gatorade", "price": 2.25} x4],
     "total": 9.00
   }'
   {"id":"50fd2d35-6adf-4fd5-92d9-ccae6f00a045"}
   curl http://localhost:8080/receipts/{50fd2d35-6adf-4fd5-92d9-ccae6f00a045}/points
   {"points":109}
   ```

## Points Calculation Rules
- One point for every alphanumeric character in the retailer name.
- 50 points if the total is a round dollar amount with no cents.
- 25 points if the total is a multiple of $0.25.
- 5 points for every two items on the receipt.
- Additional points based on item description length and purchase timing specifics.

## All Test Case Commands
### Receipt 109: M&M Corner Market
```bash
curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "total": 9.00,
  "items": [
    {"shortDescription": "Gatorade", "price": 2.25},
    {"shortDescription": "Gatorade", "price": 2.25},
    {"shortDescription": "Gatorade", "price": 2.25},
    {"shortDescription": "Gatorade", "price": 2.25}
  ]
}'
```

### Receipt 15: Walgreens
```bash
curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
  "retailer": "Walgreens",
  "purchaseDate": "2022-01-02",
  "purchaseTime": "08:13",
  "total": 2.65,
  "items": [
    {"shortDescription": "Pepsi - 12-oz", "price": 1.25},
    {"shortDescription": "Dasani", "price": 1.40}
  ]
}'
```

### Receipt 18: Walmart
```bash
curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
  "retailer": "Walmart",
  "purchaseDate": "2022-11-11",
  "purchaseTime": "16:01",
  "total": 6.08,
  "items": [
    {"shortDescription": "Milk", "price": 3.09},
    {"shortDescription": "Bread", "price": 2.99}
  ]
}'
```

### Receipt 19: Costco
```bash
curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
  "retailer": "Costco",
  "purchaseDate": "2022-12-24",
  "purchaseTime": "13:55",
  "total": 42.95,
  "items": [
    {"shortDescription": "Tissue Box", "price": 4.99},
    {"shortDescription": "Detergent", "price": 13.49},
    {"shortDescription": "Eggs", "price": 10.49},
    {"shortDescription": "Juice", "price": 7.99},
    {"shortDescription": "Cookies", "price": 5.99}
  ]
}'
```

### Receipt 32: CVS Pharmacy
```bash
curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
  "retailer": "CVS Pharmacy",
  "purchaseDate": "2022-07-31",
  "purchaseTime": "15:45",
  "total": 2.49,
  "items": [
    {"shortDescription": "Chips", "price": 1.50},
    {"shortDescription": "Soda", "price": 0.99}
  ]
}'
```

### Receipt 28: Target
```bash
curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "total": 35.35,
  "items": [
    {"shortDescription": "Mountain Dew 12PK", "price": 6.49},
    {"shortDescription": "Emils Cheese Pizza", "price": 12.25},
    {"shortDescription": "Knorr Creamy Chicken", "price": 1.26},
    {"shortDescription": "Doritos Nacho Cheese", "price": 3.35},
    {"shortDescription": "Klarbrunn 12-PK 12 FL OZ", "price": 12.00}
  ]
}'
```

### Receipt 31: Target
```bash
curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
  "retailer": "Target",
  "purchaseDate": "2022-01-02",
  "purchaseTime": "13:13",
  "total": 1.25,
  "items": [
    {"shortDescription": "Pepsi -

 12-oz", "price": 1.25}
  ]
}'
```

### Receipt 38: 7-Eleven
```bash
curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{
    "retailer": "7-Eleven",
    "purchaseDate": "2022-05-02",
    "purchaseTime": "10:20",
    "items": [
      {"shortDescription": "Ice Cream", "price": 4.25},
      {"shortDescription": "Water", "price": 1.75},
      {"shortDescription": "Candy", "price": 0.75}
    ],
    "total": 6.75
}'
```
