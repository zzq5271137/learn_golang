/**
 * 匿名变量
 */

package main

import "fmt"

func testReturnMultiple() (int, int, int) {
	return 1, 3, 5
}

func main() {
	/*
	 * 以 _ 命名的变量为匿名变量, 会丢弃数据不处理
	 */
	num, _ := 7, 8
	fmt.Println(num)

	a, _, c := testReturnMultiple() // 例如当函数返回多个值时, 我们只需要其中一部分
	fmt.Println(a, c)
}
