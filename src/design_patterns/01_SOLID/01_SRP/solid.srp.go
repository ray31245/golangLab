// https://www.udemy.com/course/design-patterns-go/learn/lecture/16912628#overview
package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) {
	entry := fmt.Sprint(text)
	j.entries = append(j.entries, entry)
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// 違反單一職責(SRP)
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

// 不違反單一職責的做法(1)
var lineSeparator = "\n"

func SaveToFile(s []string, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(s, lineSeparator)), 0644)
}

// 不違反單一職責的做法(2)
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(s []string, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(s, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug")
	fmt.Println(strings.Join(j.entries, "\n"))
	j.Save("journal.txt")

	// 作法一
	SaveToFile(j.entries, "journal.txt")

	// 作法二
	p := Persistence{"\n"}
	p.saveToFile(j.entries, "journal.txt")
}
