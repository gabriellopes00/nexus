package block

import (
	"bytes"
	"crypto/sha512"
	"errors"
	"math"
	"math/big"
	"nexus/env"
	"nexus/utils"
)

const (
	maxNonce = math.MaxInt64
)

const hashBytes = 512

var difficulty = env.BLOCKCHAIN_MINING_DIFFICULTY

// ProofOfWork represents a proof-of-work algorithm
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProofOfWork builds and returns a new POW
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)

	if difficulty > hashBytes {
		utils.HandleException(
			errors.New("difficulty must be less then the hash bytes quantity"),
		)
	}

	target.Lsh(target, uint(hashBytes-difficulty))

	return &ProofOfWork{
		Block:  block,
		Target: target,
	}
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			utils.IntToByte(pow.Block.Timestamp.Unix()),
			utils.IntToByte(int64(pow.Block.Index)),
			utils.IntToByte(int64(difficulty)),
			utils.IntToByte(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [64]byte
	var nonce = 0

	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha512.Sum512(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

// Validate validates block's PoW.
// It means that it returns either a block hash is valid or not.
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.Block.Nonce)

	hash := sha512.Sum512(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.Target) == -1 // return either hash has the proof of work requirements or nots
	return isValid
}
