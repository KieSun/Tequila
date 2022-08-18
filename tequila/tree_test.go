package tequila

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	root := treeNode{
		name:     "/",
		children: make([]*treeNode, 0),
	}
	root.Put("/user/:id")

	node := root.Get("/user/1")
	fmt.Print(node.routerName)
}
