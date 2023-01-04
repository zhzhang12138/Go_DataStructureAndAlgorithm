package main

import "fmt"

/*
队列介绍
	队列是一个有序列表，可以用数组或是链表来实现
	遵循先入先出的原则，即：先存入队列的数据，要先取出，后存入的要后取出
*/

type Queue struct {
	items   []int // 队列
	head    int   // 队列的头
	tail    int   // 队列的尾
	maxsize int   // 队列的最大容量
	size    int   // 队列的大小
}

// Init 初始化队列
func (q *Queue) Init(maxsize int) {
	q.items = make([]int, maxsize)
	q.head = 0
	q.tail = 0
	q.maxsize = maxsize
	q.size = 0
}

// Enqueue 入队
func (q *Queue) Enqueue(item int) bool {
	if q.size == q.maxsize {
		return false
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.maxsize
	q.size++
	return true
}

// Dequeue 出队
func (q *Queue) Dequeue() (int, bool) {
	if q.size == 0 {
		return 0, false
	}
	item := q.items[q.head]
	q.head = (q.head + 1) % q.maxsize
	q.size--
	return item, true
}

// IsFull 返回队列是否已满
func (q *Queue) IsFull() bool {
	return q.size == q.maxsize
}

// IsEmpty 返回队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Size 返回队列的大小
func (q *Queue) Size() int {
	return q.size
}

func main() {
	// 初始化队列，设置容量为 10
	q := Queue{}
	q.Init(10)

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
