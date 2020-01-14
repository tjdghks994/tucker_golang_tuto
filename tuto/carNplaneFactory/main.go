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

func MakeTire(carChan, outCarChan chan Car, planeChan, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Tire_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Tire_P, "
			outPlaneChan <- plane
		}
	}
}

func MakeEngine(carChan, outCarChan chan Car, planeChan, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Engine_P, "
			outPlaneChan <- plane
		}
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

// StartPalneWork chan1에 값을 넣어줌
func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane" + strconv.Itoa(i) + ": "}
		i++
	}
}

func main() {
	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)
	// chan1 <- Car{val: "deadLock"} // chan1의 size가 0인데 값을 할당하려고 해서 err 발생

	go StartCarWork(carChan1)
	go StartPlaneWork(planeChan1)

	go MakeTire(carChan1, carChan2, planeChan1, planeChan2)
	go MakeEngine(carChan2, carChan3, planeChan2, planeChan3)

	for {
		select {
		case result := <-carChan3:
			fmt.Println(result.val)
		case result := <-planeChan3:
			fmt.Println(result.val)
		}
	}

}
