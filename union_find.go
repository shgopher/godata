package godata
type UnionFind struct {
	roots []int
}
// new a union & find.
func NewUnionFind(roots int)*UnionFind{
	ma := make([]int,roots)
	for i:= 0;i < roots;i++ {
		ma[i]=i
	}
	return &UnionFind{
		roots: ma,
	}
}
// find root(his boss)
func(u *UnionFind)FindRoot(i int)int{
	root := i
	// find the boss
	for u.roots[root] != root {
		root = u.roots[root]
	}
	// union all node to root.
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