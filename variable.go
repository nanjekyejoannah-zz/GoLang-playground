package main 



import (
"fmt"
)



func main(){
	power := getPower()
	println(power)
	fmt.Printf("The power is %d",power)
}

func getPower() int {

	return 9001
}