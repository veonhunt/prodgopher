// A program that declares whether a sentence reads the same
// backward and forward, word for word
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// getInput prompts the user for some text,
// and then reads a line of input from standard input.
// This line of text is then returned.
func getInput() string {
	fmt.Print("Enter a sentence: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()        // reads until end of line
	return scanner.Text() // returns a string containing line
}

func isNotLetter(c rune) bool {
	return !unicode.IsLetter(c)
}

// isPalindromicSentence returns whether or not the given
// sentence is palindromic. To calculate this, it splits
// the string into words, and then creates a reversed copy
// of the word slice. It then checks whether the reversal
// is equal (ignoring case) to the original.
// It also ignores any non-alphanumeric characters.
func isPalindromicSentence(s string) bool {
	// split into words and remove non-alphanumeric characters
	// in one operation by using FieldsFunc and passing in
	// isNotLetter as the function to split on.
	w := strings.FieldsFunc(s, isNotLetter)

	fmt.Println(w)

	// iterate over the words from front to back
	// simultaneously. If we find a word that is not the same
	// as the word at its matching from the back, the
	// sentence is not palindromic.
	l := len(w)
	for i := 0; i < l/2; i++ {
		fw := w[i]     // front word
		bw := w[l-i-1] // back word
		fmt.Printf("%d: (fw %s) (bw %s) (%d | %d)\n", i, fw, bw, i, l-i-1)
		if !strings.EqualFold(fw, bw) { // equals ignore case
			return false
		}
	}

	// all words matched, so the sentence must be palindromic.
	return true
}

func main() {
	for l := getInput(); l != ""; l = getInput() {
		if isPalindromicSentence(l) {
			fmt.Println("... is palindromic!")
		} else {
			fmt.Println("... is not palindromic.")
		}
	}
}
