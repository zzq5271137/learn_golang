package main

import (
	"fmt"
	"strconv"
)

/**
 * 循环语句: for
 */

func sum100() int {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return sum
}

func convertToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	fmt.Println(sum100())

	fmt.Printf(
		"%q, %q, %q, %q",
		convertToBinary(5),  // 101
		convertToBinary(13), // 1101
		convertToBinary(762189372),
		convertToBinary(0),
	)
}
