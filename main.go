// We often want to execute Go code at some point in the
// future, or repeatedly at some interval. Go's built-in
// _timer_ and _ticker_ features make both of these tasks
// easy. We'll look first at timers and then
// at [tickers](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	go tickerFunc(ticker)

	time.Sleep(11 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}

func tickerFunc(ticker *time.Ticker) {
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}
