package main

import (
	"fmt"

	"github.com/poethera/gotestcom/pkg/test"
)

func main() {
	fmt.Printf("module import test : 22 == %d\n", test.MathSum(11, 11))
}
