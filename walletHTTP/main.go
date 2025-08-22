package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var mtx = sync.Mutex{}
var money = 1000
var bank = 0

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}
	mtx.Lock()
	if money >= paymentAmount {
		money -= paymentAmount
		msg := fmt.Sprintf("успешная оплата, текущий баланс: %d\n", money)
		fmt.Println(msg)
		w.Write([]byte(msg))

	} else {
		msg := "оплата не произведена, недостаточно средств"
		fmt.Println(msg)
		w.Write([]byte(msg))
	}
	mtx.Unlock()

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "fail to read HTTP body:" + err.Error()
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}

	mtx.Lock()
	if money >= saveAmount {
		// снять с кошелька
		money -= saveAmount

		// положить в копилку
		bank += saveAmount

		msg := "перенос в копилку успешно произведен\n"
		fmt.Println(msg)
		w.Write([]byte(msg))
		msg = fmt.Sprintf("Денег в кошельке: %d\n", money)
		fmt.Println(msg)
		w.Write([]byte(msg))
		msg = fmt.Sprintf("Денег в копилке: %d\n", bank)
		fmt.Println(msg)
		w.Write([]byte(msg))
	} else {
		msg := "перенос в копилку не произведен, недостаточно средств"
		fmt.Println(msg)
		w.Write([]byte(msg))
	}
	mtx.Unlock()

}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error:", err.Error())
	}
}
