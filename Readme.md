# Identity Server (Go)

A learning-focused identity server implemented in Go, providing core authentication and identity management capabilities such as user registration, authentication, logout, and token revocation.

The project is designed with **maintainability, clean architecture, and domain-driven design (DDD) principles** in mind, with special attention to **token lifecycle and revocation strategies**.

> ⚠️ This project is not intended for production use.  
> It was built to deeply understand authentication flows, security tradeoffs, and backend system design in Go.

---

## ✨ Features

- User registration
- User authentication (login)
- Secure password hashing using bcrypt
- Logout support
- Token revocation API
- Revoked token garbage collection (GC)
- DDD-inspired layered architecture
- Repository abstractions to decouple business logic from persistence

---

## 🧱 Architecture Overview

The project follows a **DDD-inspired layered architecture**:


### Layer Responsibilities

- **Domain**
  - Core entities (e.g. User)
  - Repository interfaces
  - Business concepts independent of infrastructure

- **Application**
  - Orchestrates use cases (registration, login, logout, token revocation)
  - Coordinates domain logic via interfaces

- **Infrastructure**
  - MySQL persistence
  - Token storage & revocation logic
  - Password hashing implementation

- **Presentation**
  - HTTP handlers
  - Request validation and response mapping

Dependencies flow **inward**, keeping the domain isolated from external concerns.

---

## 🔐 Token Revocation Design

Token revocation is implemented to allow immediate invalidation of issued tokens (e.g. on logout).

Revoked tokens are persisted in the database and checked during authentication to ensure invalid tokens cannot be reused.

---

## 🗑️ Revoked Tokens GC & Partitioning Design

### Problem Statement

Token revocation tables can grow very quickly at scale.

Example scenario:
- 1M daily active users (DAU)
- ~96 token rotations per user per day
- Millions of revoked tokens inserted daily

Without proper garbage collection, a single revocation table can grow to **hundreds of millions of rows**, negatively impacting query performance and cleanup operations.

---

### Initial Approach: Row-Based GC

The initial idea was:
- Store all revoked tokens in a single table
- Periodically delete expired tokens in batches
- Tune batch size and GC frequency to keep up with insertion rate

While workable at smaller scales, this approach introduces:
- Increasing `SELECT` cost as table size grows
- Complex throughput math to ensure GC keeps up
- Hard limits at very large DAU where deletion cost exceeds feasible GC capacity

---

### Final Approach: Time-Based Partitioning

To simplify GC and ensure predictable performance, a **time-based partitioning strategy** was chosen.

#### Design

- Revoked tokens are stored in **daily tables**:
- Tokens are written to the table corresponding to their issuance date
- Revocation checks query only the relevant daily table
- Tables older than the token expiration window (e.g. 7 days) are **dropped entirely**

#### Benefits

- Constant-time revocation checks
- No table scans for GC
- Garbage collection becomes an **O(1) operation** (`DROP TABLE`)
- Performance does not degrade with total historical token volume
- Scales linearly with DAU

This approach avoids complex GC throughput calculations and keeps database behavior predictable under load.

---

### Security Considerations

- Table names are derived from parsed and validated token dates
- Dynamic SQL is limited to table selection only
- Token values and user identifiers are always parameterized
- Optional validation can enforce strict table name patterns (e.g. `revoked_tokens_\d{8}`)

---

## 🚀 Running the Project

### Prerequisites

- Go (1.20+ recommended)
- MySQL

### Steps

1. Clone the repository
2. Configure the database connection
3. Run the application:

```bash
go run main.go

---

