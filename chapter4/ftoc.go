package main

import "fmt"

func main() {
  fmt.Println("The program will convert a value from Fahrenheit into Celsius.")
  fmt.Println("Please input a Fahrenheit value")

  var fahrenheit float64
  fmt.Scanf("%f", &fahrenheit)

  var celsius = (fahrenheit - 32) * (5.0 / 9.0)

  fmt.Println("Celsius value is ", celsius)

}