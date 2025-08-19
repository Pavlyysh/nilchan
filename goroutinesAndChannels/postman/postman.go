package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- string, postmanID int, mail string) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Postmans'", postmanID, "work day ended")
			return
		default:
			fmt.Println("Postman", postmanID, "start his job")
			time.Sleep(1 * time.Second)
			fmt.Println("Postman", postmanID, "get a letter", mail)

			transferPoint <- mail
			fmt.Println("Postman", postmanID, "send letter", mail)
		}
	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmanToMail(postmanID int) string {
	ptm := map[int]string{
		1: "family",
		2: "friend",
		3: "auto",
	}

	mail, ok := ptm[postmanID]
	if !ok {
		return "lottery"
	}

	return mail
}
