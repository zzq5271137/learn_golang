package main

import (
	"fmt"
	"unicode/utf8"
)

/**
 * string和rune
 */

func chineseIntro() {
	fmt.Println("chineseIntro():")
	s := "Yes@你好世界"
	fmt.Printf("s: %s, len(s): %d\n", s, len(s)) // len(s) 方法获得的是字节数, 而不是字符数
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) // %X打印出来的, 是utf-8编码; 可以看到, 每个中文字占3个字节 (utf-8编码采用可变长, 英文字符占1字节, 中文字符占3字节)
	}
	fmt.Println()

	for i, ch := range s { // ch是一个rune (rune类型和int32是互为别名, 占4字节)
		fmt.Printf("(%d, %X, %c) ", i, ch, ch) // %X打印出来的, 是unicode编码
	}
	fmt.Println()

	fmt.Println("utf8.RuneCountInString(s):", utf8.RuneCountInString(s)) // 可以看到, s中有8个rune, 英文和中文字符都各占一个rune

	sBytes := []byte(s)
	for len(sBytes) > 0 {
		ch, size := utf8.DecodeRune(sBytes)
		sBytes = sBytes[size:]
		fmt.Printf("(%X, %c) ", ch, ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) { // 可以直接转成rune数组(slice), 再遍历 (注意与上面的 for i, ch := range s 的遍历做对比)
		fmt.Printf("(%d, %X, %c) ", i, ch, ch)
	}
	fmt.Println()
}

func main() {
	chineseIntro()
}

/*
 * Tips: "strings"包中, 定义了很多关于字符串的操作, 例如:
 * 1. strings.Fields(), strings.Split(), strings.Join()
 * 2. strings.Contains(), strings.Index()
 * 3. strings.ToLower(), strings.ToUpper()
 * 4. strings.Trim(), strings.TrimLeft(), strings.TrimRight()
 * 等等, 遇到需要处理字符串的情况, 优先使用这个包
 */
