package main

import "fmt"

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

func main() {
	fmt.Printf("%q: %d\n", "abcabcbb", lengthOfNoneRepeatingSubstring2("abcabcbb"))
	fmt.Printf("%q: %d\n", "bbbbb", lengthOfNoneRepeatingSubstring2("bbbbb"))
	fmt.Printf("%q: %d\n", "pwwkew", lengthOfNoneRepeatingSubstring2("pwwkew"))
	fmt.Printf("%q: %d\n", "", lengthOfNoneRepeatingSubstring2(""))
	fmt.Printf("%q: %d\n", "b", lengthOfNoneRepeatingSubstring2("b"))
	fmt.Printf("%q: %d\n", "abcdefg", lengthOfNoneRepeatingSubstring2("abcdefg"))

	// 解决了中文字符串的问题
	fmt.Printf("%q: %d\n", "一二三一二", lengthOfNoneRepeatingSubstring2("一二三一二"))
	fmt.Printf("%q: %d\n", "一二三四五", lengthOfNoneRepeatingSubstring2("一二三四五"))
}
