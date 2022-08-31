package binary_tree

import (
	"fmt"
	"strconv"
)

type Node struct {
	Value       int
	Left, Right *Node
	res         []string
}

func NewNode(Value int) *Node {
	return &Node{
		Value: Value,
		Left:  nil,
		Right: nil,
	}
}

func (r *Node) Add(value int) {
	if r.Value == 0 {
		if value == 0 {
			return
		}
		r.Value = value
		return
	}
	if value < r.Value {
		r.addToLeft(value)
	} else {
		r.addToRight(value)
	}

}

func (r *Node) PrintPreOrder() {
	if r == nil {
		return
	}
	fmt.Printf("%d ", r.Value)
	r.Right.PrintPreOrder()
	r.Left.PrintPreOrder()
}

func (r *Node) PrintNode(n *Node) {
	if n == nil {
		return
	}
	r.PrintNode(n.Left)
	fmt.Printf("%d ", n.Value)
	r.PrintNode(n.Right)
}

func (r *Node) String(n *Node) []string {
	if n == nil {
		return nil
	}
	r.String(n.Left)
	r.res = append(r.res, strconv.Itoa(n.Value))
	r.String(n.Right)
	return r.res
}

func (r *Node) addToLeft(value int) {
	if r.Left != nil {
		r.Left.Add(value)
	} else {
		r.Left = NewNode(value)
	}
}

func (r *Node) addToRight(value int) {
	if r.Right != nil {
		r.Right.Add(value)
	} else {
		r.Right = NewNode(value)
	}
}
