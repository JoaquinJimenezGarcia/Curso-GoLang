package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}
type BankPayment struct{}
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Payment using bank account %d\n", bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	bank := &BankPaymentAdapter{
		bankAccount: 12345,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bank)
}
