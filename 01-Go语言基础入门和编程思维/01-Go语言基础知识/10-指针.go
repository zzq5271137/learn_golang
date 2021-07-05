package main

import "fmt"

/**
 * 指针
 *
 * 注:
 * 1. golang的指针比c的指针简单很多, 原因在于golang的指针不能参与运算;
 * 2. 关于参数的值传递和引用传递, 在很多语言中, 大部分自定义类型的参数都是引用传递(例如java和python);
 *    但是在在golang中, 只有值传递这一种方式;
 */

func pointerIntro() {
	fmt.Println("pointerIntro():")
	var a int = 2
	fmt.Printf("Before: %d\n", a)
	var pa *int = &a
	*pa = 3
	fmt.Printf("After: %d\n", a)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	pointerIntro()

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
