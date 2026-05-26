# Animal Shelter — Go Practice Project

A small CLI program that models an animal shelter system in Go. Built as a practice exercise covering core Go concepts: custom types, methods, interfaces, the Stringer interface, and generics.

---

## Concepts covered

| Concept | Where it appears |
|---|---|
| Custom types & structs | `Dog` and `Cat` structs with typed fields |
| Methods & receivers | `Sound()` and `Fee()` on both structs |
| Interfaces | `Animal` interface constraining both types |
| Stringer interface | `String()` on `Dog` and `Cat` for formatted output |
| Generics | `Cheapest[T Animal]` works across any slice of animals |

---

## Project structure

```
animal-shelter/
└── main.go
```

---

## Running the project

**Prerequisites:** Go 1.21 or later (generics require Go 1.18+).

```bash
go run main.go
```

**Expected output:**

```
Cheapest dog:  Dog: Bounce (Pit Bull) - $197.54
Cheapest cat:  Cat: Whiskers [indoor] - $64.34
```

---

## Key design decisions

### Bool for `Cat.Indoor`
The `Indoor` field is stored as a `bool` (1 byte) rather than a string. This keeps the data model minimal and enforces only two valid states at the type level. The `String()` method handles translating `true`/`false` into `"indoor"`/`"outdoor"` at the display layer.

### Stringer for display
Both `Dog` and `Cat` implement `fmt.Stringer` (`String() string`). This means any call to `fmt.Println` on these types automatically uses the formatted output without any extra arguments.

### Generic `Cheapest` function
```go
func Cheapest[T Animal](animals []T) T
```
The `Animal` interface doubles as the generic type constraint. The function works for a homogeneous slice of any type that satisfies `Animal` — no interface boxing or type assertions needed.

---

## What the `Animal` interface looks like

```go
type Animal interface {
    Sound() string
    Fee()   float64
}
```

Both `Dog` and `Cat` satisfy this interface, which is what allows `Cheapest` to call `.Fee()` generically.

---

## Possible extensions

- Add more animal types (e.g. `Rabbit`, `Bird`) without changing `Cheapest`
- Implement a `Describe()` method that prints the animal's sound alongside its info
- Write a `CheapestMixed(animals []Animal) Animal` variant and compare it to the generic version