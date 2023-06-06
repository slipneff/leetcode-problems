package main 
import "fmt"

func finalValueAfterOperations(operations []string) int {
  x := 0
  for _,val := range operations{
    if val[1] == '+'{
      x++
    } else {
      x--
    }
  }
  return x
}

func main(){
  example := []string{"--X","X++","X++"}
  fmt.Printf("%v",finalValueAfterOperations(example))
}
