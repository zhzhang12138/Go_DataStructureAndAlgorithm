package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
选择排序（Selection sort）是一种简单直观的排序算法。
工作原理：第一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，
		然后再从剩余的未排序元素中寻找到最小（大）元素，然后放到已排序的序列的末尾。
		以此类推，直到全部待排序的数据元素的个数为零。
		选择排序是不稳定的排序方法。
*/

func selectionSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 已排序区间初始化为空，未排序区间初始化待排序切片
	for i := 0; i < len(nums); i++ {
		// 未排序区间最小值初始化为第一个元素
		min := i
		// 从未排序区间第二个元素开始遍历，直到找到最小值
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		// 将最小值与未排序区间第一个元素互换位置（等价于放到已排序区间最后一个位置）
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
}

func main() {
	//nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	//selectionSort(nums)
	//fmt.Println(nums)

	var arr []int
	for i := 0; i < 80000; i++ {
		arr = append(arr, rand.Intn(900000))
	}
	start := time.Now().Unix()
	selectionSort(arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时=%d秒", end-start)
}
