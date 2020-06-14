package module02

import (
	"sort"
)

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
func InsertionSortInt(list []int) {
	sortedList := []int{list[0]}
	for idx := 1; idx < len(list); idx++ {
		index := FindTargetWithBinaryCheck(sortedList, list[idx])
		movedElements := append([]int{list[idx]}, sortedList[index:]...)
		sortedList = append(sortedList[:index], movedElements...)
	}
	for idx, _ := range list {
		list[idx] = sortedList[idx]
	}
}

// func FindTarget(list []int, target int) int {
// 	postion := len(list)
// 	for idx := postion - 1; idx >= 0; idx-- {
// 		// fmt.Println(list[idx], target)
// 		if list[idx] < target {
// 			return idx + 1
// 		}
// 	}
// 	return 0
// }

func FindTargetWithBinaryCheck(list []int, target int) int {
	left := 0
	right := len(list) - 1
	for left <= right {
		mid := (left + right) / 2
		if list[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
func InsertionSortString(list []string) {
	sortedList := []string{list[0]}
	for idx := 1; idx < len(list); idx++ {
		index := FindTargetStringWithBinaryCheck(sortedList, list[idx])
		movedElements := append([]string{list[idx]}, sortedList[index:]...)
		sortedList = append(sortedList[:index], movedElements...)
	}
	for idx, _ := range list {
		list[idx] = sortedList[idx]
	}
}
func FindTargetStringWithBinaryCheck(list []string, target string) int {
	left := 0
	right := len(list) - 1
	for left <= right {
		mid := (left + right) / 2
		if list[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
	sortedList := []Person{people[0]}
	for idx := 1; idx < len(people); idx++ {
		index := FindTargetPersonWithBinaryCheck(sortedList, people[idx])
		movedElements := append([]Person{people[idx]}, sortedList[index:]...)
		sortedList = append(sortedList[:index], movedElements...)
	}
	for idx, _ := range people {
		people[idx] = sortedList[idx]
	}
}

func FindTargetPersonWithBinaryCheck(list []Person, target Person) int {
	left := 0
	right := len(list) - 1
	for left <= right {
		mid := (left + right) / 2
		if list[mid].Age < target.Age {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
	lenOfList := list.Len()
	for sortedLen := 1; sortedLen < lenOfList; sortedLen++ {
		swapIndex := BinarySearchIndex(list, sortedLen, sortedLen)
		for idx := sortedLen; idx > swapIndex; idx-- {
			list.Swap(idx, idx-1)
		}
	}
}

func BinarySearchIndex(list sort.Interface, sortedLen, targetIndex int) int {
	left := 0
	right := sortedLen - 1
	for left <= right {
		mid := (left + right) / 2
		if list.Less(mid, targetIndex) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
