package decrypt

func Nimbus(str string) string {
	decryptedStr := ""

	for _, c := range str {
		asciiCode := int(c)
		char := string(asciiCode - 3)
		decryptedStr += char
	}

	return decryptedStr
}
