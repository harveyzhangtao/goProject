package tree

import "fmt"


func create( value int) *TreeNode  {
	return &TreeNode{Value:value}
}
//为结构体定义方法
func (node *TreeNode) SetValue ( value int)  {
	node.Value = value
}

func (node *TreeNode) Stra()  {
	if node == nil{
		return
	}
	node.Left.Stra()
	node.Print()
	node.Right.Stra()
}

func (node TreeNode) Print()  {
	fmt.Print(node.Value)
}

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}
