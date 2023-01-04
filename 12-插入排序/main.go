package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
插入排序，一般也被称为直接插入排序。
	对于少量元素的排序，它是一个有效的算法。
	插入排序是一种最简单的排序方法，它的基本思想是将一个记录插入到已经排好序的有序表中，
	从而一个新的、记录数增1的有序表。在其实现过程使用双层循环，外层循环对除了第一个元素之外的所有元素，
	内层循环对当前元素前面有序表进行待插入位置查找，并进行移动 [2]  。
*/

func insertionSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		// 每次从未排序区间取一个数据 value
		value := nums[i]
		// 在已排序区间找到插入位置
		j := i - 1
		for ; j >= 0; j-- {
			// 如果比 value 大后移
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		// 插入数据 value
		nums[j+1] = value
	}

	return nums
}

func main() {
	//nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	//nums = insertionSort(nums)
	//fmt.Println(nums)

	var arr []int
	for i := 0; i < 80000; i++ {
		arr = append(arr, rand.Intn(900000))
	}
	start := time.Now().Unix()
	insertionSort(arr)
	end := time.Now().Unix()
	fmt.Printf("插入排序耗时=%d秒", end-start)
}
