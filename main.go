package main

import (
	"algorithm/crawler"
	"algorithm/twosum"
	"fmt"
	"time"
)

func main() {
	nums := []int{9, 11, 0, 15}
	target := 9
	resultTwoSum := twosum.TwoSum(nums, target)
	fmt.Println(resultTwoSum)
	go crawler.Main()
	time.Sleep(time.Second)
}
