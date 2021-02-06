// Brando
// Learn how to program in go
//  https://learn.go.dev/
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	var keyWord = "BIKE"
	var position = 5
	var plainText = "Lorem Ipsum is simply dummy text of the printing and typesetting industry Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book It has survived not only five centuries Random character"
	var cipherText = ""
	alphabet := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var size = len(alphabet) // Also used in modulo calculcations
	var cipher [len(alphabet)][2]string
	var plainTextIndex = 0
	var cipherTextIndex = 1
	var tempChar = ""
	var index = 0
	var inOrderAlphabetIndex = 0
	var offsetAlphabetIndex = 0
	var offsetIndex = position
	var lengthOfKeyWord = len(keyWord)

	// Print out the variables
	fmt.Println("Keyword: ", keyWord)
	fmt.Println("Position: ", position)
	fmt.Println("Length of sample text: ", len(plainText))
	fmt.Println("Sample text: ")
	fmt.Println(plainText)
	fmt.Println("Alphabet: ", alphabet)

	fmt.Println(("\nCalculating key...\n"))

	// Get the secret key translation
	for index < size {
		tempChar = alphabet[inOrderAlphabetIndex] // load the current alpha bet

		// Load Alphabet regularly
		cipher[index][plainTextIndex] = tempChar

		// Load the cipher text translation
		if index < lengthOfKeyWord {
			// We can use the index variable to index the keyword string
			// so long as we are under the length of the string
			// We are going to fill in the other column starting at
			// the offset position defined by the user
			tempChar = string(keyWord[index])
			cipher[offsetIndex][cipherTextIndex] = tempChar
		} else {
			// In this case we will continue filling in the rest of the
			// cipher with the rest of the alphabet
			// We will only insert the characters that were not in the
			// the keyword

			// If the current letter is contained in the keyword string
			// then skip this index
			if strings.Contains(keyWord, alphabet[offsetAlphabetIndex]) {
				offsetAlphabetIndex++
			}

			tempChar = alphabet[offsetAlphabetIndex]
			cipher[offsetIndex][cipherTextIndex] = tempChar

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

	fmt.Println("Key: ", cipher)

	fmt.Println(("\nEncrypting...\n"))

	// TODO encrypt plaintext
	size = len(plainText)
	index = 0
	tempChar = ""
	for index < size {
		tempChar = string(plainText[index]) // Get the current string
		cipherText = cipherText + tempChar
		index++
	}
}
