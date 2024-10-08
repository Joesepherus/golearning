package minHeapTuple

import (
	"fmt"
)

var Heap = []HeapNode{}

// Run demonstrates how to use the Priority queue.
func Run() {
	fmt.Println(Heap)

	Heap = Insert(HeapNode{2, "B"}, Heap)
	Heap = Insert(HeapNode{1, "C"}, Heap)
	fmt.Println("Heap after Insertion:", Heap)

	Heap = Remove(Heap)
	fmt.Println("Heap after removing the root:", Heap)
}

// HeapNode represents a node in the Priority queue with a Priority and Value.
type HeapNode struct {
	Priority int
	Value    string
}

// Insert adds a new element to the Priority queue (min-Heap).
func Insert(node HeapNode, Heap []HeapNode) []HeapNode {
	fmt.Println("Inserting node:", node)
	Heap = append(Heap, node)
	index := len(Heap) - 1

	// Bubble up to maintain Heap property
	for index > 0 {
		parentIndex := (index - 1) / 2
		if Heap[index].Priority >= Heap[parentIndex].Priority {
			break
		}
		// Swap the child node with its parent
		Heap[index], Heap[parentIndex] = Heap[parentIndex], Heap[index]
		index = parentIndex
	}

	return Heap
}

// Remove Removes the root element from the Priority queue (min-Heap) and maintains the Heap property.
func Remove(Heap []HeapNode) []HeapNode {
	if len(Heap) == 0 {
		fmt.Println("Heap is empty, nothing to Remove.")
		return Heap
	}

	fmt.Println("Removing root node:", Heap[0])
	lastIndex := len(Heap) - 1
	Heap[0] = Heap[lastIndex] // Move the last element to the root
	Heap = Heap[:lastIndex]   // Remove the last element
	lastIndex--               // Adjust the last index
	index := 0

	// Bubble down to maintain Heap property
	for index < lastIndex {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		smallerChildIndex := leftChildIndex

		// Check if right child exists and is smaller than the left child
		if rightChildIndex <= lastIndex && Heap[rightChildIndex].Priority < Heap[leftChildIndex].Priority {
			smallerChildIndex = rightChildIndex
		}

		// If the current node is smaller or no children exist, stop bubbling down
		if smallerChildIndex > lastIndex || Heap[index].Priority <= Heap[smallerChildIndex].Priority {
			break
		}

		// Swap the node with its smaller child
		Heap[index], Heap[smallerChildIndex] = Heap[smallerChildIndex], Heap[index]
		index = smallerChildIndex
	}

	return Heap
}
