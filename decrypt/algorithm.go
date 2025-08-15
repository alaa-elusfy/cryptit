// Decrypt package consists of all the decryption algorithms
package decrypt

// Decrypts by reducing the ASCII code by 3 for each character
func Nimbus(str string) string {
	decryptedStr := ""

	for _, c := range str {
		asciiCode := int(c)
		char := string(asciiCode - 3)
		decryptedStr += char
	}

	return decryptedStr
}
