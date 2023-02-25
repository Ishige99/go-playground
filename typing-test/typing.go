package main

import "fmt"

func main() {
	result := 0

	result += TypingTest(1, "dog")
	result += TypingTest(2, "cat")
	result += TypingTest(3, "fish")
	result += TypingTest(4, "god")

	fmt.Printf("--- result: %d point --- \n", result)
}

func TypingTest(n int, q string) int {
	var input string

	fmt.Printf("第%d問目", n)
	fmt.Printf("この値を入力してください：%s\n", q)
	fmt.Scan(&input)

	if q == input {
		fmt.Println("true! +10 point")
		return 10
	} else {
		fmt.Println("false.. +0 point")
		return 0
	}
}
