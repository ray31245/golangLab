// https://www.udemy.com/course/design-patterns-go/learn/lecture/16912634#overview
package main

import "fmt"

// combination of OCP and Repository demo

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct{}

func (f *Filter) filterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) filterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) filterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// filterBySize, filterBySizeAndColor

type Specitication interface {
	IsSatisfiled(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfiled(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfiled(p *Product) bool {
	return p.size == spec.size
}

type AndSpecitication struct {
	first, second Specitication
}

func (spec AndSpecitication) IsSatisfiled(p *Product) bool {
	return spec.first.IsSatisfiled(p) && spec.second.IsSatisfiled(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specitication) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfiled(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	fmt.Print("Green products (old):\n")
	f := Filter{}
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	// ↑↑↑ BEFORE

	// ↓↓↓ AFTER
	fmt.Print("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}

	largeGreenSpec := AndSpecitication{largeSpec, greenSpec}
	fmt.Print("Large blue items:\n")
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
