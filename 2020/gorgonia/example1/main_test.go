package main

import "testing"

func Test_add(t *testing.T) {
	if add(40, 2) != 42 {
		t.Fail()
	}
}

func Benchmark_add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(40, 2)
	}
}
