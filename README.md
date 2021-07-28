# godata
[![Goproxy.cn](https://goproxy.cn/stats/github.com/shgopher/godata/badges/download-count.svg)](https://goproxy.cn)

Basic data structures and operations written in Go which  golang.org/pkg/container don't support such as stack, queue, etc.

|list|done|
|:---:|:---:|
|trie|:ok:|
|bitmap|:ok:|
|bloom filter|:ok:|
|lru|:ok:|
|heap|:ok:|
|stack|:ok:|
|queue|:ok:|
|dfs|:ok:|
|bfs|:ok:|
|union and find|:ok:|
|skiplist|:x:|
|b|:x:|
|graph|:x:|
|dancing Links|:x:|
|segment tree|:x:|
|huffman tree|:x:|
## Quick start
```go
package main
import(
  "github.com/googege/godata"
)
func main(){
// stack
    stack := godata.NewStack(5)
    stack.Push(1)
    stack.Length()
    stack.Top()
    stack.Pop()
// trie
    trie := NewTrie()
    trie.Insert("hello world")
    trie.Insert("hello China")
    trie.Insert("hello My dear")
    trie.Search("hello world")
    trie.StartWith("he")
    trie.Image("he")
}

```
## HERE
|项目|介绍|
|:---:|:---:|
|便宜服务器推荐|[阿里云](https://www.aliyun.com/minisite/goods?userCode=ol87kpmz)，[梯子服务器](https://app.cloudcone.com/?ref=2525):支持支付宝|
|微信公众号|![p](https://raw.githubusercontent.com/basicExploration/Demos/master/pluspro.png)|
|我的社交平台|[b站](https://space.bilibili.com/23170151)|
