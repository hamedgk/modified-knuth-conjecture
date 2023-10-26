package main

import (
	"fmt"
	"knuth/data_structures"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	goal, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err.Error()) // 3.1415927410125732
		return
	}

	explored := map[float64]bool{}
	frontier := data_structures.NewQueue()
	frontier.Enqueue(data_structures.Node{Parent: nil, Operator: data_structures.None, Result: 4.0})

	for {
		if frontier.IsEmpty() {
			fmt.Println("empty frontier")
			return
		}
		value, ok := frontier.Dequeue()
		if node, found := value.IsGoal(goal); found {
			data_structures.TraceBack(node)
			return
		}
		if ok {
			value.Expand(frontier, explored)
		}
	}
}
