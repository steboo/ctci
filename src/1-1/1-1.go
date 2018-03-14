package main

import (
	"fmt"
	"os"
	"sort"
)

// Problem Statement
// Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?

func is_unique(s string) bool {
	// Let's first solve the first half of that problem... using a data structure.
	// Going to assume ASCII, but I'm not sure that it matters...
	var m map[rune]int
	m = make(map[rune]int)
	for _, char := range s {
		m[char] += 1

		if m[char] > 1 {
			return false
		}
	}

	// I don't think there are any constraints on the input with this approach (including character set).

	return true
}

func is_unique2(s string) bool {
	// To not solve this problem without using a data structure, the immediate idea that comes to mind is to sort the string, and look for duplicate adjacent characters.
	// I assume Go has immutable strings (furious Googling... it does indeed), so it seems reasonable that we would be permitted to use a character array (or perhaps a rune array).
	// (Imaginary interviewer nods slowly. Was that a yes? Let's implement it anyway...)

	// If the string has 0 or 1 characters, then all characters are trivially unique.
	if len(s) <= 1 {
		return true
	}

	array := []rune(s)

	sort.Slice(array, func(i, j int) bool { return array[i] < array[j] })

	for i := 0; i < len(array)-1; i++ {
		if array[i] == array[i+1] {
			return false
		}
	}

	return true
}

// Solution 1 has a worst-case runtime complexity of O(n)
// Solution 2 has a worst-case runtime complexity of O(nlgn)

func is_unique3(s string) bool {
	// So hint #117 suggests the route of a bit vector which probably better satisfies the "without additional data structures" better than my previous attempt.
	// Presumably each bit in the vector represents a character. The question is, how many characters do we want to support? The advantage of the previous solutions was that we could support any character set. To use a bit vector, we could probably reasonably restrict our input to ASCII characters (which has 128 characters), so our bit vector could be 128 bits.
	// However, Golang does not support 128 bit integers, so I have separate high and low 64 bit integers.
	var high uint64 = 0
	var low uint64 = 0
	var bv *uint64
	for _, char := range s {
		// Assuming ASCII inputs...
		code := uint8(char)

		if char > 64 {
			bv = &high
			code -= 64
		} else {
			bv = &low
		}

		// Check to see if the bit is set
		if *bv&(1<<code) != 0 {
			return false
		}

		// Set the bit
		*bv |= 1 << code
	}

	return true
}

// Finally, hint #132 suggests that we can do this in O(nlgn) time. It's now pretty clear that this could have been solved in O(n^2) in a brute force manner with no data structures, so maybe that's what the author was hinting at to avoid. Interestingly, the solution section in the book considers the bit vector solution to use a data structure and mentions the possibility of modifying the string in place.

// Golang questions to follow up on:
// 1. When to use var or not?
// 2. When to use := or =?
// 3. Was rune (which is apparently an int32) the right choice here? Would strings have been a better choice?

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Missing argument. %s STRING\n", os.Args[0])
		os.Exit(2)
	}

	fmt.Printf("Provided string has all unique characters: %t\n", is_unique(os.Args[1]))
	fmt.Printf("Provided string has all unique characters: %t\n", is_unique2(os.Args[1]))
	fmt.Printf("Provided string has all unique characters: %t\n", is_unique3(os.Args[1]))
}
