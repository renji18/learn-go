package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}

	fmt.Println(twoSum(nums, 6))
}

func twoSum(nums []int, target int) []int {
	ans := []int{0, 0}

	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				ans[0] = i
				ans[1] = j
			}
		}
	}

	return ans
}


/*
	We range over all elements in the array for the first element. After fixing the first element in the loop, we create a new loop for the first element. This second element loops over the array and we do index 0 + all index of array, then index 1 + all index of array. Until we either exaust the outer loop or we get our desired output.	
*/