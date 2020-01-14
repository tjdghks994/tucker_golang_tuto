package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

type Plane struct {
	val string
}

func MakeTire(carChan, outCarChan chan Car) {
	for {
		car := <-carChan
		car.val += "Tire_C, "
		outCarChan <- car

	}
}

func MakeEngine(carChan, outCarChan chan Car) {
	for {
		car := <-carChan
		car.val += "Engine_C, "
		outCarChan <- car

	}
}

// StartCarWork chan1에 값을 넣어줌
func StartCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car" + strconv.Itoa(i) + ": "}
		i++
	}
}

func main() {
	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	// chan1 <- Car{val: "deadLock"} // chan1의 size가 0인데 값을 할당하려고 해서 err 발생

	go StartCarWork(carChan1)
	go MakeTire(carChan1, carChan2)
	go MakeEngine(carChan2, carChan3)

	for {
		result := <-carChan3
		fmt.Println(result.val)
	}

}
