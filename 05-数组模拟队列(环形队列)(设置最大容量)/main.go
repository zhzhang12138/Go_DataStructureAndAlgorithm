package main

import "fmt"

/*
队列介绍
	队列是一个有序列表，可以用数组或是链表来实现
	遵循先入先出的原则，即：先存入队列的数据，要先取出，后存入的要后取出

环形队列
	是一种特殊的队列，其特点在于可以将队列的尾部与队列的头部相连，形成一个环。
	在设置了最大容量的环形队列中，实际可以存储的元素数量是最大容量减 1
	因为，在环形队列中，假如最大容量为5，队列的尾部元素的下一个位置是队列的头部元素。当队列的尾部元素为 4 时，队列的尾部元素的下一个位置是队列的头部元素0。因此，在最大容量为 5 的情况下，队列只能存储 4 个元素。

环形队列 VS 一般队列
	可以避免出队操作中的数组拷贝操作。在一般队列中，如果想要将队列的头部元素出队，就必须将剩余元素向前移动，这样就需要执行数组拷贝操作。而在环形队列中，可以通过改变指针的位置来避免这个问题。
	可以更高效的利用存储空间。在一般队列中，如果队列满了，就必须再开一个更大的数组来存储元素，然后将原来的元素拷贝到新的数组中。而在环形队列中，可以通过将队列的尾部与头部相连的方式来利用存储空间，使得环形队列可以更高效的利用存储空间。
*/

type CircularQueue struct {
	items   [5]int // 队列
	head    int    // 队列的头
	tail    int    // 队列的尾
	maxsize int    // 队列的最大容量
	size    int    // 队列的大小
}

// Init 初始化环形队列
func (q *CircularQueue) Init(maxsize int) {
	q.maxsize = maxsize
	q.head = 0
	q.tail = 0
	q.size = 0
}

// Enqueue 入队
func (q *CircularQueue) Enqueue(item int) bool {
	if (q.tail+1)%q.maxsize == q.head {
		return false
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
	return (q.tail+1)%q.maxsize == q.head
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
func (q *CircularQueue) Head() (int, bool) {
	if q.head == q.tail {
		return 0, false
	}
	return q.items[q.head], true
}

// Tail 返回队列的尾部元素
func (q *CircularQueue) Tail() (int, bool) {
	if q.head == q.tail {
		return 0, false
	}
	return q.items[(q.tail-1+q.maxsize)%q.maxsize], true
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
	q.Enqueue(5) // 返回 false，因为队列已满

	// 出队
	fmt.Println(q.Dequeue()) // 输出: 1
	fmt.Println(q.Dequeue()) // 输出: 2
	fmt.Println(q.Dequeue()) // 输出: 3

	// 入队
	q.Enqueue(5)
	q.Enqueue(6)

	// 显示队列中的元素
	q.Show()

	// 返回队列的大小、容量、头部元素和尾部元素
	fmt.Println("------------------")
	fmt.Println(q.Size()) // 输出: 3
	fmt.Println(q.Cap())  // 输出: 5
	fmt.Println(q.Head()) // 输出: 4
	fmt.Println(q.Tail()) // 输出: 6
	fmt.Println("------------------")

	// 出队
	fmt.Println(q.Dequeue()) // 输出: 4
	fmt.Println(q.Dequeue()) // 输出: 5
	fmt.Println(q.Dequeue()) // 输出: 6
	// 队列为空
	fmt.Println(q.IsEmpty()) // 输出: true
	fmt.Println(q.Size())    // 输出: 0
}
