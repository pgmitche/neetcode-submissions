type KthLargest struct {
	// implement a min heap, where if len(minHeap) is greater than K
	// the min (root) is removed
	pq *pq
	k int
}


func Constructor(k int, nums []int) KthLargest {
    // initialise
	kth := KthLargest{
		pq: &pq{minHeap: []int{0}},
		k: k,
	}

	for _, num := range nums {
		kth.pq.Enqueue(num)
	}

	for len(kth.pq.minHeap)-1 > k {
		kth.pq.Dequeue()
	}

	return kth
}


func (this *KthLargest) Add(val int) int {
    this.pq.Enqueue(val)
	if len(this.pq.minHeap)-1 > this.k {
		this.pq.Dequeue()
	}

	return this.pq.Peek()
}

type pq struct {
    minHeap []int
}

func (pq *pq) Enqueue(val int) {
	// push to the last space in the array
	pq.minHeap = append(pq.minHeap, val)

	// location of entered value
    cursor := len(pq.minHeap) - 1

	// track the cursor of the new value as it percolates up,
	// comparing it to it's parent node (cursor/2)
	for cursor > 1 && pq.minHeap[cursor] < pq.minHeap[cursor/2] {
		// temp the current value
		temp := pq.minHeap[cursor]
		// overwrite it's position with the parent
		pq.minHeap[cursor] = pq.minHeap[cursor/2]
		// then write it to the parent's position
		pq.minHeap[cursor/2] = temp
		// move the cursor to the parent's position
		cursor = cursor/2
	}
}

func (pq *pq) Dequeue() int {
	if len(pq.minHeap) == 1 {
		return -1
	}

	if len(pq.minHeap) == 2 {
		var val int
		val, pq.minHeap = pq.minHeap[len(pq.minHeap)-1], pq.minHeap[:len(pq.minHeap)-1]
		return val
	}

	// take the root value
	res := pq.minHeap[1] 
	// replace with the last value
	pq.minHeap[1], pq.minHeap = pq.minHeap[len(pq.minHeap)-1], pq.minHeap[:len(pq.minHeap)-1]
	// and percolate it down, comparing with it's children
	// where lhs child is (2*i) and rhs is (2*i)+1
	root := 1
	for 2*root < len(pq.minHeap) {
		lhs := 2*root
		rhs := 2*root+1
		// if the right child is smaller than the left child
		// and the right child is smaller than the val
		if rhs < len(pq.minHeap) && 
			pq.minHeap[rhs] < pq.minHeap[lhs] && 
			pq.minHeap[rhs] < pq.minHeap[root] {
			
			temp := pq.minHeap[root]
			pq.minHeap[root] = pq.minHeap[rhs]
			pq.minHeap[rhs] = temp

			root = rhs
		} else if pq.minHeap[root] > pq.minHeap[lhs] {
			// if the left child is smaller then shift it
			temp := pq.minHeap[root]
			pq.minHeap[root] = pq.minHeap[lhs]
			pq.minHeap[lhs] = temp
			root = lhs
		} else {
			// escape
			break
		}
	}
	return res
}

func (pq *pq) Peek() int {
	return pq.minHeap[1]
}