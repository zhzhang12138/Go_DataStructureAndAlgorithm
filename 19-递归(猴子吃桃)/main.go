package main

import "fmt"

/*
猴子吃桃问题
	有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个，以后每天猴子都吃其中的一半，然后再多吃一个，当到第十天时，想再吃时(还没吃)，
	发现只有1个桃子了。问题：最初共有多少个桃子？
思路分析：
	(1)第十天只有一个桃子
	(2)第九天有几个桃子 = (第十天的桃子数量+1) * 2 = 4个
	(3)规律：第n天的桃子数量 peach(n) = (peach(n+1)+1) *2
*/

// n 范围是1-10之间
func expression(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数不对")
		return 0 // 返回0表示不对
	}
	if n == 10 {
		return 1
	} else {
		return (expression(n+1) + 1) * 2
	}
}

func main() {
	res := expression(1)
	fmt.Print(res)
}
