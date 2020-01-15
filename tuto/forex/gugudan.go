package main

import "fmt"

func main() {
	//gugu()
	star()
	star1()
	star2()
	starsol()
}

func gugu() {
	for dan := 1; dan < 10; dan++ {
		/*
			if i == 5 {
				continue
			}
		*/
		fmt.Println(dan, "단")

		for i := 1; i < 10; i++ {
			if dan == 3 && i == 2 {
				continue
			}

			fmt.Println(dan, "x", i, "=", dan*i)
		}

	}
}

func star() {
	num := 10

	for i := 1; i < num; i += 2 {
		for k := 1; k < num-i; k += 2 {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func star1() {
	num := 10

	for i := 1; i < num; i++ {

		if i < 5 {
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
		} else {
			for j := 0; j < num-i; j++ {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func star2() {
	num := 10

	for i := 1; i < num; i += 2 {
		for k := 0; k < num-i; k += 2 {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := num + 1; i > 0; i -= 2 {
		for k := 0; k < num-i; k += 2 {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func starsol() {

	line := 4
	for i := 0; i < line; i++ {
		for space := 0; space < line-i-1; space++ {
			fmt.Print(" ")
		}
		for star := 0; star < 2*i+1; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
