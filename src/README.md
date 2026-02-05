# Go

## Starting a new project

Initialise using the `go` CLI, e.g. `go mod init hello`. This creates a `go.mod` file.

## Basic project types

Executables are delineated from libraries by the `package main` keyword.

## Basic types

* `bool`
* `string`
* `int  int8  int16  int32  int64`
* `uint uint8 uint16 uint32 uint64 uintptr`
* `byte` // alias for uint8
* `rune` // alias for int32, represents a Unicode code point
* `float32 float64`
* `complex64 complex128`

## Variables

Variables can be set at function or package level.

```
var something int` // zero initialisation
something := 1`. // type inferred
var i, j int = 1, 2`
var c, python, java = true, false, "no!"`
c, python, java := true, false, "no!" // alternative of the above
// "block" initialisation
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
const Pi = 3.14 // constant, cannot be declated with := syntax
// "block" constants
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
```

## Functions

```
func add(x int, y int) int {
	return x + y
}

// Shortened form of the above
func add2(x, y int) int {
	return x + y
}

// Can return multiple values
func swap(x, y string) (string, string) {
	return y, x
}

// With named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // "naked" return
}
```

## Type conversions

Assignment between items of different type requires an explicit conversion.

```
i := 42
f := float64(i)
u := uint(f)
```