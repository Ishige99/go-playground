package main

import (
	"fmt"
	"strings"
)

func main() {
	mode := 0

	fmt.Println("select mode")
	fmt.Printf("1: easy,\n2: difficult,\n")

	fmt.Scan(&mode)
	TestFunc(mode)
}

func TestFunc(n int) {
	result := 0

	switch n {
	case 1:
		fmt.Println("easy mode!")
		result += TypingTestEasy(1, "dog")
		result += TypingTestEasy(2, "cat")
		result += TypingTestEasy(3, "fish")
		result += TypingTestEasy(4, "god")
	case 2:
		fmt.Println("difficult mode!")
		result += TypingTestDifficult(1, "Computer")
		result += TypingTestDifficult(2, "Keyboard")
		result += TypingTestDifficult(3, "Software")
		result += TypingTestDifficult(4, "Internet")
	default:
		fmt.Println("under development....")
	}

	fmt.Printf("--- result: %d point --- \n", result)
}

func TypingTestEasy(n int, q string) int {
	var input string

	fmt.Printf("第%d問目", n)
	fmt.Printf("この値を入力してください：%s\n", q)
	fmt.Scan(&input)

	if strings.EqualFold(q, input) {
		fmt.Println("true! +10 point")
		return 10
	} else {
		fmt.Println("false.. +0 point")
		return 0
	}
}

func TypingTestDifficult(n int, q string) int {
	var input string

	fmt.Printf("第%d問目", n)
	fmt.Printf("この値を入力してください：%s\n", q)
	fmt.Scan(&input)

	if strings.EqualFold(q, input) {
		fmt.Println("true! +10 point")
		return 10
	} else {
		fmt.Println("false.. +0 point")
		return 0
	}
}
