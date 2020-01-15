package main

import "fmt"

func main() {
	initMap()
	Test() // 테스트를 먼저하고, 실패를 일단 성공하도록 코딩, 성공 강화 (개선 작업)
}

var opMap map[string]func(int, int) int

func initMap() {
	opMap = make(map[string]func(int, int) int)

	opMap["+"] = add
	opMap["-"] = sub
	opMap["*"] = mul
	opMap["/"] = div
	opMap["**"] = pow
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}
func pow(a, b int) int {
	rst := 1
	for i := 0; i < b; i++ {
		rst *= a
	}
	return rst
}

func Calculate(op string, a, b int) int {

	/*if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "*" {
		return a * b
	}
	return 0
	*/
	// else if 가 계속 나오니까 switch로 수정

	/*
		rst := 0
		switch op {
		case "+":
			rst = a + b
		case "-":
			rst = a - b
		case "*":
			rst = a * b
		case "/":
			rst = a / b
		}
		return rst
	*/
	// Map을 사용하여 refactoring

	if v, ok := opMap[op]; ok { // 성공 강화
		return v(a, b)
	}
	return 0
}

func Test() {
	/*
		o := Calculate("+", 3, 2)
		if o != 5 {
			fmt.Println("Test failed! expected: ", 5, "output: ", o)
			return
		}
		fmt.Println("Success!") // return 5 면 성공

		o = Calculate("+", 5, 4)
		if o != 9 {
			fmt.Println("Test failed! expected: ", 9, "output: ", o)
			return
		}
		fmt.Println("Success!") // return a + b 면 성공
	*/

	if !testCalculate("Test1", "+", 3, 2, 5) {
		return
	} // 성공 강화: refactoring

	if !testCalculate("Test2", "+", 5, 4, 9) {
		return
	}

	if !testCalculate("Test3", "-", 5, 4, 1) {
		return
	}

	if !testCalculate("Test4", "-", 3, 6, -3) {
		return
	}

	if !testCalculate("Test5", "*", 3, 6, 18) {
		return
	}

	if !testCalculate("Test6", "*", 3, 0, 0) {
		return
	}

	if !testCalculate("Test7", "*", 3, -3, -9) {
		return
	}

	if !testCalculate("Test8", "/", 9, 3, 3) {
		return
	}

	if !testCalculate("Test9", "**", 2, 3, 8) {
		return
	}

	if !testCalculate("Test10", "**", 2, 0, 1) {
		return
	}

}

func testCalculate(testcase, op string, a, b, expected int) bool {
	o := Calculate(op, a, b)
	if o != expected {
		fmt.Println(testcase, "Test failed! expected: ", expected, "output: ", o)
		return false
	}
	fmt.Println(testcase, "Success!")
	return true
}
