package main

import "fmt"

func buildArray(nums []int) []int {
  newArray := make([]int,len(nums))
  for _,b := range nums {
    newArray[b] = nums[nums[b]]
  }
  return newArray
}

func main() {
  example := []int{0,2,1,5,3,4}
  fmt.Printf("%v",buildArray(example))
}
