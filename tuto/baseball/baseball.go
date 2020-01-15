package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 랜덤 시드 설정

	game := MakeNum() // 숫자 3개 선언
	cnt := 0

	for {
		cnt++

		users := InputNum()

		result := CompareNum(game, users)

		if result { // result == true 일 때 승리
			fmt.Println(cnt, "번 만에 맞췄습니다!")
			break
		}
	}
}

// MakeNum 0~9 사이의 겹치지 않는 숫자를 생성
func MakeNum() [3]int {
	var game [3]int

	for i := 0; i < 3; i++ {

		for {
			rnd := rand.Intn(10)
			duplicated := false

			for j := 0; j < i; j++ {
				if game[j] == rnd { //겹칠 때 다시 뽑을 수 있도록
					duplicated = true
					break
				}
			}

			if !duplicated { // 겹치지 않을 때 값 대입
				game[i] = rnd
				break
			}
		}
	}

	fmt.Println(game)
	return game

}

// InputNum 사용자로부입력을 받, 단, 숫자는 겹치지 않게 입력한다.
func InputNum() [3]int {
	var input [3]int
	var res [3]int

	for {
		success := false
		fmt.Println("0 ~ 9까지의 숫자를 중복 되지 않게 3개 입력하세요!")
		_, err := fmt.Scanln(&input[0], &input[1], &input[2])
		if err != nil {
			fmt.Println("잘 못 입력하셨습니다")
		}

		for i := 0; i < 3; i++ { //중복 검사
			duplicated := false
			for j := i + 1; j < 3; j++ {
				if input[i] == input[j] {
					fmt.Println("숫자를 중복 되게 입력하였습니다")
					duplicated = true
					break
				}
			}

			if !duplicated {
				res = input
				success = true
				break
			} else {
				break
			}
		}
		if success {
			break
		}
	}
	//fmt.Println(input)
	return res
}

// CompareNum 사용자의 input을 바탕으로 정답과 동일한지 검사
func CompareNum(game, users [3]int) bool {
	res := false

	strikes := 0
	balls := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game[i] == users[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}

	if strikes < 3 {
		fmt.Println(strikes, "S", balls, "B")
	} else {
		res = true
	}
	return res
}
