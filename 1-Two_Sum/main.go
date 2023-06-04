package main

import "fmt"
func contains(s map[int]int, e int) (int,bool) {
  for k, v := range s {
    if v == e {
      return k,true
    }
  }
  return 1,false
}

func twoSum(nums []int, target int) []int {
  ans := make([]int,0)
  length := len(nums) 
  hashmap := make(map[int]int)
  for i:=0;i<length;i++ {
    bufInt,bufBool := contains(hashmap,nums[i])
    if bufBool == true {
      ans = append(ans,bufInt) 
      ans = append(ans,i)
      return ans
    }
    hashmap[i] = target-nums[i]
  }
  return ans
}
func main(){
  example := []int{2,7,11,15}
  target := 9

  fmt.Printf("%v\n",twoSum(example,target)) 
}
