package paymentprocessingsystem

import (
	"errors"
	"fmt"
)

// Base Account
// Structure
type BaseAccount struct {
	Owner   string
	Balance float64
}

// Methods
func (base *BaseAccount) Deposit(amount float64) {
	base.Balance += amount
}

func (base *BaseAccount) Withdraw(amount float64) error {
	if base.Balance < amount {
		return errors.New("insufficient funds")
	}
	return nil
}


// Withdrawal Methods structure
type DebitCard struct {
	BaseAccount
	CardNumber string
	DailyLimit float64
}

type CryptoWallet struct {
	BaseAccount
	WalletAddress string
	GasFee float64
}

// Custumize the type Debitcard & CryptoWallet output
func (card DebitCard) String() string {
	lastFour := card.CardNumber[len(card.CardNumber)-4:]
	return fmt.Sprintf( "DebitCard [**** %s] Owner: %s | Balance: $%.2f | Daily limit: $%.2f", lastFour, card.Owner, card.Balance, card.DailyLimit)
}

func (wallet CryptoWallet) String() string {
	lastFour := wallet.WalletAddress[len(wallet.WalletAddress)-4:]
	return fmt.Sprintf( "WalletAddress [**** %s] Owner: %s | Balance: $%.2f | Network Fee: $%.2f", lastFour, wallet.Owner, wallet.Balance, wallet.GasFee)
}

// Payment contract
type PaymentInstruction interface {
	Send(amount float64) error
	AvailableBalance() float64
}

// Exposing DebitCard to PaymeentInstruction 
func (card DebitCard) Send(amount float64) error {
	if amount > card.DailyLimit {
		return errors.New("daily limit exceeded")
	}
	return card.Withdraw(amount)
}

func (card DebitCard) AvailableBalance() float64 {
	return card.Balance
}

// Exposing CryptoWallet to PaymeentInstruction 
func (wallet CryptoWallet) Send(amount float64) error {
	amount += amount*0.03
	return wallet.Withdraw(amount) 	
}

func (wallet CryptoWallet) AvailableBalance() float64 {
	return wallet.Balance
}


