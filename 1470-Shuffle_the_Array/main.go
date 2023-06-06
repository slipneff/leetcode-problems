package main
import "fmt"

func shuffle(nums []int, n int) []int {
    counter := 0
    newArray := make([]int,n*2)
    for i:=0;i<n*2;i=i+2{
        newArray[i]=nums[counter]
        newArray[i+1]=nums[counter+n]
        counter++    
    }
    return newArray
}

func main(){
  example := []int{2,5,1,3,4,7}
  n := 3
  fmt.Printf("%v",shuffle(example,n))
}
