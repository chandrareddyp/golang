package main

import "fmt"

type Person struct{
	Name string
	Age int
}

func main() {
	p := new(Person)
	fmt.Println(p)
	//p2 := make(Person)
	//fmt.Println(p2)
	p3 := &Person{}
	fmt.Println(p3)
	m := make(map[string]int)
	fmt.Println(&m)
	mn := new(map[string]int)
	fmt.Println(mn)

	fmt.Println(mn)
	myMap := map[string]string{"name": "Bob"}
	  newMap := map[string]string(myMap) // Copy existing map
	newMap["age"] = "25"
	fmt.Println("newMap:", newMap)
	
}