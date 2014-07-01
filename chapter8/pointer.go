package main

import "fmt"

func main() {
  z := 5
  zero(&z)
  fmt.Println(z)

  o := new(int)
  one(o)
  fmt.Println(*o)

  x := 1.5
  square(&x)
  fmt.Println(x)

  a := 1
  b := 2
  swap(&a, &b)
  fmt.Println(a, b)
}

func zero(xPtr *int) {
  *xPtr = 0
}

func one(xPtr *int) {
  *xPtr = 1
}

func square(x *float64) {
  *x = *x * *x
}

func swap(x *int, y *int) {
  tem := *x
  *x = *y
  *y = tem
}