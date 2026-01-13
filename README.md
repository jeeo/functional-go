# Lazy iterators evaluation

## About
This poc (that probably contain bugs) presents the idea of Lazy Iterators Evaluation with functional operations.

The core idea is to enhance in memory operations over big lists. Instead of doing multiple iterations over a data-set, we iterate over the operations that we want to perform. This allows that filtered elements in a previous operation (evaluation) won't be accounted anymore.
In this project, we only have Map and Filter operations, but it can easily be extended to any other desired operation, we just need to implement the Iterator interface.

Also this project doesn't implement the new Go iter constraint, but I'll do it soon (and remove this Crappy interface with the `HasNext()` and `Next()`)


## Example

Let's say we have a list with 10 elements [1-10], we want to: 
1. Filter out the prime numbers
2. transform the number to a struct (let's say { number, double })
3. filter all the numbers with double less than 10
4. any other op 

The regular imperative approach would be

```go
// 1st step: 10 iterations
var filtered []int
for _, x := range elements {
    // logic to filter out the prime numbers
}

// 2nd step: 6 iterations
var transformed []struct{ number, double int }
for {
    // logic to transform the elements
}

// 3rd step: 6 iterations
for  {
    // logic to filter all the numbers with the double less than 10 
}

// anything else will be 2 iterations
for {}
```


with the lazy evaluation

```go
type NumberPair struct {
    Number int
    Double int
}

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Build the lazy pipeline - no iterations happen yet
iter := iterable.NewIterable(elements)

// 1. Filter out prime numbers (keep non-primes)
filtered := tools.Filter(iter, func(x int) bool {
    return !isPrime(x)
})

// 2. Transform to struct {Number, Double}
transformed := tools.Map(filtered, func(x int) *NumberPair {
    return &NumberPair{Number: x, Double: x * 2}
})

// 3. Filter numbers where double >= 10
finalFiltered := tools.Filter(transformed, func(p NumberPair) bool {
    return p.Double >= 10
})

// 4. Any other operation - e.g., extract just the number
result := tools.Map(finalFiltered, func(p NumberPair) *int {
    return &p.Number
})

// Only now, when we consume the iterator, all operations are evaluated
// in a single pass through the data
for result.HasNext() {
    fmt.Println(*result.Next())
}
```
