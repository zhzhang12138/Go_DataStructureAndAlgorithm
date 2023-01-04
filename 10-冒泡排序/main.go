package main

import "fmt"

/*
冒泡排序
	它重复地走访过要排序的元素列，依次比较两个相邻的元素，如果顺序（如从大到小、首字母从Z到A）错误就把他们交换过来。
	走访元素的工作是重复地进行，直到没有相邻元素需要交换，也就是说该元素列已经排序完成。
*/

func bubbleSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 冒泡排序核心实现代码
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	return nums
}

func main() {
	nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	nums = bubbleSort(nums)
	fmt.Println(nums)
}
