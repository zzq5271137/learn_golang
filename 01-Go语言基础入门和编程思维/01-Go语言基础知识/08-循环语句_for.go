package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * 循环语句: for
 */

/*
 * 完整的for循环演示
 *
 * for 起始条件; 终止条件; 递增表达式 {
 *     ...
 * }
 */
func sum100() int {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return sum
}

/*
 * 省略起始条件的for循环演示
 */
func convertToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

/*
 * 省略起始条件和递增表达式的for循环演示
 */
func printFile(filename string) {
	fmt.Println("printFile():")
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	//for ; scanner.Scan(); { // 没有初始条件和递增表达式, 只有结束条件的for循环, 相当于其他语言的while循环 (golang中没有while循环)
	for scanner.Scan() { // 也可以写成这样 (省略分号 ;)
		fmt.Println(scanner.Text())
	}
}

/*
 * 起始条件、终止条件、递增表达式都省略的for循环演示 (死循环)
 */
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(sum100())

	fmt.Printf(
		"%q, %q, %q, %q\n",
		convertToBinary(5),  // 101
		convertToBinary(13), // 1101
		convertToBinary(762189372),
		convertToBinary(0),
	)

	printFile("/Users/zzq/Documents/workspaces/GoLand_workspace/learn_golang/01-Go语言基础入门和编程思维/01-Go语言基础知识/08-循环语句_for.txt")

	//forever()
}
