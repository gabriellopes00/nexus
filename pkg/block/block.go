package block

import (
	"bytes"
	"crypto/sha512"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index     int
	Timestamp time.Time
	PrevHash  []byte
	Hash      []byte
	Data      []byte
	Nonce     int
}

func NewBlock(index int, prevHash []byte, data []byte) *Block {
	block := &Block{
		Index:     index,
		Data:      data,
		PrevHash:  prevHash,
		Timestamp: time.Now(),
		Nonce:     rand.Int(),
	}

	block.Hash = block.GenHash()
	return block
}

func (b *Block) GenHash() []byte {
	payload := bytes.Join(
		[][]byte{
			b.Data,
			b.PrevHash,
			[]byte(strconv.Itoa(b.Nonce)),
			[]byte(b.Timestamp.String()),
		},
		[]byte{},
	)
	hash := sha512.Sum512([]byte(payload))
	return hash[:]
}

func (b *Block) Mine(difficulty int) {

	hashChan := make(chan []byte)

	driveHash := func(channel chan<- []byte) {
		b.Nonce++
		HASH := b.GenHash()
		channel <- HASH
	}

	go func() {
		for !strings.HasPrefix(string(b.Hash), strings.Repeat("0", difficulty)) {
			driveHash(hashChan)
		}
		close(hashChan)
	}()

	for hash := range hashChan {
		b.Hash = hash
	}
}

type BlockParse struct {
	Data      interface{}
	Timestamp string
	Hash      string
	PrevHash  string
}

// func (b *Block) Parse() (*BlockParse, error) {
// 	// var buf bytes.Buffer
// 	// var decodeData []byte = b.Data
// 	// fmt.Println(decodeData)

// 	// dencoder := gob.NewDecoder(&buf)
// 	// err := dencoder.Decode(&b.Data)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return nil, err
// 	// }

// 	blockParse := BlockParse{
// 		Timestamp: b.Timestamp.String(),
// 		Hash:      hex.EncodeToString(b.Hash),
// 		PrevHash:  hex.EncodeToString(b.PrevHash),
// 		Data:      "decodeData",
// 	}

// 	fmt.Println(blockParse)

// 	return &blockParse, nil
// }
