package main

import (
	"fmt"

	"github.com/alaa-elusfy/cryptit/decrypt"
	"github.com/alaa-elusfy/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Alaa Yousfi")
	decryptedStr := decrypt.Nimbus("Dodd#\\rxvil")

	fmt.Println("Encrypt:", encryptedStr)
	fmt.Println("Decrypt:", decryptedStr)
}
