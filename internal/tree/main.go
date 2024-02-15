package tree

import (
	"fmt"
	"internal/utils"
	"math"
)

type Node[T interface{}] struct {
	Data   T
	Left   *Node[T]
	Right  *Node[T]
	Parent *Node[T]
}

func NewNode[T any](data T, nodes ...*Node[T]) *Node[T] {
	return &Node[T]{
		Data:   data,
		Left:   utils.Ternary[*Node[T]](nodes[0] != nil, nodes[0], nil),
		Right:  utils.Ternary[*Node[T]](nodes[1] != nil, nodes[1], nil),
		Parent: utils.Ternary[*Node[T]](nodes[2] != nil, nodes[2], nil),
	}
}

func (n *Node[T]) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node[T]) IsRoot() bool {
	return n.Parent == nil
}

func (n *Node[T]) IsLeftChild() bool {
	return n.Parent != nil && n.Parent.Left == n
}

func (n *Node[T]) IsRightChild() bool {
	return n.Parent != nil && n.Parent.Right == n
}

func (n *Node[T]) Sibling() *Node[T] {
	if n.IsRoot() {
		return nil
	}
	if n.IsLeftChild() {
		return n.Parent.Right
	}
	return n.Parent.Left
}

func (n *Node[T]) GetRoot() *Node[T] {
	if n.IsRoot() {
		return n
	}
	return n.Parent.GetRoot()
}

func (n *Node[T]) GetHeight() float64 {
	if n.IsLeaf() {
		return 0
	}
	return 1 + math.Max(n.Left.GetHeight(), n.Right.GetHeight())
}

func (node *Node[T]) Print() { // initially root must be passed here
	visited := []*Node[T]{}
	print(node, visited)
}

func print[T any](node *Node[T], visited []*Node[T]) {
	if !utils.Includes(node, visited) {
		fmt.Println(node.Data)
		visited = append(visited, node)
	}

	if node.Left != nil && !utils.Includes[Node[T]](node.Left, visited) {
		print(node.Left, visited)
	}

	if node.Right != nil && !utils.Includes[Node[T]](node.Right, visited) {
		print(node.Right, visited)
	}
}
