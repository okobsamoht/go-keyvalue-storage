## Go Key-Value Storage API

This is a simple key-value storage application built with Go, using the Gin web framework and LevelDB for data storage. It provides a RESTful API to create, retrieve, and delete key-value pairs.

### Features

- Store key-value pairs in a LevelDB database.
- RESTful API with the following endpoints:
    - POST /kv: Create a new key-value pair.
    - GET /kv/:key: Retrieve the value associated with a given key.
    - DELETE /kv/:key: Delete a key-value pair.

### Prerequisites

- Go (v1.16 or higher)
- LevelDB Go bindings
- Gin web framework

### Installation

1. Clone the repository:

   git clone https://github.com/okobsamoht/go-keyvalue-storage.git

   cd go-keyvalue-storage

2. Install the required dependencies:

   go mod tidy

   or

   go get github.com/syndtr/goleveldb/leveldb

   go get -u github.com/gin-gonic/gin


### Usage

1. Start the server:

   go run main.go

   The server will run on http://localhost:8080.

2. Use a tool like Postman or curl to interact with the API.

### API Endpoints

#### Create a Key-Value Pair

- Endpoint: POST /kv
- Request Body:

  {
  "key": "your_key",
  "value": "your_value"
  }

- Example:

  `curl -X POST http://localhost:8080/kv -H "Content-Type: application/json" -d '{"key": "name", "value": "John Doe"}'`

- Response:

  {
  "status": "success"
  }

#### Retrieve a Value

- Endpoint: GET /kv/:key
- Example:

  `curl http://localhost:8080/kv/name`

- Response:

  {
  "key": "name",
  "value": "John Doe"
  }

#### Delete a Key-Value Pair

- Endpoint: DELETE /kv/:key
- Example:

  `curl -X DELETE http://localhost:8080/kv/name`

- Response:

  {
  "status": "deleted"
  }

### Error Handling

The API returns appropriate HTTP status codes and error messages for various scenarios, such as:

- 400 Bad Request: If the request body is missing required fields.
- 404 Not Found: If the requested key does not exist.
- 500 Internal Server Error: If there is an issue with the database operations.

### License

This project is unlicensed.

## Contributing

Feel free to submit issues or pull requests if you have suggestions or improvements!

Acknowledgments

- Go
- Gin
- LevelDB
