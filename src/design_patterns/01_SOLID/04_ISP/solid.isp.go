package main

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// ok if you need a multifunction device
type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print(d Document) {

}

func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
}

func (o OldFashionedPrinter) Print(d Document) {
	// ok
}

func (o OldFashionedPrinter) Fax(d Document) {
	panic("opreation not supported")
}

func (o OldFashionedPrinter) Scan(d Document) {
	panic("opreation not supported")
}

// better appraoch: split into several interfaces
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// printer only
type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {

}

// combine interface
type Photocopier struct{}

func (p Photocopier) Scan(d Document) {

}

func (p Photocopier) Print(d Document) {

}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

// interface combintion + decorator
type MultiFunctionMatchine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMatchine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMatchine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {

}
