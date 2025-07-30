# 🛠️ Order Management API – Golang

This is a personal project I built in Go to explore and apply backend development concepts, particularly around concurrency, clean architecture, and idiomatic Go code. It implements a simple in-memory API for managing customer orders.

The solution includes RESTful endpoints, structured request validation, basic error handling, and extensibility for future persistence or integration.

---

## 💡 Project Purpose

The main goals I focused on:

- Practicing Go standard libraries and patterns.
- Applying concurrency with goroutines and channels in a meaningful way (e.g., simulating order processing).
- Writing clean, testable, and maintainable code.
- Designing a flexible and minimal architecture that can scale with minor changes.

---

## 📦 What the API Does

This API allows:

- Creating new orders.
- Retrieving a specific order by ID.
- Listing all orders.
- Updating the status of an order (`pending`, `paid`, `delivered`, or `cancelled`).

---

## 🧱 Order Structure

```json
{
  "id": "uuid",
  "customer": "John Doe",
  "items": [
    {
      "product": "Keyboard",
      "quantity": 2
    }
  ],
  "total": 120.00,
  "status": "pending",
  "created_at": "2025-06-25T12:00:00Z"
}
