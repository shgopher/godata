package godata

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue(3)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	t.Log(queue.Top())
	t.Log(queue.Pop())
	t.Log(queue.Pop())
	t.Log(queue.Pop())
	// nil
	t.Log(queue.Pop())
	t.Log(queue.Length())
	t.Log(queue.val)
	t.Log("top",queue.Top())
	queue.Push(11)
	queue.Push(12)
	queue.Push(13)
	t.Log(queue.PopBack())
	t.Log(queue.PopBack())
	t.Log(queue.PopBack())


}
