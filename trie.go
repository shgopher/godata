package godata

// 仅支持 小写字符
type Trie struct { // 字典树使用数组不如hash table 好使。还有 一个节点的值存在于他的children的k上 所以其实字典树没有节点 所谓值
	// 都是存在于children中的
	Children map[rune]*Trie
	IsWord bool
}


/** Initialize your data structure here. */
func NewTrie() *Trie {
	return &Trie{
		Children: map[rune]*Trie{},
		IsWord : false,
	}
}


/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
	for _,v := range word {
		if t.Children[v] == nil { // 使用 v - 'a' 是为了获取0值
			t.Children[v] = &Trie { /// 插入的时候看是否存在这个k了，如果有了就往下遍历，如果没有就新建。
				IsWord:false,
				Children: map[rune]*Trie{},
			}
		}
		t = t.Children[v] // 往下遍历 类似 node= node.next
	}
	t.IsWord = true // 将已经 文字结束设置为true
}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {

	for _,v := range word {
		if t.Children[v] == nil {
			return false
		}
		t = t.Children[v]
	}
	return t.IsWord // 为了看是不是前缀还是说是真的全部
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) (bool,*Trie) {

	for _,v := range prefix {
		if t.Children[v] == nil  {
			return false,nil
		}
		t = t.Children[v]
	}
	return true,t // 只要匹配直接返回true，反正都符合了。
}
func(t *Trie)Image(prefix string) [][]rune {
	result := make([][]rune,0)
	re := make([]rune,0,len(prefix))
	re = append(re,[]rune(prefix)...)

	if is,tr := t.StartsWith(prefix);is {
		rangeTrie(tr,&result,re)
	}
	return result
}
func rangeTrie(t *Trie,result *[][]rune,re []rune){
	if t.IsWord == true {
		ma := make([]rune,len(re))
		copy(ma,re)
		*result = append(*result,ma)
	}
	for k,v := range t.Children {
		rangeTrie(v,result,append(re,k))
	}
}