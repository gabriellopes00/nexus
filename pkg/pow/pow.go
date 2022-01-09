package pow

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
	"nexus/pkg/block"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 16

// POW represents a proof-of-work
type POW struct {
	block  *block.Block
	target *big.Int
}

// NewProofOfWork builds and returns a new POW
func NewPOW(b *block.Block) *POW {
	target := big.NewInt(1)
	target.Lsh(target, uint(512-targetBits))

	return &POW{b, target}
}

func (pow *POW) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevHash,
			pow.block.Data,
			IntToByte(pow.block.Timestamp.Unix()),
			IntToByte(int64(targetBits)),
			IntToByte(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run performs a proof-of-work
func (pow *POW) Mine() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		if math.Remainder(float64(nonce), 100000) == 0 {
			fmt.Printf("\r%x", hash)
		}
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

// Validate validates block's PoW
func (pow *POW) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha512.Sum512(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

func IntToByte(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
