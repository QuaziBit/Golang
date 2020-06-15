// Author: Olexandr Matveyev
// Date: 08-05-2019
// Palindrome check with recursive approach

package main

import "fmt"

func main() {
	var slice []rune
	var palindrome string
	var is bool

	palindrome = "aaebcbeaa"
	slice = []rune(palindrome)
	is = IsPalindrome(slice)

	fmt.Printf("\nIs the word [%s] palindrom [%t]\n\n", palindrome, is)
}

// IsPalindrome function recursively checks if word is palindrome.
// accepts argument: []rune type.
// return: bool type.
func IsPalindrome(p []rune) bool {

	// will return 'is' variable
	var is bool

	// this 'slice' variable is used to store temporary result of sliced word
	var slice []rune
	slice = p

	// 'first' and 'last' characters/runes of a word
	var first rune
	var last rune

	first = slice[0]
	last = slice[len(slice)-1]

	// if 'first' and 'last' runes are equals keep going slicing word and making recursive calls
	if first == last {

		// if slice size is '1' return true
		// it means we successfully sliced an entire word, tested out each rune,
		// and word is palindrome
		if len(slice) == 1 {
			return true
		}

		// getting lest rune from the slice
		x := len(slice) - 1

		// removing from slice first and very last rune
		slice = append(slice[1:2], slice[2:x]...)

		// making recursive call
		is = IsPalindrome(slice)

	} else {

		// if runes on both ends of a word are not equals, than word is not a palindrome
		return false
	}

	// return final result
	return is
}
