package block_test

import (
	"fmt"
	"nexus/pkg/block"
	"testing"
)

func TestProofOfWork(t *testing.T) {
	t.Run("run", func(t *testing.T) {
		b := block.NewBlock(0, []byte{}, []byte{})
		pow := block.NewProofOfWork(b)
		fmt.Println(pow.Run())
	})

	t.Run("run-workers", func(t *testing.T) {
		b := block.NewBlock(0, []byte{}, []byte{})
		pow := block.NewProofOfWork(b)
		fmt.Println(pow.Run())
	})
}
