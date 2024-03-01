package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i, num := range nums {
		res[i] = prefix
		prefix *= num
	}
	fmt.Println("prefix multiplication completed")
	fmt.Println("Actual: ", nums)
	fmt.Println("res : ", res)

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
		// fmt.Printf("postfix: %v - index : %v - Value : %v \n", postfix, i, nums[i])
	}
	return res
}

func main() {
	// Test Case1 : []int{1, 1, 2}
	testCase1 := []int{5, 2, 3, 4}
	// Test Case2 : []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	testCase2 := []int{-1, 1, 0, -3, 3}

	fmt.Println(productExceptSelf(testCase1))
	fmt.Println(productExceptSelf(testCase2))
}
