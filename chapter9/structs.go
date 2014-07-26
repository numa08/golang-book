package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1
  return math.Sqrt(a * a + b * b)
}

type Circle struct {
  x float64
  y float64
  r float64
}

type Rectangle struct {
  x1 ,y1 ,x2 ,y2 float64
}

func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
  return math.Pi * c.r * 2
}

func (r *Rectangle) area() float64 {
  w := distance(r.x1, r.y1, r.x1, r.y2)
  l := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func (r *Rectangle) perimeter() float64 {
  w := distance(r.x1, r.y1, r.x1, r.y2)
  l := distance(r.x1, r.y1, r.x2, r.y1)
  return w * 2 + l * 2
}

type Person struct {
  Name string
}

func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
  Person
  Model string
}

type Shape interface {
  area() float64
  perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}

type MultiShape struct {
  shapes []Shape
}

func (m *MultiShape) area() float64 {
  var area float64
  for _, s := range m.shapes {
    area += s.area()
  }
  return area
}

func main() {
  // c := Circle{x : 0, y : 0, r : 5}
  c := Circle{0, 0, 5}
  fmt.Println(c.x , c.y, c.r)
  c.x = 10
  c.y = 5

  fmt.Println(cicleArea(&c))
  fmt.Println(c.area())
  fmt.Println(c.perimeter())

  a := new(Android)
  a.Person.Talk()
  a.Talk()

  r := Rectangle{0, 0, 10, 20}
  fmt.Println(r.area())
  fmt.Println(r.perimeter())
  total := totalArea(&r, &c)
  fmt.Println(total)

}

func cicleArea(c *Circle) float64 {
  return math.Pi * c.r * c.r
}