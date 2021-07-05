package main

import "fmt"

/**
 * 切片介绍
 *
 * 切片本身并不是一个值类型, 切片(slice)是对数组(array)的一个视图(view)
 */

func sliceIntro() {
	fmt.Println("sliceIntro():")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr=", arr)
	fmt.Println("arr[2:6]=", arr[2:6]) // 左开右闭
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:]=", arr[:])
}

func updateSlice(s []int) { // 以切片为参数
	s[0] = 100
}

func sliceUpdateIntro() {
	fmt.Println("sliceUpdateIntro():")
	arr := [...]int{0, 1, 2, 3, 4}
	s1 := arr[2:]
	s2 := arr[:]
	fmt.Printf("Before calling updateSlice(s1), arr=%v, s1=%v, s2=%v\n", arr, s1, s2)
	updateSlice(s1)
	fmt.Printf("After calling updateSlice(s1), arr=%v, s1=%v, s2=%v\n", arr, s1, s2)
	updateSlice(s2) // 可以以这种方式, 代替数组指针的使用, 实现数组的引用传递的效果
	fmt.Printf("After calling updateSlice(s2), arr=%v, s1=%v, s2=%v\n", arr, s1, s2)
}

func reslice() {
	fmt.Println("reslice():")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[:]
	s2 := s1[:5]
	s3 := s2[2:]
	fmt.Println("arr=", arr)
	fmt.Println("s1(arr[:])=", s1)
	fmt.Println("s2(s1[:5])=", s2)
	fmt.Println("s3(s2[2:])=", s3)
}

func sliceExtend() {
	/*
	 * slice有三个属性:
	 * 1. ptr: 指向slice开头的那个元素
	 * 2. len: 表明slice的长度, 使用 s[i] 的方式取值时, 只能取到len里面的值
	 * 3. cap: 表明slice的容量, 使用 s[n1:n2] 的切片的方式时, 可以向后扩展
	 *
	 * 注:
	 * 1. 以下方为例, s1的len为{2, 3, 4, 5}, 而s1的cap为{2, 3, 4, 5, 6, 7},
	 *    所以 s1[4] 会报错, 而 s1[3:5] 却可以取出值; 但是slice只能向后扩展, 不能向前扩展;
	 * 2. s[i] 不可以超越 len(s), 向后扩展不可以超越底层数组 cap(s);
	 */
	fmt.Println("sliceExtend():")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("arr=", arr)
	fmt.Println("s1(arr[2:6])=", s1)
	fmt.Println("s2(s1[3:5])=", s2) // slice可以向后扩展, 但不可以向前扩展

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
}

func main() {
	sliceIntro()
	sliceUpdateIntro()
	reslice()
	sliceExtend()
}
