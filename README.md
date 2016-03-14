# Semaphore
Quick and dirty semaphore for safe concurrent access.

## Usage example
```
package main

import (
	"fmt"
	"time"

	"github.com/skybon/semaphore"
)

func main() {
	concurrentActions := 1
	printNum := 10

	exitSem := semaphore.MakeSemaphore(concurrentActions)
	mySem := semaphore.MakeSemaphore(concurrentActions)

	for i := 0; i < printNum; i++ {
		go mySem.Exec(func() {
			fmt.Println("Hello world!")
			exitSem.Release()
			time.Sleep(1 * time.Second)
		})
	}

	for i := 0; i < exitSem.LimitCount(); i++ {
		exitSem.Acquire()
	}
}
```
