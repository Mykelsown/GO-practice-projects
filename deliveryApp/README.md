# Go Mock Delivery App

A mock delivery application built in Go to practice core language concepts including **interfaces**, **structs**, **methods**, and project structuring.

---

## Purpose

This project was built as a hands-on exercise to deepen understanding of:

- Defining and implementing **interfaces**
- Organizing logic with **structs** and **methods**
- Composing behavior through **interface polymorphism**
- Structuring a Go program cleanly within a single package

---

## Architecture Overview

The app simulates a checkout flow for a delivery order. Three core contracts (interfaces) drive the system, each with multiple interchangeable implementations.

```
Order
  └── CheckoutSystem
        ├── PaymentProcessor  →  Card | Wallet | Cash
        ├── DeliveryPattern   →  Rider | Drone
        └── Notifier          →  Email | SMS | Push
```

---

## Interfaces (Contracts)

| Interface | Method(s) | Description |
|---|---|---|
| `PaymentProcessor` | `Pay(amount float64) string` | Handles payment processing |
| `DeliveryPattern` | `Delivery(order Order) string`, `TrackLocation() string` | Manages order dispatch and tracking |
| `Notifier` | `Notify(order Order) string` | Sends delivery notifications to the customer |

---

## Structs & Implementations

### Payment (`PaymentProcessor`)
| Struct | Field | Behaviour |
|---|---|---|
| `PayWithCard` | `CardNumber string` | Pays using a card number |
| `PayWithWallet` | `TxnID string` | Pays using a wallet transaction ID |
| `PayWithCash` | — | Pays with cash, no extra fields needed |

### Delivery (`DeliveryPattern`)
| Struct | Fields | Behaviour |
|---|---|---|
| `RiderDelivery` | `BikersId int`, `Location string` | Dispatches a bike rider; tracks at **10km away** |
| `DroneDelivery` | `DroneId int`, `Location string` | Dispatches a drone; tracks at **5km away** |

### Notification (`Notifier`)
| Struct | Field | Behaviour |
|---|---|---|
| `EmailNotification` | `EmailAddy string` | Sends an email notification |
| `SmsNotification` | `PhoneNumber string` | Sends an SMS notification |
| `PushNotification` | `PhoneIp string` | Sends a push notification |

---

## Checkout System

`CheckoutSystem` is the central coordinator. It holds one implementation of each interface and runs the full checkout pipeline when `Checkout(order Order)` is called.

```go
type CheckoutSystem struct {
    payment  PaymentProcessor
    delivery DeliveryPattern
    notify   Notifier
}
```

**Checkout flow:**
1. Processes payment via `payment.Pay()`
2. Dispatches the order via `delivery.Delivery()` and fetches location via `delivery.TrackLocation()`
3. Sends a notification via `notify.Notify()`

---

## Running the App

**Prerequisites:** Go 1.18+

```bash
# Clone the repo
git clone <your-repo-url>
cd go-mock-delivery-app

# Run the app
go run main.go
```

### Sample Output

Using the default `main()` configuration (card payment, drone delivery, email notification):

```
Payment done with Card
Card Number: 5436 8687 3432 0097
Earphone of ID: 001, on drone 46
5km away
                mykelsamuel512@gmail.com
Hi Micheal, your order of Earphone 30000.00 has been shipped for delivery
```

---

## Swapping Implementations

Because everything is interface-driven, you can mix and match any combination of payment, delivery, and notification without changing any other logic:

```go
// Rider delivery + wallet payment + SMS notification
call := CheckoutSystem{
    payment:  PayWithWallet{TxnID: "TXN-98765"},
    delivery: RiderDelivery{BikersId: 12, Location: "Lagos"},
    notify:   SmsNotification{PhoneNumber: "+234 800 000 0000"},
}
call.Checkout(order)
```

---

## Key Go Concepts Practiced

- **Interfaces** — decoupling behaviour from implementation, enabling polymorphism
- **Structs** — grouping related data fields into coherent types
- **Methods** — attaching behaviour to types using receiver functions
- **Composition** — building `CheckoutSystem` by composing multiple interfaces rather than inheriting
- **Single-package structuring** — organising types and logic clearly within one `main` package
- **Type-safe duck typing** — no type explicitly declares that it implements an interface; Go's compiler infers satisfaction automatically from method signatures. If `PayWithCard` has a `Pay(float64) string` method, it is a `PaymentProcessor` — no declaration needed. If the signature is wrong or missing, the program fails to compile. This gives you the flexibility of duck typing with the safety of static type checking, and this project demonstrates it across all three interfaces and their nine total implementations.


---

## Project Structure

```
go-mock-delivery-app/
└── main.go       # All types, interfaces, implementations, and entry point
```

---

## Possible Extensions

- [ ] Split into multiple packages (`payment`, `delivery`, `notification`)
- [ ] Add error handling to interface methods
- [ ] Persist orders using a database or JSON file
- [ ] Build a REST API layer on top of the checkout system
- [ ] Write unit tests for each interface implementation

---

## Author

Built as a personal Go learning project — feel free to fork and extend it!