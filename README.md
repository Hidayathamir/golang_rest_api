# golang_rest_api

Checkout the [task](https://github.com/Hidayathamir/golang_rest_api/blob/main/task/interview_golang_backend.md). This project aims to implement clean code principles by separating concerns into different modules. The following is a brief overview of the project's structure:

- `entity`: Contains the database schema and Data Transfer Objects (DTOs) used for manipulating entities.
- `repository`: Implements the database access layer.
- `usecase`: Contains the business logic of the application.
- `handler`: Links the routers with the corresponding use cases while also check the request.
- `router`: Lists all the routers used in the application.

## 1. Project Setup and Execution

To run the project, you can choose to either use Docker or run it manually.

### 1.1 Running Using Docker

1. Use Docker Compose to set up the database:
   ```
   docker compose up
   ```
2. Run the server using Docker images:
   ```
   docker run --net=host hidayathamir/golang_rest_api
   ```

### 1.2 Running Manually

1. Edit the configuration in `config/config.json` according to your database setup.
2. Run the server using the following command:
   ```
   go run .
   ```

## 2. API Endpoints

The following are the available endpoints in the API.

### 2.1 Adding a Product

To add a product, use the following curl command:

```
curl --location 'http://localhost:8081/products' \
--header 'Content-Type: application/json' \
--data '{
    "name": "product a",
    "price": 123,
    "description": "desc product a",
    "quantity": 2
}'
```

### 2.2 Getting Products

To get products, use the following curl command.

1. Latest product.
   ```
   curl --location 'http://localhost:8081/products?sort_by=created_at&sort=desc'
   ```
2. Price low to high.
   ```
   curl --location 'http://localhost:8081/products?sort_by=price&sort=asc'
   ```
3. Price high to low.
   ```
   curl --location 'http://localhost:8081/products?sort_by=price&sort=desc'
   ```
4. Product name a to z.
   ```
   curl --location 'http://localhost:8081/products?sort_by=name&sort=asc'
   ```
5. Product name z to a.
   ```
   curl --location 'http://localhost:8081/products?sort_by=name&sort=desc'
   ```
