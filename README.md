# Go API Server

This project is a simple Go API server that connects to a PostgreSQL database. It is designed to be run in Docker containers, making it easy to deploy and manage.

## Project Structure

```
library-go-api
├── src
│   ├── book
│   │   ├── bookcommon.go
│   │   ├── bookops.go
│   │   └── lendingbook.go
│   ├── config
│   │   ├── constant.go
│   │   ├── db.go
│   │   └── init.go
│   ├── handler
│   │   ├── book.go
│   │   ├── handlercommon.go
│   │   └── user.go
│   ├── models
│   │   ├── bookmodals.go
│   │   ├── usermodels.go
│   │   └── validator.go
│   ├── routes
│   │   └── routesmain.go
│   ├── sql
│   │   └── v1.0
│   │       └── initdb.sql
│   ├── user
│   │   ├── usercommon.go
│   │   └── userops.go
│   ├── utils
│   │   └── common.go
│   └── main.go
├── .dockerignore
├── .gitignore
├── docker-compose.debug.yml
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md            # Project documentation
```

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Setup

1. Clone the repository:
   ```
   git clone <repository-url>
   cd go-api-server
   ```

2. Build and run the containers:
   ```
   docker-compose up --build
   ```

3. (Optional) If you add or update dependencies, run:
   ```
   go mod tidy
   ```

4. Access the API server at `http://localhost:8080`.

## Example API Usage (cURL)

### User APIs

**Create a new user**
```sh
curl -X POST http://localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","mobileNo":"1234567890"}'
```

**Fetch user by mobile number**
```sh
curl "http://localhost:8080/api/v1/user?mobilenum=1234567890"
```

**Update user information**
```sh
curl -X PUT http://localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"userID":"user_xxx","name":"Jane Doe","mobileNo":"1234567890"}'
```

**Delete user by mobile number**
```sh
curl -X DELETE "http://localhost:8080/api/v1/user?mobilenum=1234567890"
```

---

### Book APIs

**Add a new book**
```sh
curl -X POST http://localhost:8080/api/v1/book \
  -H "Content-Type: application/json" \
  -d '{"name":"Book Title","author":"Author Name","genre":"Fiction","desc":"Description here","sku":"BOOKSKU123"}'
```

**Remove a book by book ID**
```sh
curl -X DELETE "http://localhost:8080/api/v1/book?bookid=BOOKID123"
```

**Fetch available books (paginated)**
```sh
curl -X POST http://localhost:8080/api/v1/allbook \
  -H "Content-Type: application/json" \
  -d '{"pageNum":1,"pageSize":10}'
```

---

### Lending APIs

**Lend a book to a user**
```sh
curl -X POST http://localhost:8080/api/v1/rent \
  -H "Content-Type: application/json" \
  -d '{"userID":"user_xxx","sku":"BOOKSKU123"}'
```

**Return a book**
```sh
curl -X POST http://localhost:8080/api/v1/return \
  -H "Content-Type: application/json" \
  -d '{"userID":"user_xxx","bookID":"BOOKID123"}'
```
