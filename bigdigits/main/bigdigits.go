package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The goal for this small application is to print out equivalent numbers of entered numbers
// only in much bigger format, in the console

// Here must be two 2 options:
// 1: Define big numbers with multidimensional slices
// 2: Read from text file big numbers

// bigDigits function will return big number and as argument will accept regular number
func bigDigits(n int) []string {
	var tmp []string

	// Slice of numbers
	digits := [][]string{
		{
			"  000  ",
			" 0   0 ",
			"0     0",
			"0     0",
			"0     0",
			" 0   0 ",
			"  000  "},
		{
			"   1   ",
			"  11   ",
			"   1   ",
			"   1   ",
			"   1   ",
			"   1   ",
			"  111  "},
		{
			"  222  ",
			"2    2 ",
			"    2  ",
			"   2   ",
			"  2    ",
			" 2     ",
			"22222  "},
		{
			"  333  ",
			" 3   3 ",
			"     3 ",
			"   33  ",
			"     3 ",
			" 3   3 ",
			"  333  "},
		{
			"   4   ",
			"  44   ",
			" 4 4   ",
			"4  4   ",
			"444444 ",
			"   4   ",
			"   4   "},
		{
			" 55555 ",
			" 5     ",
			" 5     ",
			"  555  ",
			"     5 ",
			" 5   5 ",
			"  555  "},
		{
			"  666  ",
			" 6     ",
			" 6     ",
			"  666  ",
			" 6   6 ",
			" 6   6 ",
			"  666  "},
		{
			" 77777 ",
			"     7 ",
			"    7  ",
			"   7   ",
			"  7    ",
			" 7     ",
			" 7     "},
		{
			" 8888  ",
			"8    8 ",
			"8    8 ",
			" 8888  ",
			"8    8 ",
			"8    8 ",
			" 8888  "},
		{
			" 9999  ",
			"9    9 ",
			"9    9 ",
			" 9999  ",
			"     9 ",
			"     9 ",
			"     9 "}}

	// Test
	/*
		for _, col := range digits[0] {
			fmt.Printf("%s\n", col)
		}
	*/

	// Find specific sub-slice
	tmp = digits[n]
	return tmp
}

// buildDigits function is used to build big digit or a set of big digits
// currently used to build only one big digit
func buildDigits(n int) []string {

	var size int

	// the size of slice should not be 0, at minimum it should ne 7
	// ------------------------------------------------------------ //
	if n == 0 {
		size = 7
	} else {
		size = n
	}
	// ------------------------------------------------------------ //

	var str string
	var tmp = make([]string, size, size*2)

	// Generating single digit
	// ------------------------------------------------------------ //
	for _, col := range bigDigits(n) {
		str += col + "\n"
	}
	tmp[0] = str
	// ------------------------------------------------------------ //

	return tmp
}

func printDigits(n int) {
	str := buildDigits(n)
	fmt.Printf("%s", str[0])
}

func main() {

	/*
		for i := 0; i < 10; i++ {
			printDigits(i)
		}
	*/

	// Get arguments
	// ------------------------------------------------------------ //
	var arg string
	arg = strings.Join(os.Args[1:], "")

	if len(arg) == 1 {
		fmt.Printf("\nOne digit: %s\n\n", arg)
	} else if len(arg) > 1 {
		fmt.Printf("\nString of digits: %s\n\n", arg)
	}
	// ------------------------------------------------------------ //

	// Convert string to int
	// ------------------------------------------------------------ //
	var n int64
	n, _ = strconv.ParseInt(arg, 10, 64)

	var num int
	num = int(n)
	// ------------------------------------------------------------ //

	printDigits(num)

}
