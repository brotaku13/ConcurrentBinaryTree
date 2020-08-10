package ConcurrentBinarytree

import (
	_"fmt"
)

type Node struct {
	Value int
	parent *Node
	right *Node
	left *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree)Insert(value int) bool {
	if tree.root == nil {
		tree.root = &Node{value, nil, nil, nil}
		return true
	}
	cursor := tree.root
	for {
		if value < cursor.Value {
			if cursor.left == nil {
				cursor.left = &Node{value, cursor, nil, nil}
				return true
			} else {
				cursor = cursor.left
			}
		} else if value > cursor.Value {
			if cursor.right == nil {
				cursor.right = &Node{value, cursor, nil, nil}
				return true
			} else {
				cursor = cursor.right
			}
		} else {
			return false
		}
	}
}

func (tree *BinaryTree) remove(n *Node, value int) bool {
	cursor := n
	// find the node to be deleted first
	for cursor != nil {
		if cursor.Value == value {
			break
		} else if cursor.Value < value {
			cursor = cursor.right
		} else {
			cursor = cursor.left
		}
	}
	// if the cursor is nil, then we couldn't find it
	if cursor == nil {
		return false
	}

	// Case 1: node to be deleted has no children
	if(cursor.right == nil && cursor.left == nil) {
		if cursor != tree.root {
			// Change parent's pointer to point to new 
			if(cursor.parent.left == cursor) {
				cursor.parent.left = nil
			} else {
				cursor.parent.right = nil
			}
		} else {
			tree.root = nil
		}
	} else if (cursor.left != nil && cursor.right != nil) {
		// Case 2: node to be delted has two children	
		// get the minimum key
		minNode := cursor.right		
		for minNode.left != nil {
			minNode = minNode.left
		}
		// take a copy of the miniimum data
		minData := minNode.Value
		// recursively remove the successor
		tree.remove(cursor, minNode.Value)

		cursor.Value = minData

	} else {
		// Case 3: node to be deleted has only one node
		// first find the child that exists
		var child *Node
		if cursor.right != nil {
			child = cursor.right
		} else {
			child = cursor.left
		}
		// if node to be deleted is not a root node, then set it's parent to
		// it's child
		if cursor != tree.root {
			if cursor == cursor.parent.left {
				cursor.parent.left = child
			} else {
				cursor.parent.right = child
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
			cursor = cursor.right
		} else {
			cursor = cursor.left
		}
	}
	return nil
}