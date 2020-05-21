package basic

import (
	"fmt"
	"testing"
)

func TestRBTree(t *testing.T) {
	tree := NewRBTree()
	node := NewNode2(8, Black)
	fmt.Println(node)
	tree.Insert(node).
		Insert(NewNode2(2, Red)).
		Insert(NewNode2(16, Black)).
		Insert(NewNode2(54, Red)).
		Insert(NewNode2(464, Red)).
		Insert(NewNode2(2, Red))

	fmt.Println(tree)

	PrintTree(tree)
}

func Test2(t *testing.T) {
	tree := NewRBTree()
	tree.Insert(NewNode2(2, Red)).
		Insert(NewNode2(16, Black)).
		Insert(NewNode2(54, Red)).
		Insert(NewNode2(464, Red)).
		Insert(NewNode2(2, Red))
	PrintTree(tree)

	tree.LeftRotate(tree.root)
	fmt.Println("----------------------------------------------------------------")
	PrintTree(tree)
	tree.RightRotate(tree.root)
	fmt.Println("----------------------------------------------------------------")
	PrintTree(tree)
	tree.RightRotate(tree.root.Right)
	fmt.Println("----------------------------------------------------------------")
	PrintTree(tree)

}
