# n2o: nitrous boost for Go

## Quick overview

You start by writing ordinary Go code.
Then you spot some execution path that requires optimizations, the hot one.

```go
func array8sum(xs *[8]int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}
```

Suppose you think that loop unrolling might help.
Instead of doing this:

```go
func array8sum(xs *[8]int) int {
	total := 0
	total += xs[0]
	total += xs[1]
	total += xs[2]
	total += xs[3]
	total += xs[4]
	total += xs[5]
	total += xs[6]
	total += xs[7]
	return total
}
```

You can do this:

```go
func array8sum(xs *[8]int) int {
	total := 0
	//opt: unroll
	for _, x := range xs {
		total += x
	}
	return total
}
```

And you get best of two worlds:
1. Readability of the initial version.
2. Performance of the optimized (unrolled) code.

There are other optimizations that can be performed by the `n2o`.
For example, you may ask to apply all sensible optimizations to the function:

```go
// Can also use func/space to optimize for code size.
//opt: func/speed
func array8sum(xs *[8]int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}
```

Read docs to know more about supported directives and their meaning.

Basically, you write Go and then improve performance by optimizer hints, instead of doing dirty work by yourself.
