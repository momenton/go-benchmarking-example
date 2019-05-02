package main

import "testing"

func BenchmarkSelfConcatOperator1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		selfConcatOperator("test", 1000)
	}
}

func BenchmarkSelfConcatBuffer1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		selfConcatBuffer("test", 1000)
	}
}

func BenchmarkSelfConcatOperator100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		selfConcatOperator("test", 100000)
	}
}

func BenchmarkSelfConcatBuffer100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		selfConcatBuffer("test", 100000)
	}
}
