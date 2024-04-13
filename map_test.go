package main_test

import "testing"

const mapSize = 1000

func BenchmarkMap1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		for i := 0; i < mapSize; i++ {
			m[i] = i
		}
	}
}

func BenchmarkMap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int, mapSize)
		for i := 0; i < mapSize; i++ {
			m[i] = i
		}
	}
}
