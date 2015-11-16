package main 

import (
	"fmt"
	"reflect"
	)

var wordList = []string {"cat", "abstentious", "arterious", "toaster", "placentious"}
var letterList = []string {" a", "e", "i", "o", "u"} 
//Added a space in front of the a to try to quick-and-dirty my way to a solution before going deeper into the problem.
var letterMap = map[string]string{"a": "a", "e": "e", "i": "i", "o": "o", "u": "u"}
//Could use index to cut looping short when one letter is in the wrong place
//Could use length to determine quick True or False
func main() {
	fmt.Println(letterList)
	for _, word := range wordList {
		vowelList := make([]string, 1)
		for _, letter := range word {
			if _, ok := letterMap[string(letter)]; ok {
				vowelList = append(vowelList, string(letter))
			}
		}
		fmt.Println(vowelList)
		fmt.Println(letterList)
		//They appear to be the same, but output disagrees.
		if reflect.DeepEqual(vowelList, letterList) {
			fmt.Println(word, "has all the vowels in order, once each!")
		} else {
			fmt.Println(word, "definitely doesn't have all the vowels in order.")
		}
	}
}

// vowelList = []
// for word in wordList:
// 	for letter in word:
// 		if letter in vowelList:
// 			return False
// 		if letter in letterList:
// 			vowelList.append(letter)
// if vowelList == letterList:
// 	return True