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

	//var root tree.TreeNode
	//root = treeNode{value:3}
	//root.left = &treeNode{}
	//root.right = &treeNode{5, nil, nil}
	//root.right.left = new(treeNode)
	//root.left.right = create(2)
	//root.left.setValue(4)
	////root.left.print()
	////fmt.Println(root)
	//root.stra()
}

func (node TreeNode) Print()  {
	fmt.Print(node.Value)
}

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}
