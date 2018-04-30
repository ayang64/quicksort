package quicksort

func NaivePartition(a []int, lo, hi int) int {
	pivot := a[hi]

	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	i++
	a[i], a[hi] = a[hi], a[i]
	return i
}

func NaiveSort(a []int) {
	naiveSort(a, 0, len(a)-1)
}

func naiveSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	p := NaivePartition(a, lo, hi)
	naiveSort(a, lo, p-1)
	naiveSort(a, p+1, hi)
}
