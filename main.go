package main

import (
	"algorithm/twosum"
	"fmt"
)


func main(){
	nums := []int{9,11,0,15}
	target := 9
	resultTwoSum := twosum.TwoSum(nums, target)
	fmt.Println(resultTwoSum)
}
