package main

import "fmt"

/**
 * map
 */

func mapCreate1() {
	fmt.Println("mapCreate1():")

	/*
	 * 创建map:
	 * m := map[key-type]value-type{}
	 */
	m := map[string]string{
		"name":   "zzqgo",
		"age":    "100",
		"course": "golang",
	}
	fmt.Println(m)
}

func mapCreate2() {
	fmt.Println("mapCreate2():")

	// 使用 make() 函数创建map
	m1 := make(map[string]int) // empty map
	fmt.Printf("(make) m1: %v, m1 == nil: %t\n", m1, m1 == nil)

	var m2 map[string]int // nil
	fmt.Printf("(var) m2: %v, m2 == nil: %t\n", m2, m2 == nil)
}

func mapTraverse() {
	fmt.Println("mapTraverse():")
	m := map[string]string{
		"name":   "zzqgo",
		"age":    "100",
		"course": "golang",
	}

	/*
	 * 和遍历数组一样, map的遍历可以使用range,
	 * 并且, 依然可以省略 k 或 v, 或者使用 _ 进行占位
	 */
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v) // 从打印的结果可以看出, map是无序的, 每次打印的顺序可能都不一样
	}
}

func getValueFromMap() {
	fmt.Println("getValueFromMap():")
	m := map[string]string{
		"name":   "zzqgo",
		"age":    "100",
		"course": "golang",
	}

	fmt.Printf("m=%v\n", m)
	fmt.Printf("m[\"course\"]=%q\n", m["course"])

	// 若map中不存在此key, 不会报错, 会返回一个zero value (根据value的类型, 会返回不同的zero value, 对于string类型来说, 是空字符串 "")
	fmt.Printf("m[\"whatever\"]=%q\n", m["whatever"])

	key := "whatever"
	if myValue, ok := m[key]; ok { // 可以以这种方式判断key存不存在
		fmt.Printf("%s=%s\n", key, myValue)
	} else {
		fmt.Printf("Not exist key: %s\n", key)
	}
}

func deleteFromMap() {
	fmt.Println("deleteFromMap():")
	m := map[string]string{
		"name":   "zzqgo",
		"age":    "100",
		"course": "golang",
	}
	fmt.Printf("Before delete, m=%v\n", m)
	delete(m, "course")
	fmt.Printf("After delete, m=%v\n", m)
}

func main() {
	mapCreate1()
	mapCreate2()
	mapTraverse()
	getValueFromMap()
	deleteFromMap()
}

/*
 * 关于map的key, golang中什么类型可以作为map的key:
 * 1. map使用哈希表, key必须可以比较是否相等
 * 2. 对于golang的内建类型, 除了slice、map、function的内建类型都可以作为map的key
 * 3. 对于自定义类型(Struct类型), 若不包含类型为slice、map、function的字段, 则也可以作为map的key
 */
