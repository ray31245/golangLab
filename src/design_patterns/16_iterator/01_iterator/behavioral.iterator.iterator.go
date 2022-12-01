package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) NmaeGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

type PersonNameIterator struct {
	person  *Person
	current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person, -1}
}

func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("We should be here")
}

func (p *Person) Name() []string {
	return []string{p.FirstName, p.MiddleName, p.LastName}
}

func main() {
	// nomral iterator ↓↓↓
	// p := Person{"Alexander", "Graham", "Bell"}
	// for _, name := range p.Name() {
	// 	fmt.Println(name)
	// }

	// generator ↓↓↓
	// p2 := Person{"Alexander", "", "Bell"}
	// for namechan := range p2.NmaeGenerator() {
	// 	fmt.Println(namechan)
	// }

	// construct(pointer move) ↓↓↓
	p3 := Person{"Alexander", "Graham", "Bell"}
	for it := NewPersonNameIterator(&p3); it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
