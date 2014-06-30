package main

import "fmt"

func main() {
  elements := map[string]string {
  "H"  :  "Hydrogen",
  "He" : "Helium",
  "Li" : "Lithium",
  "Be" : "Beryllium",
  "B"  :  "Boron",
  "C"  :  "Carbon",
  "N"  :  "Nitrogen",
  "O"  :  "Oxgen",
  "F"  :  "Fluorine",
  "Ne" : "Neon",
  }

  name, ok := elements["Un"]
  fmt.Println(name, ok)
}