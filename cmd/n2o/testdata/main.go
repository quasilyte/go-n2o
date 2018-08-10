// +build !n2o

package main

func testUnroll() {
	sum := 0
	//n2o: unroll
	for _, x := range [3]int{1, 2, 3} {
		sum += x
	}
	println(sum)
}

//go:noinline
func add1(x int) int { return x + 1 }

//go:noinline
func add1ifNotZero(x int) int {
	if x != 0 {
		return x + 1
	}
	return x
}

func addN(x, n int) int {
	// Do it with a loop to make function non-inlineable.
	for i := 0; i < n; i++ {
		x++
	}
	return x
}

func testInline() {
	var v = 10
	//n2o: inline
	v = add1(v)
	v = add1ifNotZero(v) //n2o: inline
	v = add1ifNotZero(v) // Not inlined
	v = addN(v, v)       //n2o: inline
	v = addN(v, v)       // Not inlined
	println(v)
}

func main() {
	testUnroll()
	testInline()
}
