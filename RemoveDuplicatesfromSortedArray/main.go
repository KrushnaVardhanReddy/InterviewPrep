package main

import "fmt"

func removeDuplicates(nums []int) int {
	fmt.Println("Input :", nums)
	lp := 1 // Left Pointer

	// loop through each item
	for index := 1; index < len(nums); index++ {
		// compare current with previous item
		// if they are not equal move current item to left pointer
		// increase the left pointer
		if nums[index] != nums[index-1] {
			nums[lp] = nums[index]
			lp += 1
		}
	}

	fmt.Println("nums :", nums)

	return lp
}

func main() {
	// Test Case1 : []int{1, 1, 2}
	testCase1 := []int{1, 1, 2}
	// Test Case2 : []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	testCase2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(testCase1))
	fmt.Println(removeDuplicates(testCase2))
}
