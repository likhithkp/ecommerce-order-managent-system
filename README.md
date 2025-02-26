# E-Commerce Order Management System

## Overview
This project is a microservices-based **Order Management System** built using **Golang, PostgreSQL, Kafka**, and **net/http**. It efficiently handles order processing, payment verification, and inventory updates while ensuring **scalability** and **real-time communication** between services.

## Architecture
The system is divided into three microservices:
1. **Order Service** - Handles order creation and forwards payment processing.
2. **Payment Service** - Processes payments and confirms transactions.
3. **Inventory Service** - Updates inventory once payment is successful and sends confirmation.

### Workflow:
1. The **Order Service** receives an order request via REST API.
2. It produces an event to Kafka for the **Payment Service**.
3. The **Payment Service** processes the payment and produces an event for the **Inventory Service**.
4. The **Inventory Service** updates stock and confirms the order.

## Tech Stack
- **Golang** (net/http for APIs)
- **PostgreSQL** (Database)
- **Kafka** (Event streaming for async processing)
- **pgx** (PostgreSQL driver for Go)

## Setup & Installation
### Prerequisites:
- Go 1.20+ installed
- Docker & Docker Compose (optional for Kafka & PostgreSQL)

### Steps:
1. Clone the repository:
   ```sh
   git clone https://github.com/likhithkp/ecommerce-order-management-system.git
   cd order-management-system
   ```
2. Run Kafka and PostgreSQL via Docker:
   ```sh
   docker-compose up -d
   ```
3. Start the services:
   ```sh
   go run order/main.go
   go run payment/main.go
   go run inventory/main.go
   ```

## API Endpoints (Order Service)
- **POST /order** - Create a new order

## Future Enhancements
- Implement real-time notifications.
- Add gRPC for inter-service communication.
- Improve database optimizations.

---
**Author**: Likhith K.P  
