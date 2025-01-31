package main

import "fmt"

type PaymentMethod interface {
	ProcessPayment(amount float64) string
}

type CreditCard struct{}

func (cre CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf("신용카드로 %.2f원 결제 완료!", amount)
}

type PayPal struct{}

func (pay PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("페이팔로 %.2f원 결제 완료!", amount)
}

type Cash struct{}

func (cas Cash) ProcessPayment(amount float64) string {
	return fmt.Sprintf("현금으로 %.2f원 결제 완료!", amount)
}

type BankTransfer struct{}

func (ban BankTransfer) ProcessPayment(amount float64) string {
	return fmt.Sprintf("계좌이체로 %.2f원 결제 완료!", amount)
}

func HandlePayment(pm PaymentMethod, amount float64) {
	result := pm.ProcessPayment(amount)
	fmt.Println(result)
}

func main() {
	creditCard := CreditCard{}
	paypal := PayPal{}
	cash := Cash{}
	bank := BankTransfer{}

	HandlePayment(creditCard, 50000) // "카드로 50000원 결제 완료!"
	HandlePayment(paypal, 30000)     // "페이팔로 30000원 결제 완료!"
	HandlePayment(cash, 45000)       // "현금으로 45000원 결제 완료!"
	HandlePayment(bank, 100000)      // "계좌이체로 100000원 결제 완료!"
}
