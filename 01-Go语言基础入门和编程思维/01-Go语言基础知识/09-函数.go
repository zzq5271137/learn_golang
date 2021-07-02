package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
 * 函数
 *
 * 注: golang中的函数, 没有java中的重载, 没有python中的默认参数、可选参数等花哨的概念
 */

//func div(a, b int) (int, int) {
func div(a, b int) (q, r int) { // 可以给返回值起名字
	//return a / b, a % b
	q = a / b
	r = a % b
	return // 给返回值起名字后, 直接return, 就可以把相应的返回值返回出去
}

func eval2(a, b int, op string) (int, error) { // 定义一个返回值加上一个error, 是golang中的常用做法
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b) // 使用 _ 进行占位
		return q, nil
	default:
		//panic("unsupported operator: " + op) // 使用panic(), 程序会中断
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
}

func apply(op func(int, int) int, a, b int) int { // 函数可以作为参数
	p := (reflect.ValueOf(op)).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int { // 可变参数列表
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	q, r := div(13, 3)
	fmt.Printf("商: %d, 余数: %d\n", q, r)

	if result, err := eval2(13, 3, "$"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(pow, 2, 3))
	fmt.Println(apply(
		func(a, b int) int { // 匿名函数
			return int(math.Pow(float64(a), float64(b)))
		}, 2, 3))

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
}
