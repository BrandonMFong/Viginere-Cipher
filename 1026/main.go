// Brando
// Learn how to program in go
//  https://learn.go.dev/
//

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var ArgC = len(os.Args)
	if ArgC != 3 {
		// handle error
		fmt.Println("Please pass: Arg1=KeyWord & Arg2=Position")
		os.Exit(2)
	}

	var keyWord = strings.ToUpper(os.Args[1])
	var position = os.Args[2]

	var plainText = "WHRNDJSFKVWKJCZKUCXXVDCVNJRFKEWKJBDJXQRNOKWLNJRFKGSCNRDFSWUELHKKZKRFKVWMRJNOMRUKGNIGAGTRNDJZLJSGVMTCWQUNZLVYYFGHLMSGGYVOWMLKWLEGXRLRFKSJHCJSGKMUJVNUTZFCKGJWNSZVZZRYGWVYYBGLGTFLNZASASUGRLNMJZJJRURWDYTXLGGTFLGCTRZDJUNCDBGSLGCYHVDQUELGCCDDKYTCFNROBWCRNZLSFKXODPKEAKJKCOHRNBMOZUZJCQGMVAMUJKGCRUWRFKQWZLJSZDPKRZDQGVEZNYZFCNOBLTPKRZTLMTHNLVDYRQNDLNMQCGVLGISQDXNENLKNXSFKRZDJBDKZQYGWOYYRWCGZVSRJGAWKJKCGQYTFWLYXLSKYJDTTRZNZDPMQWZRJHKZNVNAMRSDFSGZVSRCSOLXQNDVHBTNLKGQDLNBXNHSFKISQDUQXDYXNXJGRKAMEYNEDZUCQRMSZFZEKCLNNASASGTSGNLKNXSFKBMOZUZJCQGRKGCLDDKNGRLHR"
	var cipherText = ""
	alphabet := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var size = 0 // Also used in modulo calculcations
	var cipherArray [len(alphabet)][2]string
	var key map[string]string
	var plainTextIndex = 0
	var cipherTextIndex = 1
	var tempChar = ""
	var index = 0
	var inOrderAlphabetIndex = 0
	var offsetAlphabetIndex = 0
	var offsetIndex int
	var lengthOfKeyWord = len(keyWord)
	var frequencyDistribution map[string]int // We are counting how many times the letter appears in cipher text
	var tempSize int
	var tempIndex int
	var tempCipherText string
	var cipherTextArray []string
	var garbage int
	var err error

	// Convert from string to int
	offsetIndex, err = strconv.Atoi(position)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	// Print out the variables
	fmt.Println("PROGRAM VARIABLES: ")
	fmt.Println("Keyword: ", keyWord)
	fmt.Println("Position: ", position)
	fmt.Println("Length of sample text: ", len(plainText))
	fmt.Println("Sample text: ")
	fmt.Println(plainText)
	fmt.Println("Alphabet: ", alphabet)

	fmt.Println(("\nCalculating key map...\n"))

	// Get the secret key translation
	size = len(alphabet)
	index = 0
	tempChar = ""
	for index < size {
		tempChar = alphabet[inOrderAlphabetIndex] // load the current alpha bet

		// Load Alphabet regularly
		cipherArray[index][plainTextIndex] = tempChar

		// Load the cipherArray text translation
		if index < lengthOfKeyWord {
			// We can use the index variable to index the keyword string
			// so long as we are under the length of the string
			// We are going to fill in the other column starting at
			// the offset position defined by the user
			tempChar = string(keyWord[index])
			cipherArray[offsetIndex][cipherTextIndex] = tempChar
		} else {
			// In this case we will continue filling in the rest of the
			// cipherArray with the rest of the alphabet
			// We will only insert the characters that were not in the
			// the keyword

			// If the current letter is contained in the keyword string
			// then skip this index
			if strings.Contains(keyWord, alphabet[offsetAlphabetIndex]) {
				offsetAlphabetIndex++
			}

			tempChar = alphabet[offsetAlphabetIndex]
			cipherArray[offsetIndex][cipherTextIndex] = tempChar

			// Incrememnt the offset alphabet if we are done with
			// inserting the keyword
			offsetAlphabetIndex++
		}

		index++
		inOrderAlphabetIndex++

		// Stay within modulo ${size}
		offsetIndex++
		offsetIndex = offsetIndex % size
	}

	// Create a go map (dictionary)
	size = len(cipherArray)
	index = 0
	key = make(map[string]string) // initialize map
	for index < size {
		key[cipherArray[index][plainTextIndex]] = cipherArray[index][cipherTextIndex]
		index++
	}

	fmt.Println("Key: ", key) // Show key

	// Encrypt plain text
	fmt.Println(("\nEncrypting...\n"))

	size = len(plainText)
	index = 0
	tempChar = ""
	for index < size {
		tempChar = string(plainText[index]) // Get the current string

		// In this case, I do not care for cases
		// I just worry about encrypting the text
		tempChar = strings.ToUpper(tempChar) // Convert to upper case

		tempChar = key[tempChar] // encrypt
		cipherText = cipherText + tempChar

		index++
	}

	cipherText = "WHRNDJSFKVWKJCZKUCXXVDCVNJRFKEWKJBDJXQRNOKWLNJRFKGSCNRDFSWUELHKKZKRFKVWMRJNOMRUKGNIGAGTRNDJZLJSGVMTCWQUNZLVYYFGHLMSGGYVOWMLKWLEGXRLRFKSJHCJSGKMUJVNUTZFCKGJWNSZVZZRYGWVYYBGLGTFLNZASASUGRLNMJZJJRURWDYTXLGGTFLGCTRZDJUNCDBGSLGCYHVDQUELGCCDDKYTCFNROBWCRNZLSFKXODPKEAKJKCOHRNBMOZUZJCQGMVAMUJKGCRUWRFKQWZLJSZDPKRZDQGVEZNYZFCNOBLTPKRZTLMTHNLVDYRQNDLNMQCGVLGISQDXNENLKNXSFKRZDJBDKZQYGWOYYRWCGZVSRJGAWKJKCGQYTFWLYXLSKYJDTTRZNZDPMQWZRJHKZNVNAMRSDFSGZVSRCSOLXQNDVHBTNLKGQDLNBXNHSFKISQDUQXDYXNXJGRKAMEYNEDZUCQRMSZFZEKCLNNASASGTSGNLKNXSFKBMOZUZJCQGRKGCLDDKNGRLHR"
	fmt.Println("Cipher text: \n", cipherText)

	size = len(keyWord) // offset
	index = 0           // defines starting point for offest
	for index < size {

		// Go through each letter of the alphabet and count how many times it occurs in cipher text
		tempSize = len(alphabet)
		tempIndex = 0
		tempChar = ""
		frequencyDistribution = make(map[string]int)
		for tempIndex < tempSize {

			// Get only the offset characters
			// start index offset is defined by 'index' var
			// and the offset increments are defined by size
			cipherTextArray = strings.Split(cipherText, "")
			var result []string // init empty slice
			for i := index; i < len(cipherTextArray); i = i + size {
				result = append(result, cipherTextArray[i])
			}
			tempCipherText = strings.Join(result, "")

			tempChar = alphabet[tempIndex]
			frequencyDistribution[tempChar] = strings.Count(tempCipherText, tempChar) // Count how often the current letter appears in ciphertext

			// The way I increment needs to correspond to the key length
			// Let's start off with just indexing through the whole table
			tempIndex++
		}

		// Show the frequency analysis of the cipher text in the respective offset
		fmt.Println("\nFrequency analysis for offset", index, ":")
		for i, index := range rankMapStringInt(frequencyDistribution) {
			fmt.Printf("[%s : %d]\n", index, frequencyDistribution[index])
			garbage = i
		}
		fmt.Println()

		index++
	}

	fmt.Println(garbage) // garbage variable
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
