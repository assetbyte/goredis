package main 

import "fmt"
func TwoSum(nums []int, target int)[]int{
	seen := make(map[int]int) // value <-> index
	for i, num := range nums {
		difference := target - num
		if j, ok := seen[difference]; ok { // j - value, ok is either true or false
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    fmt.Println("TwoSum result:", TwoSum(nums, 9)) // expect: [0 1]
}