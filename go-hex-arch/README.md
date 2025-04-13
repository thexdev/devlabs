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
    "ID": "dc9701bf-830c-4f2c-8022-3cb533ddc9c0",
    "Expression": {
      "A": 10,
      "B": 20,
      "Operator": "-"
    },
    "Result": -10
  },
  {
    "ID": "69ef713c-f343-461d-b22b-e7b9246dc625",
    "Expression": {
      "A": 10,
      "B": 20,
      "Operator": "+"
    },
    "Result": 30
  },
  {
    "ID": "5826d45e-529e-4011-b0f4-13363a17a851",
    "Expression": {
      "A": 10,
      "B": 20,
      "Operator": "+"
    },
    "Result": 30
  }
]
```
