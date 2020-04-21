package godata

type DNode struct {
	Value interface{}
	Child []*DNode
}
var (
	ma = make(map[*DNode]int)
)

func(d *DNode)DFS(ma map[*DNode]int,value interface{})*DNode{
	ma[d]++
	var t *DNode
	if d.Child == nil && d.Value != value {
		return nil
	}
	if d.Value == value {
		return d
	}
	for _,v := range d.Child {
		if _,ok := ma[v];!ok {
			t = v.DFS(ma,value)
			if t != nil {
				return t
			}
		}
	}
	return t
}
