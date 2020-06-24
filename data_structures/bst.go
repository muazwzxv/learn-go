package main

type node struct {
	data  int
	left  *node
	right *node
}

type tree struct {
	root *node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
