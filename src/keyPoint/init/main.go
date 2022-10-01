package main

import (
	"fmt"
	"strings"
	"sync"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	var str string = "Hi, I'm Marc"
	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))
	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))
	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))
	array := map[string]bool{
		"aaa": true,
	}
	sa := "aaa"
	sb := "bbb"
	fmt.Println(array[sa])
	fmt.Println(array[sb])
	fmt.Println(sa == "baa")
	fmt.Println(len(sa))
	fmt.Println(sa[0:2])
	fmt.Println(sa + sb)
	array2 := map[string]string{
		"zzz": "xxx",
	}
	fmt.Println(len(array2["ccc"]))
	num1 := 10
	var wg sync.WaitGroup
	intChan := make(chan int, 100)
	wg.Add(1)
	for i := 0; i < num1; i++ {
		fmt.Println(i)
		go func(i0 int) {
			for j := 0; j < 10; j++ {
				wg.Add(1)
				fmt.Println("send", i0+j)
				intChan <- i0 + j
				wg.Done()
			}
		}(i * 10)
	}
	wg.Done()
	// wg.Wait()
	close(intChan)

	for i := range intChan {
		fmt.Println("recive", i)
	}
	wg.Wait()
}
