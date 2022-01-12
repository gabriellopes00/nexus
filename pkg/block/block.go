package block

type Block struct {
	// Index     int
	// Timestamp time.Time
	PrevHash []byte
	Hash     []byte
	Data     []byte
	Nonce    int
}

func NewBlock(index int, prevHash []byte, data []byte) *Block {
	block := &Block{
		// Index:     index,
		Data:     data,
		PrevHash: prevHash,
		// Timestamp: time.Now(),
		Nonce: 0,
	}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}
