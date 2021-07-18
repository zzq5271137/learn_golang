/**
 * 自动推导类型
 */

package main

import "fmt"

func typeDeduction() {
	fmt.Println("typeReduction():")

	/*
	 * 使用 := 声明变量并赋值, 编译器会自动推导出变量类型
	 */
	myVar := 10
	fmt.Printf("(%d, %T)\n", myVar, myVar) // %T: 打印类型

	a, b, c, d, e := 2.5, "hello", true, [...]int{1, 2, 3}, []string{"hello", "world"} // 可以同时声明多个不同类型的变量
	fmt.Printf("(%.2f, %T) (%s, %T) (%t, %T) (%v, %T) (%v, %T)\n", a, a, b, b, c, c, d, d, e, e)
}

func valueSwap() {
	fmt.Println("valueSwap():")

	i, j := 10, 20
	fmt.Printf("Before swap, i=%d, j=%d\n", i, j)
	i, j = j, i
	fmt.Printf("After swap, i=%d, j=%d\n", i, j)
}

func main() {
	typeDeduction()
	valueSwap()
}
