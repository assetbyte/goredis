package main 

import "fmt"

func majorityElement(nums []int) int {
	counts := make(map[int]int)
	limit := len(nums) / 2 
	for _, num := range nums {
		counts[num] ++
		if counts[num] > limit {
			return num
		}
	}
	return -1 
}

func testMajority() {
	nums := []int{3, 2, 3, 3, 1}
	fmt.Println("Majority element:", majorityElement(nums))
}

