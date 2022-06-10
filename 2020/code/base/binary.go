package main


import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1,11,-5,7,2,0,12,16}
	sort.Ints(numbers)
	fmt.Println(numbers)
	index := sort.SearchInts(numbers,7)
	fmt.Println(index)

	ind := sort.Search(len(numbers),func(i int) bool {
		return numbers[i] >= 7 
	})
	fmt.Println("The first number >= 7 is at index:", ind)
    fmt.Println("The first number >= 7 is:", numbers[ind])
}