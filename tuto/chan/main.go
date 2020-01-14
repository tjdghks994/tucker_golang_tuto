package main

import "fmt"

func pop(c chan int) {
	fmt.Println("Pop func")
	v := <-c // c 에 값이 들어올 때까지 대기 중, 값이 없으면 대기
	fmt.Println(v)
}

func main() {
	var c chan int
	c = make(chan int) // channel 초기화
	// c = make(chan int, 1) //chan의 size를 1로 초기화

	// c <- 10 : deadlock이 발생함, channel의 size가 0으로 유지 되어야 하기 때문에 어디선가는 이걸 빼줘야함 그런데 빼는 동작 전이 호출 되기 전에 넣기 때문에 deadlock이 발생함

	go pop(c)
	c <- 10 // 길이가 0이므로 뺄 때까지 기다림

	fmt.Println("end of program")
}
