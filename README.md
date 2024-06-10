# Receipt Processor

## Getting Started

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

#### Running with Docker
1. Ensure Docker is installed and running on your machine.

2. Build the Docker image:
   ```bash
   docker build -t receipt-processor-app .
   ```

3. Run the application in a Docker container:
   ```bash
   docker run -d -p 8080:8080 receipt-processor-app
   ```
   This will start the service in the background, available at http://localhost:8080.

### API Usage

#### Process Receipts
- **Endpoint**: `POST /receipts/process`
- **Payload**: JSON object representing the receipt.
- **Response**: JSON containing a unique ID for the processed receipt.

   Example Request:
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

   Another Test Case:
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
