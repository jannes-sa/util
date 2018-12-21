package benchtest

func newManual() map[int]int {
	return make(map[int]int, 0)
}

func newInterface() map[interface{}]interface{} {
	return make(map[interface{}]interface{}, 0)
}

func simulateMapManual(m *map[int]int, key int, val int) {
	(*m)[key] = val
}

func simulateMapInterface(m *map[interface{}]interface{}, key interface{}, val interface{}) {
	(*m)[key] = val
}
