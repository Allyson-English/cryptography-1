package week1

import (
	"encoding/hex"
	"fmt"
)

func OTP() {
	message1 := fmt.Sprintf("%x", "Can you please send me the secret recipe for your grandma's famous chocolate chip cookies?")
	message2 := fmt.Sprintf("%x", "Sure, all you need is: butter, eggs, salt, flour, chocolate chips and a whole lot of love.")

	xorResult, err := xor(message1, message2)
	if err != nil {
		panic(err)
	}

	possiblePlaintext1 := " secret recipe "
	plaintext3 := applyPossiblePlaintexts(xorResult, possiblePlaintext1)

	for _, x := range plaintext3 {
		fmt.Println(x)
	}

}

func applyPossiblePlaintexts(xordPlaintexts, plaintextGuess string) []string {
	afterXorWGuess := make([]string, 0)
	for i := 0; i < len(xordPlaintexts); i++ {
		pt := ""
		slide := 0
		for x := 0; x < len(xordPlaintexts); x++ {
			ix := i + x
			if slide >= len(plaintextGuess) {
				slide = 0
			}
			if ix >= len(xordPlaintexts) {
				break
			}
			xor := string(xordPlaintexts[ix] ^ plaintextGuess[slide])
			pt += xor
			slide++
		}

		afterXorWGuess = append(afterXorWGuess, pt)
	}
	return afterXorWGuess
}

func xor(msg1, msg2 string) (string, error) {
	bytes1, err := hex.DecodeString(msg1)
	if err != nil {
		return "", err
	}
	bytes2, err := hex.DecodeString(msg2)
	if err != nil {
		return "", err
	}

	length := len(bytes1)
	if len(bytes2) < length {
		length = len(bytes2)
	}

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = bytes1[i] ^ bytes2[i]
	}

	return string(result), nil
}

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}
