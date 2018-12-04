package heaptest

// import "test/heaptest/pkg1"

// go build -gcflags '-m'
// go build -gcflags '-m -l'
func main() {
	var nm map[string]int
	nm["1"] = 1
	nm["2"] = 2
	MapCalc(&nm)

	v1 := nm["1"]
	v2 := nm["2"]
	add(&v1, &v2)
}

func add(x, y *int) {
	m := multply(x, y)
	n := mulWithoutPointer(*x, *y)
	o := multplyReturnPointer(x, y)
	*x = m + n + (*o)
}

func multply(x, y *int) int {
	return (*y) * (*y)
}

func mulWithoutPointer(x, y int) int {
	return x * y
}

func multplyReturnPointer(x, y *int) *int {
	z := (*x) * (*y)
	return &z
}

// func ValAddInter(x interface{}) int {
// 	return x.(int) + x.(int)
// }

func MapCalc(xMap *map[string]int) {
	(*xMap)["1"] = (*xMap)["1"] + (*xMap)["2"]
	(*xMap)["2"] = (*xMap)["1"] * (*xMap)["2"]
}

// func CheckHeapAdd(x *int) *int {
// 	return x
// }

// type S struct {
// 	M *int
// }

// func DoLeak() {
// 	var i int
// 	refStruct(&i)
// }

// func refStruct(y *int) (z S) {
// 	z.M = y
// 	return z
// }

// func escapeToHeap() *int {
// 	x := 12
// 	return &x
// }

// func noEscapeToHeap() int {
// 	x := 12
// 	return x
// }

// func printlnNotEscape(st string) {
// 	fmt.Println(st)
// }
