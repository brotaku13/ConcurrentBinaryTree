package main

import (
	"fmt"
)

type Node struct {
	Value int
	Right *Node
	Left *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("Node@%p {%v, %p, %p}", n, n.Value, n.Right, n.Left)
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree)insert(value int) bool {
	if tree.root == nil {
		tree.root = &Node{value, nil, nil}
		return true
	}
	cursor := tree.root
	for {
		if value < cursor.Value {
			if cursor.Left == nil {
				cursor.Left = &Node{value, nil, nil}
				return true
			} else {
				cursor = cursor.Left
			}
		} else if value > cursor.Value {
			if cursor.Right == nil {
				cursor.Right = &Node{value, nil, nil}
				return true
			} else {
				cursor = cursor.Right
			}
		} else {
			return false
		}
	}
}

func (tree *BinaryTree) remove(n *Node, value int) bool {
	var parent *Node = nil
	cursor := n
	// find the node to be deleted first
	for cursor != nil {
		if cursor.Value == value {
			break
		} else if cursor.Value < value {
			parent = cursor
			cursor = cursor.Right
		} else {
			parent = cursor
			cursor = cursor.Left
		}
	}
	// if the cursor is nil, then we couldn't find it
	if cursor == nil {
		return false
	}
	// Case 1: node to be deleted has no childredn
	if(cursor.Right == nil && cursor.Left == nil) {
		if cursor != tree.root {
			if(parent.Left == cursor) {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		} else {
			tree.root = nil
		}
	} else if (cursor.Left != nil && cursor.Right != nil) {
		// Case 2: node to be delted has two children	
		// get the minimum key
		minNode := cursor.Right		
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		// take a copy of the miniimum data
		minData := minNode.Value
		// recursively remove the successor
		tree.Remove(minNode.Value)

		cursor.Value = minData

	} else {
		// Case 3: node to be deleted has only one node
		var child *Node
		if cursor.Right != nil {
			child = cursor.Right
		} else {
			child = cursor.Left
		}
		// if node to be deleted is not a root node, then set it's parent to
		// it's child
		if cursor != tree.root {
			if cursor == parent.Left {
				parent.Left = child
			} else {
				parent.Right = child
			} 
		} else {
			tree.root = child
		}
	}
	return true
}

func (tree *BinaryTree) Remove(value int) bool {
	return tree.remove(tree.root, value)
}


func (tree *BinaryTree) Search(value int) *Node {
	cursor := tree.root
	for cursor != nil {
		if cursor.Value == value {
			return cursor
		} else if cursor.Value < value {
			cursor = cursor.Right
		} else {
			cursor = cursor.Left
		}
	}
	return nil
}

func main() {

}
