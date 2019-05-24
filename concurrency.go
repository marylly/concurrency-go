package main

import(
	"fmt"
	"sync"
)

var semaphore sync.WaitGroup

func a(counter *int) {
	*counter = *counter + 1
	fmt.Printf("statement a%v: A arrived\n", *counter)
	semaphore.Done()
}
func b(counter *int) {
	*counter = *counter + 1
	fmt.Printf("statement b%v: B arrived\n", *counter)
	semaphore.Done()
}

func main() {
	var counterA int = 0
	var counterB int = 0
	
	semaphore.Add(2)
	go a(&counterA)
	go b(&counterB)
	
	go b(&counterB)
	go a(&counterA)
	semaphore.Wait()
}