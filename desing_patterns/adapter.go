package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying  with back account %d\n", bankAccount)
}

// Create Struct Adapter  for  apply the adapter pattern above Payment Interface
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	// ERROR: this error is because the BankPayment isn't implementing the inteface Payment correctly
	// bank := &BankPayment{}
	// ProcessPayment(bank)

	bank := &BankPaymentAdapter{
		BankPayment: &BankPayment{},
		bankAccount: 99966676,
	}
	ProcessPayment(bank)
}
