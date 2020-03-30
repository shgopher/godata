package godata

import (
	"fmt"
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
	fmt.Println(PopMin(&a))
	fmt.Println(PopMin(&a))
	fmt.Println(PopMin(&a))
	fmt.Println(PopMin(&a))
	fmt.Println(PopMin(&a))
	fmt.Println(PopMin(&a))
}
