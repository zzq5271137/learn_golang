package main

import "fmt"

/**
 * 条件语句: switch
 */

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b // golang的 switch 会自动 break, 除非使用 fallthrough
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator: " + op) // panic() 是报错打印
	}
	return result
}

func grade(score int) string {
	g := ""
	switch { // switch 关键字后面也可以没有表达式, 将判断放入 case 里面
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("wrong score: %d", score))
	}
	return g
}

func main() {
	//fmt.Println(eval(2, 3, "+"))
	//fmt.Println(eval(2, 3, "^"))

	fmt.Println(
		grade(20),
		grade(62),
		grade(88),
		grade(93),
	)
	fmt.Println(grade(120))
}
