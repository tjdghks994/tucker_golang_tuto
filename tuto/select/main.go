package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func main() {

	c := make(chan int)

	go push(c)

	timer := time.After(10 * time.Second) // 타이머 channel
	tickTimer := time.Tick(2 * time.Second)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
			// 어떤 값이 들어왔을 때만 코드가 실행 될 수 있도록 함
		case <-timer:
			fmt.Println("Time Out!")
			return
		case <-tickTimer:
			fmt.Println("tictok")
		default:
			fmt.Println("Idle")
			time.Sleep(1 * time.Second)
			// 대기 모드라고 생각하면 됨, 1초마다 입력이 있는지 파악함!
		}
	}
}
