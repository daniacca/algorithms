# Problem: Integer Multiplication

**Input**:
Two _n_-digit nonnegative integers, _x_ and _y_

**Output**:
The product _x $\cdot$ y_

## Solutions

### 1. Language Multiply

Simplier solution, use the language multiplier operator

```golang
func Multiply(x int, y int) int {
	return x * y
}
```

### 2. Middle School Algorithm

The algorithms will work as follow:

```{r, tidy=FALSE, eval=FALSE }
result := 0

for each _digit_ of _y_
    multiply _x_, _digit_
    sum to the result, adding _n_ trailing zero
end for

return result
```
