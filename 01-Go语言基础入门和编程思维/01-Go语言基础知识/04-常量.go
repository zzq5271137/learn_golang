package main

import (
	"fmt"
	"math"
)

/**
 * 常量
 *
 * 注意点:
 * 1. 在go语言中定义常量时, 并不会把名字全部大写, 因为大小写在go语言中是由含义的; 常量的命名遵循普通变量的命名即可(驼峰)
 * 2. 声明常亮时, 如果不声明类型, 则此常量可以作为各种类型来使用
 */

func consts() {
	//const test  // 使用 const 定义常量, 必须给初始值
	const fileName string = "abc.txt" // 定义常量时, 可以声明出类型
	const a, b = 3, 4                 // 也可以不声明类型, 不声明类型时, 常量的类型是不确定的
	const (                           // 一次定义多个常量也可以这样写
		clentId string = "c2005"
		x, y           = 2.3, 4.5
	)
	var c int
	c = int(math.Sqrt(a*a + b*b)) // 这里不把a和b转成float64也是可以的, 因为a和b是没有声明类型的常量
	fmt.Println(fileName, c, clentId, x, y)
}

const studentId = "820064" // 常量也可以定义在函数的外面 (即包内), 所有的函数都能用

func main() {
	consts()
	fmt.Println(studentId)
}
