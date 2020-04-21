package godata

import (
	"fmt"
	"testing"
)

func TestNode_BFS(t *testing.T) {
	n := new(Node)
	n.Value= 12
	n.Child = []*Node{
		&Node{
			Value: 1,
			Child: []*Node{
				&Node{
					Value: 100,
					Child: nil,
				},
				&Node{
					Value: 189,
					Child: nil,
				},
			},
		},
		&Node{
			Value: 187,
			Child: []*Node{
				&Node{
					Value: 10320,
					Child: nil,
				},
				&Node{
					Value: 1849,
					Child: nil,
				},
			},
		},
		&Node{
			Value: 541,
			Child: []*Node{
				&Node{
					Value: 10320,
					Child: nil,
				},
				&Node{
					Value: 18787659,
					Child: nil,
				},
			},
		},

	}
	fmt.Println(n.BFS(18787659))
}