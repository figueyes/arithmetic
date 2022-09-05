package main

import (
	"code/go/testing/arithmetic/concurrency/id"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}

	c := make(chan string)
	wg.Add(2)
	//go id.GenerateID(wg, c)
	go id.BulkGenerateID(wg, c)
	go id.Log(wg, c)
	wg.Wait()

	//rand.Seed(time.Now().UnixNano())

	/*
		Return receive-only channels as results
		In the following example, the values of two arguments of the sumSquares function
		call are requested concurrently. Each of the two channel receive operations will
		block until a send operation performs on the corresponding channel. It takes about
		three seconds instead of six seconds to return the final result.
	*/
	//a, b := future.LongTimeRequestReceiveOnly(), future.LongTimeRequestReceiveOnly()
	//result := future.SumSquares(<-a, <-b)
	//fmt.Println(result)

	/*
		Pass send-only channels as arguments
		Same as the last example, in the following example, the values of two arguments
		of the sumSquares function call are requested concurrently. Different to the last
		example, the longTimeRequest function takes a send-only channel as parameter instead
		of returning a receive-only channel result.
	*/
	//ra, rb := make(chan int32), make(chan int32)
	//go future.LongTimeRequestSendOnly(ra)
	//go future.LongTimeRequestSendOnly(rb)
	//result = future.SumSquares(<-ra, <-rb)
	//fmt.Println(result)

	/*
		The first response wins

		Sometimes, a piece of data can be received from several sources to avoid high latencies.
		For a lot of factors, the response durations of these sources may vary much. Even for a
		specified source, its response durations are also not constant. To make the response
		duration as short as possible, we can send a request to every source in a separated
		goroutine. Only the first response will be used, other slower ones will be discarded.

		Note, if there are N sources, the capacity of the communication channel must be at
		least N-1, to avoid the goroutines corresponding the discarded responses being blocked
		for ever.
	*/
	//startTime := time.Now()
	//// c must be a buffered channel
	//c := make(chan int32, 5)
	//for i := 0; i < cap(c); i++ {
	//	go first.Source(c)
	//}
	//
	//// Only the first response will be used.
	//rnd := <-c
	//fmt.Println("start ", time.Since(startTime))
	//fmt.Println("finish ", rnd)

	//values := make([]byte, 32*1024*1024)
	//if _, err := rand.Read(values); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//done := make(chan struct{}) // can be buffered or not
	//
	//// The sorting goroutine
	//go func() {
	//	sort.Slice(values, func(i, j int) bool {
	//		return values[i] < values[j]
	//	})
	//	// Notify sorting is done.
	//	done <- struct{}{}
	//}()
	//
	//// do some other things ...
	//
	//<-done // waiting here for notification
	//fmt.Println(values[0], values[len(values)-1])

	/*
		Use Channels as Mutex Locks
	*/

	// The capacity must be one.
	//mutex := make(chan struct{}, 1)
	//
	//counter := 0
	//increase := func() {
	//	mutex <- struct{}{} // lock
	//	counter++
	//	<-mutex // unlock
	//}
	//
	//increase1000 := func(done chan<- struct{}) {
	//	for i := 0; i < 1000; i++ {
	//		increase()
	//	}
	//	done <- struct{}{}
	//}
	//
	//done := make(chan struct{})
	//go increase1000(done)
	//go increase1000(done)
	//<-done
	//<-done
	//fmt.Println(counter) // 2000

}
