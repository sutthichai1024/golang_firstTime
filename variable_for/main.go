package main

import "fmt"

func init() {
	fmt.Println("Init ...")
}

func main() {
	fmt.Println("#2 start ...")

	// declare go type 1
	// var vTBoolean bool = true
	// var vTInteger int = 5
	// var vTFloat float64 = 1.25
	// var vTString string = "Hello Go #2"
	// var vTFunction func() = func() {
	// 	fmt.Println("In Function")
	// }

	// declare variable go type 2
	// var vTBoolean = true
	// var vTInteger = 5

	// // declare variable go type 3
	// vTFloat := 1.25
	// vTString := "Hello Go #2"
	// vTFunction := func() {
	// 	fmt.Println("In Function")
	// }

	//default variable and not assign data
	// var variableTypeBoolean bool    //false
	// var variableTypeInteger int     //0
	// var variableTypeFloat float64   //0
	// var variableTypeString string   //
	// var variableTypeFunction func() // invalid memory address or nil pointer dereference

	// fmt.Println(variableTypeBoolean)
	// fmt.Println(variableTypeInteger)
	// fmt.Println(variableTypeFloat)
	// fmt.Println(variableTypeString)
	// variableTypeFunction()

	//Const
	// const x int = 4
	// const y int = 5
	// const z int = x * y
	// fmt.Println(z)

	// const text string = "Hi"

	// fmt.Println(text)

	//for in go
	for i := 0; i <= 100; i++ {
		if i == 5 {
			continue
		} else if i == 10 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("--------------------")
	//while in go
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("--------------------")
	for {
		fmt.Println(i)
		break
	}

	//do while in go
	for ok := true; ok; ok = i > 0 {
		fmt.Println(i)
		i--
	}
	fmt.Println("--------------------")
	j := 0
	for {
		fmt.Println(j)

		if !(j > 0) {
			break
		}
	}

}
