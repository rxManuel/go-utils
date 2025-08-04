# How to use it

## range.go
> Provides a *declarative*, eager integer range generator that supports _chaining functions_ such as:
> - Filter: Filters elements in the range according to a condition.
> - Map: Transforms elements according to a rule.
> - MapOnly: Transforms only the elements that meet a condition.
> - Reduce: Combines all elements into a single value using a specified rule or operation.
> - Any: Verify if at least one element meets the specified condition.
> - All: Verify if all elements meet the specified condition.
> - Count: Count the total of elements in the range

**You can find more examples in [range_test.go](fp-utils/range_test.go)**
### Exemple 1: print a range from 1 to 5 not inclusive
```
package main

import(
    "fmt"
    "github.com/rxManuel/go-utils/fp-utils"
)

func main() {
	r := fputils.GenerateIntRange(1, 5) //package name is fputils
	fmt.Println("Range:", r)
}

```
#### output:
> Range: [1 2 3 4]

### Example 2: print a range from 1 to 5 inclusive
```
...
	r := fputils.GenerateIntRange(1, 5) //package name is fputils
	fmt.Println("Range:", r)
...
```
#### output:
> Range: [1 2 3 4 5]
### Example 3: print all odd numbers in a range from 1 to 10
```
...
	isOdd := func (x int) bool { return x%2 != 0}
	r := fputils.GenerateInclusiveIntRange(1,10).Filter(isOdd)
	fmt.Println("Odds:", r)
...
```
#### output:
> Odds: [1 3 5 7 9]

### Example 4: print all odd numbers and then all even numbers in a range from 1 to 10
```
...
	isOdd := func (x int) bool { return x%2 != 0}
	
	naturals := fputils.GenerateIntRange(1, 10)
	
	fmt.Println("Odds: ", naturals.Filter(isOdd))
	fmt.Println("Evens: ", naturals.Not(isOdd))
...
```
#### output:
> ```
> Odds:  [1 3 5 7 9]
> Evens:  [2 4 6 8]
> ```
