package godata

import (
	"fmt"
	"testing"
)

func TestDNode_DFS(t *testing.T) {
	fmt.Println(ddd().DFS(ma,1))

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
