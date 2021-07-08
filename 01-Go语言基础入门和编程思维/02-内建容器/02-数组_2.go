package main

import "fmt"

/**
 * 数组
 *
 * 数组是值类型:
 * 1. [10]int 和 [20]int, 对于golang来说是不同的类型
 * 2. 调用 func f(arr [10]int) 会拷贝数组, 即在函数内改变数组元素的值, 并不会影响函数外的数组
 *    想要实现数组的引用传递, 必须使用指针 (或者使用数组切片)
 *    (记得, 参数传递只有一种方式, 就是值传递)
 */

func printArray1(arr [3]int) { // 参数为一个int类型的数组, 其长度必须为3
	arr[0] = 1000
	fmt.Printf("In function printArray1, change arr[0]=%d, then traverse the array: \n", arr[0])
	for i, v := range arr {
		fmt.Printf("arr[%d]=%d\n", i, v)
	}
}

func printArray2(arr *[3]int) {
	//(*arr)[0] = 1000
	arr[0] = 1000 // 省略写法
	fmt.Printf("In function printArray2, change arr[0]=%d, then traverse the array: \n", arr[0])
	for i, v := range arr {
		fmt.Printf("arr[%d]=%d\n", i, v)
	}
}

func main() {
	arr1 := [3]int{3, 5, 7}
	fmt.Printf("Before calling printArray1, arr1=%v\n", arr1)
	printArray1(arr1)
	fmt.Printf("After calling printArray1, arr1=%v\n", arr1)

	fmt.Println("################################################")

	arr2 := [3]int{3, 5, 7}
	fmt.Printf("Before calling printArray2, arr2=%v\n", arr2)
	printArray2(&arr2)
	fmt.Printf("After calling printArray1, arr2=%v\n", arr2)
}
