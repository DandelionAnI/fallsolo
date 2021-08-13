
package main

import (
	"fmt"
)

func code() {
	nums := []int{1,2,43,5,1,1,3,5}
	for i, v := range nums {
		if v == 1 {
			nums = append(nums[:i], nums[i+1:]...)
		}
		if v == 5 {
			nums = append(nums,0)
		}
		fmt.Println(i,nums)
	}
}