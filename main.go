package main

import "code/go/testing/arithmetic/concurrency"

func main() {
	//rand.Seed(time.Now().UnixNano())
	//
	//a, b := concurrency.LongTimeRequest(), concurrency.LongTimeRequest()
	//fmt.Println(concurrency.SumSquares(<-a, <-b))

	concurrency.GoRoutine()
}
