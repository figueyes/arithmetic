package id

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

func GenerateID(wg *sync.WaitGroup, c chan string) {
	uid := uuid.New()
	c <- uid.String()

	wg.Done()
}

func BulkGenerateID(wg *sync.WaitGroup, c chan<- string) {
	for i := 0; i < 1000; i++ {
		uid := uuid.New()
		c <- fmt.Sprintf("%d %s", i+1, uid.String())
	}
	close(c)
	wg.Done()
}

func Log(wg *sync.WaitGroup, c <-chan string) {
	for {
		select {
		case uid, ok := <-c:
			if ok {
				fmt.Printf("id: %s\n", uid)
			}
		}
	}
}
