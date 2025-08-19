package main

import (
	"pavlyysh/nilchan/payments"
	"pavlyysh/nilchan/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewCrypto() // поменяв эту строчку кода на NewBank или NewPayPal код будет работать, но уже с другим платежным методом

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("FLStudio", 10)
	gameID := paymentModule.Pay("Game", 5)
	paymentModule.Pay("Food", 3)

	paymentModule.Cancel(gameID)

	allInfo := paymentModule.AllInfo()

	pp.Println(allInfo)
}
