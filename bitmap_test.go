package godata

import (
	"fmt"
	"testing"
)

func TestNewMap(t *testing.T) {
	b := NewBitMap(100000000)
	for i := 1; i < 1000000001; i++ {
		b.Add(uint64(i))
	}
	fmt.Println(b.IsExit(0))
	fmt.Println(b.IsExit(1000000001))
	fmt.Println("0",b.IsExit(1))//true
	fmt.Println("1", b.IsExit(100000000))  // true
	fmt.Println("2", b.IsExit(99999998))   // true
	fmt.Println("3", b.IsExit(1000000000)) //false
	b.Delete(100000000)
	fmt.Println("4", b.IsExit(100000000)) //false
	b.Add(100000000)
	fmt.Println("5", b.IsExit(100000000)) // true
	b.Init()
	fmt.Println("6", b.IsExit(1)) //false
	b.Add(1)
	fmt.Println("7", b.IsExit(1)) //true
}
func BenchmarkBitMap_IsExit(b *testing.B) {
	bii := NewBitMap(100000000)
	for i := 1; i < 1000000001; i++ {
		bii.Add(uint64(i))
	}
	for i := 0; i < b.N; i++ {
		bii.IsExit(10)
	}
}