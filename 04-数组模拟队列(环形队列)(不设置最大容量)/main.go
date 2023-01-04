package main

import "fmt"

/*
队列介绍
	队列是一个有序列表，可以用数组或是链表来实现
	遵循先入先出的原则，即：先存入队列的数据，要先取出，后存入的要后取出

环形队列
	是一种特殊的队列，其特点在于可以将队列的尾部与队列的头部相连，形成一个环。

环形队列 VS 一般队列
	可以避免出队操作中的数组拷贝操作。在一般队列中，如果想要将队列的头部元素出队，就必须将剩余元素向前移动，这样就需要执行数组拷贝操作。而在环形队列中，可以通过改变指针的位置来避免这个问题。
	可以更高效的利用存储空间。在一般队列中，如果队列满了，就必须再开一个更大的数组来存储元素，然后将原来的元素拷贝到新的数组中。而在环形队列中，可以通过将队列的尾部与头部相连的方式来利用存储空间，使得环形队列可以更高效的利用存储空间。
*/

type CircularQueue struct {
	items   []int // 队列
	head    int   // 队列的头
	tail    int   // 队列的尾
	maxsize int   // 队列的最大容量
	size    int   // 队列的大小
}

// Init 初始化环形队列
func (q *CircularQueue) Init(maxsize int) {
	q.maxsize = maxsize
	q.head = 0
	q.tail = 0
	q.size = 0
	q.items = make([]int, maxsize)
}

// Enqueue 入队
func (q *CircularQueue) Enqueue(item int) bool {
	if q.size == q.maxsize {
		// 扩展切片容量
		newItems := make([]int, q.maxsize*2)
		// 复制原切片元素
		copy(newItems, q.items[q.head:])
		copy(newItems[q.maxsize-q.head:], q.items[:q.head])
		// 更新切片和 head、tail
		q.items = newItems
		q.head = 0
		q.tail = q.maxsize
		q.maxsize *= 2
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.maxsize
	q.size++
	return true
}

// Dequeue 出队
func (q *CircularQueue) Dequeue() (int, bool) {
	if q.head == q.tail {
		return 0, false
	}
	item := q.items[q.head]
	q.head = (q.head + 1) % q.maxsize
	q.size--
	return item, true
}

// IsFull 返回队列是否已满
func (q *CircularQueue) IsFull() bool {
	return q.size == q.maxsize
}

// IsEmpty 返回队列是否为空
func (q *CircularQueue) IsEmpty() bool {
	return q.head == q.tail
}

// Size 返回队列的大小
func (q *CircularQueue) Size() int {
	return q.size
}

// Cap 返回队列的容量
func (q *CircularQueue) Cap() int {
	return q.maxsize
}

// Head 返回队列的头部元素
func (q *CircularQueue) Head() int {
	if q.head == q.tail {
		return 0
	}
	return q.items[q.head]
}

// Tail 返回队列的尾部元素
func (q *CircularQueue) Tail() int {
	if q.head == q.tail {
		return 0
	}
	return q.items[(q.tail-1+q.maxsize)%q.maxsize]
}

// Show 显示队列中的元素
func (q *CircularQueue) Show() {
	fmt.Print("queue:")
	for i := q.head; i != q.tail; i = (i + 1) % q.maxsize {
		fmt.Printf("%d ", q.items[i])
	}
	fmt.Println()
}

func main() {
	// 初始化环形队列，设置容量为 5
	q := CircularQueue{}
	q.Init(5)

	// 入队
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6) // 扩展切片容量，入队成功

	// 出队
	fmt.Println(q.Dequeue()) // 输出: 1
	fmt.Println(q.Dequeue()) // 输出: 2
	fmt.Println(q.Dequeue()) // 输出: 3

	// 入队
	q.Enqueue(7)
	q.Enqueue(8)

	// 返回队列的大小、容量、头部元素和尾部元素
	fmt.Println(q.Size()) // 输出: 4
	fmt.Println(q.Cap())  // 输出: 10
	fmt.Println(q.Head()) // 输出: 4
	fmt.Println(q.Tail()) // 输出: 8

	// 出队
	fmt.Println(q.Dequeue()) // 输出: 4
	fmt.Println(q.Dequeue()) // 输出: 5
	fmt.Println(q.Dequeue()) // 输出: 6
	fmt.Println(q.Dequeue()) // 输出: 7
	fmt.Println(q.Dequeue()) // 输出: 8
	// 队列为空
	fmt.Println(q.IsEmpty()) // 输出: true
	fmt.Println(q.Size())    // 输出: 0
}
