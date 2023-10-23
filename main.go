package main

import (
	"fmt"
	"knuth/data_structures"
)

func main() {
	queue := data_structures.NewQueue()
	goal := 111.0

	for {
		if queue.IsEmpty() {
			fmt.Println("empty frontier")
			return
		}
		value, ok := queue.Dequeue()
		if node, ok := value.IsGoal(goal); ok {
			data_structures.TraceBack(node)
			return
		}
		if ok {
			value.Expand(queue)
		}
	}
}
