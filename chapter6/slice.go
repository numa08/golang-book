package main

import "fmt"

func main() {
  slice1 := []int{1, 2, 3}
  slice2 := make([]int, 2)
  copy(slice2, slice1)
  fmt.Println(slice1, slice2)

  slice3 := make([]int, 3, 9)
  fmt.Println(len(slice3))
}