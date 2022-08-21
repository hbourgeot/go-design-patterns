package main

import "fmt"

/*
	Aquí la estructura BankPayment NO cumple con la siguiente interfaz, por lo que
  se debe implementar un adaptador
*/
type Payment interface {
	Pay()
}

type CashPayment struct{}

func (p CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

func (b BankPayment) Pay(bankAccount int) {
	fmt.Printf("Payment using bankaccount %d\n", bankAccount)
}

// Adaptador para que BankPayment cumpla la interfaz

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	bank := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}

	ProcessPayment(cash)
	ProcessPayment(bank)
}
