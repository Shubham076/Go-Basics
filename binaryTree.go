package main

import (
	"container/list"
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

func generateTree(arr []int) *node {
	var root *node = nil
	q := list.New() //doubly linked list
	q.PushFront(&node{arr[0], nil, nil})
	i := 1
	for q.Len() > 0 && i < len(arr) {
		rm := q.Front()  //type assertion in go
		el := rm.Value.(*node)
		q.Remove(rm);

		if root == nil {
			root = el  
		}

		// fmt.Print(el, arr[i]);
		if arr[i] != -1 {
			el.left = &node{arr[i], nil, nil}
			q.PushFront(el.left)
		}
		i++;

		if i >= len(arr) {
			break;
		} 

		if(arr[i] != -1) {
			el.right = &node{arr[i], nil, nil}
			q.PushFront(el.right)
		}
		i++;
	}
	return root
}

func preOrder(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.data);
	preOrder(root.left);
	preOrder(root.right);
}

func main() {
	arr := [] int {1, 2, 3, 4, 5};
	root := generateTree(arr);
	// root := node{}
	preOrder(root);
}
