package binarySearchTree

import (
	"fmt"
	"golearning/queue"
	"log"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

var tree = &BST{}

var myQueue = &queue.Queue{}

func Run() {
	log.Println("Binary Search")
	// BST Example initialize nodes
	//       8
	//      /   \
	//     3     10
	//    / \      \
	//   1   6     14
	//      / \   /
	//     4  7  13
	root := &Node{value: 8}
	root.left = &Node{value: 3}
	root.right = &Node{value: 10}

	root.left.left = &Node{value: 1}
	root.left.right = &Node{value: 6}
	root.left.right.left = &Node{value: 4}
	root.left.right.right = &Node{value: 7}

	root.right.right = &Node{value: 14}
	root.right.right.left = &Node{value: 13}

	// Initialize the BST with the root
	tree.root = root

	// Example: print the root value
	fmt.Println("Root:", tree.root.value)
	fmt.Println("Left Traversal:")
	leftTraversal(tree.root)

	fmt.Println("Right Traversal:")
	rightTraversal(tree.root)

	fmt.Println("PreOrder Traversal:")
	preOrderTraversal(tree.root)

	fmt.Println("PostOrder Traversal:")
	postOrderTraversal(tree.root)

	fmt.Println("LevelOrder Traversal:")
	myQueue.Enqueue(tree.root)
	levelOrderTraversal()

	fmt.Println("Delete 1:")

	deleteNode(6, &tree.root)

	myQueue.Enqueue(tree.root)

	fmt.Println("levelOrderTraversal START")
	levelOrderTraversal()
	fmt.Println("levelOrderTraversal END")

	insert(&Node{value: 2}, tree.root)
	insert(&Node{value: 5}, tree.root)
	fmt.Println("Root:", tree.root.left.left.right)

	search2 := search(51, tree.root)
	fmt.Println("search2", search2)

}

func insert(node *Node, parent *Node) {
	fmt.Println("node.value, parent.value", node.value, parent.value)
	if node.value < parent.value {
		if parent.left != nil {
			insert(node, parent.left)
		} else {
			fmt.Println("Adding node to parent left", parent)
			parent.left = node
		}
	}
	if node.value > parent.value {
		if parent.right != nil {
			insert(node, parent.right)
		} else {
			fmt.Println("Adding node to parent right", parent)

			parent.right = node
		}
	}
}

func findMin(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

//	   8
//	  /   \
//	 3     10
//	/ \      \
//
// 1   6     14
//
//	 / \   /
//	4  7  13
//
// case one - no children:
// delete 1
//
//		   8
//		  /   \
//		 3     10
//		 \      \
//
//	      6     14
//
//		 / \   /
//		4  7  13
//
// case two - one children:
// delete 14
//
//	   8
//	  /   \
//	 3     10
//	/ \      \
//
// 1   6     13
//
//	 / \
//	4  7
//
// case three - two children:
// delete 6
//
//	   8
//	  /   \
//	 3     10
//	/ \      \
//
// 1   4     14
//
//	    \   /
//	_	7  13
//
// case four - delete root:
// delete 8
//
//		    3
//		  /   \
//		 1     10
//		 \      \
//
//	 	   4     14
//
//		    \   /
//		_	7  13
func delete(value int, node *Node) *Node {
	// If the tree is empty, return nil
	if node == nil {
		return nil
	}

	// Step 1: Search for the node to delete
	if value < node.value {
		node.left = delete(value, node.left) // Recur to the left subtree
	} else if value > node.value {
		node.right = delete(value, node.right) // Recur to the right subtree
	} else {
		// Node found
		// Case 1: No Children (Leaf Node)
		if node.left == nil && node.right == nil {
			return nil // Remove the node
		}
		// Case 2: One Child
		if node.left == nil {
			return node.right // Return the right child
		}
		if node.right == nil {
			return node.left // Return the left child
		}
		// Case 3: Two Children
		// Find the inorder successor (minimum in the right subtree)
		successor := findMin(node.left)
		node.value = successor.value                   // Replace value with successor's value
		node.left = delete(successor.value, node.left) // Delete the successor
	}
	return node // Return the (possibly updated) node
}

func findLeftMost(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return findLeftMost(node.left)
}

func reconnect(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	if node.right == nil {
		return node.left
	}
	var point *Node = findLeftMost(node.right)
	point.left = node.left

	return node.right
}

func deleteNode(value int, node **Node) {
	if node == nil {
		return
	}

	if value < (*node).value {
		deleteNode(value, &(*node).left)
	} else if value > (*node).value {
		deleteNode(value, &(*node).right)
	} else {
		*node = reconnect(*node)
	}
}

func search(value int, node *Node) *Node {
	fmt.Println("search value, node.value", value, node.value)
	if value == node.value {
		return node
	}
	if value < node.value {
		if node.left != nil {
			return search(value, node.left)
		}
	}
	if value > node.value {
		if node.right != nil {
			return search(value, node.right)
		}
	}
	return nil
}

//	   8
//	  /   \
//	 3     10
//	/ \      \
//
// 1   6     14
//
//	 / \   /
//	4  7  13
// Left, Root, Right
// also known are inorder traversal
// 1, 3, 4, 6, 7, 8, 10, 13, 14
// DFS search

func leftTraversal(parent *Node) {
	if parent.left != nil {
		leftTraversal(parent.left)
	}
	fmt.Println("value", parent.value)
	if parent.right != nil {
		leftTraversal(parent.right)
	}
}

//	   8
//	  /   \
//	 3     10
//	/ \      \
//
// 1   6     14
//
//	 / \   /
//	4  7  13
// Right, Root, Left
// 14, 13, 10, 8, 7, 6, 4, 3, 1
// DFS search

func rightTraversal(parent *Node) {
	if parent.right != nil {
		rightTraversal(parent.right)
	}
	fmt.Println("value", parent.value)
	if parent.left != nil {
		rightTraversal(parent.left)
	}
}

//	   8
//	  /   \
//	 3     10
//	/ \      \
//
// 1   6     14
//
//	 / \   /
//	4  7  13
// Root, Left, Right
//  8, 3, 1, 6, 4, 7, 10, 14, 13
// DFS search

func preOrderTraversal(parent *Node) {
	fmt.Println("value", parent.value)
	if parent.left != nil {
		preOrderTraversal(parent.left)
	}
	if parent.right != nil {
		preOrderTraversal(parent.right)
	}
}

//	   8
//	  /   \
//	 3     10
//	/ \      \
//
// 1   6     14
//
//	 / \   /
//	4  7  13
//
// Left, Right, Root
// 1, 4, 7, 6, 3, 13, 14, 10, 8
// DFS search
func postOrderTraversal(parent *Node) {
	if parent.left != nil {
		postOrderTraversal(parent.left)
	}
	if parent.right != nil {
		postOrderTraversal(parent.right)
	}
	fmt.Println("value", parent.value)
}

//	   8
//	  /   \
//	 3     10
//	/ \      \
//
// 1   6     14
//
//	 / \   /
//	4  7  13
//
// 8, 3, 10, 1, 6, 14, 4, 7, 13
// BFS search
func levelOrderTraversal() {
	queueItem, _ := myQueue.Dequeue()
	parent := queueItem.(*Node)
	fmt.Println("value", parent.value)
	if parent.left != nil {
		myQueue.Enqueue(parent.left)
	}
	if parent.right != nil {
		myQueue.Enqueue(parent.right)
	}
	if !myQueue.IsEmpty() {
		levelOrderTraversal()
	}
}
