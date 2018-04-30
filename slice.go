package quicksort

func SlicePartition(a []int) int {
	pivot := a[len(a)-1]

	i := -1
	for j := range a {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	i++
	a[i], a[len(a)-1] = a[len(a)-1], a[i]
	return i
}

func SliceSort(a []int) {
	if len(a) == 0 {
		return
	}

	p := SlicePartition(a)
	SliceSort(a[:p])
	SliceSort(a[p+1:])
}
