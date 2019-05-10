package main

import (
	"github.com/gogf/gf/g/container/gtree"
)

func main() {
	tree := gtree.NewRedBlackTree(func(v1, v2 interface{}) int {
		return v1.(int) - v2.(int)
	})
	for i := 0; i < 20; i++ {
		tree.Set(i, i*10)
	}
	tree.Print()
	tree.Flip()
	tree.Print()
}