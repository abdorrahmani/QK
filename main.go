package main

import (
	"QK/bb84"
	"fmt"
)

func main() {
	// Initialize the BB84 with 128 bits
	protocol := bb84.NewBB84Protocol(128)

	// Run the bb84
	if err := protocol.Run(); err != nil {
		fmt.Printf("Error running BB84: %v\n", err)
		return
	}

	//Example: Encrypt and decrypt a message
	secureMessage, _ := protocol.SecureComm.Encrypt("Hello, ANOPHEL!", "Alice")
	fmt.Printf("Encrypted Message: %+v\n", secureMessage)

	decryptedMessage, _ := protocol.SecureComm.Decrypt(secureMessage)
	fmt.Printf("Decrypted Message: %s\n", decryptedMessage)
}
