package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("숫자를 입력 하세요~")          // print
	reader := bufio.NewReader(os.Stdin) // scan
	line, _ := reader.ReadString('\n')  // 개행문자까지 읽기
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	if (n1 < 36656) && (n2 < 36656) {
		fmt.Printf("입력하신 숫자는 %d, %d 입니다 \n", n1, n2)
	} else {
		fmt.Println("입력 범위를 초과했습니다.")
		os.Exit(0)
	}
	fmt.Println("연산자를 입력하세요!")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	/*
		if line == "+" {
			fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
		} else if line == "-" {
			fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
		} else if line == "*" {
			fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
		} else if line == "/" {
			fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
		} else {
			fmt.Println("연산자가 아닙니다~")
		}
	*/

	switch line {
	case "+":
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	default:
		fmt.Println("연산자가 아닙니다~")
	}
}
