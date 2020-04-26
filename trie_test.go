package godata

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("😄 嘿嘿嘿")
	trie.Insert("😄 嘟嘟嘟")
	trie.Insert("😄 你看看我是谁")
	trie.Insert("😄 我tm的真帅")
	fmt.Println("test search",trie.Search("😄"))
	fmt.Println(trie.StartsWith("😄"))
	fmt.Println("test search2",trie.Search("😄 嘿嘿嘿"))
	for _, v := range trie.Image("😄") {
			fmt.Println("测试结果",string(v))
	}
}
func BenchmarkTrie_Insert(b *testing.B) {
	trie := NewTrie()
	for i := 0; i < b.N; i++ {
		trie.Insert("😄 嘿嘿嘿")
	}
}
func BenchmarkTrie_Search(b *testing.B) {
	trie := NewTrie()
	trie.Insert("😄 嘿嘿嘿")
	for i := 0; i < b.N; i++ {
		trie.Search("😄")
	}
}
func BenchmarkTrie_StartsWith(b *testing.B) {
	trie := NewTrie()
	trie.Insert("😄 嘿嘿嘿")
	for i := 0; i < b.N; i++ {
		trie.StartsWith("😄")
	}
}
func BenchmarkTrie_Image(b *testing.B) {
	trie := NewTrie()
	trie.Insert("1234")
	trie.Insert("12345678")
	trie.Insert("2468")
	for i := 0; i < b.N; i++ {
		trie.Image("1")
	}
}
// 测试多字符长字符下的插入和搜索性能
func BenchmarkTrie_Insert2(b *testing.B) {
	trie := NewTrie()

	for i := 0; i < b.N; i++ {
		trie.Insert(w1)
		trie.Insert(w2)
		trie.Insert(w3)
		trie.Insert(w4)
		trie.Insert(w5)
		trie.Insert(w6)
		trie.Insert(w7)
		trie.Insert(w8)
		trie.Insert(w9)
	}
}
func BenchmarkTrie_Search2(b *testing.B) {
	trie := NewTrie()
	trie.Insert(w1)
	trie.Insert(w2)
	trie.Insert(w3)
	trie.Insert(w4)
	trie.Insert(w5)
	trie.Insert(w6)
	trie.Insert(w7)
	trie.Insert(w8)
	trie.Insert(w9)
	for i := 0; i < b.N; i++ {
	trie.Search("google is the greatest company arroud the world")
	}
}
func BenchmarkTrie_StartsWith2(b *testing.B) {
	trie := NewTrie()
	trie.Insert(w1)
	trie.Insert(w2)
	trie.Insert(w3)
	trie.Insert(w4)
	trie.Insert(w5)
	trie.Insert(w6)
	trie.Insert(w7)
	trie.Insert(w8)
	trie.Insert(w9)
	for i := 0; i < b.N; i++ {
		trie.StartsWith("google")
	}
}
func BenchmarkTrie_Image2(b *testing.B) {
	trie := NewTrie()
	trie.Insert(w1)
	trie.Insert(w2)
	trie.Insert(w3)
	trie.Insert(w4)
	trie.Insert(w5)
	trie.Insert(w6)
	trie.Insert(w7)
	trie.Insert(w8)
	trie.Insert(w9)
	for i := 0; i < b.N; i++ {
		trie.Image("google")
	}
}
func TestTrie_Image(t *testing.T) {
	trie := NewTrie()
	trie.Insert(w1)
	trie.Insert(w2)
	trie.Insert(w3)
	trie.Insert(w4)
	trie.Insert(w5)
	trie.Insert(w6)
	trie.Insert(w7)
	trie.Insert(w8)
	trie.Insert(w9)
	for _,v := range trie.Image("google is my") {
		fmt.Println(string(v))
	}
}

var (
	w1 = "google is my favorite company"
	w2 = "google is my love company"
	w3 = "google is very cool"
	w4 = "google is so good i like it"
	w5 = "google is too big "
	w6 = "google is the greatest company arroud the world"
	w7 = "google is amazing"
	w8 = "google is my lover"
	w9 = "i like google"
)
