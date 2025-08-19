package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (c Bank) Pay(usd int) int {
	fmt.Println("Оплата картой")
	fmt.Println("Размер оплаты:", usd, "долларов")
	return rand.Int()
}

func (c Bank) Cancel(id int) {
	fmt.Println("Операция оплаты картой:", id, "отменена")
}
