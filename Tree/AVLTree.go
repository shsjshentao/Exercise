package Tree

import "fmt"

/*
Usage:
avlTree := new(Tree.AVLTree)
avlTree.Insert([]int{65,11,91,6,50,9,14,44,27}...)
fmt.Println(avlTree.FindMin())
fmt.Println(avlTree.Count())
avlTree.Remove(14)
fmt.Println(avlTree.Find(10))
fmt.Println(avlTree.Count())
fmt.Println(avlTree)
 */
type AVLTree struct {
	node *TreeNode
}

type TreeNode struct{
	value int
	lChild *TreeNode
	rChild *TreeNode
}


func (avl *AVLTree) Insert(value ... int) {
	for _, v := range value {
		if avl.node == nil {
			avl.node = new(TreeNode)
			avl.node.value = v
		}
		node := avl.node
		for node != nil {
			if v < node.value {
				if node.lChild == nil {
					node.lChild = new(TreeNode)
					node.lChild.value = v
					break
				} else {
					node = node.lChild
				}
			} else if v > node.value {
				if node.rChild == nil {
					node.rChild = new(TreeNode)
					node.rChild.value = v
					break
				} else {
					node = node.rChild
				}
			} else {
				break
			}
		}
	}
}

func (avl *AVLTree) Find(value int) bool{
	node := avl.node
	for {
		if node != nil {
			if value < node.value {
				node = node.lChild
			} else if value > node.value {
				node = node.rChild
			} else {
				return true
			}
		} else {
			return false
		}
	}
}

func (avl *AVLTree) FindMin() int{
	if avl.node == nil {
		return 0
	}
	node := avl.node
	for node.lChild != nil {
		node = node.lChild
	}
	return node.value
}

func (avl *AVLTree) findMax() int{
	if avl.node == nil {
		return 0
	}
	node := avl.node
	for node.rChild != nil {
		node = node.rChild
	}
	return node.value
}

func (avl *AVLTree) Remove(value int) bool {
	node := avl.node
	var parentNode *TreeNode
	for {
		if node != nil {
			if value < node.value {
				parentNode = node
				node = node.lChild
			} else if value > node.value {
				parentNode = node
				node = node.rChild
			} else {
				break
			}
		} else {
			return false
		}
	}
	if node.rChild !=nil {
		lNode := node.rChild
		var rParentNode *TreeNode
		for lNode.lChild != nil {
			rParentNode = lNode
			lNode = lNode.lChild
		}
		node.value = lNode.value
		rParentNode.lChild = nil
	} else if node.lChild != nil && node.rChild == nil {
		node = node.lChild
	} else {
		if parentNode.lChild == node {
			parentNode.lChild =nil
		} else {
			parentNode.rChild =nil
		}
	}
	return true
}

func(avl *AVLTree) Count() int {
	node := avl.node
	var count func(node *TreeNode) int
	count = func(node *TreeNode) int {
		num := 0
		if node == nil {
			return 0
		} else {
			num = 1 + count(node.lChild) + count(node.rChild)
		}
		return num
	}
	return count(node)
}

func (avl *AVLTree) String() string{
	node := avl.node
	var list string
	var output func(node *TreeNode) string
	output = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		list += fmt.Sprint(node.value, " ")
		output(node.lChild)
		output(node.rChild)
		return list
	}
	return output(node)
}