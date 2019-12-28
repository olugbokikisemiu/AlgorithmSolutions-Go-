package main

import "testing"

func BenchmarkSortAsc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortAsc([]int{20, 4, 12, 1, 5, 3, 19})
	}
}
