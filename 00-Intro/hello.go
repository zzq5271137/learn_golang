/**
 * golang介绍
 */

/*
 * 1. golang是以包(package)为组织单位的
 * 2. 每个.go源文件必须声明包
 * 3. 程序的入口为"main"包下的"main()"函数
 */
package main

import "fmt"

/*
 * main()函数为程序的入口, 此函数不能传参, 也不能定义返回值
 */
func main() {
	fmt.Println("hello world")
}
