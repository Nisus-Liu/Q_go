package basic

import (
	"fmt"
)

type NodeColor uint8

const (
	Red   NodeColor = iota
	Black NodeColor = iota
)

// EMPTY 空节点, 叶子节点标识
// 空节点是黑色
type NilNode = Node

var EMPTY = &NilNode{Color: Black}

type Node struct {
	Key    int64
	Color  NodeColor
	Left   *Node
	Right  *Node
	Parent *Node
}

func (n *Node) GetLeft() interface{} {
	return n.Left
}
func (n *Node) GetRight() interface{} {
	return n.Right
}

func NewNode(key int64, color NodeColor, left *Node, right *Node, parent *Node) *Node {
	if &color == nil {
		color = Red
	}
	if left == nil {
		left = EMPTY
	}
	if right == nil {
		right = EMPTY
	}
	if parent == nil {
		parent = EMPTY
	}
	return &Node{Key: key, Color: color, Left: left, Right: right, Parent: parent}
}

func NewNode2(key int64, color NodeColor) *Node {
	return NewNode(key, color, EMPTY, EMPTY, EMPTY)
}

func (t *Node) String() string {
	return fmt.Sprintf("key: %d, color: %d, left.key: %d, right.key: %d",
		t.Key, t.Color, t.Left.Key, t.Right.Key)
	// return fmt.Sprintf("key: %d, color: %d, left: %s, right: %s",
	// 	t.Key, t.Color, t.Left.String(), t.Right.String())
}

// RBTree 红黑树
type RBTree struct {
	root *Node
}

func NewRBTree() *RBTree {
	return &RBTree{root: EMPTY}
}

func (r *RBTree) GetRoot() *Node {
	return r.root
}

func (t *RBTree) Insert(node *Node) *RBTree {
	t.insert0(t.root, node)
	t.insertFixUp(node)
	return t
}

func (t *RBTree) insert0(root *Node, node *Node) *RBTree {
	var nextNode *Node = nil
	if root == EMPTY {
		// 还是个空树, 当前节点作为根节点
		t.root = node
		return t
	} else if node.Key < root.Key {
		if root.Left == EMPTY {
			node.Parent = root
			root.Left = node
			return t
		}
		nextNode = root.Left
	} else if node.Key >= root.Key {
		if root.Right == EMPTY {
			node.Parent = root
			root.Right = node
			return t
		}
		nextNode = root.Right
	}

	return t.insert0(nextNode, node)
}

func (t *RBTree) LeftRotate(x *Node) {
	if x == EMPTY {
		return
	}
	// x - y 构成旋转支轴
	y := x.Right
	x.Right = y.Left
	if y.Left != EMPTY {
		y.Left.Parent = x // 互指
	}

	y.Parent = x.Parent
	// y 是作为左还是右孩子? 还是根? 取决于x之前是什么
	if x.Parent == EMPTY { // x 是根 => y 要作为新根
		t.root = y
	} else if x == x.Parent.Left { // x 是左 => y 左
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

func (t *RBTree) RightRotate(x *Node) {
	if x == EMPTY {
		return
	}
	// x - y 构成旋转支轴 (? 主动点和其左子 构成支轴)
	y := x.Left
	x.Left = y.Right
	if y.Right != EMPTY {
		y.Right.Parent = x // 互指
	}

	y.Parent = x.Parent
	// y 是作为左还是右孩子? 还是根? 取决于x之前是什么
	if x.Parent == EMPTY { // x 是根 => y 要作为新根
		t.root = y
	} else if x == x.Parent.Left { // x 是左 => y 左
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Right = x
	x.Parent = y
}

// 修复因插入破坏的红黑性质
// 复杂, 暂照抄
func (t *RBTree) insertFixUp(z *Node) {
	for z.Parent.Color == Red {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right // z的叔节点
			if y.Color == Red {
				z.Parent.Color = Black
				y.Color = Black
				z.Parent.Parent.Color = Red
				z = z.Parent.Parent
			} else if z == z.Parent.Right { // 先左旋, 再右旋
				z = z.Parent
				t.LeftRotate(z)
			}
			z.Parent.Color = Black
			z.Parent.Parent.Color = Red
			t.RightRotate(z.Parent.Parent)
		} else { // 对称
			y := z.Parent.Parent.Left
			if y.Color == Red {
				z.Parent.Color = Black
				y.Color = Black
				z.Parent.Parent.Color = Red
				z = z.Parent.Parent
			} else if z == z.Parent.Left {
				z = z.Parent
				t.RightRotate(z)
			}
			z.Parent.Color = Black
			z.Parent.Parent.Color = Red
			t.LeftRotate(z.Parent.Parent)
		}
	}
	t.root.Color = Black
}

func (t *RBTree) String() string {
	return t.root.String()
}
