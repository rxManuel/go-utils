package fputils

type IntRange []int
type IntPredicate func(x int) bool
type IntTransformer func(x int) int
/***
 * An eager int generator to filter, map, reduce, count and find in a declarative way.
 *
 * HOW TO USE IT?
 *
 * what if you want to get all odd numbers between 1 to 100
 * then:
 * 	isOdd := func(x int) bool { return x%2 != 0}
 * 	hundredInts := GenerateInclusiveIntRange(1, 100)
 *  	odds := hundredInts.Filter(isOdd)
 * and then you want all pairs (not odd)
 * 	pairs := hundredInts.Not(isOdd)
 *
 * Yes you can reuse the range it is not consumed or muted, but what if you want a new copy or all pairs until 150?
 * then:
 *
 * 	pairsUntil150 := GenerateInclusiveIntRange(1, 150).Not(isOdd)
 * or how many pairs?
 * 	howManyPairs149 := GenerateIntRange(1, 150).Not(isOdd).Count()
 * or total of odd numbes
 * 	sum := func(a, b int) int {return a+b}
 * 	oddSum := GenerateIntRange(1,150).Filter(isOdd).Reduce(0, sum)
 * */
func GenerateIntRange(start, end int) IntRange {
	if end <= start {
		panic("Invalid range, start must be less than end")
	}
	r := make([]int, (end - start))
	for i := 0; i < len(r); i++ {
		r[i] = start + i
	}
	return IntRange(r)
}

func GenerateInclusiveIntRange(start, end int) IntRange {
	return GenerateIntRange(start, end+1)
}

func (r IntRange) Filter(predicate IntPredicate) IntRange {
	filtered := make([]int, 0, len(r))
	for _, i := range r {
		if predicate(i) {
			filtered = append(filtered, i)
		}
	}
	return IntRange(filtered)
}

func (r IntRange) Map(mapper IntTransformer) IntRange {
	mapped := make([]int, 0, len(r))
	for _, i := range r {
		mapped = append(mapped, mapper(i))
	}
	return IntRange(mapped)
}

func (r IntRange) MapOnly(mapper IntTransformer, predicate IntPredicate) IntRange {
	mapped := make([]int, 0, len(r))
	for _, i := range r {
		if predicate(i) {
			mapped = append(mapped, mapper(i))
		} else {
			mapped = append(mapped, i)
		}
	}
	return IntRange(mapped)
}


func (r IntRange) Not(predicate IntPredicate) IntRange {
	negated := func(x int) bool {
		return !predicate(x)
	}
	return r.Filter(negated)
}

func (r IntRange) Any(predicate IntPredicate) bool {
	for _, i := range r {
		if predicate(i) {
			return true
		}
	}
	return false
}

func (r IntRange) All(predicate IntPredicate) bool {
	for _, i := range r {
		if !predicate(i) {
			return false
		}
	}
	return true
}

func (r IntRange) Find(predicate IntPredicate) (int, bool) {
	for _, i := range r {
		if predicate(i) {
			return i, true
		}
	}
	return 0, false
}

func (r IntRange) Reduce(initial int, f func(acc, curr int) int) int {
	acc := initial
	for _, i := range r {
		acc = f(acc, i)
	}
	return acc
}

func (r IntRange) Count() int {
	return len(r)
}
