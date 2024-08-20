package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string
	var age int
	lang := "go"

	fmt.Scanf("%s", &name) //for scan - format specifier of the input in "", and the variable assignment using &
	fmt.Scanf("%d", &age)
	fmt.Printf("Hello, %s ", name) //while printing too we have to specify the type of text and then a , and the variable name without anything (, name)
	fmt.Printf("your age is: %d! \n", age)
	fmt.Printf("This is %s \n", lang)
	fmt.Printf("Type is %s", reflect.TypeOf(name)) //this kind of reflect.Typeof() is used to print the type of variable

}
