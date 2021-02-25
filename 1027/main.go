// Brando
// References:
// https://stackoverflow.com/questions/19249588/go-programming-generating-combinations
// https://yourbasic.org/golang/generate-permutation-slice-string/
/*
	Possible words:
		ADOPT
*/
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

const columnsToConsider uint = 1

func main() {
	var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var keyLength = 6
	var popularLettersInText string
	// var popularLettersInText = "EHSTDHIX"
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
	var plainTextArray []string
	var indexPossibleKeys int
	var possibleKeysSize int
	var tempString string
	var wordsText []string
	var xIndex int
	var yIndex int
	var plainTextArrayIndex int
	var dontInsert bool

	// Read file for cipher text
	fileData, err = ioutil.ReadFile("cipher3.txt")
	if err != nil {
		log.Panicf("Failed to read file: %s", err)
	}

	cipherText = string(fileData)
	cipherText = strings.ReplaceAll(cipherText, " ", "")
	fmt.Println("Ciphertext:\n ")
	fmt.Println(cipherText, "\n ")

	// Read file for words
	wordsText, err = readLines("words.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("\n\nBased on English text, the following is mapping each ciphertext to the appropriate key letters with 'e'")
	fmt.Println("Mappings (this fulfills 10.2.7c): ")

	popularLettersInText = getPopularLettersInText(cipherText, keyLength)

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
	fmt.Println("Final: ", stringCandidate)

	fmt.Println("\nCalculating string Candidates...\n ")

	index = 0
	possibleKeys = make(map[int]string)
	tempIndex = 0
	dontInsert = false
	Perm([]rune(stringCandidate), func(a []rune) {
		tempString = string(a)[0:keyLength]

		// Test if we already have this word
		tempIndex = 0
		size = len(possibleKeys)
		dontInsert = false
		for tempIndex < size {
			if tempString == possibleKeys[tempIndex] {
				dontInsert = true
				break
			}
			tempIndex++
		}

		// If not, then include as possible keys
		if !dontInsert {
			possibleKeys[index] = tempString
			index++
		}
	})

	fmt.Println("\nRunning decryption... \n ")

	indexPossibleKeys = 0
	possibleKeysSize = len(possibleKeys)
	plainTextArrayIndex = 0
	for indexPossibleKeys < possibleKeysSize {

		keyCandidate = possibleKeys[indexPossibleKeys]
		tempString = strings.ToLower(keyCandidate)
		for i, word := range wordsText {
			if word == tempString {
				fmt.Println("Decrypting '", word, "' in line ", i)

				// Decrypt the cipher text
				indexCipher = 0
				indexKey = 0
				cipherSize = len(cipherText)
				keySize = len(keyCandidate)
				tempChar = ""
				plainText = ""
				for indexCipher < cipherSize {
					// Get index of the current key letter
					// decrypt
					xIndex = strings.Index(alphabet, string(cipherText[indexCipher]))
					yIndex = strings.Index(alphabet, string(keyCandidate[indexKey]))
					tempIndex = xIndex - yIndex

					// apply mod26
					tempIndex = tempIndex % sizeOfAlphabet
					if tempIndex < 0 {
						tempIndex = tempIndex + sizeOfAlphabet
					}

					plainText += string(alphabet[tempIndex])

					indexCipher++
					indexKey++
					indexKey %= keySize
				}
				plainText = "[" + keyCandidate + "] " + plainText
				plainTextArray = append(plainTextArray, plainText)
				plainTextArrayIndex++
			} else {
				// fmt.Println(keyCandidate, " is not in words text")
			}
		}

		indexPossibleKeys++
	}

	fmt.Println()
	for i, text := range plainTextArray {
		fmt.Println(i, ": ", text, "\n ")
	}
}

func getPopularLettersInText(cipherText string, keyLength int) string {
	// var length uint = 4
	var columns uint = columnsToConsider
	var result string
	var size uint
	var index uint
	var tempSize uint
	var tempIndex uint
	var tempChar string
	var frequencyDistribution map[string]int
	var cipherTextArray []string
	var tempCipherText string // Contains ciphertext at offset
	var temp []string         // init empty slice

	size = uint(keyLength) // offset
	index = 0              // defines starting point for offest
	for index < size {
		temp = temp[:0] // Clear array

		// Get only the offset characters
		// start index offset is defined by 'index' var
		// and the offset increments are defined by size
		cipherTextArray = strings.Split(cipherText, "")
		for i := index; i < uint(len(cipherTextArray)); i = i + size {
			temp = append(temp, cipherTextArray[i])
		}
		tempCipherText = strings.Join(temp, "")

		// Go through each letter of the alphabet and count how many times it occurs in cipher text
		tempSize = uint(len(alphabet))
		tempIndex = 0
		tempChar = ""
		frequencyDistribution = make(map[string]int)
		for tempIndex < tempSize {

			tempChar = alphabet[tempIndex]
			frequencyDistribution[tempChar] = strings.Count(tempCipherText, tempChar) // Count how often the current letter appears in ciphertext

			// The way I increment needs to correspond to the key length
			// Let's start off with just indexing through the whole table
			tempIndex++
		}

		// Get the popular text into a string
		tempIndex = 0
		for _, index := range rankMapStringInt(frequencyDistribution) {
			if tempIndex < columns {
				result = result + index
			} else {
				break
			}
			tempIndex++
			// break // just get one
		}
		index++
	}

	return result
}

// Credits to: https://stackoverflow.com/a/60513740/12135693
func rankMapStringInt(values map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	ranked := make([]string, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

/* Credits to:  https://yourbasic.org/golang/generate-permutation-slice-string/ */

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
