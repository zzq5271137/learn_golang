/**
 * 格式化输出总结
 */

package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func formatPrint() {
	fmt.Println("formatPrint():")

	/*
	 * %v: 输出结构体, 例如, {10 30}
	 * %+v: 输出结构体显示字段名, 例如, {x:10 y:30}
	 * %#v: 输出结构体源代码片段, 例如, main.Point{x:10, y:30}
	 * %T: 输出值的类型, 例如, main.Point
	 * %t: 输出格式化布尔值, 例如, true
	 * %d: 输出标准的十进制格式化, 例如, 100
	 * %b: 输出标准的二进制格式化, 例如, 99 对应 1100011
	 * %c: 输出定整数的对应字符, 例如, 99 对应 c
	 * %x: 输出十六进制编码, 例如, 99 对应 63
	 * %f: 输出十进制格式化, 例如, 99 对应 63
	 * %e: 输出科学技科学记数法表示形式, 例如, 123400000.0 对应 1.234000e+08
	 * %E: 输出科学技科学记数法表示形式, 例如, 123400000.0 对应 1.234000E+08
	 * %s: 进行基本的字符串输出, 例如, "\"string\""  对应 "string"
	 * %q: 源代码中那样带有双引号的输出, 例如, "\"string\""  对应 "\"string\""
	 * %p: 输出一个指针的值, 例如, &jgt 对应 0xc00004a090
	 * %: 后面使用数字来控制输出宽度, 默认结果使用右对齐并且通过空格来填充空白部分
	 * %5.2f: 指定浮点型的输出宽度, 例如, %5.2f 表示总长度5(包含小数点), 小数点后2位
	 * %*.*f: 宽度与精度均可用字符 '*' 表示
	 */
	jgt := Point{10, 30}
	fmt.Printf("%v\n", jgt)          // {10 30}
	fmt.Printf("%+v\n", jgt)         // {x:10 y:30}
	fmt.Printf("%#v\n", jgt)         // main.Point{x:10, y:30}
	fmt.Printf("%T\n", jgt)          // main.Point
	fmt.Printf("%t\n", true)         // true
	fmt.Printf("%d\n", 100)          // 100
	fmt.Printf("%b\n", 99)           // 99 对应 1100011
	fmt.Printf("%c\n", 99)           // 99 对应 c
	fmt.Printf("%x\n", 99)           // 99 对应 63
	fmt.Printf("%f\n", 99.9)         // 99.9 对应 99.900000
	fmt.Printf("%e\n", 123400000.0)  // 123400000.0 对应 1.234000e+08
	fmt.Printf("%E\n", 123400000.0)  // 123400000.0 对应 1.234000E+08
	fmt.Printf("%s\n", "\"string\"") // "\"string\""  对应 "string"
	fmt.Printf("%q\n", "\"string\"") // "\"string\""  对应 "\"string\""
	fmt.Printf("%p\n", &jgt)         // &jgt 对应 0xc00004a090
	fmt.Printf("%6d\n", 8)           // 打印数字8, 前面有5个空格, 即总长度6
	fmt.Printf("%6s\n", "abc")       // 打印acb, 前面有3个空格, 即总长度6
	fmt.Printf("%5.2f\n", 1.2)       // 打印1.20, 前面有1个空格, 即总长度5, 小数点后2位
	fmt.Printf("%*.*f\n", 5, 2, 1.2) // 打印1.20, 前面有1个空格, 即总长度5, 小数点后2位

	// Sprintf() 格式化并返回一个字符串而不带任何输出
	s := fmt.Sprintf("是字符串: %s", "string")
	fmt.Println(s)
}

func main() {
	formatPrint()
}
