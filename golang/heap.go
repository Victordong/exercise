package golang

func partBigHeap(arr []int, begin int, end int) {
	father, son := begin, begin*2+1
	for son <= end {
		if son+1 <= end && arr[son] < arr[son+1] {
			son = son + 1
		}
		if arr[father] >= arr[son] {
			return
		} else {
			arr[father], arr[son] = arr[son], arr[father]
			father, son = son, son*2+1
		}
	}
}

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	n := len(arr)
	if n <= k {
		return arr
	}
	heap := make([]int, k)
	copy(heap, arr)
	for i := k/2 - 1; i >= 0; i++ {
		partBigHeap(heap, i, k-1)
	}
	for i := k; i < n; i++ {
		if heap[0] <= arr[i] {
			continue
		} else {
			heap[0] = arr[i]
			partBigHeap(heap, 0, k-1)
		}
	}
	return heap
}
