package main

import (
	"fmt"
)

func main() {
	var i any

	i = 7
	fmt.Println(i)

	i = "hi"
	fmt.Println(i)

	s := i.(string)
	fmt.Println("s:", s)

	n, ok := i.(int)
	if ok {
		fmt.Println("n:", n)
	} else {
		fmt.Println("not an int")
	}

	switch i.(type) {
	case int:
		fmt.Println("an int")
	case string:
		fmt.Println("a string")
	default:
		fmt.Printf("unknwon type: %T\n", i)
	}

	fmt.Println(max([]int{3, 1, 2}))
	fmt.Println(max([]float64{3, 1, 2}))

}

type Number interface {
	int | float64
}

func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max

}
