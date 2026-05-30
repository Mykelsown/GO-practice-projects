# Payment Processing System — Go Practice Project

A fintech-themed CLI program that simulates a basic payment processing system. It models two payment instruments — debit cards and crypto wallets — and demonstrates core Go concepts through real-world domain logic.

---

## Concepts covered

| Concept | Where it appears |
|---|---|
| Custom types & structs | `BaseAccount`, `DebitCard`, `CryptoWallet` |
| Composition & embedding | `DebitCard` and `CryptoWallet` embed `BaseAccount` |
| Methods & receivers | `Deposit`, `Withdraw`, `Send`, `AvailableBalance` |
| Interfaces | `PaymentInstrument` interface constrains both types |
| Stringer interface | `String()` on `*DebitCard` and `*CryptoWallet` |
| Generics | `Richest[T PaymentInstrument]` works across any account slice |

---

## Project structure

```
payment-processor/
└── main.go
```

---

## How to run

**Prerequisites:** Go 1.21 or later.

```bash
go run main.go
```

**Expected output:**

```
======================== Debit Card Transaction ======================================
Withdrawal Successful
<nil>
Deposited 250000.00 successfully
Richest debit card:  DebitCard [****6457] Owner: Micheal | Balance: $379000.00 | Daily limit: $20000.00

======================== Crypto Wallet Transaction ======================================

Withdrawal Successful
<nil>
Deposited 150000.00 successfully
Richest Crypto Wallet:  WalletAddress [****12pw] Owner: Sultan | Balance: $332000.00 | Network Fee: $332.00
```

---

## Design decisions and why

### 1. Embedding instead of a named field

`DebitCard` and `CryptoWallet` both embed `BaseAccount` directly:

```go
type DebitCard struct {
    BaseAccount       // embedded — not a named field
    CardNumber string
    DailyLimit float64
}
```

This promotes `BaseAccount`'s fields and methods (`Deposit`, `Withdraw`) onto `DebitCard` itself. You can call `card.Deposit(100)` directly instead of `card.BaseAccount.Deposit(100)`. This avoids duplicating shared logic across two types and reflects how real payment systems share a common account layer.

If a named field had been used instead (`account BaseAccount`), method promotion would not happen and every call would need to go through the field name explicitly.

### 2. Pointer receivers everywhere

Every method uses a pointer receiver (`*BaseAccount`, `*DebitCard`, `*CryptoWallet`):

```go
func (base *BaseAccount) Withdraw(amount float64) error { ... }
func (card *DebitCard) String() string { ... }
```

There are two reasons for this:

- **Mutation**: `Deposit` and `Withdraw` change the `Balance` field. A value receiver would operate on a copy, so the original struct would never be updated.
- **Consistency**: Once any method on a type uses a pointer receiver, all methods should follow suit. Mixing pointer and value receivers on the same type creates an inconsistent method set and can cause subtle bugs when the type is used through an interface.

### 3. `Withdraw` returns `nil` on success

```go
func (base *BaseAccount) Withdraw(amount float64) error {
    if base.Balance < amount {
        return errors.New("Send failed: insufficient funds")
    }
    base.Balance -= amount
    fmt.Println("Withdrawal Successful")
    return nil                         // correct: no error means nil
}
```

In Go, the `error` return type signals that something went wrong — a non-nil error is a failure. Returning a success message as an error would cause every caller to treat a successful transaction as a failure. Confirmation output is done via `fmt.Println` separately, keeping error signalling clean.

### 4. `GasFee` instead of `NetworkFee`

The `CryptoWallet` field is named `GasFee`:

```go
type CryptoWallet struct {
    BaseAccount
    WalletAddress string
    GasFee float64     // domain-accurate name
}
```

In crypto networks (Ethereum, etc.), the fee paid to validators for processing a transaction is specifically called a *gas fee*. Using domain-accurate naming makes the code more readable to anyone familiar with the fintech space.

### 5. Generic `Richest` function

```go
func Richest[T PaymentInstrument](accounts []T) T {
    highestAcc := accounts[0]
    for _, account := range accounts[1:] {
        if account.AvailableBalance() > highestAcc.AvailableBalance() {
            highestAcc = account
        }
    }
    return highestAcc
}
```

The `PaymentInstrument` interface doubles as the generic type constraint. This means `Richest` works for `[]*DebitCard` or `[]*CryptoWallet` without duplicating the comparison logic. The loop starts at `accounts[1:]` to skip comparing the baseline element against itself — a small but deliberate efficiency.

### 6. Storing pointers in slices

```go
cards := []*DebitCard{&card1, &card2}
```

The slice stores pointers to the original `card1` and `card2` variables rather than copies. This means:

- Any mutation (e.g. after a `Send` or `Deposit`) is reflected in the original variable.
- The struct is not copied into the slice — only an 8-byte pointer is stored per element.

---

## Possible extensions

- Add a `Transaction` struct that logs each `Send` with a timestamp and status.
- Implement a `MonthlyLimit` on `DebitCard` that resets every 30 days.
- Add a `Convert(toCurrency string) float64` method on `CryptoWallet` using a live exchange rate.
- Write a `Poorest[T PaymentInstrument]` function and compare its implementation to `Richest`.