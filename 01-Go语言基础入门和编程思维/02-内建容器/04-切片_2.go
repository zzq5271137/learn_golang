package main

import "fmt"

/**
 * 切片的操作
 */

func sliceAppend() {
	fmt.Println("sliceAppend():")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Printf("s1=%v\n", s1)   // [2 3 4 5]
	fmt.Printf("s2=%v\n", s2)   // [5 6]
	fmt.Printf("s3=%v\n", s3)   // [5 6 10]
	fmt.Printf("s4=%v\n", s4)   // [5 6 10 11]
	fmt.Printf("s5=%v\n", s5)   // [5 6 10 11 12]
	fmt.Printf("arr=%v\n", arr) // [0 1 2 3 4 5 6 10]

	/*
	 * 可以看到, s3的append, 将arr的最后一个元素7替换成了10,
	 * s4和s5的append, 没有影响到arr, s4和s5不再是arr的view;
	 *
	 * 即, 添加元素时如果超越cap, 系统会重新分配更大的底层数组;
	 * 由于值传递的关系, 必须接收append的返回值;
	 */
}

func printSlice(s []int) {
	fmt.Printf("s=%v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))
}

func sliceCreate1() {
	fmt.Println("sliceCreate1():")

	/*
	 * 不一定非得从array中创建slice, 可以直接定义slice变量;
	 * Zero value for slice is nil
	 */
	var s []int
	fmt.Println("var s []int, s == nil:", s == nil)
	for i := 0; i < 18; i++ {
		printSlice(s) // 可以看到, 装不下时, 每次扩充容量cap会乘以2
		s = append(s, i*2+1)
	}
	fmt.Printf("%v\n", s)
}

func sliceCreate2() {
	fmt.Println("sliceCreate2():")
	s := []int{1, 3, 5, 7} // 执行这行代码, golang在底层创建了一个array {1,3,5,7}, 然后又创建了一个slice, 去view这个array (arr[:])
	fmt.Printf("s=%v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))
}

func sliceCreate3() {
	fmt.Println("sliceCreate3():")

	// 使用 make() 函数创建切片, make(type, len, cap)
	s1 := make([]int, 16)
	s2 := make([]int, 10, 32)
	printSlice(s1)
	printSlice(s2)
}

func sliceCopy() {
	fmt.Println("sliceCopy():")
	s1 := []int{1, 3, 5, 7}
	s2 := make([]int, 16)

	// 使用 copy() 函数复制切片
	copy(s2, s1)
	printSlice(s2)
}

func deleteElementFromSliceMiddle() {
	fmt.Println("deleteElementFromSliceMiddle():")
	s := []int{0, 1, 2, 3, 4, 5, 6, 7} // 想要删掉3这个元素 (删除中间的元素)
	printSlice(s)
	s = append(s[:3], s[4:]...) // 特殊语法, 以可变参数的形式传递参数
	printSlice(s)
}

func popElementFromSliceFront() {
	fmt.Println("popElementFromSliceFront():")
	s := []int{0, 1, 2, 3, 4, 5, 6, 7} // 想删掉第一个元素
	printSlice(s)
	s = s[1:]
	printSlice(s)
}

func popElementFromSliceBack() {
	fmt.Println("popElementFromSliceBack():")
	s := []int{0, 1, 2, 3, 4, 5, 6, 7} // 想删掉最后一个元素
	printSlice(s)
	s = s[:len(s)-1]
	printSlice(s)
}

func main() {
	sliceAppend()
	sliceCreate1()
	sliceCreate2()
	sliceCreate3()
	sliceCopy()
	deleteElementFromSliceMiddle()
	popElementFromSliceFront()
	popElementFromSliceBack()
}
