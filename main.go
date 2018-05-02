package main

import (
	"fmt"
	"github.com/shsjshentao/Exercise/Tree"
)


func main() {
	avlTree := new(Tree.AVLTree)
	avlTree.Insert([]int{65,11,91,6,50,9,14,44,27}...)
	fmt.Println(avlTree.FindMin())
	fmt.Println(avlTree.Count())
	avlTree.Remove(14)
	fmt.Println(avlTree.Find(10))
	fmt.Println(avlTree.Count())
	fmt.Println(avlTree)
}