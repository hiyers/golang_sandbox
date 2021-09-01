package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	 bar()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

func foo(){
	for i:=0;i<10;i++{
		fmt.Println("foo",i)
	}
	wg.Done()
}

func bar(){
	for i:=0;i<10;i++{
		fmt.Println("bar",i)
	}
	//wg.Done()
}