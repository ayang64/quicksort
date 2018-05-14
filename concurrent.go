package quicksort

import (
	"sync"
)

func ConcurPartition(a []int) int {
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

func ConcurSort(a []int) {
	w := sync.WaitGroup{}
	w.Add(1)
	go concurSort(a, &w)
	w.Wait()
}

func concurSort(a []int, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(a) == 0 {
		return
	}

	p := ConcurPartition(a)

	w := sync.WaitGroup{}
	w.Add(2)
	go concurSort(a[:p], &w)
	go concurSort(a[p+1:], &w)
}
