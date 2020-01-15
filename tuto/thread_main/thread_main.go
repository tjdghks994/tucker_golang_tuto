package main

import (
	"fmt"
	"time"
)

func main() {

	go func1() // 이렇게 하면 go thread로 실행되어 아래 함수와 go 가 붙은 함수가 동시에 실행 될 수 있음
	//func1() // 이렇게 하면 func1()이 모두 실행 된 후에 아래 코드가 실행됨
	// go의 thread 는 os thread 하나에서 여러개의 go thread를 실행 할 수 있도록 하여 context switching이
	// 발생하지 않도록 해줌
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("main :", i)
	}
	fmt.Scanln()
}

func func1() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("func1 :", i)
	}
}
