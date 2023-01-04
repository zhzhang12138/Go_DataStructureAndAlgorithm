package main

import "fmt"

/*
环形链表
	环形链表是一种特殊的链表，它的末尾的下一个节点指向链表的第一个节点，形成一个环的结构。

与普通链表相比，环形链表具有以下几个特点：
	它的末尾的下一个节点指向链表的第一个节点，而不是指向 nil。
	由于链表的末尾与开头相连，因此遍历环形链表时，必须记录遍历的起点和终点。
	在环形链表中，每个节点都有一个 next 指针，指向下一个节点。
*/

// Node 定义一个节点类型，包含数据域和指向下一个节点的指针域。
type Node struct {
	data int
	next *Node // 表示指向下一个结点
}

// 在环形链表的末尾添加新节点
func appendNode(tail *Node, data int) *Node {
	newNode := &Node{data: data, next: tail.next}
	tail.next = newNode
	return newNode
}

// 在环形链表的指定位置插入新节点
func insertAfter(tail *Node, index int, data int) {
	curr := tail.next
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	newNode := &Node{data: data, next: curr.next}
	curr.next = newNode
}

// 删除环形链表中指定位置的节点
func deleteNode(tail *Node, index int) {
	curr := tail.next
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	curr.next = curr.next.next
}

// 修改环形链表中指定位置的节点的值
func updateNode(tail *Node, index int, data int) {
	curr := tail.next
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	curr.data = data
}

// 查询环形链表中指定位置的节点的值
func queryNode(tail *Node, index int) int {
	curr := tail.next
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.data
}

// 打印单向环形链表
func displayList(tail *Node) {
	curr := tail.next
	for curr != tail {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func main() {
	// 创建环形链表
	tail := &Node{data: 0, next: nil}
	tail.next = tail

	// 在环形链表的末尾添加新节点
	tail = appendNode(tail, 1)
	tail = appendNode(tail, 2)
	tail = appendNode(tail, 3)

	//displayList(tail) // 输出 0 -> 1 -> 2

	// 在环形链表的第 2 个位置插入新节点
	insertAfter(tail, 1, 4)

	// 删除环形链表的第 3 个位置的节点
	deleteNode(tail, 2)

	// 修改环形链表的第 2 个位置的节点的值
	updateNode(tail, 1, 5)

	// 查询环形链表的第 2 个位置的节点的值
	fmt.Println(queryNode(tail, 1)) // 输出 5

	displayList(tail) // 输出 0 -> 5 -> 4

}
