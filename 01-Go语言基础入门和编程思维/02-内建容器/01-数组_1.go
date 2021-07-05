package main

import "fmt"

/**
 * 数组
 *
 * 数组的简单介绍;
 */

func arrIntro() {
	fmt.Println("arrIntro():")
	// 数组定义方式一
	var arr1 [3]int

	// 数组定义方式二, 此种方式必须定义初始值
	arr2 := [3]int{1, 3, 5}

	// 数组定义方式三 (效果同方式一)
	arr3 := [3]int{}

	// 数组定义方式四, 可以不定义数组大小, 编译器自己推导大小 (但是要在[]中加上...)
	arr4 := [...]int{2, 4, 6, 8}

	fmt.Printf("arr1: %v\narr2: %v\narr3: %v\narr4: %v\n",
		arr1, arr2, arr3, arr4)
}

func twoDimensionalArrIntro() {
	fmt.Println("twoDimensionalArrIntro():")
	// 二维数组
	grid := [4][5]int{} // 4行5列
	fmt.Printf("grid: %v\n", grid)
}

func traverse1() {
	fmt.Println("traverse1():")
	arr1 := [3]int{1, 3, 5}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
}

func traverse2() {
	fmt.Println("traverse2():")
	arr1 := [3]int{1, 3, 5}
	for i := range arr1 { // 使用range关键字
		fmt.Println(arr1[i])
	}
}

func traverse3() {
	fmt.Println("traverse3():")
	arr1 := [3]int{1, 3, 5}
	for i, v := range arr1 { // 可以直接把值取出来
		fmt.Printf("arr[%d]=%d\n", i, v)
	}
}

func traverse4() {
	fmt.Println("traverse4():")
	arr1 := [3]int{1, 3, 5}
	for _, v := range arr1 { // 只要值 (使用 _ 省略变量)
		fmt.Println(v)
	}

	/*
	 * 在golang中, 使用 _ 省略变量的方式, 在任何地方都适用
	 */
}

func maxValue() {
	fmt.Println("maxValue():")
	numbers := [...]int{6, 10, 4, 5, 7, 13}
	maxi := 0
	maxv := numbers[maxi]
	for i, v := range numbers {
		if v > maxv {
			maxi, maxv = i, v
		}
	}
	fmt.Printf("arr=%v, maxi=%d, maxv=%d\n", numbers, maxi, maxv)
}

func main() {
	arrIntro()
	twoDimensionalArrIntro()
	traverse1()
	traverse2()
	traverse3()
	traverse4()
	maxValue()
}
