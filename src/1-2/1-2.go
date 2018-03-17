package main

import (
	"fmt"
	"os"
	"sort"
)

// Problem Statement
// Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.

func check_permutation(a string, b string) bool {
	// Initial solution: sort the two strings and compare.
	// A second solution that comes to mind is to build a hash table, and compare the counts. This solution seems messier for the white board.
	// Worst-case runtime complexity of solution #1: O(nlgn)
	// Solution #2: O(n)

	// Fast exit case:
	if len(a) != len(b) {
		return false
	}

	arr1 := []rune(a)
	arr2 := []rune(b)
	sort.Slice(arr1, func(i, j int) bool { return arr1[i] < arr1[j] })
	sort.Slice(arr2, func(i, j int) bool { return arr2[i] < arr2[j] })

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func check_permutation2(a string, b string) bool {
	// Fast exit case:
	if len(a) != len(b) {
		return false
	}

	// Assumes ASCII
	var m1, m2 map[byte]int
	m1 = make(map[byte]int)
	m2 = make(map[byte]int)

	for i := 0; i < len(a); i++ {
		m1[a[i]] += 1
		m2[b[i]] += 1
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}

		// This catches all the cases because we checked for length previously
		// i.e., if m2 has more keys than m1, one of the checked values will differ and return false
	}

	return true
	// O(n)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Missing argument. %s STRING STRING\n", os.Args[0])
		os.Exit(2)
	}

	fmt.Printf("Provided strings are permutations of one another: %t\n", check_permutation(os.Args[1], os.Args[2]))
	fmt.Printf("Provided strings are permutations of one another: %t\n", check_permutation2(os.Args[1], os.Args[2]))
}
