package godata

type Node struct {
	Value interface{}
	Child []*Node
}

// 寻找到目标值，并且返回目标struct. node是要加入搜索的root节点，value是目标值
func (n *Node) BFS(value interface{}) *Node {
	// 定义
	ma := make(map[*Node]int)
	result := &Node{}
	queue := NewQueue(1)
	// 首行辅助任务完成
	queue.Push(n)
	ma[n]++
	// 开始循环寻找
L:
	for queue.Length() > 0 {
		v := queue.Pop()
		val := v.(*Node)
		ma[val]++
		if val.Value == value {
			result = val
			break L
		}
		for _, v := range val.Child {
			if _,ok := ma[v];!ok {
				queue.Push(v)
			}
		}

	}
	return result
}

