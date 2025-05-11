package hash

import (
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// Commitment is an instance of a commitment.
type Commitment struct {
	hash  [32]byte
	nonce [32]byte
}

// NewCommitment creates a new instance of a commitment.
func NewCommitment(hash [32]byte, nonce [32]byte) *Commitment {
	return &Commitment{
		hash:  hash,
		nonce: nonce,
	}
}

// Commit creates a commitment that commits to arbitrary data.
// Returns an error if the nonce generation fails.
func Commit(data ...[]byte) (*Commitment, error) {
	var nonce [32]byte
	_, err := io.ReadFull(rand.Reader, nonce[:])
	if err != nil {
		return nil, ErrSampleNonce
	}

	hash := generateCommitment(nonce, data...)

	pair := NewCommitment(hash, nonce)

	return pair, nil
}

// Verify verifies the correctness of a commitment to data.
func Verify(pair *Commitment, data ...[]byte) bool {
	hash := generateCommitment(pair.nonce, data...)

	return hash == pair.hash
}

// generateCommitment creates the commitment by hashing the data and nonce
// using SHA-256.
func generateCommitment(nonce [32]byte, data ...[]byte) [32]byte {
	var bz []byte

	// Data.
	for _, d := range data {
		bz = append(bz, d...)
	}

	// Nonce.
	bz = append(bz, nonce[:]...)

	return sha256.Sum256(bz)
}
