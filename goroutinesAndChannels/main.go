package main

import (
	"context"
	"fmt"
	"pavlyysh/concurrency/miner"
	"pavlyysh/concurrency/postman"
	"sync"
	"time"
)

func main() {
	var coal int
	var mails []string
	var mtx sync.Mutex
	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("----->>>РАБОЧИЙ ДЕНЬ ШАХТЕРОВ ОКОНЧЕН")
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("----->>>РАБОЧИЙ ДЕНЬ ПОЧТАЛЬОНОВ ОКОНЧЕН")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 2)
	mailTransferPoint := postman.PostmanPool(postmanContext, 2)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			mtx.Lock()
			coal += v // можно вместо мьютекса использовать атомик
			mtx.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	mtx.Lock()
	fmt.Println("coal:", coal)
	mtx.Unlock()

	mtx.Lock()
	fmt.Println("mails count", len(mails))
	mtx.Unlock()
}
