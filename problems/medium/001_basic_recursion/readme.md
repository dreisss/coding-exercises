# Some Basic Recursion

Create a function that works like the Fibonacci function in math.

<details>
  <summary>Tips?</summary>

1. What is the Fibonacci function?
2. Knowing what this is, how it will look like when you define this function?
3. Try to transpose the function schema to steps, that use variables and concepts from coding.
4. Implement!
5. (Go) Problems running this code? remember that go is copiled, it requires a main function, all your code must run inside this function.

</details>

<details>
  <summary>Solution?</summary>

Google it.

Or... Based in the Tips section.

1. Is a function that looks like this:

$F(n)=F(n-1)+F(n-2)$

In other words: The current term is the last 2 summed together.
And when the current is equal or less than 1, it results in 1.

2. In go, it will look like:

```go
func fib(n int) int {
  ...
}
```

In other words: it receive a "n" from type integer and returns a value with type integer too.

3. First, if the current is equal or less than to 1, the result is 1, but if its greater, return the result of the previous and before previous summed together.

4. So, the implementation looks like:

```go
func fib(n int) int {
  if n <= 1 {
    return 1
  } else {
    return fib(n - 1) + fib(n - 2)
  }
}
```

</details>
