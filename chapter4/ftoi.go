package main

import "fmt"

func main() {
  fmt.Println("This program will convert a value from Feet into Meters.")
  fmt.Println("Please input a Feet value")

  var feet float64
  fmt.Scanf("%f", &feet)

  var meter = feet / 0.3048

  fmt.Println("Meter value is ", meter)
}