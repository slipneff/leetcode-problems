package main 
import "fmt"


func solution1(nums1 []int, m int, nums2 []int, n int) []int {
  for value := range nums2 {
    i:=m-1
    for i=m-1;i>=0 && nums1[i]>nums2[value];i--{
      nums1[i+1]=nums1[i]
    }
    m++
    nums1[i+1]=nums2[value]
    fmt.Printf("%v",nums1)
  }
  return nums1
}


func solution2(nums1 []int, m int, nums2 []int, n int) []int {   
  for m >0 || n>0{
    if m>0 && ( n==0 || nums1[m-1] >= nums2[n-1]) {
      nums1[m+n-1] = nums1[m-1]
      m--
    } else {
      nums1[m+n-1] = nums2[n-1]
      n--
    }
  }
  return nums1
}
func main(){
  example1 := []int{1,2,3,0,0,0}
  example2 := []int{2,5,6}
  m,n := 3,3
  fmt.Printf("%v",solution2(example1,m,example2,n))
}
