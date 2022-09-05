package first

import (
	"math/rand"
	"time"
)

func Source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// Sleep 1s/2s/3s
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}
