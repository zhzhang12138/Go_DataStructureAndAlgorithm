package main

import "fmt"

/*
链表
	链表是有序的列表。
	一般来说，为了比较好的对单链表进行增删改查的操作，我们都会给它设置一个头结点。头结点的作用主要是用来标识链表头，本身这个结点不存放数据。

*/

// Node 定义一个节点类型，包含数据域和指向下一个节点的指针域。
type Node struct {
	data int
	next *Node // 表示指向下一个结点
}

// 首先，它定义了一个变量 curr，初始值为单链表的头结点的指针。
// 然后，使用循环遍历单链表，找到最后一个节点。
// 最后，将最后一个节点的 next 指针指向新创建的节点，完成插入操作。
func appendNode(head *Node, data int) {
	curr := head
	for curr.next != nil {
		curr = curr.next // 让curr不断的指向下一个结点
	}
	curr.next = &Node{data: data}
}

// 首先，它定义了一个变量 curr，初始值为单链表的头结点的指针。
// 然后，使用循环遍历单链表，找到第 index 个节点。
// 最后，将第 index 个节点的 next 指针指向新创建的节点，并将新节点的 next 指针指向第 index+1 个节点，完成插入操作。
func insertAfter(head *Node, index int, data int) {
	curr := head
	for i := 1; i < index; i++ {
		curr = curr.next
	}
	node := &Node{data: data, next: curr.next}
	curr.next = node
}

// 首先，它定义了一个变量 curr，初始值为单链表的头结点的指针。
// 然后，使用循环遍历单链表，找到第 index-1 个节点。
// 最后，将第 index-1 个节点的 next 指针指向第 index+1 个节点，完成删除操作。
func deleteNode(head *Node, index int) {
	curr := head
	for i := 1; i < index-1; i++ {
		curr = curr.next
	}
	curr.next = curr.next.next
}

// 首先，它定义了一个变量 curr，初始值为单链表的头结点的指针。
// 然后，使用循环遍历单链表，找到第 index 个节点。
// 最后，修改第 index 个节点的 data 域的值。
func updateNode(head *Node, index int, data int) {
	curr := head
	for i := 1; i < index; i++ {
		curr = curr.next
	}
	curr.data = data
}

// 首先，它定义了一个变量 curr，初始值为单链表的头结点的指针。
// 然后，使用循环遍历单链表，找到第 index 个节点。最后，返回该节点的指针。
func queryNode(head *Node, index int) *Node {
	curr := head
	for i := 1; i < index; i++ {
		curr = curr.next
	}
	return curr
}

// 首先，它定义了一个变量 curr，初始值为单链表的头结点的指针。
// 然后，使用循环遍历单链表，在每个节点处打印出其 data 域的值。
func displayList(head *Node) {
	curr := head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func main() {
	// 先创建一个头结点
	var head *Node
	head = &Node{data: 0}

	// 在单链表的末尾插入新节点
	appendNode(head, 1)
	appendNode(head, 2)
	appendNode(head, 3)

	// 在单链表的第 2 个节点之后插入新节点
	insertAfter(head, 2, 999)

	// 显示链表的所以结点信息
	displayList(head)

	// 删除第 3 个节点
	deleteNode(head, 3)

	// 修改第 2 个节点的数据
	updateNode(head, 2, 5)

	// 查询第 2 个节点的数据
	node := queryNode(head, 2)
	fmt.Println(node.data)
}
