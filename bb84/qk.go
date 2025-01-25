package bb84

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
)

// Basis represents a measurement basis (Z = 0, X = 1)
type Basis int

const (
	ZBasis Basis = iota // Computational basis {|0⟩, |1⟩}
	XBasis              // Hadamard basis {|+⟩, |-⟩}
)

// Participant represents Alice or Bob in the protocol
type Participant struct {
	bits  []int
	bases []Basis
	name  string
}

// BB84Protocol coordinates the execution of the BB84 key distribution
type BB84Protocol struct {
	alice        Participant
	bob          Participant
	sharedKey    []int
	numberOfBits int
	channel      QuantumChannel
	SecureComm   SecureCommunication
}

// QuantumChannel simulates the quantum transmission
type QuantumChannel struct {
	transmittedBits []int
}

// SecureCommunication handles encryption and decryption
type SecureCommunication struct {
	sharedKey []int
	messages  []Message
}

// Message represents an encrypted message
type Message struct {
	Ciphertext string
	Sender     string
}

// NewBB84Protocol initializes the protocol
func NewBB84Protocol(bits int) *BB84Protocol {
	return &BB84Protocol{
		alice:        Participant{name: "Alice"},
		bob:          Participant{name: "Bob"},
		numberOfBits: bits,
		channel:      QuantumChannel{},
	}
}

// generateRandomSequence generates random bits or bases for a participant
func (p *Participant) generateRandomSequence(n int, isBit bool) error {
	if n <= 0 {
		return errors.New("number of bits must be greater than zero")
	}

	if isBit {
		p.bits = make([]int, n)
	} else {
		p.bases = make([]Basis, n)
	}

	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(2))
		if err != nil {
			return fmt.Errorf("failed to generate random sequence: %v", err)
		}
		if isBit {
			p.bits[i] = int(num.Int64())
		} else {
			p.bases[i] = Basis(num.Int64())
		}
	}
	return nil
}

// simulateQuantumTransmission handles qubit preparation and transmission
func (q *QuantumChannel) simulateTransmission(alice Participant, bob Participant, n int) []int {
	transmittedBits := make([]int, n)
	for i := 0; i < n; i++ {
		if alice.bases[i] == bob.bases[i] {
			transmittedBits[i] = alice.bits[i]
		} else {
			randomBit, _ := rand.Int(rand.Reader, big.NewInt(2))
			transmittedBits[i] = int(randomBit.Int64())
		}
	}
	return transmittedBits
}

// siftKey compares bases and generates the shared key
func siftKey(alice Participant, bob Participant, transmittedBits []int, n int) []int {
	sharedKey := []int{}
	for i := 0; i < n; i++ {
		if alice.bases[i] == bob.bases[i] {
			sharedKey = append(sharedKey, transmittedBits[i])
		}
	}
	return sharedKey
}

// InitializeSecureChannel sets up secure communication using the shared key
func (s *SecureCommunication) InitializeSecureChannel(sharedKey []int) {
	s.sharedKey = sharedKey
	s.messages = []Message{}
}

// Encrypt encrypts a message using the shared key
func (s *SecureCommunication) Encrypt(plaintext, sender string) (*Message, error) {
	keyBytes := convertKeyToBytes(s.sharedKey)
	plaintextBytes := []byte(plaintext)

	cipherBytes := xorBytes(plaintextBytes, keyBytes)
	ciphertext := base64.StdEncoding.EncodeToString(cipherBytes)

	msg := &Message{
		Ciphertext: ciphertext,
		Sender:     sender,
	}
	s.messages = append(s.messages, *msg)
	return msg, nil
}

// Decrypt decrypts a message using the shared key
func (s *SecureCommunication) Decrypt(msg *Message) (string, error) {
	keyBytes := convertKeyToBytes(s.sharedKey)

	cipherBytes, err := base64.StdEncoding.DecodeString(msg.Ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %v", err)
	}

	plaintextBytes := xorBytes(cipherBytes, keyBytes)
	return string(plaintextBytes), nil
}

// Run executes the BB84 protocol and establishes a secure channel
func (bb84 *BB84Protocol) Run() error {
	if err := bb84.alice.generateRandomSequence(bb84.numberOfBits, true); err != nil {
		return fmt.Errorf("failed to generate Alice's bits: %v", err)
	}
	if err := bb84.alice.generateRandomSequence(bb84.numberOfBits, false); err != nil {
		return fmt.Errorf("failed to generate Alice's bases: %v", err)
	}
	if err := bb84.bob.generateRandomSequence(bb84.numberOfBits, false); err != nil {
		return fmt.Errorf("failed to generate Bob's bases: %v", err)
	}

	bb84.channel.transmittedBits = bb84.channel.simulateTransmission(bb84.alice, bb84.bob, bb84.numberOfBits)
	bb84.sharedKey = siftKey(bb84.alice, bb84.bob, bb84.channel.transmittedBits, bb84.numberOfBits)

	bb84.SecureComm.InitializeSecureChannel(bb84.sharedKey)
	return nil
}

// convertKeyToBytes converts a bit array to a byte array
func convertKeyToBytes(key []int) []byte {
	byteLen := (len(key) + 7) / 8
	bytes := make([]byte, byteLen)
	for i := 0; i < len(key); i++ {
		if key[i] == 1 {
			bytes[i/8] |= 1 << uint(7-i%8)
		}
	}
	return bytes
}

// xorBytes performs XOR operation on byte slices
func xorBytes(a, b []byte) []byte {
	result := make([]byte, len(a))
	for i := range a {
		result[i] = a[i] ^ b[i%len(b)]
	}
	return result
}
