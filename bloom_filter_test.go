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