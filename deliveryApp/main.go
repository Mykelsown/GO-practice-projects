package main

import "fmt"

// ============= Contracts --------------
type PaymentProcessor interface {
    Pay(amount float64) string
}

type DeliveryPattern interface {
    Delivery(order Order) string
    TrackLocation() string
}

type Notifier interface {
    Notify(order Order) string
}

// ============= Order Details -----------------
type Order struct {
    ID string
    Item string
    Amount float64
    CustomerName string
}

// ============= Payment Structure -------------
type PayWithCard struct { CardNumber string }

type PayWithWallet struct { TxnID string }

type PayWithCash struct {}

// ============= Delivery Structure --------------
type RiderDelivery struct {
    BikersId int
    Location string
}

type DroneDelivery struct {
    DroneId int
    Location string
}

// ========== Notification Structure ------------
type EmailNotification struct { EmailAddy string }
 
type SmsNotification struct { PhoneNumber string }

type PushNotification struct { PhoneIp string }

//============ Payment Implementation --------- 
func (cardPay PayWithCard) Pay(amount float64 ) string {
    return fmt.Sprintf("Payment done with Card\nCard Number: %s", cardPay.CardNumber)
}

func (walletPay PayWithWallet) Pay(amount float64 ) string {
    return fmt.Sprintf("Payment done with Wallet\nTxnID: %s", walletPay.TxnID)
}

func (cashPay PayWithCash) Pay(amount float64 ) string {
    return "Payment done with cash"
}

//=========== Delivery Implementaion ----------
func (rider RiderDelivery) Delivery(order Order) string {
    return fmt.Sprintf("%s of ID: %s, on rider %d", order.Item, order.ID, rider.BikersId)
}

func (rider RiderDelivery) TrackLocation() string {
    return "10km away"
}

func (drone DroneDelivery) Delivery(order Order) string {
    return fmt.Sprintf("%s of ID: %s, on drone %d", order.Item, order.ID, drone.DroneId)
}

func (drone DroneDelivery) TrackLocation() string {
    return "5km away"
}

// ========== Notification Implementaion ----------
func (email EmailNotification) Notify(order Order) string {
    return fmt.Sprintf("\t\t\t\t%s\nHi %s, your order of %s %.2f has been shipped for delivery", email.EmailAddy, order.CustomerName, order.Item, order.Amount)
}

func (sms SmsNotification) Notify(order Order) string {
    return fmt.Sprintf("\t\t\t\t%s\nHi %s, your order of %s %.2f has been shipped for delivery", sms.PhoneNumber, order.CustomerName, order.Item, order.Amount)
}

func (push PushNotification) Notify(order Order) string {
    return fmt.Sprintf("\t\t\t\t%s\nHi %s, your order of %s %.2f has been shipped for delivery", push.PhoneIp, order.CustomerName, order.Item, order.Amount)
}

// ========== checkout -----------
type CheckoutSystem struct {
    payment PaymentProcessor
    delivery DeliveryPattern
    notify Notifier
}

func (cs CheckoutSystem) Checkout(order Order) {
    paymentInfo := cs.payment.Pay(order.Amount)
    fmt.Println(paymentInfo)

    deliveryInfo, locationInfo := cs.delivery.Delivery(order), cs.delivery.TrackLocation()
    fmt.Println(deliveryInfo)
    fmt.Println(locationInfo)

    notificationMsg := cs.notify.Notify(order)
    fmt.Println(notificationMsg)
}

func main() {
    order := Order{
        ID: "001",
        Item: "Earphone",
        Amount: 30000.00,
        CustomerName: "Micheal",
    }

    call := CheckoutSystem{
        payment: PayWithCard{CardNumber: "5436 8687 3432 0097"},
        delivery: DroneDelivery{DroneId: 0056},
        notify: EmailNotification{EmailAddy: "mykelsamuel512@gmail.com"},
    }
    call.Checkout(order)

}