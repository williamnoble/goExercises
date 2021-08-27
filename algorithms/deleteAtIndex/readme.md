# Removing an element from a slice

## Fast (Changes Order)

```go
package main

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2
	// Remove the element at index i from a. 
	a[i] = a[len(a)-1] // Copy last element to index i. 
	a[len(a)-1] = ""   // Erase last element (write zero value).
	a = a[:len(a)-1]   // Truncate slice. fmt.Println(a) // [A B E D]
}

```

## Slow Version (Maintains Order)

```go
package main

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2
	// Remove the element at index i from a. 
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index
	a[len(a)-1] = ""     // Erase last element (write zero value)
	a = a[:len(a)-1]     // Truncate slice.
}

```

## Append

```go
package main

func main() {
	index := 1
	s := []string{"one", "two", "three"}
	// 1) append everything upto but not including the index
	// 2) append everything after the index.
	x := append(s[:index], s[index+1:]...)
	_ = x
}
```

End.

