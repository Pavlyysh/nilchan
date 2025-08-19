package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func NewPayPal() PayPal {
	return PayPal{}
}

func (c PayPal) Pay(usd int) int {
	fmt.Println("Оплата через PayPal")
	fmt.Println("Размер оплаты:", usd, "USD")
	return rand.Int()
}

func (c PayPal) Cancel(id int) {
	fmt.Println("Операция оплаты PayPal:", id, "отменена")
}
