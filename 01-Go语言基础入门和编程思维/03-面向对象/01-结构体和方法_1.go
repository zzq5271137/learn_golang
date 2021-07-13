package main

import "fmt"

/**
 * 结构体和方法
 *
 * 注:
 * 1. 关于golang的面向对象, golang仅支持封装, 不支持继承和多态
 * 2. golang没有class, 只有struct
 */

type treeNode struct {
	value       int
	left, right *treeNode
}

/*
 * golang中, 给对象(结构体)定义方法的方式和其他的面向对象语言不太一样;
 * 结构体的方法的定义, 和普通的函数一样, 都是func;
 * 只不过, 结构体的方法在方法名前, 多了一个"接收者", 例如下面的 (node treeNode)
 * "接收者"类似于java中的this, 或者python中的self;
 * 所谓"接收者", 其实也是一种参数传递, 而且依然是值传递;
 */
func (node treeNode) setValue(v int) {
	node.value = v
}

/*
 * 值接收者 VS 指针接收者:
 * 1. 要改变内容, 必须使用指针接收者
 * 2. 结构体过大, 也考虑使用指针接收者 (由于值接收者是拷贝, 考虑性能, 所以结构体过大时建议使用指针接收者)
 * 3. (建议)一致性: 如有指针接收者, 最好都用指针接收者
 * 4. 值/指针接收者均可接收值/指针 (即, 不论方法的接收者定义成值或者指针, 对象或者对象的指针均可调用该方法)
 */
func (node *treeNode) setValue2(v int) {
	if node == nil { // 接收者允许是nil指针(空指针)
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.value = v // 不论是结构体本身还是地址, 一律使用"."来访问成员, 编译器会自动解析
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	fmt.Println(node.value)
	node.right.traverse()
}

/*
 * golang中没有构造方法的概念, 可以使用工厂函数来创建构造体
 */
func createNode(value int) *treeNode {
	/*
	 * 这里是创建了一个局部变量, 并返回了该局部变量的地址, 在golang中, 这没问题;
	 *
	 * 这其实涉及到, 这里的结构体treeNode是创建在堆上的还是栈上的?
	 * 这个问题对于其他的编程语言来说很重要, 比如c++或者java;
	 * 一般来说, 函数体的局部变量是分配在栈上的, 函数一旦退出, 局部变量就会立刻被销毁;
	 * 在java中, 对象都是在堆上的(使用new关键字创建对象), 所以才会有垃圾回收机制;
	 * 那么在golang中, 函数内的局部变量是分配在堆上的还是栈上的呢?
	 * 答案是, 不需要知道, 或者, 不一定; 分配在堆上或者栈上, 取决于编译器;
	 * 例如, 下面的treeNode, 当编译器看到这个结构体最后被取地址并且返回出去给别人用了, 那么编译器就会将它分配到堆上(当然, 被分配到堆上, 就会参与垃圾回收);
	 * 反之, 如果这个treeNode只是在函数体内被使用, 那么编译器就会将它分配到栈上, 当函数退出就会被销毁;
	 */
	return &treeNode{value: value}
}

func structCreate() {
	fmt.Println("structCreate():")
	var root treeNode // 只声明不赋值, 则结构体内成员的值为其类型的Zero Value
	fmt.Println(root)

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) // 不论是结构体本身还是地址, 一律使用"."来访问成员, 编译器会自动解析
	root.left.right = createNode(2)

	nodes := []treeNode{
		{value: 3},
		{},
		{5, nil, &root},
	}
	fmt.Println(nodes)
}

func structMethod() {
	fmt.Println("structMethod():")
	node := treeNode{value: 1}
	fmt.Println("Before node.setValue(100), node.value =", node.value)
	node.setValue(100)
	fmt.Println("After node.setValue(100), node.value =", node.value)
	fmt.Println("Before node.setValue2(100), node.value =", node.value)
	node.setValue2(100)
	fmt.Println("After node.setValue2(100), node.value =", node.value)

	var node2 *treeNode // nil
	node2.setValue2(20) // nil指针依然可以调用方法
}

func traverseTest() {
	fmt.Println("traverseTest():")
	root := treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) // 不论是结构体本身还是地址, 一律使用"."来访问成员, 编译器会自动解析
	root.left.right = createNode(2)
	root.traverse()
}

func main() {
	structCreate()
	structMethod()
	traverseTest()
}
