/**
 * rune类型
 */

package main

import (
	"fmt"
	"unicode/utf8"
)

/*
 * 寻找最长不含有重复字符的子串, 例如:
 * abcabcbb -> abc
 * bbbbb -> b
 * pwwkew -> wke
 */
func lengthOfNoneRepeatingSubstring1(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) { // 将string字符串转成byte数组(slice), 再遍历
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

/*
 * 寻找最长不含有重复字符的子串, 例如:
 * abcabcbb -> abc
 * bbbbb -> b
 * pwwkew -> wke
 *
 * 加强版, 使用rune数组(slice), 解决中文字符串的问题
 */
func lengthOfNoneRepeatingSubstring2(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) { // 将string字符串转成rune数组(slice), 而不是byte数组(slice)
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func runeTest() {
	fmt.Println("runeTest():")

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
	fmt.Printf("%q: %d\n", "abcabcbb", lengthOfNoneRepeatingSubstring1("abcabcbb"))
	fmt.Printf("%q: %d\n", "bbbbb", lengthOfNoneRepeatingSubstring1("bbbbb"))
	fmt.Printf("%q: %d\n", "pwwkew", lengthOfNoneRepeatingSubstring1("pwwkew"))
	fmt.Printf("%q: %d\n", "", lengthOfNoneRepeatingSubstring1(""))
	fmt.Printf("%q: %d\n", "b", lengthOfNoneRepeatingSubstring1("b"))
	fmt.Printf("%q: %d\n", "abcdefg", lengthOfNoneRepeatingSubstring1("abcdefg"))

	// 中文字符会有问题
	fmt.Printf("%q: %d\n", "一二三一二", lengthOfNoneRepeatingSubstring1("一二三一二"))
	fmt.Printf("%q: %d\n", "一二三四五", lengthOfNoneRepeatingSubstring1("一二三四五"))

	// 解决了中文字符串的问题
	fmt.Printf("%q: %d\n", "一二三一二", lengthOfNoneRepeatingSubstring2("一二三一二"))
	fmt.Printf("%q: %d\n", "一二三四五", lengthOfNoneRepeatingSubstring2("一二三四五"))

	runeTest()
}
