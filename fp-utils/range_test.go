package fputils

import "testing"

func compareRangeContent(got IntRange, want []int ) bool {
	if( len(got) != len(want)) {
		return false
	}

	for i := 0; i < len(want) ; i++ {
		if(got[i] != want[i]) {
			return false
		}
	}
	return true
}

func TestDeclarativeIntRangeGenerator (t *testing.T ) {
/// GenerateIntRange
    t.Run("Generte ints no inclusive", func(t *testing.T){
        want := []int{1,2,3,4,5,6,7,8,9}
        got := GenerateIntRange(1,10)
        if !compareRangeContent(got, want) {
            t.Errorf("got = %+v; want = %+v", got, want)
        }
    })

    t.Run("Generte ints inclusive", func(t *testing.T){
	    want := []int{2,3,4,5,6,7,8,9,10}
	    got := GenerateInclusiveIntRange(2,10)
	    if !compareRangeContent(got, want) {
		    t.Errorf("got = %+v; want = %+v", got, want)
	    }
    })
// Filter - NOT
    t.Run("Filter pairs", func(t *testing.T){
	    isEven := func(x int) bool {return x%2 == 0}
	    want := []int{1,3,5,7,9}
	    got := GenerateIntRange(1,10).Not(isEven)

	    if !compareRangeContent(got, want) {
		    t.Errorf("got = %+v; want = %+v", got, want)
	    }
    })

    t.Run("Filter odds", func(t *testing.T){
	    isEven := func(x int) bool {return x%2 == 0}
	    want := []int{2,4,6,8}
	    got := GenerateIntRange(1,10).Filter(isEven)

	    if !compareRangeContent(got, want) {
		    t.Errorf("got = %+v; want = %+v", got, want)
	    }
    })
// MAP
    t.Run("Transform to doubles", func(t *testing.T){
	    duplicate := func(x int) int {return x*2}
	    want := []int{2,4,6,8,10,12,14,16,18}
	    got := GenerateIntRange(1,10).Map(duplicate)

	    if !compareRangeContent(got, want) {
		    t.Errorf("got = %+v; want = %+v", got, want)
	    }
    })


    t.Run("Duplicate only the odds", func(t *testing.T){
	    duplicate := func(x int) int {return x*2}
	    isOdd := func(x int) bool {return !(x%2 == 0)}
	    want := []int{2,2,6,4,10,6,14,8,18}
	    got := GenerateIntRange(1,10).MapOnly(duplicate, isOdd)

	    if !compareRangeContent(got, want) {
		    t.Errorf("got = %+v; want = %+v", got, want)
	    }
    })
//ANY
    t.Run("Any returns true if any value is even", func(t *testing.T) {
	    isEven := func(x int) bool { return x%2 == 0 }
	    got := GenerateIntRange(1, 5).Any(isEven)
	    want := true
	    if got != want {
		    t.Errorf("got = %v; want = %v", got, want)
	    }
    })

    t.Run("Any returns false if no value is greater than 10", func(t *testing.T) {
	    gt10 := func(x int) bool { return x > 10 }
	    got := GenerateIntRange(1, 5).Any(gt10)
	    want := false
	    if got != want {
		    t.Errorf("got = %v; want = %v", got, want)
	    }
    })
// ALL
    t.Run("All returns true when all values are less than 10", func(t *testing.T) {
	    lt10 := func(x int) bool { return x < 10 }
	    got := GenerateIntRange(1, 9).All(lt10)
	    want := true
	    if got != want {
		    t.Errorf("got = %v; want = %v", got, want)
	    }
    })

    t.Run("All returns false if any value is odd", func(t *testing.T) {
	    isEven := func(x int) bool { return x%2 == 0 }
	    got := GenerateIntRange(2, 6).All(isEven)
	    want := false
	    if got != want {
		    t.Errorf("got = %v; want = %v", got, want)
	    }
    })
// Find
    t.Run("Find returns first odd number", func(t *testing.T) {
	    isOdd := func(x int) bool { return x%2 != 0 }
	    got, ok := GenerateIntRange(2, 6).Find(isOdd)
	    want := 3
	    if !ok || got != want {
		    t.Errorf("got = %v (ok=%v); want = %v", got, ok, want)
	    }
    })

    t.Run("Find returns false if no match", func(t *testing.T) {
	    gt10 := func(x int) bool { return x > 10 }
	    _, ok := GenerateIntRange(1, 5).Find(gt10)
	    if ok {
		    t.Errorf("expected no match, but got ok = true")
	    }
    })
// Reduce
    t.Run("Reduce: sums elements", func(t *testing.T) {
	    sum := func(a, b int) int { return a + b }
	    got := GenerateIntRange(1, 5).Reduce(0, sum)
	    want := 10 // 1 + 2 + 3 + 4
	    if got != want {
		    t.Errorf("got = %v; want = %v", got, want)
	    }
    })

    t.Run("Reduce: multiplies elements", func(t *testing.T) {
	    mult := func(a, b int) int { return a * b }
	    got := GenerateIntRange(1, 4).Reduce(1, mult)
	    want := 6 // 1 * 2 * 3
	    if got != want {
		    t.Errorf("got = %v; want = %v", got, want)
	    }
    })

    t.Run("Reduce: multiplies all elements", func(t *testing.T) {
	    mult := func(a, b int) int { return a * b }
	    got := GenerateInclusiveIntRange(1, 4).Reduce(1, mult)
	    want := 24 // 1 * 2 * 3 * 4
	    if got != want {
		    t.Errorf("got = %v; want = %v", got, want)
	    }
    })
// Count
    t.Run("Count returns correct length", func(t *testing.T) {
	    got := GenerateIntRange(1, 10).Count()
	    want := 9
	    if got != want {
		    t.Errorf("got = %v; want = %v", got, want)
	    }
    })
////// Method chaining examples //////////////////////////////////////////////
    t.Run("Count all odd numbers in a range", func(t *testing.T) {
	    isOdd := func(x int) bool {return !(x%2 == 0)}
	    got := GenerateIntRange(1, 10).Filter(isOdd).Count()
	    want := 5
	    if got != want {
		    t.Errorf("got = %v; want = %v", got, want)
	    }
    })

    t.Run("Sum range with conditional mapping on odd value", func(t *testing.T) {
	    isOdd := func(x int) bool {return !(x%2 == 0)}
	    sum := func(a, b int) int { return a + b }
	    duplicate := func(x int) int {return x*2}
	    // if n is odd then duplicate n (e.g. 1,2,3,4 -> 2,2,6,4)
	    got := GenerateIntRange(1, 6).MapOnly(duplicate, isOdd).Reduce(0, sum)
	    want := 24 // 2 + 2 + 6 + 4 + 10
	    if got != want {
		    t.Errorf("got = %v; want = %v", got, want)
	    }
    })
}
