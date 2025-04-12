# Go Hexagonal Architecture

Hexagoal Architecture implementations in Go. Here the service is serving to handle calculator operations. From basic arithmetic to advanced problem.

## Core Functionality

### Basic Arithmetic Operations

- Addition (`+`), Subtraction (`-`), Multiplication (`*`), Division (`/`)
- Input validation (e.g., division by zero, invalid operators).

### Calculation History

- Automatically logs every successful calculation.
- In-memory storage (temporary, resets on server restart).
- Thread-safe concurrency handling.

## API Endpoints

### `POST /calculate`

**Request**

```json
{
  "expression": {
    "a": 1,
    "b": 5,
    "operator": "+"
  },
  "result": 6
}
```

**Response**

```json
{
  "expression": {
    "a": 1,
    "b": 5,
    "operator": "?"
  },
  "result": 0,
  "error": "invalid operator"
}
```

### `POST /history`

**Response**

```json
[
  {
    "id": "...", // UUID
    "request": { "a": 10, "b": 5, "operator": "+" },
    "response": { "result": 15, "error": "" },
    "timestamp": 1710000000
  }
]
```
