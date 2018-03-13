package main

import "goProject/tree"

func main() {

	var root tree.TreeNode
	root = TreeNode{value:3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = create(2)
	root.left.setValue(4)
	//root.left.print()
	//fmt.Println(root)
	root.stra()
}
