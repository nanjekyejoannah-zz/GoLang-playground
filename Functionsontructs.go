package main

import (
	"fmt"
)

type Saiyan struct {
	Name string
	Power int
}

func (s *Saiyan) Super(){
	s.Power += 10000
}

func main(){
	goku := &Saiyan("goku",400)
	goku.Super()
	fmt.Printf(goku.Power)

}

//construtor..just write a method that acts as a factory
func newSaiyan {name string , power int} *Saiyan {
	return &Saiyan {
		Name : name
		Power : power
	}
}

//or

func newSaiyan {name string , power int} Saiyan {
	return Saiyan {
		Name : name
		Power : power
	}
}