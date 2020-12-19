package main

import (
	"fmt"

	"github.com/vtesin/aws-lambda-developer-guide/sample-apps/blank-go/gobook/ch2/popcount"
)

func init() {
	fmt.Println("Main init done.")
}

func main() {
	var x uint64 = 65535
	fmt.Println(popcount.Count(x))
	fmt.Println(popcount.Shift(x))
	fmt.Println(popcount.Precalc(x))
}
