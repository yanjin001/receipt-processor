# Receipt Processor

## Overview
The Receipt Processor is a simple web service designed to process shopping receipts and calculate points based on predefined rules.

## Features
- Processes shopping receipts in JSON format.
- Calculates points based on the content of the receipts.
- Generates unique identifiers for receipts and provides an endpoint to query the points awarded.

## Installation

Ensure you have the Go programming environment installed on your system. You can download and install it from the [official Go website](https://golang.org/dl/).

## Running the Service

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/your-username/receipt-processor.git
   cd receipt-processor
   ```

2. Run the service:

   ```bash
   go run main.go
   ```

   The service will start on `http://localhost:8080`.

   ***Example***:
   <img width="528" alt="Screen Shot 2024-01-21 at 9 19 47 PM" src="https://github.com/yanjin001/receipt-processor/assets/62823935/f92bad41-631a-4749-a5a3-378ed81dcde9">


## Usage

- Send a POST request to `/receipts/process` to process a receipt.
- Send a GET request to `/receipts/{id}/points` to retrieve the points for a specific receipt.

## API Documentation

- **POST `/receipts/process`**

  **Payload**: Receipt JSON
  
  ```json
  {
      "retailer": "Store Name",
      "purchaseDate": "YYYY-MM-DD",
      "purchaseTime": "HH:MM",
      "items": [
          {
              "shortDescription": "Item 1",
              "price": "Price 1"
          },
          // ... more items ...
      ],
      "total": "Total Amount"
  }
  ```

  **Response**: JSON containing an id for the receipt.

  ***Example***:
  <img width="1047" alt="Screen Shot 2024-01-21 at 9 21 41 PM" src="https://github.com/yanjin001/receipt-processor/assets/62823935/aa1b673e-f78d-48ea-9f2e-b044e94f24e7">


- **GET `/receipts/{id}/points`**

  Retrieve the number of points awarded for a specific receipt.

  **Response**: JSON object containing the number of points.

  ***Example***:
  <img width="1015" alt="Screen Shot 2024-01-21 at 9 23 02 PM" src="https://github.com/yanjin001/receipt-processor/assets/62823935/9dc33bdf-9d6e-4ad3-b32a-62db87bcc186">

