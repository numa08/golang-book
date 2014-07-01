### How do you get the memory address of a variable?

append symbole which is `&` to variable.

### How do you assign a value to pointer?

append symbole which is `*` to valiable, and assign a value.

### How do you create a new pointer?

call `new` function. that function's argument is a type.

### What is the value of x after running this program:


```go
func square(x *float64) {
  *x = *x * *x
}
func main() {
  x := 1.5
  square(&x) 
}
```

2.25(1.5 ^ 2)

### Write a program that can swap two integers `(x := 1, y := 2; swap(&x, &y))` should give you `x=2` and `y=1`.