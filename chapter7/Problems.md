### `sum` is a function which takes a slice of numbers and add them together. What would its function sinature look like in Go?


```go
func sum(arr... int) int 
```


### Write a function which taks an integer and halves it and returns true if it was even or false if it was odd. For example `half(1)` should return`(0, false)` and `hals(2)` should retrun `(1, true)`.

### Using `makeEvenGenerator` as an example, write a `makeOddGenerator` function that generates odd numbers.

### The Fibonacci sequence is defined as : `fib(0) = 0`, `fib(1) = 1`, `fib(n) = fib(n - 1) + fib(n - 2)`.Write a recursive function which can find `fib(n)`

### What are defer, panic and recover? How do you recover from a run-time panic?

defer schedules a function call to be run after the function completes. panic is function to cause a run time error. recover is function too to stops the panic and returns the value that was passed to the call to panic.

When program calls panic, then programs call recover so will recover from a run-time panic.