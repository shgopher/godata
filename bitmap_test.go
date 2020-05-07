package godata

import (
	"testing"
)

func TestNewMap(t *testing.T) {
	b := NewBitMap(100000000)
	for i := 1; i < 1000000001; i++ {
		b.Add(uint64(i))
	}
	t.Log(b.IsExit(0))
	t.Log(b.IsExit(1000000001))
	t.Log("0",b.IsExit(1))//true
	t.Log("1", b.IsExit(100000000))  // true
	t.Log("2", b.IsExit(99999998))   // true
	t.Log("3", b.IsExit(1000000000)) //false
	b.Delete(100000000)
	t.Log("4", b.IsExit(100000000)) //false
	b.Add(100000000)
	t.Log("5", b.IsExit(100000000)) // true
	b.Init()
	t.Log("6", b.IsExit(1)) //false
	b.Add(1)
	t.Log("7", b.IsExit(1)) //true
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