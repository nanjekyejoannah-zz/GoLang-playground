package main

import (
	"fmt"
)
type Person struct {
	Name string
}

func (p *Person) introduce(){
	fmt.Println("hey my name is %s",p.Name)

}

type Saiyan struct {
	*Person
	Power int
}

func main(){
	goku := &Saiyan {
		Person : &Person{"goku"},
		Power : 9000,

	}
	goku.introduce()
}