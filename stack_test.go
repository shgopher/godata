package godata

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	fmt.Println(stack.Length()) // 0
	fmt.Println(stack.Push(1))  //1
	fmt.Println(stack.Pop())    //1
	fmt.Println(stack.Pop())   // nil
	fmt.Println(stack.Top())   // nil
	fmt.Println(stack.Push(1)) // 1
	fmt.Println(stack.Push(2)) // 2
	fmt.Println(stack.Push(3)) // 3
	fmt.Println(stack.Length()) // 3
	fmt.Println(stack.Top()) // 3
	fmt.Println(stack.Pop()) // 3
	fmt.Println(stack.Pop()) // 2
	fmt.Println(stack.Pop()) // 1
	fmt.Println(stack.Pop()) // nil
	fmt.Println(stack.Length()) // 0
}
