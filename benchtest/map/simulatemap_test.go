package benchtest

import "testing"

func BenchmarkSimulateMapInterface(b *testing.B) {
	m := newInterface()
	for i := 0; i < b.N; i++ {
		simulateMapInterface(&m, i, i)
	}
}

func BenchmarkSimulateMapManual(b *testing.B) {
	m := newManual()
	for i := 0; i < b.N; i++ {
		simulateMapManual(&m, i, i)
	}
}
