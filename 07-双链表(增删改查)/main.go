package main

import "fmt"

/*
双链表
	是一种特殊的链表数据结构，它的每个节点都有两个指针：一个指向前驱节点，另一个指向后继节点。
	这样，双链表就可以从前向后遍历，也可以从后向前遍历。

双链表通常用来解决单链表不能解决的问题，例如：
	(1)在单链表中，如果要在一个节点之前插入新节点，就需要遍历到该节点的前一个节点，然后修改指针。而在双链表中，可以直接通过前驱节点的指针插入新节点。
	(2)在单链表中，如果要删除一个节点，就需要遍历到该节点的前一个节点，然后修改指针。而在双链表中，可以直接通过前驱节点和后继节点的指针删除节点。
*/

type Node struct {
	data int   // 存储节点的数据
	prev *Node // 指向前驱节点的指针
	next *Node // 指向后继节点的指针
}

// 在双链表的末尾插入新节点
// 该函数首先遍历到双链表的最后一个节点，
// 然后将其 next 指针指向新创建的节点，并将新节点的 prev 指针指向最后一个节点。
func appendNode(head *Node, data int) {
	curr := head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &Node{data: data, prev: curr}
}

// 在双链表的第 index 个节点之后插入新节点
// 该函数首先遍历到双链表的第 index 个节点，
// 然后创建一个新节点，并将新节点的 prev 指针指向第 index 个节点，将新节点的 next 指针指向第 index+1 个节点。
// 然后，将第 index+1 个节点的 prev 指针指向新节点，将第 index 个节点的 next 指针指向新节点。

// 如果 index 小于 1，就直接退出函数，不执行任何操作。
// 如果在遍历双链表时，curr 已经指向了 nil，也就是遍历到了双链表的末尾，那么就直接退出函数，不执行任何操作。
// 如果插入的是双链表的最后一个节点之后，curr.next 会指向 nil。
// 在这种情况下，我们不能通过 curr.next.prev 来修改新节点的前驱节点的指针，否则会出现空指针异常。
// 因此，我们需要先判断 curr.next 是否为 nil，然后再修改新节点的前驱节点的指针。
func insertAfter(head *Node, index int, data int) {
	if index < 1 {
		return
	}
	curr := head
	for i := 1; i < index; i++ {
		if curr == nil {
			return
		}
		curr = curr.next
	}
	if curr == nil {
		return
	}
	node := &Node{data: data, prev: curr, next: curr.next}
	if curr.next != nil {
		curr.next.prev = node
	}
	curr.next = node
}

// 该函数首先遍历到双链表的第 index 个节点，
// 然后将其前驱节点的 next 指针指向其后继节点，将其后继节点的 prev 指针指向其前驱节点，完成删除操作。

// 如果在遍历双链表时，curr 已经指向了 nil，也就是遍历到了双链表的末尾，那么就直接退出函数，不执行任何操作。
// 如果删除的是双链表的第一个节点，curr.prev 会指向 nil。在这种情况下，我们不能通过 curr.prev.next 来修改后继节点的指针，否则会出现空指针异常。因此，我们需要先判断 curr.prev 是否为 nil，然后再修改后继节点的指针。
// 如果删除的是双链表的最后一个节点，curr.next 会指向 nil。同样的，我们也需要先判断 curr.next 是否为 nil，然后再修改前驱节点的指针。
func deleteNode(head *Node, index int) {
	if index < 1 {
		return
	}
	curr := head
	for i := 1; i < index; i++ {
		if curr == nil {
			return
		}
		curr = curr.next
	}
	if curr == nil {
		return
	}
	if curr.prev != nil {
		curr.prev.next = curr.next
	}
	if curr.next != nil {
		curr.next.prev = curr.prev
	}
}

//该函数首先遍历到双链表的第 index 个节点，然后修改其 data 域的值。
func updateNode(head *Node, index int, data int) {
	curr := head
	for i := 1; i < index; i++ {
		curr = curr.next
	}
	curr.data = data
}

// 该函数首先遍历到双链表的第 index 个节点，然后返回该节点的指针。
func queryNode(head *Node, index int) *Node {
	curr := head
	for i := 1; i < index; i++ {
		curr = curr.next
	}
	return curr
}

// 顺序打印
// 该函数首先定义了一个变量 curr，初始值为双链表的头结点的指针。
// 然后，使用循环遍历双链表，在每个节点处打印出其 data 域的值。
func displayList(head *Node) {
	curr := head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

// 逆序打印
// 该函数首先定义了一个变量 curr，初始值为双链表的最后一个节点的指针。
// 然后，使用循环遍历双链表，在每个节点处打印出其 data 域的值。
func reverseList(head *Node) {
	curr := head
	for curr.next != nil {
		curr = curr.next
	}
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.prev
	}
}

func main() {
	// 1、创建一个双链表，并将其头结点的指针赋值给一个变量 head
	head := &Node{data: 1}

	// 2、使用 appendNode 函数在双链表的末尾插入新节点
	appendNode(head, 2)
	appendNode(head, 3)
	appendNode(head, 4)

	//displayList(head)  // 输出 1 -> 2 -> 3 -> 4

	// 3、使用 insertAfter 函数在双链表的第 2 个节点之后插入新节点
	insertAfter(head, 2, 5)

	//displayList(head) // 输出 1 -> 2 -> 5 -> 3 -> 4

	// 4、使用 updateNode 函数修改双链表的第 3 个节点的数据
	updateNode(head, 3, 6)

	//displayList(head) // 输出 1 -> 2 -> 6 -> 3 -> 4

	// 5、使用 deleteNode 函数删除双链表的第 4 个节点
	deleteNode(head, 4)

	//displayList(head) // 输出 1 -> 2 -> 6 -> 4

	// 6、使用 queryNode 函数查询双链表的第 3 个节点的数据
	node := queryNode(head, 3)
	fmt.Println(node.data) // 输出 6

	// 7、使用 displayList 函数遍历双链表并显示其所有节点的信息
	displayList(head) // 输出 1 -> 2 -> 6 -> 4

	// 8、使用 reverseList 函数逆序遍历双链表并显示其所有节点的信息
	reverseList(head) // 输出 4 -> 6 -> 2 -> 1
}
