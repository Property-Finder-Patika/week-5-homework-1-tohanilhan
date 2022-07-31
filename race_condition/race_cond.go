// Create a sample program that simulates a race condition.
// So it should run fine when run sequentially but cause trouble when run in two or more goroutines.
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mtx sync.Mutex
	wg  sync.WaitGroup
)

type Calculator struct {
	Result int
}

func (c *Calculator) Sum(num1, num2 int) {
	c.Result = num1 + num2
}

func (c *Calculator) Subtract(num1, num2 int) {
	c.Result = num1 - num2

}

func (c *Calculator) SumWithMutex(num1, num2 int, waitG *sync.WaitGroup) {
	//add lock
	mtx.Lock()

	//add unlock
	defer mtx.Unlock()

	c.Result = num1 * num2

	// Wait for the goroutines to finish
	waitG.Done()
}

func (c *Calculator) SubtractWithMutex(num1, num2 int, waitG *sync.WaitGroup) {

	//add lock
	mtx.Lock()

	c.Result = num1 - num2

	//add unlock
	mtx.Unlock()

	// Wait for the goroutines to finish
	waitG.Done()
}
func main() {

	// Expected output when you run it with:
	/*
		/* -go run -race race_cond.go --> To be able to debug the race condition if there are any.
	*/

	// - Found 1 data race(s)
	// Which means that if two or more functions tries to change the same data asyncroniously,
	// it causes race condition.

	RaceCondition()

	// To fix that problem, you have to use mutex lock, or not run these functions as go routines.
	// If you want to see the results, you can uncomment these functions.

	// Solved the problem by using mutex lock.
	// NoRaceConditionWithMutex()

	// Solved the problem without using go routines.
	// NoRaceConditionWithoutGoRoutines()
}

func RaceCondition() {

	c := Calculator{}

	// Create two go routines
	go c.Sum(1, 2)
	go c.Subtract(3, 2)

	// Wait for the goroutines to finish
	time.Sleep(1 * time.Second)

	fmt.Println(c.Result)
}

func NoRaceConditionWithoutGoRoutines() {

	c := Calculator{}

	c.Sum(1, 2)
	c.Subtract(3, 2)

	fmt.Println(c.Result)
}

func NoRaceConditionWithMutex() {
	c := Calculator{}

	// Run in two goroutines
	wg.Add(2)
	go c.SumWithMutex(1, 2, &wg)
	go c.SubtractWithMutex(3, 2, &wg)
	wg.Wait()

	fmt.Println(c.Result)
}
