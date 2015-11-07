package main 

import "fmt"

func main() {
	var count = make(map[int]int)
	// var lst = []int {3, 4, 1, 15, 12, 8, 8, 9, 8, 3, 10, 4, 8}
	var lst = []int {2}
	var highCount = 0
	var commonInt = 0

	for _, each := range lst {
		if _, ok := count[each]; ok { 
			count[each] += 1
			} else { 
			count[each] = 1
		}
		if count[each] > highCount {
			highCount = count[each]
			commonInt = each
		}
	}

	fmt.Println(commonInt)

}

