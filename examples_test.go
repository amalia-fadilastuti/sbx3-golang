package sum

import "fmt"

func ExampleInts() {
	fmt.Println(Ints([]int{1, 2, 3, 4, 5}...))
	// Output:
	// 15
}
