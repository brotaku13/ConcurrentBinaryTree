package ConcurrentBinarytree

import "testing"

func TestInsertEmptyTree(t *testing.T) {
	tree := BinaryTree{}
	if tree.root != nil {
		t.Error("Root is not nil")
	}
	tree.Insert(10);

	root := tree.root;
	if root == nil {
		t.Error("Root does not exist")
	}
	if root.Value != 10 {
		t.Error("Root value is not correct")
	}
	if root.parent != nil {
		t.Error("Root parent is not nil")
	}
}

func TestInsertRightWithOnlyRoot(t *testing.T) {
	tree := BinaryTree{}
	if got := tree.Insert(10); !got {
		t.Error("Could not insert value 10")
	}
	if got := tree.Insert(20); !got {
		t.Error("Coudl not insert value 20")
	}

	newNode := tree.root.right
	if newNode == nil || newNode.Value != 20 || newNode.parent != tree.root {
		t.Error("New node is not correct")
	}
}

func TestInsertLeftWithOnlyRoot(t *testing.T) {
	tree := BinaryTree{}
	if got := tree.Insert(10); !got {
		t.Error("Could not insert value 10")
	}
	if got := tree.Insert(5); !got {
		t.Error("Coudl not insert value 5")
	}

	newNode := tree.root.left
	if newNode == nil || newNode.Value != 5 || newNode.parent != tree.root {
		t.Error("New node is not correct")
	}
}

func TestInsertLeftWithMultipleLevel(t *testing.T) {
	tree := BinaryTree{}
	if got := tree.Insert(10); !got {
		t.Error("Could not insert value 10")
	}
	if got := tree.Insert(5); !got {
		t.Error("Could not insert value 20")
	}
	if got := tree.Insert(3); !got {
		t.Error("Could not insert value 3")
	}

	newNodeParent := tree.root.left
	newNode := tree.root.left.left
	if newNode == nil || newNode.Value != 3 || newNode.parent != newNodeParent {
		t.Error("Error inserting new node")
	}
}

func TestInsertRightWithMultipleLevel(t *testing.T) {
	tree := BinaryTree{}
	if got := tree.Insert(10); !got {
		t.Error("Could not insert value 10")
	}
	if got := tree.Insert(15); !got {
		t.Error("Could not insert value 20")
	}
	if got := tree.Insert(18); !got {
		t.Error("Could not insert value 7")
	}

	newNodeParent := tree.root.right
	newNode := tree.root.right.right
	if newNode == nil || newNode.Value != 18 || newNode.parent != newNodeParent {
		t.Error("Error inserting new node")
	}
}

func TestInsertNodeAlreadyExists(t *testing.T) {
	tree := BinaryTree{}
	if got := tree.Insert(10); !got {
		t.Error("Could not insert value 10")
	}
	if got := tree.Insert(10); got {
		t.Error("Accidentally insert value 10")
	}
}


func TestSearchRightSubtree(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(25)

	if got := tree.Search(25); got == nil {
		t.Error("Could not find node")
	}
}

func TestSearchLeftSubtree(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(25)

	if got := tree.Search(15); got == nil {
		t.Error("Could not find node")
	}
}

func TestSearchRoot(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)

	if got := tree.Search(10); got == nil {
		t.Error("Could not find node")
	}
}

func TestSearchNonExist(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(25)

	if got := tree.Search(45); got != nil {
		t.Error("Could not find node")
	}	
}

func TestRemoveRootNoChildren(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)
	if removed := tree.Remove(10); !removed {
		t.Error("Could not remove root")
	}
	
	if tree.root != nil {
		t.Error("Tree root not deleted")
	}
}

func TestRemoveRootWithOneChild(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)
	tree.Insert(20)
	inserted := tree.root.right

	if removed := tree.Remove(10); !removed {
		t.Error("Could not remove root")
	}

	if tree.root != inserted && tree.root.Value != 20 {
		t.Error("New root is not correct")
	}
}

func TestRemoveRootWithTwoChildren(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)

	insertedLeft := tree.root.left
	insertedRight := tree.root.right

	if removed := tree.Remove(10); !removed {
		t.Error("Could not remove root")
	}

	if tree.root != insertedRight && tree.root.Value != 20 && tree.root.left != insertedLeft && tree.root.left.Value != 5 {
		t.Error("New root is not correct")
	}
}

func TestRemoveLeafNode(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)

	tree.Insert(5)
		tree.Insert(7)
		tree.Insert(3)

	tree.Insert(20)
		tree.Insert(12)
		tree.Insert(25)

	if removed := tree.Remove(25); !removed {
		t.Error("Could not remove root")
	}
}

func TestRemoveNotFound(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)

	tree.Insert(12)
	tree.Insert(25)

	if removed := tree.Remove(56); removed {
		t.Error("Removed non-existant node")
	}
}

func TestRemoveInternalNodeTwoChildren(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)

	tree.Insert(5)
		tree.Insert(7)
		tree.Insert(3)

	tree.Insert(20)

		tree.Insert(12)
			tree.Insert(11)
			tree.Insert(13)

		tree.Insert(25)
			tree.Insert(22)
			tree.Insert(27)

	if removed := tree.Remove(20); !removed {
		t.Error("Could not remove internal node")
	}

	replacement := tree.root.right
	if replacement.left.parent != replacement || replacement.right.parent != replacement{
		t.Error("Parent node incorrect on child")
	}
	if replacement.Value != 22 {
		t.Error("replaced with wrong node")
	}
	if replacement.parent != tree.root {
		t.Error("parent node not set")
	}
}

func TestRemoveInteriorNodeWithOneChild(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(10)

	tree.Insert(5)
		tree.Insert(7)
		tree.Insert(3)

	tree.Insert(20)

		tree.Insert(12)		 // <- removing this node
			tree.Insert(11) 

		tree.Insert(25)
			tree.Insert(27)
	
	if removed := tree.Remove(12); !removed {
		t.Error("could not remove existing node")
	}

	if tree.root.right.left.Value != 11 {
		t.Error("Incorrect value in removed replacement")
	}

	if removed := tree.Remove(25); !removed {
		t.Error("could not remove existing node")
	}

	if tree.root.right.right.Value != 27 {
		t.Error("Incorrect value in removed replacement")
	}
}