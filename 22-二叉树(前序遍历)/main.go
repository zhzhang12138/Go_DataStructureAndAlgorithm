package main

import "fmt"

/*
二叉树
	二叉树（Binary tree）是树形结构的一个重要类型。

	二叉树是n个有限元素的集合，该集合或者为空、或者由一个称为根（root）的元素及两个不相交的、
	被分别称为左子树和右子树的二叉树组成，是有序树。
	当集合为空时，称该二叉树为空二叉树。在二叉树中，一个元素也称作一个节点。

前序遍历
	前序遍历是指先访问根节点，然后遍历左子树，最后遍历右子树。
中序遍历
	中序遍历是指先遍历左子树，然后访问根节点，最后遍历右子树。
后序遍历
	后序遍历是指先遍历左子树，然后遍历右子树，最后访问根节点。
*/

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

// PreOrder 前序遍历
// 前序遍历是指先访问根节点，然后遍历左子树，最后遍历右子树。
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

func main() {
	// 构建一个二叉树
	/*
				 宋江
				 / \
			  吴用   卢俊义
			  / \      \
		   李逵 鲁智深   林冲
	*/
	root := &Hero{
		No:   1,
		Name: "宋江",
	}
	left1 := &Hero{
		No:   2,
		Name: "吴用",
	}
	left2 := &Hero{
		No:   10,
		Name: "李逵",
	}
	left3 := &Hero{
		No:   12,
		Name: "鲁智深",
	}
	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}
	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}

	root.Left = left1
	left1.Left = left2
	left1.Right = left3
	root.Right = right1
	right1.Right = right2

	PreOrder(root)

}
