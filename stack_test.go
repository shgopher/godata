package godata

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack(3)
	t.Log(stack.Length()) // 0
	t.Log(stack.Push(1))  //1
	t.Log(stack.Pop())    //1
	t.Log(stack.Pop())    // nil
	t.Log(stack.Top())    // nil
	t.Log(stack.Push(1))  // 1
	t.Log(stack.Push(2))  // 2
	t.Log(stack.Push(3))  // 3
	t.Log(stack.Length()) // 3
	t.Log(stack.Top())    // 3
	t.Log(stack.Pop())    // 3
	t.Log(stack.Pop())    // 2
	t.Log(stack.Pop())    // 1
	t.Log(stack.Pop())    // nil
	t.Log(stack.Length()) // 0
}
