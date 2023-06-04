package main

import "fmt"
func getConcatenation(nums []int) []int {
  lenArray := len(nums)
  newArray := make([]int,lenArray*2)
  for a,b := range nums {
    newArray[a]=b
    newArray[a+lenArray]=b
  }
  return newArray
}
func main(){
  example := []int{1,2,1}
  fmt.Printf("%v",getConcatenation(example))
}
