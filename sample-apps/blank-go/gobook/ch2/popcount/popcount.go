package popcount

import "fmt"

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	fmt.Println("Lookup table initialised.")
}

// Shift uses bit shift instead of lookup
func Shift(x uint64) int {
	r := 0
	for i := 0; i < 63; i++ {
		r += int((x >> i) & 1)
	}

	return r
}

// Count returns the popultation count  (number of set bits) of x
func Count(x uint64) int {
	r := 0
	for i := 0; i <= 7; i++ {
		r += int(pc[byte(x>>(i*8))])
	}

	return r
}

// Precalc uses look-up table that was in init method
func Precalc(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
