## Hexagonal Architecture — Practical Notes

This note captures insights and implementations from discussions around building applications using pure Hexagonal Architecture (HexArch), without mixing in Domain-Driven Design (DDD), unless mentioned. It provides context and practical code snippets to solidify understanding.

---

### 1. **Basic Structure of HexArch**

**Context**: Understanding the separation of concerns

- **Domain Layer**: Contains the business logic.
- **Application Layer**: Orchestrates the business logic using domain entities and services.
- **Adapters (Primary/Secondary)**:
  - **Primary Adapters**: Input interfaces like HTTP, gRPC, CLI.
  - **Secondary Adapters**: Output interfaces like databases, external APIs, messaging systems.

```
/internal
  /domain
    order.go
  /application
    place_order.go
  /adapter
    /http
      handler.go
    /repo
      order_repo.go
```

---

### 2. **Port and Use Case Design**

**Context**: Who defines the port, and where should it live?

- **Ports are defined in the application layer**.
- **Domain layer should never know about application or ports**.
- Ports act as contracts for what the application can do.

```go
// application/port/order.go
package port

type PlaceOrder interface {
    Execute(cmd PlaceOrderCommand) error
}
```

Use case implementation:

```go
// application/usecase/place_order.go
package usecase

type PlaceOrderUseCase struct {
    repo domain.OrderRepository
}

func (uc *PlaceOrderUseCase) Execute(cmd PlaceOrderCommand) error {
    // orchestrate domain
    return uc.repo.Save(...)
}
```

---

### 3. **General Interface Strategy**

**Context**: What if we have multiple use cases with similar patterns?

You can introduce a command pattern with a general interface:

```go
type UseCase[C any, R any] interface {
    Execute(command C) (R, error)
}
```

But be cautious: avoid over-abstraction. Start concrete; extract generic later.

---

### 4. **Aggregating Use Cases**

**Context**: What if handler uses multiple use cases?

Instead of injecting each use case individually:

```go
type OrderApplication struct {
    PlaceOrder usecase.PlaceOrder
    CancelOrder usecase.CancelOrder
}
```

Now the handler can depend on `OrderApplication`.

---

### 5. **Ports Reflect Features, Not Technical Components**

**Context**: Instead of thinking in technical layers, think like the actor.

Ports should reflect the **actor’s behavior**, e.g.:

```go
type Arithmetic interface {
    Add(a, b int) int
    Sub(a, b int) int
}
```

Later, you can group into `ArithmeticApplication` if needed.

---

### 6. **MVP vs Scale-Up**

**MVP Phase**:

- Use concrete implementations.
- Avoid premature abstractions.
- Focus on shipping features.

**Scale-Up Phase**:

- Introduce clear port interfaces.
- Split apps into modules per domain.
- Use grouped services (e.g., `ArithmeticApplication`).

---

### 7. **Repository Design**

**Context**: Working with one or multiple DBs

- When multiple DBs involved, avoid general `Repository` interfaces.
- Instead, separate per domain or persistence boundary:

```go
type OrderRepository interface {
    Save(order Order) error
    FindByID(id string) (Order, error)
}

// if using both Mongo and PostgreSQL
// create separate impls in adapter layer
```

---

### 8. **History or Stateful Domain Objects**

**Context**: Can domain hold state like stacks?

Yes, domain entities can have in-memory state:

```go
type History struct {
    stack []string
}

func (h *History) Push(entry string) {
    h.stack = append(h.stack, entry)
}
```

However, use case must orchestrate fetching from and persisting to the repository.

---

### 9. **Summary**

- Ports are defined by application, not domain.
- Adapters call application via ports.
- Domain contains pure business rules.
- Application layer orchestrates domain and infra.
- Concrete implementations live in the adapter layer.
- Think features first; group later when scaling.
- Avoid over-engineering in MVP.

Let the system evolve.
