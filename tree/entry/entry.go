package main

import (
	"goProject/tree"
	"fmt"
	"goProject/tree/queue"
)

func main() {

	//var root tree.Node
	//root = tree.Node{Value:3}
	//root.Left = &tree.Node{}
	//root.Right = &tree.Node{5, nil, nil}
	//root.Right.Left = new(tree.Node)
	//root.Left.Right = tree.Create(2)
	//root.Left.SetValue(4)
	//root.Left.Print()
	////fmt.Println(root)
	//root.Stra()
	//fmt.Println("")
	//myNode := MyTreeNode{&root}
	//myNode.postOrder()

	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	q.Pop()
	fmt.Println(q.IsEmpty())
	q.Pop()
	q.Pop()

	fmt.Println(q.IsEmpty())

}

type MyTreeNode struct {
	node *tree.Node
}

func (myNode * MyTreeNode) postOrder()  {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.Left}
	left.postOrder()
	right := MyTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()

}
