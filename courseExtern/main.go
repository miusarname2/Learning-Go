package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hola, GO")

	//variable

	var myString string = "abc"

	fmt.Println(myString)

	myString = "abc"

	// Error myString = 6

	var myInt = 7

	fmt.Println(myInt)

	var myNInt int = 12

	fmt.Println(myNInt)

	fmt.Println(myNInt - 1)

	// concatenar

	fmt.Printf("%s %d", myString, myNInt)

	fmt.Println(reflect.TypeOf(myInt))

	var bool bool

	bool = true

	fmt.Println(bool)

	motoMoto := "HOla mundo"

	fmt.Println(motoMoto)

	//constante

	const pi = 3.14159

	fmt.Println(pi)

	if myInt == 10 {
		fmt.Println(":(")
	} else if myInt == 11 || myString != "hola" {
		fmt.Println(":(1")
	} else {
		fmt.Println(":)")
	}

	// Arrs

	var myArr [3]int
	myArr[0] = 1
	myArr[1] = 2
	myArr[2] = 3

	fmt.Println(myArr)
	fmt.Println(myArr[1])

	// Map
	myMap := make(map[string]int)
	myMap["a"] = 1

	fmt.Println(myMap)

	//list

	myList := list.New()

	myList.PushBack(1)

	fmt.Println(myList.Back())

}
