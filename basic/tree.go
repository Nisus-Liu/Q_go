package basic

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type BiNode interface {
	GetLeft() interface{}
	GetRight() interface{}
}

type Tree interface {
	GetRoot() interface{}
}



/*
== 二叉树树形打印 ==
                                       8
                   2                                      16
                             2                                      54
                                                                        464
1 准备二维切片, 第一维度表示行. 第二维表示每行的打印内容.
2 求得树的最大高度
3 每层根据自身所在高度, 计算占宽, 并追加到对应行中.
4 逐行输出内容

每个节点key占位 5 个字符， 简单起见， 节点间隔和key宽一致。
节点所在度设为 d ， w 节点宽度和间隔。
i = （h-d+1）
单个节点占宽 = (2^i - 1) * w

每个节点居中， 两边空格填充 。
*/
const W = 5
const ONE_W_SPACE = "     "
func PrintTree(tree *RBTree) {
	h := maxDepth(tree.GetRoot())
	var lines = make([]bytes.Buffer, h)
	curNode := tree.GetRoot()
	printTree0(h, curNode, 1, lines)

	// 输出
	for _, line := range lines {
		fmt.Println(line.String())
	}
}

func printTree0(h int, curNode *Node, d int, lines []bytes.Buffer) {
	var buffer *bytes.Buffer = &lines[d-1]
	k := math.Pow(2, float64(h-d+1)) - 1
	var left, right *Node
	// 空节点时, 只需要占位符
	if(curNode==EMPTY) {
		// buffer.WriteString(strings.Repeat(ONE_W_SPACE, int(k)))
		left = EMPTY
		right = EMPTY
	}else {
		left = curNode.Left
		right = curNode.Right
	}
	pad := strings.Repeat(ONE_W_SPACE, int(k)/2)
	buffer.WriteString(pad)
	// key 居右对齐
	strKey := formatNodeContent(curNode)
	i := W - len(strKey)
	if i > 0 {
		keyPad := strings.Repeat(" ", i)
		buffer.WriteString(keyPad)
	} else if i < 0 {
		fmt.Printf("WARN: `key` 长度超过了 %d \n", W)
	}
	buffer.WriteString(strKey)
	// 间隔
	buffer.WriteString(ONE_W_SPACE)
	buffer.WriteString(pad)

	if d < h {
		printTree0(h, left, d+1, lines)
		printTree0(h, right, d+1, lines)
	}
}

// 格式化节点打印内容
// 黑节点 () 包裹
func formatNodeContent(node *Node) string {
	var cnt = ""
	if node != EMPTY {
		cnt = strconv.FormatInt(node.Key, 10)
	}else {
		cnt = "nil"
	}
	if node.Color == Black{
		cnt = "("+ cnt +")"
	}
	return cnt
}

func maxDepth(curNode *Node) int {
	if curNode == EMPTY {
		return 0
	}

	ld := maxDepth(curNode.Left)
	rd := maxDepth(curNode.Right)

	var i = 0
	if(ld> rd) {
		i = ld
	}else {
		i = rd
	}
	return i + 1
}
