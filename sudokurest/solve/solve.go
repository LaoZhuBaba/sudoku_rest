package solve

import (
	"fmt"
)

func isIntInSlice(i int, s []int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

// findCandidates looks up relatedElementsIndexes to determine what values
// would clash.  Then returns a list of numbers 1-9 which don't clash.
func findCandidates(s []int, relatedElementsIndexes []int, maxValue int) (candidates []int) {
	clashes := []int{}
	candidates = []int{}

	for _, v := range relatedElementsIndexes {
		clashes = append(clashes, s[v])
	}
	for count := 1; count <= maxValue; count++ {
		if !isIntInSlice(count, clashes) {
			candidates = append(candidates, count)
		}
	}
	return candidates
}

// recursiveSolve takes i as an index into  slice s of max length where related_elements
// is a slice max length containing lists of related elements.  For in slice s, element x
// will have related_elements[x] which must all be different from x.  Any number in the
// range 1 - 9 which isn't on the related_element list is a possible solution candidate.
func recursiveSolve(i int, s []int, maxValue int, maxIndex int, relatedElements [][]int) bool {
	// start at i and skip ahead past elements which already have a number in the range 1 - 9.
	// fmt.Printf("recursiveSolve() called with i: %d maxValue: %d maxIndex: %d, s: %v\n", i, maxValue, maxIndex, s)
	// fmt.Printf("relatedElements: %v\n", relatedElements)
	count := i
	for s[count] != 0 {
		if count >= maxIndex {
			return true
		}
		count++
	}
	candidates := findCandidates(s, relatedElements[count], maxValue)
	for _, candidate := range candidates {
		s[count] = candidate
		if count == maxIndex {
			return true
		}
		if recursiveSolve(count+1, s, maxValue, maxIndex, relatedElements) {
			return true
		}
	}
	// if no candidate value provides a solution then reset s[i] back to 0 and
	// return false.  This path is no good.
	s[count] = 0
	return false
}

func SolveSudoku(s []int, maxValue int, maxIndex int, relatedElements [][]int) (solution []int, err error) {
	// Make a copy of s because recursiveSolve updates its second parameter as a side effect
	sCopy := make([]int, len(s))
	copy(sCopy, s)

	result := recursiveSolve(0, sCopy, maxValue, maxIndex, relatedElements)
	if !result {
		return nil, fmt.Errorf("failed to solve sudoku")
	}
	return sCopy, nil
}
