package main 

import ("strconv"
	"fmt")

func main() {
	var count = make(map[string]int)
	var lst = []int {3, 4, 1, 15, 12, 8, 8, 9, 8, 3, 10, 4, 8}


	for _, each := range lst {
		eachStr := strconv.Itoa(each)
		if _, ok := count[eachStr]; ok { 
			count[eachStr] += 1
			} else { 
			count[eachStr] = 1
		}
	}
//need a second bit that iterates through the map, changing the returned
//value based on which has the highest count

	var keys []string
	for k := range count {
		keys = append(keys, k)
	}

	var highCount int
	var commonInt string
	for _, k := range keys {
		if count[k] > highCount {
			commonInt = k
			highCount = count[k]
		}
	}

	fmt.Println(commonInt)

}



