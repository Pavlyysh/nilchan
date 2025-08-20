package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	str := "Hello, world!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("error writing response")
	} else {
		fmt.Println("success")
	}
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "pay"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("error while pay operation")
	} else {
		fmt.Println("success pay operation")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "cancel"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("error while cancel operation")
	} else {
		fmt.Println("success cancel operation")
	}
}

func main() {
	http.HandleFunc("/default", defaultHandler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	fmt.Println("Starting HTTP-server")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	}

	fmt.Println("Server finished his jov")
}
