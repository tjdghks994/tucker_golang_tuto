package main

import (
	hash "tucker_golang_tuto/tuto/datastruct"
	"fmt"
)

func main() {
	fmt.Println("tjdghks994 = ", hash.RollingHash("tjdghks994"))

	m := hash.CreateMap()
	m.Add("AAA", "01077777777")
	m.Add("BBB", "01088888888")
	m.Add("CCC", "01011111111")
	m.Add("DSGSADGDS", "01055555555")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("BBB = ", m.Get("BBB"))
	fmt.Println("DDD = ", m.Get("DDD"))
}
