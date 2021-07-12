package main

import "fmt"

/*
 * 寻找最长不含有重复字符的子串, 例如:
 * abcabcbb -> abc
 * bbbbb -> b
 * pwwkew -> wke
 */

func lengthOfNoneRepeatingSubstring(s string) int {
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

func main() {
	fmt.Printf("%q: %d\n", "abcabcbb", lengthOfNoneRepeatingSubstring("abcabcbb"))
	fmt.Printf("%q: %d\n", "bbbbb", lengthOfNoneRepeatingSubstring("bbbbb"))
	fmt.Printf("%q: %d\n", "pwwkew", lengthOfNoneRepeatingSubstring("pwwkew"))
	fmt.Printf("%q: %d\n", "", lengthOfNoneRepeatingSubstring(""))
	fmt.Printf("%q: %d\n", "b", lengthOfNoneRepeatingSubstring("b"))
	fmt.Printf("%q: %d\n", "abcdefg", lengthOfNoneRepeatingSubstring("abcdefg"))

	// 中文字符会有问题
	fmt.Printf("%q: %d\n", "一二三一二", lengthOfNoneRepeatingSubstring("一二三一二"))
	fmt.Printf("%q: %d\n", "一二三四五", lengthOfNoneRepeatingSubstring("一二三四五"))
}
