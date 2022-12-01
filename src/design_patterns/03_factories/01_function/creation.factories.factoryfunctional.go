package main

import "fmt"

type Person struct {
	Name string
	Age  int
	// default value is 2
	EyeCount int
}

// factory function
// func NewPerson(name string,age int) Person {
// 	return Person{name,age}
// }

func NewPerson(name string, age int) *Person {
	// validate
	if age < 16 {
		fmt.Println("不可以色色")
		return nil
	}
	return &Person{name, age, 2}
}

func main() {
	// initialize directly
	// p := Person{"John", 22}
	// fmt.Println(p)

	// use a constructor
	p2 := NewPerson("Jane", 21)
	p2.Age = 30
	fmt.Println(p2)
}
