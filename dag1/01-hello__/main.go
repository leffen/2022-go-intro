package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person)Show() {
   fmt.Printf("Navn: %s Alder: %d\n",p.name,p.age)
}




func main() {
	person := Person{name: "test2",age: 17}
	person.Show()
}
