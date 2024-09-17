/*
Explanation
Race Condition: In this example, the increment and decrement goroutines both modify the global counter variable concurrently. The counter variable 
is shared between these two goroutines and accessed without any synchronization mechanisms like mutexes. This leads to a race condition.
How It Occurs:
Concurrent Access: Both increment and decrement goroutines access and modify the counter variable at the same time.
Read-Modify-Write Sequence: Each goroutine reads the current value of counter, modifies it, and then writes the new value back. If these operations 
are not atomic, multiple goroutines can read the same initial value, modify it, and write it back, which results in lost updates.
Example: Suppose counter is 5 and both goroutines simultaneously read the value. Both see 5, increment (or decrement) it, and write back 6 (or 4). 
As a result, the counter might end up being 4 or 6 instead of the expected 0.
Demonstrating the Race Condition
To demonstrate the race condition, run the program multiple times. You will often see that the final value of counter is not zero as expected. 
This variability is due to the race condition where the operations of incrementing and decrementing are not synchronized.
Detecting Race Conditions
To detect and diagnose race conditions in Go, you can use the race detector. Run the program with the -race flag:
go run -race main.go
This will help you identify race conditions and understand where concurrent access to shared variables is causing issues.
*/

package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		counter++
	}
}

func decrement(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		counter--
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&wg)
	go decrement(&wg)

	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter)
}
