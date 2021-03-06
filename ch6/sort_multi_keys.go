///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and sort package

import (
	"fmt"
	"sort"
)

// A Commit is a record of code checkin
type Commit struct {
	username string
	lang     string
	numlines int
}

type lessFunc func(p1 *Commit, p2 *Commit) bool

// multiSorter implements the Sort interface, sorting the Commits within.
type multiSorter struct {
	Commits      []Commit
	lessFunction []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (multiSorter *multiSorter) Sort(Commits []Commit) {
	multiSorter.Commits = Commits
	sort.Sort(multiSorter)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(lessFunction ...lessFunc) *multiSorter {
	return &multiSorter{
		lessFunction: lessFunction,
	}
}

// Len is part of sort.Interface.
func (multiSorter *multiSorter) Len() int {
	return len(multiSorter.Commits)
}

// Swap is part of sort.Interface.
func (multiSorter *multiSorter) Swap(i int, j int) {
	multiSorter.Commits[i] = multiSorter.Commits[j]
	multiSorter.Commits[j] = multiSorter.Commits[i]
}

// Less is part of sort.Interface.

func (multiSorter *multiSorter) Less(i int, j int) bool {

	var p *Commit
	var q *Commit
	p = &multiSorter.Commits[i]
	q = &multiSorter.Commits[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(multiSorter.lessFunction)-1; k++ {
		less := multiSorter.lessFunction[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return multiSorter.lessFunction[k](p, q)
}

//main method
func main() {

	var Commits = []Commit{
		{"james", "Javascript", 110},
		{"ritchie", "python", 250},
		{"fletcher", "Go", 300},
		{"ray", "Go", 400},
		{"john", "Go", 500},
		{"will", "Go", 600},
		{"dan", "C++", 500},
		{"sam", "Java", 650},
		{"hayvard", "Smalltalk", 180},
	}

	var user func(*Commit, *Commit) bool
	user = func(c1 *Commit, c2 *Commit) bool {
		return c1.username < c2.username
	}

	var language func(*Commit, *Commit) bool
	language = func(c1 *Commit, c2 *Commit) bool {
		return c1.lang < c2.lang
	}

	var increasingLines func(*Commit, *Commit) bool
	increasingLines = func(c1 *Commit, c2 *Commit) bool {
		return c1.numlines < c2.numlines
	}

	var decreasingLines func(*Commit, *Commit) bool
	decreasingLines = func(c1 *Commit, c2 *Commit) bool {
		return c1.numlines > c2.numlines // Note: > orders downwards.
	}

	OrderedBy(user).Sort(Commits)
	fmt.Println("By username:", Commits)

	OrderedBy(user, increasingLines).Sort(Commits)
	fmt.Println("By username,asc order", Commits)

	OrderedBy(user, decreasingLines).Sort(Commits)
	fmt.Println("By username,desc order", Commits)

	OrderedBy(language, increasingLines).Sort(Commits)
	fmt.Println("By lang,asc order", Commits)

	OrderedBy(language, decreasingLines, user).Sort(Commits)
	fmt.Println("By lang,desc order", Commits)

}
