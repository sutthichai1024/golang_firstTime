package main

import "fmt"

// func plus(a int, b int) int {
// 	return a + b
// }

// func prism(length, width, height int) int {
// 	return length * width * height
// }

// func findFirstAndLast(msg string) (string, string) {
// 	return string(msg[0]), string(msg[len(msg)-1])
// }

// func swap(x, y int) (a, b int) {
// 	a = x
// 	b = y
// 	return
// }

// func sum(nums ...int) {
// 	fmt.Println(nums, " = ")
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	fmt.Println(total)
// }

// func intseq() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func factUsingfor(n int) int {
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial
}

func main() {
	fmt.Println("#4 start ...")
	fmt.Println("----------------------------")

	// a, b := 1, 2
	// res := plus(a, b)
	// fmt.Println("1 + 2 = ", res)

	// res = prism(4, 2, 2)
	// fmt.Println("Rectangular prism = ", res)
	// fmt.Println("----------------------------")

	// a, b := findFirstAndLast("Hello Gopher")
	// fmt.Println(a, b)

	// i, j := 0, 1
	// i, j = swap(i, j)
	// fmt.Println(i, j)
	// fmt.Println("----------------------------")

	// nums := []int{1, 2, 3, 4, 5}

	// sum(1, 2)
	// sum(1, 2, 3)
	// sum(nums[0:4]...)
	// sum(nums...)

	// fmt.Println("----------------------------")

	// nextInt := intseq()

	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// fmt.Println("----------------------------")
	// newInts := intseq()
	// fmt.Println(newInts())

	// fmt.Println("----------------------------")

	fmt.Println(fact(10))
	fmt.Println(factUsingfor(10))

}
