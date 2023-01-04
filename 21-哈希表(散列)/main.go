package main

import (
	"fmt"
)

/*
哈希表(散列)
	散列表（Hash table，也叫哈希表），是根据关键码值(Key value)而直接进行访问的数据结构。
	也就是说，它通过把关键码值映射到表中一个位置来访问记录，以加快查找的速度。
	这个映射函数叫做散列函数，存放记录的数组叫做散列表。

实际需求：
	google公司的一个上机题：
		使用hashtable来实现一个雇员的管理系统[增删改查]
		有一个公司，当有新的员工来报道时，要求该员工的信息加入(id,性别,年龄,住址...)，当输入该员工的id时，要求查找到该员工的所有信息
	要求：
		(1)不使用数据库，尽量节省内存，速度越快越好--->哈希表(散列)
		(2)添加时，保证按照id从低到高插入
	思路分析：
		(1)使用链表来实现哈希表，该链表不带表头[即：链表的第一个结点就存放雇员信息]
		(2)思路分析并画出示意图
		(3)代码实现[增删改查(显示所有员工，按id查询)]
*/

// Emp 定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 方法待定...
func (this *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员 %d \n", this.Id%7, this.Id)
}

// EmpLink 定义EmpLink
// 这里的EmpLink不带表头，即第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

// Insert 1、添加员工的方法，保证添加时，编号从小到大
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head   // 这时辅助指针
	var pre *Emp = nil // 这也是一个辅助指针 pre在cur前面
	// 如果当前的EmpLink就是一个空链表
	if cur == nil {
		this.Head = emp // 完成
		return
	}
	// 如果不是一个空链表，给emp找到对应的位置并插入
	// 思路是：让cur和emp比较，然后让pre保持在cur前面
	for {
		if cur != nil {
			// 比较
			if cur.Id > emp.Id {
				// 找到位置
				break
			}
			pre = cur // 保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	// 退出时，我们看一下是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

// ShowLink 显示当前链表的信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}
	// 遍历当前的链表并显示数据
	cur := this.Head // 辅助的指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 雇员姓名=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println() // 换行处理
}

// 根据id查找对应的雇员，如果没有，就返回nil
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

// HashTable 定义hashtable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// Insert 给HashTable编写Insert雇员的方法
func (this *HashTable) Insert(emp *Emp) {
	// 使用散列函数，确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	// 使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}

// ShowAll 编写方法 显示hashtable所有的雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

// HashFun 编写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7 // 得到一个值，就是对应的链表的下标
}

// FindById 编写一个方法，完成查找
func (this *HashTable) FindById(id int) *Emp {
	// 使用散列函数，确定将该雇员应该在哪个链表
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable

	for {
		fmt.Println("==============雇员系统菜单==============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示显示雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("exit  表示退出雇员")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员姓名")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入id号")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在 \n", id)
			} else {
				// 编写一个方法，显示雇员信息
				emp.ShowMe()
			}
		case "exit":
			break
		default:
			fmt.Println("输入有误！！！")

		}

	}
}
