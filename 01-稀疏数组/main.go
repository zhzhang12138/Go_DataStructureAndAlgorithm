package main

import "fmt"

/*
稀疏数组
	当一个数组中大部分元素为0，或者为同一个值的数组时，可以使用稀疏数组来保存该数组

稀疏数组的处理方式是
	(1)记录数组一共有几行几列，有多少个不同的值
	(2)把具有不同值的元素的行列及值记录在一个小规模的数组中，从而缩小程序的规模
*/

// SparseArrayElement 定义一个类型来表示稀疏数组中的元素
type SparseArrayElement struct {
	Row int // 行索引
	Col int // 列索引
	Val int // 元素值
}

// SparseArray 定义一个类型来表示稀疏数组本身
type SparseArray struct {
	Rows     int                  // 行数
	Cols     int                  // 列数
	Elements []SparseArrayElement // 用于存储稀疏数组中非零元素的切片
}

// Get 实现一个方法来获取某个位置的元素
func (a *SparseArray) Get(row, col int) int {
	// 遍历稀疏数组中的所有元素
	for _, e := range a.Elements {
		if e.Row == row && e.Col == col {
			// 如果找到了，就返回元素值
			return e.Val
		}
	}
	// 如果没有找到，就返回零值
	return 0
}

// Set 实现一个方法来设置某个位置的元素
func (a *SparseArray) Set(row, col, val int) {
	// 遍历稀疏数组中的所有元素
	for i, e := range a.Elements {
		if e.Row == row && e.Col == col {
			// 如果找到了，就更新元素值
			a.Elements[i].Val = val
			return
		}
	}
	// 如果没有找到，就新增一个元素
	a.Elements = append(a.Elements, SparseArrayElement{
		Row: row,
		Col: col,
		Val: val,
	})
}

func main() {
	a := &SparseArray{
		Rows:     3,
		Cols:     4,
		Elements: []SparseArrayElement{},
	}

	a.Set(0, 0, 1)
	a.Set(0, 3, 2)
	a.Set(1, 1, 3)
	a.Set(2, 2, 4)

	fmt.Println(a.Get(0, 0)) // 输出 1
	fmt.Println(a.Get(0, 1)) // 输出 0
	fmt.Println(a.Get(1, 1)) // 输出 3
	fmt.Println(a.Get(2, 2)) // 输出 4

	fmt.Println(a)
}
