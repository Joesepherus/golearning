package minHeap

import "fmt"

func Run() {
	var heap = []int{1, 3, 4, 5, 6}
	fmt.Println(heap)
	heap = insert(2, heap)
	fmt.Println(heap)
	remove(heap)
	fmt.Println(heap)

}

func insert(value int, heap []int) []int {
	fmt.Println("Inserting value: ", value)
	var newHeap = append(heap, value)
	var index = len(newHeap) - 1
	for index > 0 {

		parentIndex := (index - 1) / 2
		fmt.Println("parentIndex: ", parentIndex)
		if value > newHeap[parentIndex] {
			break
		}
		temp := newHeap[parentIndex]
		newHeap[parentIndex] = value
		newHeap[index] = temp
		index = parentIndex
	}

	return newHeap
}

func remove(heap []int) []int {
	fmt.Println("Removing first element")
	var newHeap = heap
	var lastIndex = len(newHeap) - 1
	newHeap[0] = newHeap[lastIndex]
	newHeap = newHeap[:lastIndex]
	lastIndex--
	index := 0
	value := newHeap[index]

	for index < lastIndex {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		smallerChildIndex := leftChildIndex

		if rightChildIndex <= lastIndex && newHeap[rightChildIndex] < newHeap[leftChildIndex] {
			smallerChildIndex = rightChildIndex

		}

		if smallerChildIndex > lastIndex || newHeap[index] <= newHeap[smallerChildIndex] {
			break
		}

		newHeap[index] = newHeap[smallerChildIndex]
		newHeap[smallerChildIndex] = value
		index = smallerChildIndex
	}
	return newHeap
}
