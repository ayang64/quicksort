package quicksort

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
	c := make(chan struct{}, 2)
	go concurSort(a, c)
	<-c
}

func concurSort(a []int, d chan<- struct{}) {
	if len(a) == 0 {
		d <- struct{}{}
		return
	}

	p := ConcurPartition(a)

	done := make(chan struct{})
	go concurSort(a[:p], done)
	go concurSort(a[p+1:], done)
	<-done
	<-done
	d <- struct{}{}
}
