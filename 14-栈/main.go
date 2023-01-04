package main

import "fmt"

/*
栈（Stack）又叫堆栈，
	是限定只能在一端进行插入和删除操作的线性表，并且满足后进先出（LIFO）的特点，
	即最后插入的最先被读取。我们把允许插入和删除的一端叫做栈顶，另一个端为固定的一端，叫做栈底，不含任何数据的栈叫做空栈。
*/

// Stack 使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int    // 表示我们栈最大可以存放的个数
	Top    int    // 表示栈顶，因为栈顶固定，因此我们直接使用Top
	arr    [5]int // 数组模拟栈
}

// 初始化一个栈
func (s *Stack) InitStack() {
	s.Top = -1
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	if s.Top == -1 {
		return true
	}
	return false
}

// 判断栈是否满了
func (s *Stack) IsFull() bool {
	if s.Top == s.MaxTop-1 {
		return true
	}
	return false
}

// 入栈
func (s *Stack) Push(val int) {
	if s.IsFull() {
		fmt.Println("stack full")
		return
	}
	s.Top++
	s.arr[s.Top] = val
}

// 出栈
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("stack empty")
		return 0
	}
	val := s.arr[s.Top]
	s.Top--
	return val
}

// 遍历栈
func (s *Stack) List() {
	if s.IsEmpty() {
		fmt.Println("stack empty")
		return
	}
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, s.arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
		arr:    [5]int{},
	}
	stack.InitStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.List()
	fmt.Println("出栈")
	stack.Pop()
	stack.Pop()
	stack.List()
}
