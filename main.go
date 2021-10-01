package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Use the main to run the functions created below!
	thisIsaLocalFunction()
	ThisIsaGlobalFunction()
	numericTypes()
	loopTypes()
	arraysTypes()
	sliceExample()
	ret := sumFunction(1, 1)
	if ret == 2 {
		go thisIsaGoroutine()
	} else {
		fmt.Println("The Goroutine will not be called")
	}

	// The program will not close until you press CTRL-C
	<-c
}

func thisIsaLocalFunction() {
	fmt.Println("This function lives only in the Main package!")
	fmt.Print("The same applies to variables!")
}

func ThisIsaGlobalFunction() {
	fmt.Println("This function is God like!")
	fmt.Print("The same applies to variables!")
}

func thisIsaGoroutine() {
	for {
		fmt.Println("Hi from my first Go Routine!")
	}
}

func numericTypes() {
	// You can comment out your codes like C o//
	/*										  */
	var Integer int = 127
	var Integer64 int64 = 150000000000000
	fmt.Println("Integer Value:", Integer, "Integer64 Value:", Integer64)

	var Float32 float32 = 0.12345
	var Float64 float64 = 0.0000000000003
	fmt.Println("Float32 Value:", Float32, "Floatr64 Value:", Float64)

	// Look out! This := operator is very useful
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3:", c3)
	fmt.Printf("Type of c3: %T\n", c3)

	cZero := c3 - c3
	fmt.Println("cZero:", cZero)

}

func loopTypes() {
	// Regular For loop
	for i := 0; i < 100; i++ {
	}

	// "While" loop
	/*
		for {
			fmt.Println("I'm a while loop O.o")
		}
	*/

	// "Do - While" loop
	/*
		var AnExpression bool
		for ok := true; ok; ok = AnExpression  {
			fmt.Println("I will run this until we reach false!")
		}
	*/

	// Many examples using the previous structures
	// Regular For
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, " ")
	}

	// "While"
	fmt.Println()
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()

	// "Do - While"
	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}

		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// This is a bonus! Iterating over an Array
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}
}

func arraysTypes() {
	anArray := [4]int{1, 2, 3, 4}
	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12},
		{13, 14, 15, 16}}
	threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}
	fmt.Println("The length of", anArray, "is", len(anArray))
	fmt.Println("The first element of", twoD, "is", twoD[0][0])
	fmt.Println("The length of", threeD, "is", len(threeD))

	for i := 0; i < len(threeD); i++ {
		v := threeD[i]
		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
		}
		fmt.Println()
	}

	fmt.Println("The length of", anArray, "is", len(anArray))
	fmt.Println("The first element of", twoD, "is", twoD[0][0])
	fmt.Println("The length of", threeD, "is", len(threeD))

	for i := 0; i < len(threeD); i++ {
		v := threeD[i]
		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
		}

		fmt.Println()
	}
	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
		fmt.Println()
	}
}

func sliceExample() {
	// This will create an empty slice s1 of type int with size 5
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)

	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice)

	fmt.Println("The basic is above! The rest is more advanced!")

	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)

	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -2)
	aSlice = append(aSlice, -3)
	aSlice = append(aSlice, -4)
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

	fmt.Println("That's all! We've covered a lot of basics of Go. Of course there still ")

}

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func sumFunction(a int, b int) int {
	var c int
	c = a + b
	return c
}
