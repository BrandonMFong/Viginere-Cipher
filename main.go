// Brando
// Learn how to program in go
//  https://learn.go.dev/
//

package main

import (
	"fmt"
)

func main() {
	var keyWord = "BIKE"
	var position = 5
	var sampleText = "Lorem Ipsum is simply dummy text of the printing and typesetting industry Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book It has survived not only five centuries Random character"
	alphabet := [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var size = len(alphabet)
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
	fmt.Println("Length of sample text: ", len(sampleText))
	fmt.Println("Sample text: ")
	fmt.Println(sampleText)
	fmt.Println("Alphabet: ", alphabet)

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
			cipher[offsetIndex][cipherTextIndex] = string(keyWord[index])
		} else {
			// In this case we will continue filling in the rest of the
			// cipher with the rest of the alphabet
			// We will only insert the characters that were not in the
			// the keyword

			// TODO complete filling in the cipher column
		}

		offsetIndex++
		index++
		inOrderAlphabetIndex++
		offsetAlphabetIndex++
	}

	// TODO encrypt plaintext

}
