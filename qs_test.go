package quicksort

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkSorts(b *testing.B) {
	sorters := []struct {
		Name string
		Func func([]int)
	}{
		{Name: "Naive", Func: NaiveSort},
		{Name: "Slice", Func: SliceSort},
		{Name: "Concurrent", Func: ConcurSort},
	}
	// generate random data
	randomSlice := func(c int) []int {
		rc := make([]int, c, c)
		for i := 0; i < c; i++ {
			rc[i] = rand.Intn(c)
		}
		return rc
	}

	for i := 2; i < (1 << 25); i = i << 1 {
		nums := randomSlice(i)
		td := make([]int, len(nums), len(nums))
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			for _, method := range sorters {
				b.Run(method.Name, func(b *testing.B) {
					copy(td, nums)
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						method.Func(td)
					}
				})
			}
		})
	}
}

func TestConcurSort(t *testing.T) {
	d := []int{12, 5, 8, 2, 10, 4, 6, 1, 11, 9, 3, 7}

	t.Logf("before sort: d = %#v", d)
	ConcurSort(d)
	t.Logf("after sort: d = %#v", d)
}

func TestSliceSort(t *testing.T) {
	d := []int{12, 5, 8, 2, 10, 4, 6, 1, 11, 9, 3, 7}

	t.Logf("before sort: d = %#v", d)
	SliceSort(d)
	t.Logf("after sort: d = %#v", d)
}

func TestNaieveSort(t *testing.T) {
	d := []int{12, 5, 8, 2, 10, 4, 6, 1, 11, 9, 3, 7}

	t.Logf("before sort: d = %#v", d)
	NaiveSort(d)
	t.Logf("after sort: d = %#v", d)
}
