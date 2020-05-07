package godata

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	t1 := make([]int,0,0)
	var a MinHeap
	a = t1
	PushMin(&a,11)
	PushMin(&a,211)
	PushMin(&a,32)
	PushMin(&a,43)
	PushMin(&a,51)
	PushMin(&a,63)
	t.Log(PopMin(&a))
	t.Log(PopMin(&a))
	t.Log(PopMin(&a))
	t.Log(PopMin(&a))
	t.Log(PopMin(&a))
	t.Log(PopMin(&a))
}
