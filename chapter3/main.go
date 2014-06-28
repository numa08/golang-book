package main

import "fmt"

func main() {
  fmt.Println("1 + 1 = ", 1.0 + 1.0)
  fmt.Println(len("Hello World"))
  fmt.Println("Hello World"[1])
  fmt.Println("Hello " + "World")
  fmt.Println(true && true)
  fmt.Println(true && false)
  fmt.Println(true || true)
  fmt.Println(true || false)
  fmt.Println(!true)
  fmt.Println("321325 x 424521 = ", 321325 * 424521) // problem 3
  fmt.Println((true && false) || (false && true) || !(false && false))
}