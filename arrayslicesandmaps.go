package main 

import (
"fmt"
"math/rand"
"sort"
)


func main(){

	//array declarations
	var arr1 [10] int
	arr1[0] = 67
	goals := [4]int {1,2,3,4}

	//slices declarations
	names := [] string {"joan","sara"}
	checks := make ([] bool,10)
	var days [] string
	slices := make ([] int ,0,10)

	scores := make ([] int , 100)

	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	worst := make([]int, 5)
	copy(worst, scores[:5])
	fmt.Println(worst)

	//maps

	lookup := make (map[string] string)
	lookup["goki"] = "lee"
	lookup["jiho"] = "hu"
	lookup["jemo"] = "jamjei"

	fn,ln := lookup["jemo"]
	fmt.Println(fn,ln)

	total := len(lookup)
	fmt.Println(total)

	delete (lookup,"goki")
	fmt.Println(goals)
	fmt.Println(names)
	fmt.Println(checks)
	fmt.Println(days)
	fmt.Println(slices)


}