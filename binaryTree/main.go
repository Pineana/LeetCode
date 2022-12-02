package main

import "fmt"

func main() {
	root := buildTree()
	preOrder(root)
	fmt.Println()
	preOrderWithStack(root)
	fmt.Println()
	inOrder(root)
	fmt.Println()
	inOrderWithStack(root)
	fmt.Println()
	postOrder(root)
	fmt.Println()
	postOrderWithStack(root)
	fmt.Println()

}

type Stack struct {
	Data  []*TreeNode
	Count int
}

func NewStack() *Stack {
	data := make([]*TreeNode, 0)
	return &Stack{
		Data:  data,
		Count: 0,
	}
}

func (s *Stack) push(node *TreeNode) {
	s.Data = append(s.Data, node)
	s.Count++
}

func (s *Stack) pop() *TreeNode {
	if s.Count > 0 {
		s.Count--
		popItem := s.Data[s.Count]
		s.Data = s.Data[0:s.Count]
		return popItem
	}
	panic("stack is empty")
}

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree() *TreeNode {
	root := &TreeNode{Data: 1}
	root.Left = &TreeNode{Data: 2}
	root.Right = &TreeNode{Data: 3}
	root.Left.Left = &TreeNode{Data: 4}
	root.Left.Right = &TreeNode{Data: 5}
	return root
}

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Data)
	preOrder(node.Left)
	preOrder(node.Right)
}

func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	postOrder(node.Left)
	postOrder(node.Right)
	fmt.Printf("%d ", node.Data)
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	fmt.Printf("%d ", node.Data)
	inOrder(node.Right)
}

func preOrderWithStack(node *TreeNode) {
	s := NewStack()
	for s.Count != 0 || node != nil {
		for node != nil {
			fmt.Printf("%d ", node.Data)
			s.push(node)
			node = node.Left
		}
		if s.Count != 0 {
			node = s.pop()
			node = node.Right
		}
	}
}

func inOrderWithStack(node *TreeNode) {
	s := NewStack()
	for s.Count != 0 || node != nil {
		for node != nil {
			s.push(node)
			node = node.Left
		}
		if s.Count != 0 {
			node = s.pop()
			fmt.Printf("%d ", node.Data)
			node = node.Right
		}
	}
}

func postOrderWithStack(node *TreeNode) {
	s := NewStack()
	s2 := NewStack()
	for s.Count != 0 || node != nil {
		for node != nil {
			s.push(node)
			s2.push(node)
			node = node.Right
		}
		if s.Count != 0 {
			node = s.pop()
			node = node.Left
		}
	}
	for s2.Count > 0 {
		fmt.Printf("%d ", s2.pop().Data)
	}
}
