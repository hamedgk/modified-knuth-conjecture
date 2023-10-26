package data_structures

import (
	"fmt"
	"math"
)

type Operation int

const (
	Sqrt Operation = iota
	Floor
	TimesFive
	None
)

type Node struct {
	Parent *Node
	Operator Operation
	Result   float64
}

func (node *Node) Expand(queue Queue, explored map[float64]bool) {
	sqrtNode := Node{Parent: node, Operator: Sqrt, Result: math.Sqrt(node.Result)}
	floorNode := Node{Parent: node, Operator: Floor, Result: math.Floor(node.Result)}
	multiply5Node := Node{Parent: node, Operator: TimesFive, Result: node.Result * 5.0}

	children := []Node{
		floorNode,
		sqrtNode,
		multiply5Node,
	}
	for _, node := range children {
		if !explored[node.Result] {
			queue.Enqueue(node)
			explored[node.Result] = true
		}
	}
}

func (node *Node) IsGoal(goal float64) (Node, bool) {
	if node.Result == goal {
		return *node, true
	} else {
		return *node, false
	}
}

func TraceBack(node Node) {
	if node.Parent == nil {
		fmt.Printf("init=(%.2f)", node.Result)
		return
	}

	TraceBack(*node.Parent)
	switch node.Operator {
	case Sqrt:
		fmt.Printf(" -> Sqrt=(%.2f)", node.Result)
	case Floor:
		fmt.Printf(" -> Floor=(%.2f)", node.Result)
	case TimesFive:
		fmt.Printf(" -> *5=(%.2f)", node.Result)
	default:
		return
	}
}
