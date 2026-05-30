package main

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
	fmt.Printf("Deposited %.2f successfully\n", amount)
}

func (base *BaseAccount) Withdraw(amount float64) error {
	if base.Balance < amount {
		return errors.New("Send failed: insufficient funds")
	}
	base.Balance -= amount
	fmt.Println("Withdrawal Successful")
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
func (card *DebitCard) String() string {
	lastFour := card.CardNumber[len(card.CardNumber)-4:]
	return fmt.Sprintf( "DebitCard [****%s] Owner: %s | Balance: $%.2f | Daily limit: $%.2f", lastFour, card.Owner, card.Balance, card.DailyLimit)
}

func (wallet *CryptoWallet) String() string {
	lastFour := wallet.WalletAddress[len(wallet.WalletAddress)-4:]
	return fmt.Sprintf( "WalletAddress [****%s] Owner: %s | Balance: $%.2f | Network Fee: $%.2f", lastFour, wallet.Owner, wallet.Balance, wallet.GasFee)
}

// Payment contract: valid for any type that implements the type(s) in it.
type PaymentInstrument interface {
	Send(amount float64) error
	AvailableBalance() float64
}

// Exposing DebitCard to PaymeentInstrument 
func (card *DebitCard) Send(amount float64) error {
	if amount > card.DailyLimit {
		return errors.New("Send failed: amount exceeds daily limit")
	}
	return card.Withdraw(amount)
}

func (card *DebitCard) AvailableBalance() float64 {
	return card.Balance
}

// Exposing CryptoWallet to PaymeentInstrument 
func (wallet *CryptoWallet) Send(amount float64) error {
	amount += wallet.GasFee
	return wallet.Withdraw(amount) 	
}

func (wallet *CryptoWallet) AvailableBalance() float64 {
	return wallet.Balance
}

// Gives the user account with the highest balance
func Richest[T PaymentInstrument](accounts []T) T{
	highestAcc := accounts[0] // Takes the very first account in the slice as the account with the highest balance
	for _, account := range accounts[1:] { // skipping the very first element in the slice, since we already have it stored in a variable before
		if account.AvailableBalance() > highestAcc.AvailableBalance() {// This helps to compare the initial we set to other account in the slice
			highestAcc = account
		}
	}
	return highestAcc
}


func main() {
	// declaring instance of the types
	// debit card
	card1 := DebitCard{
		BaseAccount: BaseAccount{
			Owner: "Micheal",
			Balance: 129000.00,
		},
		CardNumber: "5125122313246457",
		DailyLimit: 20000.00,
	}
	card2 := DebitCard{
		BaseAccount: BaseAccount{
			Owner: "Sultan",
			Balance: 332000.00,
		},
		CardNumber: "1642124100981276",
		DailyLimit: 50000.00,
	}

	// crypto wallet
	wallet1:= CryptoWallet{
		BaseAccount: BaseAccount{
			Owner: "Micheal",
			Balance: 129000.00,
		},
		WalletAddress: "0x126her78dm",
		GasFee: 129,
	}
	wallet2:= CryptoWallet{
		BaseAccount: BaseAccount{
			Owner: "Sultan",
			Balance: 332000.00,
		},
		WalletAddress: "0c323him12pw",
		GasFee: 332,
	}

	cards:=[]*DebitCard{&card1, &card2}
	wallets := []*CryptoWallet{&wallet1, &wallet2}

	// ====
	fmt.Println("======================== Debit Card Transaction ======================================")
	fmt.Println(card2.Send(20000))
	card1.Deposit(250000)

	
	fmt.Println("Richest debit card: ", Richest(cards))

	// +++++++++
	fmt.Println("")
	fmt.Println("======================== Crypto Wallet Transaction ======================================")
	fmt.Println("")
	
	// -----
	fmt.Println(wallet2.Send(32234))
	wallet1.Deposit(150000)
	
	fmt.Println("Richest Crypto Wallet: ", Richest(wallets))

}

