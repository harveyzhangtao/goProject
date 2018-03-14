package tree

import "fmt"


func Create( value int) *Node  {
	return &Node{Value:value}
}
//为结构体定义方法
func (node *Node) SetValue ( value int)  {
	node.Value = value
}

func (node *Node) Stra()  {
	if node == nil{
		return
	}
	node.Left.Stra()
	node.Print()
	node.Right.Stra()
}

func (node Node) Print()  {
	fmt.Print(node.Value)
}

type Node struct {
	Value int
	Left, Right *Node
}
