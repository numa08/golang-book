package main

import "fmt"

func main() {
  fmt.Println(sum(1, 2, 3))
  fmt.Println(half(1))
  fmt.Println(half(2))

  fmt.Println(gratest(4, 2, 10, 8, 4, 7))

  nextOdd := makeOddGenerator()
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())

  fmt.Println(fib(0))
  fmt.Println(fib(1))
  fmt.Println(fib(10))
}

func sum(args... int) int {
  sum := 0
  for _, v := range args {
    sum += v
  }
  return sum
}

// problem 2
func half(num int) (int, bool) {
  return num / 2, num % 2 == 0
}

// problem 3
func gratest(args... int) int {
  gratest := args[0]
  for i := 1; i < len(args) ; i++ {
    if gratest < args[i] {
      gratest = args[i]
    }
  }

  return gratest
}

func makeOddGenerator() func() uint {
  i := uint(1)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

func fib(n int) int {
  if n == 0 {
    return 0
  }

  if n == 1 {
    return 1
  }

  return fib(n - 1) + fib(n - 2)
}