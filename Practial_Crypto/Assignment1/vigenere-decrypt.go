package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func checkArgs() {
	if len(os.Args) != 3 {
		panic("THe number of arguments is incorrect")
	}
}

func checkForFile(err error) {
	if err != nil {
		panic("the name of the file is incorrect")
	}
}

func checkForKeySize(keyLength float64) {
	if keyLength > 32 {
		panic("key is too big, please try with a key smaller than 32 characters")
	}
}

func main() {
	checkArgs()

	var key, fileName string
	var keyLength float64
	key = strings.ToUpper(os.Args[1])
	fileName = os.Args[2]
	keyLength = float64(len(key))

	cipherText, err := ioutil.ReadFile(fileName)
	checkForFile(err)
	checkForKeySize(keyLength)
	byteCipherText := []byte(cipherText)

	m := make(map[int]byte)

	for i := range byteCipherText {
		keyPlace := int(math.Mod(float64(i), keyLength))
		m[i] = key[keyPlace]    // map position of plaintext to position of key
		byteCipherText[i] -= 65 // reduce to 0 - 25 code
		m[i] -= 65              // reduce to 0 - 25 code
	}

	plainText := make([]byte, len(byteCipherText))

	for i := range byteCipherText {
		var plainTextLetter byte
		if byteCipherText[i]-m[i] > 26 {
			plainTextLetter = byteCipherText[i] - m[i] + 26
		} else {
			plainTextLetter = byteCipherText[i] - m[i]
		}
		plainText[i] = plainTextLetter
		plainText[i] += 65
	}

	fmt.Println(string(plainText))

}
