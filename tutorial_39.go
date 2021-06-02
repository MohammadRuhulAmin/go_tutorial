package main

import (
	"fmt"
	"sync"
)

func counting(value *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	*value++
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	var val int = 0
	for x := 1; x <= 1000; x++ {
		wg.Add(1)
		go counting(&val, &wg, &m)
	}
	wg.Wait()
	fmt.Println(val)
	fmt.Println("More About Mutex:", "https://golangdocs.com/mutex-in-golang")
}
