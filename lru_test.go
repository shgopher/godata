package godata

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	l := NewLRU(2)
	fmt.Println(l.Get(1))
	l.Put(1,1)
	l.Put(1,2)
	l.Put(2,2)
	fmt.Println(l.Get(1))
	l.Put(3,3)
	fmt.Println(l.Get(2))
	l.Put(4,4)
	fmt.Println(l.Get(4))
	l.ReBuild()
	l.Put(5,5)
	fmt.Println(l.Get(4))
	fmt.Println(l.Get(5))
}

func BenchmarkNewLRU(b *testing.B) {

	for i := 0; i < b.N; i++ {
	 NewLRU(2)
	}
}
func BenchmarkLRUCache_Get(b *testing.B) {
	l := NewLRU(3)
	l.Put(1,1)
	l.Put(2,2)
	for i := 0; i < b.N; i++ {
		l.Get(2)
	}
}
func BenchmarkLRUCache_Put(b *testing.B) {
	l := NewLRU(2)
	for i := 0; i < b.N; i++ {
		l.Put(1,1)
	}
}
func BenchmarkLRUCache_ReBuild(b *testing.B) {
	l := NewLRU(2)
	l.Put(1,1)
	l.Put(2,2)
	l.Get(1)
	l.Put(2,3)
	l.Put(5,5)
	for i := 0; i < b.N; i++ {
		l.ReBuild()
	}
}