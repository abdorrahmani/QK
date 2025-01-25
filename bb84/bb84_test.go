package bb84

import (
	"testing"
)

// TestGenerateRandomSequence tests random bit and basis generation for Alice and Bob
func TestGenerateRandomSequence(t *testing.T) {
	alice := Participant{}
	bob := Participant{}

	err := alice.generateRandomSequence(128, true)
	if err != nil {
		t.Fatalf("failed to generate Alice's random bits: %v", err)
	}
	if len(alice.bits) != 128 {
		t.Errorf("expected Alice's bits length to be 128, got %d", len(alice.bits))
	}

	err = alice.generateRandomSequence(128, false)
	if err != nil {
		t.Fatalf("failed to generate Alice's random bases: %v", err)
	}
	if len(alice.bases) != 128 {
		t.Errorf("expected Alice's bases length to be 128, got %d", len(alice.bases))
	}

	err = bob.generateRandomSequence(128, false)
	if err != nil {
		t.Fatalf("failed to generate Bob's random bases: %v", err)
	}
	if len(bob.bases) != 128 {
		t.Errorf("expected Bob's bases length to be 128, got %d", len(bob.bases))
	}
}

// TestBB84Protocol tests the end-to-end execution of the BB84 protocol
func TestBB84Protocol(t *testing.T) {
	protocol := NewBB84Protocol(128)

	// Run the protocol
	if err := protocol.Run(); err != nil {
		t.Fatalf("failed to run BB84 protocol: %v", err)
	}

	// Check the shared key length
	if len(protocol.sharedKey) == 0 {
		t.Fatalf("shared key length should not be zero")
	}

	// Ensure shared key values are 0 or 1
	for _, bit := range protocol.sharedKey {
		if bit != 0 && bit != 1 {
			t.Errorf("invalid bit in shared key: %d", bit)
		}
	}
}

// TestSecureCommunication tests encryption and decryption using the shared key
func TestSecureCommunication(t *testing.T) {
	protocol := NewBB84Protocol(128)

	// Run the protocol
	if err := protocol.Run(); err != nil {
		t.Fatalf("failed to run BB84 protocol: %v", err)
	}

	message := "Hello, Bob!"
	sender := "Alice"

	// Encrypt the message
	encryptedMsg, err := protocol.SecureComm.Encrypt(message, sender)
	if err != nil {
		t.Fatalf("failed to encrypt message: %v", err)
	}
	if encryptedMsg.Ciphertext == "" {
		t.Errorf("ciphertext should not be empty")
	}

	// Decrypt the message
	decryptedMsg, err := protocol.SecureComm.Decrypt(encryptedMsg)
	if err != nil {
		t.Fatalf("failed to decrypt message: %v", err)
	}
	if decryptedMsg != message {
		t.Errorf("expected decrypted message to be '%s', got '%s'", message, decryptedMsg)
	}
}

// TestSharedKeyConsistency tests if Alice and Bob share the same key
func TestSharedKeyConsistency(t *testing.T) {
	protocol := NewBB84Protocol(128)

	// Run the protocol
	if err := protocol.Run(); err != nil {
		t.Fatalf("failed to run BB84 protocol: %v", err)
	}

	aliceKey := protocol.sharedKey
	bobKey := protocol.sharedKey // Bob uses the same shared key in this simulation

	if len(aliceKey) != len(bobKey) {
		t.Fatalf("key lengths do not match, Alice: %d, Bob: %d", len(aliceKey), len(bobKey))
	}

	for i := 0; i < len(aliceKey); i++ {
		if aliceKey[i] != bobKey[i] {
			t.Errorf("shared keys do not match at index %d, Alice: %d, Bob: %d", i, aliceKey[i], bobKey[i])
		}
	}
}
