package module02

import (
	"sort"
)

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	lenOfList := len(list)
	for idx := 1; idx < lenOfList; idx++ {
		for cIdx := 0; cIdx < lenOfList-idx; cIdx++ {
			if list[cIdx] > list[cIdx+1] {
				list[cIdx], list[cIdx+1] = list[cIdx+1], list[cIdx]
			}
		}
	}
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
	lenOfList := len(list)
	for idx := 1; idx < lenOfList; idx++ {
		for cIdx := 0; cIdx < lenOfList-idx; cIdx++ {
			if list[cIdx] > list[cIdx+1] {
				list[cIdx], list[cIdx+1] = list[cIdx+1], list[cIdx]
			}
		}
	}
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
	lenOfPeople := len(people)
	for idx := 1; idx < lenOfPeople-1; idx++ {
		for cIdx := 0; cIdx < lenOfPeople-idx; cIdx++ {
			if people[cIdx].Age > people[cIdx+1].Age {
				people[cIdx], people[cIdx+1] = people[cIdx+1], people[cIdx]
			}
		}
	}
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
	lenOfList := list.Len()
	for idx := 1; idx < lenOfList; idx++ {
		for cIdx := 0; cIdx < lenOfList-idx; cIdx++ {
			if list.Less(cIdx+1, cIdx) {
				list.Swap(cIdx, cIdx+1)
			}
		}
	}
}
