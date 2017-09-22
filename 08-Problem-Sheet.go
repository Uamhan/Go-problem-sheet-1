package main

import ( 
	"fmt"
	"sort"
)


//PROBELEM 8 : merge to slices into a sorted slice
func main() {
	//creats two sorted slices
	//passes these slices to mergeSort function
	//prints new merged list
	list1 := []int{1, 4, 6}
	list2 := []int{2,3,5}

	mergedList := mergeSort(list1, list2)

	fmt.Printf("%v\n", list1)
	fmt.Printf("%v\n", list2)
	fmt.Printf("%v\n", mergedList)
	
}
func mergeSort(list1, list2 []int) []int {
	list3 := []int{1} 
	list3 = append(list1,list2...)
	sort.Ints(list3)
	return list3
	}