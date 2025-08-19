package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- int, minerID int, power int) {
	defer wg.Done()
	for {

		// Кейс, когда после отмены контекста майнеры заканчивают работу
		fmt.Println("Miner", minerID, "start his job")

		select {
		case <-ctx.Done():
			fmt.Println("Miners'", minerID, "work ended for today")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Miner", minerID, "get coal", power)
		}

		select {
		case <-ctx.Done():
			fmt.Println("Miners'", minerID, "work ended for today")
			return
		case transferPoint <- power:
			fmt.Println("Miner", minerID, "send coal", power)
		}

		/* Кейс, когда после отмены контекста майнеры дорабатывают свою итерация

		select {
		case <-ctx.Done():
			fmt.Println("Miners'", minerID, "work ended for today")
			return
		default:
			fmt.Println("Miner", minerID, "start his job")
			time.Sleep(1 * time.Second)
			fmt.Println("Miner", minerID, "get coal", power)
			transferPoint <- power
			fmt.Println("Miner", minerID, "send coal", power)
		}

		*/
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
