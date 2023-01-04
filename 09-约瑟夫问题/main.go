package main

import "fmt"

/*
约瑟夫问题
	是一个出现在计算机科学和数学中的问题，在计算机编程算法中类似问题又称为约瑟夫环或者叫做“丢手帕问题”。
原理
	n个人围成一圈，选择一个人开始报数，每过m次则当前报数的人则会被踢出圈，然后挨着的下一个人作为又一轮游戏的第一个报数的人，这样重复最后只会剩下一个人。
例如:
	n = 6, m = 5,陆续被踢出圈的人的编号为：5,4,6,2,3,1。
*/

// Boy 约瑟夫问题
// 小孩
type Boy struct {
	no   int  // 小孩编号
	next *Boy // 指向下一个小孩的指针[默认值是nil]
}

// AddBoy 编写一个函数构成单向的环形链表
// num表示约瑟夫中小孩的个数，
// 返回环形链表的第一个小孩，也就是头结点
func AddBoy(num int) *Boy {
	head := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		return head
	}
	// 循环构建环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			no:   i,
			next: nil,
		}

		// 分析构成循环链表，需要一个辅助指针
		// 1、因为第一个小孩比较特殊
		if i == 1 {
			head = boy
			curBoy = boy
			curBoy.next = head
		} else {
			curBoy.next = boy
			curBoy = boy
			curBoy.next = head // 构造环形链表
		}
	}
	return head
}

// 打印
func showBoy(head *Boy) {
	if head.next == nil {
		fmt.Println("链表为空")
		return
	}

	temp := head
	for {
		fmt.Println("boy: ", temp.no)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

// 求出最后一个小孩，first为编号为1的小孩，startNo为游戏开始的小孩的编号，countNum为游戏规则的间隔
func playGame(first *Boy, startNo, countNum int) {
	if first.next == nil {
		fmt.Println("就一个小孩玩游戏，玩不了")
		return
	}
	// 找到first的前一个节点，由于为环形链表最后一个节点的next就是first头结点，用于判断整个环形链表最后只剩下一个小孩了
	tail := first
	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}
	// 找到游戏开始的小孩
	startNode := first
	for {
		if startNode.no == startNo {
			break
		}
		startNode = startNode.next
		// tail也要跟着移动
		tail = tail.next
	}
	for {
		// 一轮游戏
		for i := 1; i <= countNum; i++ {
			startNode = startNode.next
			tail = tail.next
		}
		fmt.Println("退出的小孩：", startNode.no)
		if startNode == tail {
			fmt.Println("最后剩下的小孩编号：", startNode.no)
			break
		}
		startNode = startNode.next
		tail.next = startNode
	}
}

func main() {
	// 构造五个节点的环形链表
	head := AddBoy(5)
	// 打印
	showBoy(head)
	// 游戏
	playGame(head, 2, 2)
}
