package mysolutionconcurrency

import (
	"runtime"
	"sync"
	"time"
)

func CoinChange(coins []int, amount int) int {
	return coinChange(coins, amount)
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	stack := 1
	return compute(coins, amount, stack)
}

func compute(coins []int, amount int, stack int) int {
	// fmt.Println("stack:", stack)
	// fmt.Println("amount:", amount)
	var count int
	var n int
	res := amount + 1
	wg := new(sync.WaitGroup)
	// var mu sync.Mutex
	for _, c := range coins {
		// fmt.Println(c)
		// fmt.Println(amount - c)
		n = amount - c
		if n < 0 {
			continue
		} else if n == 0 {
			return 1
		}
		maxroutine := stack * len(coins)
		// fmt.Println("NumGoroutine:", runtime.NumGoroutine())
		if runtime.NumGoroutine() < maxroutine {
			time.Sleep(time.Microsecond)
			// maxroutine += 1
			count = compute(coins, amount-c, stack+1)
			if count > 0 && count < res {
				// mu.Lock()
				res = count
				// mu.Unlock()
			}
		}
		wg.Add(1)
		// fmt.Println("Add")
		go func(scop int) {
			defer wg.Done()
			// fmt.Println(amount)
			count = compute(coins, amount-scop, stack+1)
			if count > 0 && count < res {
				// mu.Lock()
				res = count
				// mu.Unlock()
			}
			// fmt.Println("Done")
		}(c)
	}
	wg.Wait()
	// fmt.Println("return")
	if count <= 0 {
		return -1
	}
	// fmt.Println("res=", res)
	return res + 1
}
