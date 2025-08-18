# Lab 3: Network Programming with REST and gRPC

## Table of Contents

- [Table of Contents](#table-of-contents)
- [HTTP and REST APIs](#http-and-rest-apis)
- [Building a Multi-Endpoint Web Server](#building-a-multi-endpoint-web-server)
  - [1. GET /](#1-get-)
  - [2. GET /counter](#2-get-counter)
  - [3. GET /fizzbuzz?value=N](#3-get-fizzbuzzvaluen)
  - [4. GET /github](#4-get-github)
  - [5. Default Handler](#5-default-handler)
- [Tasks](#tasks)
  - [Example Usage](#example-usage)
- [Testing](#testing)
- [Troubleshooting](#troubleshooting)
  - [Debugging Tips](#debugging-tips)

## HTTP and REST APIs

Last updated: August 2025

HTTP is the foundation of data communication on the web.
A REST (Representational State Transfer) API is a popular architectural style for designing networked applications that communicate over HTTP.

Before getting started, you should familiarize yourself with the following concepts:

- **HTTP Methods**: GET, POST, PUT, DELETE, and their semantic meanings
- **Status Codes**: 200 (OK), 404 (Not Found), 301 (Moved Permanently), etc.
- **URL Routing**: How to map URLs to specific handler functions
- **Request/Response Cycle**: How HTTP requests are processed and responses generated

Useful resources:

- [HTTP/1.1 Specification](https://tools.ietf.org/html/rfc7231)
- [Go HTTP Server Documentation](https://pkg.go.dev/net/http)
- [REST API Design Best Practices](https://restfulapi.net/)

## Building a Multi-Endpoint Web Server

This assignment involves building a web server in Go using the standard `net/http` package.
Your server should handle multiple endpoints with different functionalities, demonstrating various aspects of web development.

The server should implement the following endpoints:

### 1. GET /

- Returns a simple "Hello World!" message
- Status: 200 OK

### 2. GET /counter

- Maintains a global counter that increments with each request
- Returns the current counter value in the format: `counter: X`
- Status: 200 OK

### 3. GET /fizzbuzz?value=N

- Implements the classic FizzBuzz game
- For multiples of 3: return "fizz"
- For multiples of 5: return "buzz"
- For multiples of both 3 and 5: return "fizzbuzz"
- For other numbers: return the number itself
- For invalid input: return "not an integer"
- For missing value: return "no value provided"
- Status: 200 OK

### 4. GET /github

- Redirects to `http://www.github.com`
- Status: 301 Moved Permanently

### 5. Default Handler

- For any other path, return a 404 Not Found response
- Status: 404 Not Found

## Tasks

1. **Implement the Server Structure**

   - Create a `Server` struct with necessary fields to maintain state
   - Implement `NewServer()` function to initialize the server
   - The server should not automatically start listening; that's handled by the main function

2. **Implement ServeHTTP Method**

   - The `Server` must implement the `http.Handler` interface
   - Route requests to appropriate handlers based on the URL path
   - Handle query parameters for the fizzbuzz endpoint

3. **Implement URL Routing**

   - Parse the request URL to determine which endpoint to serve
   - Use appropriate HTTP status codes for different scenarios
   - Ensure proper error handling for invalid requests

4. **Implement State Management**

   - The counter endpoint requires maintaining state between requests
   - Consider thread safety - multiple clients may access the server concurrently
   - Use appropriate synchronization mechanisms if needed

5. **Implement FizzBuzz Logic**

   - Parse the query parameter `value` from the URL
   - Implement the FizzBuzz algorithm correctly
   - Handle edge cases (non-integer values, missing parameters)

6. **Test Your Implementation**

   - Start the server and test each endpoint using curl or a web browser
   - Verify that the counter increments correctly across multiple requests
   - Test the FizzBuzz endpoint with various inputs
   - Ensure the redirect works properly

### Example Usage

```bash
# Start the server
go run .

# Test the hello endpoint (root)
curl http://localhost:8080/
# Expected: Hello World!

# Test the counter endpoint
curl http://localhost:8080/counter
# Expected: counter: 1

curl http://localhost:8080/counter
# Expected: counter: 2

# Test the fizzbuzz endpoint
curl http://localhost:8080/fizzbuzz?value=3
# Expected: fizz

curl http://localhost:8080/fizzbuzz?value=5
# Expected: buzz

curl http://localhost:8080/fizzbuzz?value=15
# Expected: fizzbuzz

curl http://localhost:8080/fizzbuzz?value=7
# Expected: 7

curl http://localhost:8080/fizzbuzz?value=abc
# Expected: not an integer

curl http://localhost:8080/fizzbuzz
# Expected: no value provided

# Test the redirect endpoint
curl -i http://localhost:8080/github
# Expected: 301 status with Location header

# Test 404 behavior
curl http://localhost:8080/nonexistent
# Expected: 404 page not found
```

## Testing

The assignment includes automated tests that verify your implementation:

- **Basic Functionality Tests**: Test each endpoint individually
- **Counter State Tests**: Verify the counter maintains state correctly
- **FizzBuzz Logic Tests**: Test various FizzBuzz scenarios
- **Error Handling Tests**: Test 404 responses and invalid inputs
- **Integration Tests**: Test combinations of requests

Run the tests using:

```bash
go test -v
```

## Troubleshooting

### Debugging Tips

- Use `fmt.Printf` or `log.Printf` to debug request handling
- Check the request URL and method using `r.URL.Path` and `r.Method`
- Use `http.Error()` to send error responses with proper status codes
- Test endpoints individually before running the full test suite
