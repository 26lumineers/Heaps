package main

import "fmt"

type MaxHeap struct {
	heapList []int
}

// add an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.heapList = append(h.heapList, key)
	//rearrange
	h.maxHeapifyUp(len(h.heapList) - 1)
}

// Extract return the largest key ,and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.heapList[0]
	l := len(h.heapList) - 1
	if len(h.heapList) == 0 {
		fmt.Println("cant extract because slice length is 0")
		return -1
	}
	//take out the last index and put it into the root
	h.heapList[0] = h.heapList[l]
	h.heapList = h.heapList[:l]
	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.heapList[parent(index)] < h.heapList[index] {
		h.swap(parent(index), index)
		index = parent(index)
		fmt.Println("index after swap : ", index)
	}
}
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.heapList) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	//loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { //when left child is the only one child
			childToCompare = l
		} else if h.heapList[l] > h.heapList[r] { //when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		//compare slice value of current index t o larger child and swap if smaller
		if h.heapList[index] < h.heapList[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// get parent index
func parent(i int) int {
	// fmt.Println("parent index :", (i-1)/2)
	return (i - 1) / 2
}

// get left child index
// left always odd number
func left(i int) int {
	return i*2 + 1
}

// get right child index
// right always even number
func right(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.heapList[i1], h.heapList[i2] = h.heapList[i2], h.heapList[i1]
}
func main() {
	m := &MaxHeap{}

	listNumber := []int{10, 20, 30, 60, 40}
	for _, v := range listNumber {
		m.Insert(v)
		// fmt.Println("current index ::", idx)
		fmt.Println(m)
	}
	for _, v := range listNumber {
		_ = v
		m.Extract()
		fmt.Println(m)
	}
}
