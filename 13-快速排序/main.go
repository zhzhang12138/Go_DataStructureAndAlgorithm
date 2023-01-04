package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
快速排序
	(1)首先设定一个分界值，通过该分界值将数组分成左右两部分。
	(2)将大于或等于分界值的数据集中到数组右边，小于分界值的数据集中到数组的左边。此时，左边部分中各元素都小于分界值，而右边部分中各元素都大于或等于分界值。
	(3)然后，左边和右边的数据可以独立排序。对于左侧的数组数据，又可以取一个分界值，将该部分数据分成左右两部分，同样在左边放置较小值，右边放置较大值。右侧的数组数据也可以做类似处理。
	(4)重复上述过程，可以看出，这是一个递归定义。通过递归将左侧部分排好序后，再递归排好右侧部分的顺序。当左、右两个部分各数据排序完成后，整个数组的排序也就完成了
*/

// 快速排序入口函数
func quickSort(nums []int, p int, r int) {
	// 递归终止条件
	if p >= r {
		return
	}
	// 获取分区位置
	q := partition(nums, p, r)
	// 递归分区（排序是在定位 pivot 的过程中实现的）
	quickSort(nums, p, q-1)
	quickSort(nums, q+1, r)
}

// 定位 pivot
func partition(nums []int, p int, r int) int {
	// 以当前数据序列最后一个元素作为初始 pivot
	pivot := nums[r]
	// 初始化 i、j 下标
	i := p
	// 后移 j 下标的遍历过程
	for j := p; j < r; j++ {
		// 将比 pivot 小的数丢到 [p...i-1] 中，剩下的 [i...j] 区间都是比 pivot 大的
		if nums[j] < pivot {
			// 互换 i、j 下标对应数据
			nums[i], nums[j] = nums[j], nums[i]
			// 将 i 下标后移一位
			i++
		}
	}

	// 最后将 pivot 与 i 下标对应数据值互换
	// 这样一来，pivot 就位于当前数据序列中间，i 也就是 pivot 值对应的下标
	nums[i], nums[r] = pivot, nums[i]
	// 返回 i 作为 pivot 分区位置
	return i
}

func main() {
	//nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	////nums := []int{-9, 78, 0, 23, -579, 70}
	//quickSort(nums, 0, len(nums)-1)
	//fmt.Println(nums)

	var arr []int
	for i := 0; i < 80000000; i++ {
		arr = append(arr, rand.Intn(9000000))
	}
	start := time.Now().Unix()
	quickSort(arr, 0, len(arr)-1)
	end := time.Now().Unix()
	fmt.Printf("快速排序耗时=%d秒", end-start)
}
