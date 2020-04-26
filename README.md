# godata
Basic data structures and operations written in Go which  golang.org/pkg/container don't support such as stack, queue, etc.
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
|items|description|
|:---:|:---:|
|Sponsor me|![p](https://raw.githubusercontent.com/basicExploration/Demos/master/donate.png)|
|aliyun vps|[aliyun vps](https://www.aliyun.com/minisite/goods?userCode=ol87kpmz)，China's largest cloud computing manufacturer, the world's top 5 cloud computing service providers|
|VPS|[vps](https://app.cloudcone.com/?ref=2525) fast and cheap just 2-3.75$/mo|
|WeChat public account|![p](https://raw.githubusercontent.com/googege/GOFamily/master/joinUsW.jpg)|
|My wechat (Please mark：form Github)|![p](https://raw.githubusercontent.com/googege/GOFamily/master/me.jpeg)|
|My bilibili|[b](https://space.bilibili.com/23170151)|
|MyYouTube|[YouTube](https://www.youtube.com/channel/UCM_-pFgD_HZDGD0yxfzguRQ?view_as=subscriber)|
