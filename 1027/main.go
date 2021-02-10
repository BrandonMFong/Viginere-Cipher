// Brando
// References:
// https://stackoverflow.com/questions/19249588/go-programming-generating-combinations
// https://yourbasic.org/golang/generate-permutation-slice-string/
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
	// var keyLength = 5
	var popularLettersInText = "EHSTMAQOPB"
	var size int
	var index int
	var tempChar string
	var tempIndex int
	var sizeOfAlphabet = len(alphabet)
	var indexOfLetterE = strings.Index(alphabet, "E")
	var stringCandidate string
	var possibleKeys map[int]string
	var fileData []byte
	var err error
	var cipherText string
	var indexCipher int
	var indexKey int
	var cipherSize int
	var keySize int
	var keyCandidate string
	var plainText string
	var wordsText string
	var indexPossibleKeys int
	var possibleKeysSize int
	var tempString string
	var tempPrevString string
	// var indexOfTheKeyLetter int

	// Read file for cipher text
	fileData, err = ioutil.ReadFile("cipher.txt")
	if err != nil {
		log.Panicf("Failed to read file: %s", err)
	}

	cipherText = string(fileData)
	cipherText = strings.ReplaceAll(cipherText, " ", "")
	fmt.Println("Ciphertext:\n ")
	fmt.Println(cipherText, "\n ")

	// Read file for words
	fileData, err = ioutil.ReadFile("words.txt")
	if err != nil {
		log.Panicf("Failed to read file: %s", err)
	}
	wordsText = string(fileData)

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
		if tempIndex < 0 {
			tempIndex = tempIndex + sizeOfAlphabet
		}

		fmt.Println(string(alphabet[tempIndex]))
		stringCandidate += string(alphabet[tempIndex]) // save into candidate string

		index++
	}

	fmt.Println("\nString Candidates: ")

	index = 0
	possibleKeys = make(map[int]string)
	Perm([]rune(stringCandidate), func(a []rune) {
		tempString = string(a)[0:5]
		if tempPrevString != tempString {
			possibleKeys[index] = tempString
			fmt.Println(tempString)
			tempPrevString = tempString
			index++
		}
	})

	indexPossibleKeys = 0
	possibleKeysSize = len(possibleKeys)
	for indexPossibleKeys < possibleKeysSize {

		keyCandidate = possibleKeys[indexPossibleKeys]
		tempString = strings.ToLower(keyCandidate)
		if strings.Contains(wordsText, tempString) {

			fmt.Println("Testing ", keyCandidate, ": \n ")

			// Decrypt the cipher text
			indexCipher = 0
			indexKey = 0
			cipherSize = len(cipherText)
			keySize = len(keyCandidate)
			tempChar = ""
			for indexCipher < cipherSize {
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
		} else {
			// fmt.Println(keyCandidate, " is not in words text")
		}

		indexPossibleKeys++
	}
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
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
