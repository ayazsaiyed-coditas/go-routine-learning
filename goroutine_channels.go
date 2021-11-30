package main
import "fmt"
func main() {
  odd := make(chan int)
  oddsquared := make(chan int)
  //odd
  go func() {
    for x := 1; x < 10; x++ {
      if x%2 != 0 {
        odd <- x
      }
    }
    close(odd)
  }()
  //oddsquared
  go func() {
    for x := range odd {
      oddsquared <- x * x
    }
    close(oddsquared)
  }()

  for x := range oddsquared {
    fmt.Println(x)
  }
}