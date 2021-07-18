/**
 * iota枚举
 *
 * 常量声明可以使用iota常量生成器初始化, 它用于生成一组以相似规则初始化的常量, 不用每一行都写一遍初始化表达式;
 * 在一个const声明语句中, 在第一个声明的常量所在的行, iota将被置为0, 然后在每一个有常量声明的行加一;
 * iota是给常量赋值使用的, 不能给变量赋值
 */

package main

import "fmt"

func iotaTest() {
	fmt.Println("defineConst():")

	const (
		a = iota // 0
		b        // 1
		c        // 2
		d        // 3
	)
	fmt.Println(a, b, c, d)

	const e = iota // 每遇到一个const关键字, iota会被重置
	fmt.Println(e)

	const (
		f, g, h = iota, iota, iota // f=0, g=0, h=0, 在同一行, iota值相同
	)
	fmt.Println(f, g, h)
}

func main() {
	iotaTest()
}
