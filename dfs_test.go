package godata

import (
	"testing"
)

func TestDNode_DFS(t *testing.T) {
	t.Log(ddd().DFS(ma,1))

}
func BenchmarkDNode_DFS(b *testing.B) {
	t := ddd()
	for i := 0; i < b.N; i++ {
		t.DFS(ma,10320)
	}
}
func ddd()*DNode{
	n := new(DNode)
	n.Value= 12
	n.Child = []*DNode{
		&DNode{
			Value: 1,
			Child: []*DNode{
				&DNode{
					Value: 100,
					Child: nil,
				},
				&DNode{
					Value: 189,
					Child: nil,
				},
			},
		},
		&DNode{
			Value: 187,
			Child: []*DNode{
				&DNode{
					Value: 10320,
					Child: nil,
				},
				&DNode{
					Value: 1849,
					Child: nil,
				},
			},
		},
		&DNode{
			Value: 541,
			Child: []*DNode{
				&DNode{
					Value: 10320,
					Child: nil,
				},
				&DNode{
					Value: 18787659,
					Child: nil,
				},
			},
		},

	}
	return n
}
