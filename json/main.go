package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"full_name"`
	Address     string `json:"address"`
	Time        time.Time
}

func (p Payment) Println() {
	fmt.Println("Description:", p.Description)
	fmt.Println("USD:", p.USD)
	fmt.Println("FullName:", p.FullName)
	fmt.Println("Address:", p.Address)
}

type HttpResponse struct {
	Money          int
	PaymentHistory []Payment
}

var mtx = sync.Mutex{}
var money = 1000
var paymentHistory = make([]Payment, 0)

func main() {
	http.HandleFunc("/pay", payHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("server error:", err)
		return
	}
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payment.Time = time.Now()

	payment.Println()

	mtx.Lock()
	if money >= payment.USD {
		money -= payment.USD
		msg := "successful payment"
		fmt.Println(msg)
		w.Write([]byte(msg))
	} else {
		msg := "not enough funds"
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}
	paymentHistory = append(paymentHistory, payment)

	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("money:", money)
	fmt.Println("Payment History:", paymentHistory)
	mtx.Unlock()
}
