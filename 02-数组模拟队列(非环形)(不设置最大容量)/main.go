package main

import "fmt"

/*
队列介绍
	队列是一个有序列表，可以用数组或是链表来实现
	遵循先入先出的原则，即：先存入队列的数据，要先取出，后存入的要后取出
*/

type Queue struct {
	items []int
}

// Init 初始化队列
func (q *Queue) Init() {
	q.items = []int{}
}

// Enqueue 入队
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue 出队
func (q *Queue) Dequeue() int {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// IsEmpty 返回队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size 返回队列的大小
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	// 初始化队列
	q := Queue{}
	q.Init()

	// 入队
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// 出队
	fmt.Println(q.Dequeue()) // 输出: 1
	fmt.Println(q.Dequeue()) // 输出: 2
	fmt.Println(q.Dequeue()) // 输出: 3

	// 队列为空
	fmt.Println(q.IsEmpty()) // 输出: true
	fmt.Println(q.Size())    // 输出: 0
}
