package main

import (
"fmt"
)

type Saiyan struct {
	Name string
	Power int

}


func main(){
	//without point
	goku := Saiyan {"Goku",9000}
	Super(goku)
	fmt.Println(goku.Power)

	//with point
	goki := &Saiyan{"Goki",9000}
	superWithPointer(goki)
	fmt.Println(goki.Power)

}

func Super(s Saiyan){
	s.Power+=10000

}

func superWithPointer(s *Saiyan){
	s.Power += 20000
}