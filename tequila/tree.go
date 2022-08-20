package tequila

import (
	"strings"
)

type treeNode struct {
	name       string
	children   []*treeNode
	routerName string
	isEnd      bool
}

func (t *treeNode) Put(str string) {
	root := t
	for index, name := range strings.Split(str, "/") {
		if index == 0 {
			continue
		}
		children := t.children
		isMatch := false
		for _, node := range children {
			if node.name == name {
				isMatch = true
				t = node
				break
			}
		}
		if !isMatch {
			node := &treeNode{name: name, children: make([]*treeNode, 0), isEnd: index == len(str)-1}
			children = append(t.children, node)
			t.children = children
			t = node
		}
	}
	t = root
}

func (t *treeNode) Get(str string) *treeNode {
	s := strings.Split(str, "/")
	routerName := ""
	for index, name := range s {
		if index == 0 {
			continue
		}
		children := t.children
		for _, node := range children {
			if node.name == name || node.name == "*" || strings.Contains(node.name, ":") {
				t = node
				routerName += "/" + node.name
				node.routerName = routerName
				if index == len(s)-1 {
					return node
				}
				break
			}
		}
	}
	return nil
}
