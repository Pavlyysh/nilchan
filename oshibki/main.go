package main

import (
	"errors"
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Balance int
}

func Pay(user *User, usd int) error {
	if user.Balance-usd < 0 {
		return errors.New("not enough funds")
	}

	user.Balance -= usd
	return nil
}

func main() {
	user := User{
		Name:    "Pasha",
		Balance: 150,
	}

	err := Pay(&user, 200)
	if err != nil {
		fmt.Println("payment error:", err.Error())
	} else {
		fmt.Println("success payment")
	}
	pp.Println(user.Balance)
}
