package main

import "testing"

func createTestTree() *BinaryTree {
	tree := BinaryTree{}
	tree.insert(20)
	tree.insert(30)
	tree.insert(10)

	tree.insert(4)
	tree.insert(15)
	
	tree.insert(25)
	tree.insert(35)
	return &tree
}

func TestSearchExisting(t *testing.T) {
	root := Node{
		3,
		&Node{
			4,
			nil,
			nil,
		},
		&Node{
			2,
			nil,
			nil,
		},
	}
	tree := BinaryTree{&root}
	got := tree.Search(3)
	if got == nil {
		t.Errorf("Search could not find 3")
	}
	got = tree.Search(4)
	if got == nil {
		t.Error("Search could not find 4")
	}
	got = tree.Search(2)
	if got == nil {
		t.Error("Search could not find 2")
	}
}

func TestSearchNotExisting(t *testing.T) {
	root := Node{
		3,
		&Node{
			4,
			nil,
			nil,
		},
		&Node{
			2,
			nil,
			nil,
		},
	}
	tree := BinaryTree{&root}
	got := tree.Search(5)
	if got != nil {
		t.Errorf("Found 5 but it doesn't exist")
	}
	got = tree.Search(1)
	if got != nil {
		t.Error("Found 1 but it doesn't exist")
	}
	got = tree.Search(45)
	if got != nil {
		t.Error("Found 45 but it doesn't exist")
	}
}

func TestInsertOnEmptyTree(t *testing.T) {
	tree := BinaryTree{}
	got := tree.insert(20)
	if !got {
		t.Error("Tree could not insert value on empty tree")
	}
}

func TestInsertOnNonEmptyTree(t *testing.T) {
	tree := createTestTree()

	if tree.root.Value != 20 {
		t.Error("Tree root was not 20")
	} 
	if tree.root.Right.Value != 30 {
		t.Error("Tree right child not 30")
	} 
	if tree.root.Left.Value != 10 {
		t.Error("Tree left child not correct")
	}

	// test left subtree
	leftRoot := tree.root.Left
	if leftRoot.Left == nil || leftRoot.Left.Value != 4 {
		t.Error("Tree left root not 4")
	}
	if leftRoot.Left == nil ||  leftRoot.Right.Value != 15 {
		t.Error("Tree left root right child not 15")
	}

	rightRoot := tree.root.Right
	if rightRoot.Left == nil || rightRoot.Left.Value != 25 {
		t.Error("Tree right root not 25")
	}
	if rightRoot.Right == nil || rightRoot.Right.Value != 35  {
		t.Error("Tree right root right child not 35")
	}
}

func TestDeleteRoot(t *testing.T) {
	tree := createTestTree()

	found := tree.Remove(20)
	if !found {
		t.Error("Could not find root to delete")
	}

	// check successor should be 25
	if tree.root.Value != 25 {
		t.Error("Successor is not correct")
	}
}


func TestRemoveFromEmptyTree(t *testing.T) {
	tree := BinaryTree{}
	got := tree.Remove(2)
	if got {
		t.Error("Found non existing element in empty tree")
	}
}

func TestDeleteLeafOnRight(t *testing.T) {
	tree := createTestTree()
	
	if got := tree.Remove(4); !got {
		t.Error("Could not remove existing leaf element")
	}
	if got := tree.Search(4); got != nil {
		t.Error("Found non existing element")
	}
}

func TestDeleteLeafOnLeft(t *testing.T) {
	tree := createTestTree()
	
	if got := tree.Remove(15); !got {
		t.Error("Could not remove existing leaf element")
	}
	if got := tree.Search(15); got != nil {
		t.Error("Found non existing element")
	}
}

func TestRemoveRootWithNoChildren(t *testing.T) {
	tree := BinaryTree{}
	tree.insert(10)
	if found := tree.Search(10); found == nil {
		t.Error("Could not find root")
	}

	if !tree.Remove(10) || tree.Search(10) != nil {
		t.Error("Delete unsuccessful")
	}
}

func TestRemoveRootWithOneChild(t *testing.T) {
	tree := BinaryTree{}
	tree.insert(10)
	tree.insert(15)

	if foundRoot, foundChild := tree.Search(10), tree.Search(15);
		foundRoot == nil || foundChild == nil {
			t.Error("Could not find root or child")
		}
	
	if !tree.Remove(10) {
		t.Error("Could not remove root")
	}

	if tree.root.Value != 15 {
		t.Error("new root value is not 15")
	}

}

func TestRemoveNonRootWithOneChild(t *testing.T) {
	tree := BinaryTree{}
	tree.insert(10)
	tree.insert(15)	
	tree.insert(20)

	if !tree.Remove(15) {
		t.Error("Could  not find existing value 15")
	}

	if tree.root.Value != 10 || tree.root.Right.Value != 20 {
		t.Error("Removed elements left incorrect tree")
	}
}