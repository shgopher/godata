package godata
type UnionFind struct {
	roots []int
}
func NewUnionFind(roots int)*UnionFind{
	ma := make([]int,roots)
	for i:= 0;i < roots;i++ {
		ma[i]=i
	}
	return &UnionFind{
		roots: ma,
	}
}
// 寻找老大的过程
func(u *UnionFind)FindRoot(i int)int{
	root := i
	// 寻找自己的大boss
	for u.roots[root] != root {
		root = u.roots[root]
	}
	// 并查集的合并
	for u.roots[i] != i{
		u.roots[i],i = root,u.roots[i]
	}
	return root
}
func(u *UnionFind)Union(x,y int) {
	xRoot := u.FindRoot(x)
	yRoot := u.FindRoot(y)
	u.roots[xRoot] = yRoot
}
func (u *UnionFind)isUnion(x,y int)bool {
	return u.FindRoot(x) == u.FindRoot(y)
}