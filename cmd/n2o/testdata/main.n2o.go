// +build n2o

package main

func testUnroll() {
	sum := 0
	//n2o: unroll
	{
		_values := [3]int{1, 2, 3}
		sum += _values[0]
		sum += _values[1]
		sum += _values[2]
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
	{
		// inlined add1(v)
		x := v
		v = x + 1
	}
	{
		// inlined add1ifNotZero(v)
		x := v
		if x != 0 {
			v = x + 1
			goto _ret0
		}
		v = x
	_ret0:
	}
	v = add1ifNotZero(v) // Not inlined
	{
		// inlined addN(v, v)
		x := v
		n := v
		for i := 0; i < n; i++ {
			x++
		}
		v = x
	}
	v = addN(v, v) // Not inlined
	println(v)
}

func main() {
	testUnroll()
	testInline()
}
