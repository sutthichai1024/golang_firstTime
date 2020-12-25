package main

import "fmt"

func main() {
	fmt.Println("#3 start ...")
	fmt.Println("----------------------------")

	// Condition (IF Else)
	// if 10-5 == 5 {
	// 	fmt.Println("Condition is true.")
	// } else {
	// 	fmt.Println("Condition is false.")
	// }

	// if x := 10; x <= 10 {
	// 	fmt.Println("Good job !!")
	// } else if x <= 7 {
	// 	fmt.Println("Cool !!")
	// } else {
	// 	fmt.Println("Awsome !!")
	// }
	// fmt.Println("------------------------")

	// switch case
	// i := 2
	// fmt.Println("write", i, " as ")
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// }

	// j := 3
	// switch j {
	// case 3:
	// 	fmt.Println("three")
	// 	fallthrough
	// case 2:
	// 	fmt.Println("two")
	// 	fallthrough
	// case 1:
	// 	fmt.Println("one")
	// default:
	// 	fmt.Println("GO !!!")
	// }

	//Array in go
	// var a [3]int
	// b := [5]int{5, 2, 1, 4, 6}

	// a[0] = 1
	// a[1] = 2
	// a[2] = 3

	// fmt.Println("Length of an array a: ", len(a))
	// fmt.Println("Value of an array a: ", a[0], a[1], a[2])
	// fmt.Println("Value of an array b: ", b)

	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2D array: ", twoD)

	// fmt.Println("----------------------")

	//Slices in Go
	// a := []string{"a", "b", "c"}
	// b := make([]string, len(a))
	// fmt.Println("value of slices b: ", b)

	// copy(b, a)
	// b = append(b, "d", "e")

	// fmt.Println("value of slices a: ", a)
	// fmt.Println("value of slices b: ", b)

	// c := b[0:len(b)]
	// fmt.Println("value of slices c: ", c)
	// fmt.Println("------- Mutate value of c -----")

	// c[0] = "x"

	// fmt.Println("Value of slice c : ", c)
	// fmt.Println("Value of slice b : ", b)

	// fmt.Println("----------------------")

	// a := []string{"a", "b", "c", "d", "e"}
	// b := a[0:5]
	// c := a[:5]
	// d := a[0:]
	// e := a[:]

	// fmt.Println("Value of slice a : ", a)
	// fmt.Println("Value of slice b : ", b)
	// fmt.Println("Value of slice c : ", c)
	// fmt.Println("Value of slice d : ", d)
	// fmt.Println("Value of slice e : ", e)

	// fmt.Println("----------------------")
	// x := a[2:4]
	// fmt.Println("Value of slice x ", x)

	// fmt.Println("----------------------")

	//maps

	// m := make(map[string]int)
	// m["math"] = 85
	// m["sci"] = 86
	// m["com"] = 28

	// fmt.Println("Length of maps m : ", len(m))
	// fmt.Println("Value of maps m : ", m)

	// fmt.Println("----------------------")

	// m := map[string]int{"com": 28, "math": 85, "sci": 86}
	// fmt.Println("Value of maps m : ", m)
	// delete(m, "math")
	// fmt.Println("Delete math from map m . ")
	// fmt.Println("Value of maps m : ", m)

	// value, prs := m["math"]
	// fmt.Println("State of math : ", prs)
	// fmt.Println("Value of maps m : ", value)

	// fmt.Println("----------------------")
	// if value, prs := m["math"]; prs {
	// 	fmt.Println("Values of math : ", value)
	// }
	// if value, prs := m["com"]; prs {
	// 	fmt.Println("Value of com: ", value)
	// }

	//Range
	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Value of a,s,m: ", a, b, m)

	fmt.Println("----------------------")
	for i, num := range a {
		fmt.Println(i, num)
	}

	fmt.Println("----------------------")
	for i, num := range b {
		fmt.Println(i, num)
	}

	fmt.Println("----------------------")
	for i, num := range m {
		fmt.Println(i, num)
	}

	fmt.Println("----------------------")

}
