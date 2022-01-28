/*
Problem #1
Create an implementation of the rotational cipher, also sometimes called the Caesar cipher.

The Caesar cipher is a simple shift cipher that relies on transposing all the letters in the alphabet using an integer key between 0 and 26. Using a key of 0 or 26 will always yield the same output due to modular arithmetic. The letter is shifted for as many values as the value of the key.

The general notation for rotational ciphers is ROT + <key>. The most commonly used rotational cipher is ROT13.

A ROT13 on the Latin alphabet would be as follows:

Plain:  abcdefghijklmnopqrstuvwxyz
Cipher: nopqrstuvwxyzabcdefghijklm

It is stronger than the Atbash cipher because it has 27 possible keys, and 25 usable keys.

Ciphertext is written out in the same formatting as the input including spaces and punctuation.
Examples

    ROT5 omg gives trl
    ROT0 c gives c
    ROT26 Cool gives Cool
    ROT13 The quick brown fox jumps over the lazy dog. gives Gur dhvpx oebja sbk whzcf bire gur ynml qbt.
    ROT13 Gur dhvpx oebja sbk whzcf bire gur ynml qbt. gives The quick brown fox jumps over the lazy dog.
------------------
-------------------------
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func RotationCipher(original string) {
	var K int
	fmt.Println("Enter a key")
	fmt.Scan(&K)
	for _, value := range original {

		if unicode.IsSpace(value) == true {
			fmt.Printf("%c", value)
		} else if unicode.IsLower(value) == true {
			fmt.Printf("%c", rune((int)(value)+K-97)%26+97)
		} else if unicode.IsUpper(value) == true {
			fmt.Printf("%c", rune((int(value)+K-65)%26+65))
		} else {
			fmt.Printf("%c", value)
		}
	}
}

func main() {
	fmt.Println(("Enter a string"))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	RotationCipher(line)
}
