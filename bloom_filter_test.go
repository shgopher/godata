package godata

import (
	"fmt"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter(3)
	bf.Add([]byte("魔都一只土拨鼠"))
	fmt.Println(bf.IsExit([]byte("魔都一只土拨鼠")))
	bf.Add([]byte("大家好"))
	fmt.Println(bf.IsExit([]byte("魔都一只土拨鼠!")))
	fmt.Println(bf.IsExit([]byte("大家好")))
}
func BenchmarkBloomFilter_Add(b *testing.B) {
	bf := NewBloomFilter(3)
	for i := 0; i < b.N; i++ {
		bf.Add([]byte("1"))
	}
}

func BenchmarkBloomFilter_IsExit(b *testing.B) {
	bf := NewBloomFilter(3)
	bf.Add([]byte("1"))
	for i := 0; i < b.N; i++ {
		bf.IsExit([]byte("1"))
	}
}
func BenchmarkHashShower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HashShower([]byte("1"),1)
	}
}