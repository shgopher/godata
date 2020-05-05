package godata

// support all utf8 code.
type Trie struct {
	// the node's value is her children's index .so the node does not have real value.
	// or you can say ,she is not have value,the child is her value.
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
		if t.Children[v] == nil {
			t.Children[v] = &Trie {
				IsWord:false,
				Children: map[rune]*Trie{},
			}
		}
		t = t.Children[v]
	}
	t.IsWord = true // it is finished.you let the IsWord = true.
}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {

	for _,v := range word {
		if t.Children[v] == nil {
			return false
		}
		t = t.Children[v]
	}
	return t.IsWord // we dont know it's the whole word,or just pre,we we return IsWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) (bool,*Trie) {

	for _,v := range prefix {
		if t.Children[v] == nil  {
			return false,nil
		}
		t = t.Children[v]
	}
	return true,t // because we dont car it's whole word or just only prefix,becase we just want prefix.
}
// return the values which match the prefix.
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