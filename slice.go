package quicksort

func SlicePartition(a []int) int {
	pivot := a[len(a)-1]
	split := 0

	for i := range a {
		if a[i] >= pivot {
			continue
		}
		a[split], a[i] = a[i], a[split]
		split++
	}

	a[split], a[len(a)-1] = a[len(a)-1], a[split]
	return split
}

func SliceSort(a []int) {
	if len(a) == 0 {
		return
	}

	p := SlicePartition(a)
	SliceSort(a[:p])
	SliceSort(a[p+1:])
}
