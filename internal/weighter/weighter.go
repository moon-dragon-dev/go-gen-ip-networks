package weighter

func binarySearch(weights []uint32, n uint32) int {
	low := 0
	high := len(weights) - 1
	for low <= high {
		mid := (low + high) / 2
		if weights[mid] < n {
			low = mid + 1
		} else if weights[mid] > n {
			high = mid - 1
		} else {
			return mid
		}
	}
	return low
}

func CreateSelector(weights []uint32) func(uint32) int {
	acc := make([]uint32, len(weights))
	sum := uint32(0)

	for i, weight := range weights {
		sum += weight
		acc[i] = sum - 1
	}

	return func(n uint32) int {
		return binarySearch(acc, n%sum)
	}
}
