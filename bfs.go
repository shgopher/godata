package godata

type Node struct {
	Value interface{}
	Child []*Node
}

// node is the root，value is target
func (n *Node) BFS(value interface{}) *Node {
	ma := make(map[*Node]int)
	result := &Node{}
	queue := NewQueue(10)
	// a mother queue.
	queue.Push(n)
	ma[n]++
	for queue.Length() > 0 {
		v := queue.Pop()
		val := v.(*Node)
		ma[val]++
		if val.Value == value {
			result = val
			break
		}
		for _, v := range val.Child {
			if _,ok := ma[v];!ok {
				queue.Push(v)
			}
		}

	}
	return result
}

