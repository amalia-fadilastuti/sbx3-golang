package sum_test

import (
	"fmt"

	"github.com/amalia-fadilastuti/sbx3-golang-level4/sum"
)

func ExampleInts() {
	fmt.Println(sum.Ints([]int{1, 2, 3, 4, 5}...))
	// Output:
	// 15
}
