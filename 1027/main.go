// Brando
// References:
// https://stackoverflow.com/questions/19249588/go-programming-generating-combinations
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"sync"
)

func main() {
	var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var keyLength = 5
	var popularLettersInText = "EHSTM"
	var size int
	var index int
	var tempChar string
	var tempIndex int
	var sizeOfAlphabet = len(alphabet)
	var indexOfLetterE = strings.Index(alphabet, "E")
	var stringCandidate string
	var possibleKeys map[int]string
	fileData, err := ioutil.ReadFile("cipher.txt")
	var cipherText string
	var indexCipher int
	var indexKey int
	var cipherSize int
	var keySize int
	var keyCandidate = "PADIO"
	var plainText string
	// var indexOfTheKeyLetter int

	if err != nil {
		log.Panicf("Failed to read file: %s", err)
	}

	// Read the file
	cipherText = string(fileData)
	cipherText = strings.ReplaceAll(cipherText, " ", "")
	fmt.Println("Ciphertext:\n")
	fmt.Println(cipherText, "\n")

	fmt.Println("Based on English text, the following is mapping each ciphertext to the appropriate key letters with 'e'")
	fmt.Println("Mappings: ")

	// Calculate mapping
	index = 0
	size = len(popularLettersInText)
	for index < size {
		tempChar = string(popularLettersInText[index])
		fmt.Print(tempChar, " => ")
		tempIndex = strings.Index(alphabet, tempChar)

		// Subtract mod26 with letter E
		tempIndex -= indexOfLetterE
		tempIndex = tempIndex % sizeOfAlphabet

		fmt.Println(string(alphabet[tempIndex]))
		stringCandidate += string(alphabet[tempIndex]) // save into candidate string

		index++
	}

	fmt.Println("\nString Candidates: ")

	index = 0
	possibleKeys = make(map[int]string)
	for str := range generate(stringCandidate) {

		// Only take the strings that are key length 5
		if len(str) == keyLength {
			fmt.Println(str)
			possibleKeys[index] = str
			index++
		}
	}

	indexCipher = 0
	indexKey = 0
	cipherSize = len(cipherText)
	keySize = len(keyCandidate)
	tempChar = ""
	for indexCipher < cipherSize {

		// Get the index of the car
		// tempChar = string(cipherText[indexCipher])
		// tempChar = string(keyCandidate[indexKey])
		// tempIndex = strings.Index(alphabet, tempChar)

		// Get index of the current key letter
		// decrypt
		tempIndex = strings.Index(alphabet, string(cipherText[indexCipher])) - strings.Index(alphabet, string(keyCandidate[indexKey]))
		tempIndex = int(math.Abs((float64(tempIndex))))

		// apply mod26
		tempIndex = tempIndex % sizeOfAlphabet

		plainText += string(alphabet[tempIndex])

		indexCipher++
		indexKey++
		indexKey %= keySize
	}

	fmt.Println(plainText)
}

func generate(alphabet string) <-chan string {
	c := make(chan string, len(alphabet))

	go func() {
		defer close(c)

		if len(alphabet) == 0 {
			return
		}

		// Use a sync.WaitGroup to spawn permutation
		// goroutines and allow us to wait for them all
		// to finish.
		var wg sync.WaitGroup
		wg.Add(len(alphabet))

		for i := 1; i <= len(alphabet); i++ {
			go func(i int) {
				// Perform permutation on a subset of
				// the alphabet.
				Word(alphabet[:i]).Permute(c)

				// Signal Waitgroup that we are done.
				wg.Done()
			}(i)
		}

		// Wait for all routines to finish.
		wg.Wait()
	}()

	return c
}

type Word []rune

// Permute generates all possible combinations of
// the current word. This assumes the runes are sorted.
func (w Word) Permute(out chan<- string) {
	if len(w) <= 1 {
		out <- string(w)
		return
	}

	// Write first result manually.
	out <- string(w)

	// Find and print all remaining permutations.
	for w.next() {
		out <- string(w)
	}
}

// next performs a single permutation by shuffling characters around.
// Returns false if there are no more new permutations.
func (w Word) next() bool {
	var left, right int

	left = len(w) - 2
	for w[left] >= w[left+1] && left >= 1 {
		left--
	}

	if left == 0 && w[left] >= w[left+1] {
		return false
	}

	right = len(w) - 1
	for w[left] >= w[right] {
		right--
	}

	w[left], w[right] = w[right], w[left]

	left++
	right = len(w) - 1

	for left < right {
		w[left], w[right] = w[right], w[left]
		left++
		right--
	}

	return true
}
