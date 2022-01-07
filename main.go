package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
)

// type Block struct {
// 	Index        uint64
// 	Nonce        int32
// 	Data         []byte
// 	PreviousHash []byte
// 	Hash         []byte
// }
type Block struct {
	Index        uint64
	Nonce        int32
	Data         string
	PreviousHash string
	Hash         string
}

type Chain struct {
	Dificulty int
	LastHash  string
	Blocks    []*Block
}

func (chain *Chain) WriteBlock(block *Block) error {
	chain.Blocks = append(chain.Blocks, block)
	return errors.New("error")
}

func (chain *Chain) PoW(block *Block) error {
	golden := false
	nonce := 0
	sum := sha256.Sum256([]byte(block.Data + block.PreviousHash + strconv.FormatUint(uint64(block.Index), 10) + strconv.Itoa(nonce)))
	var hash string = hex.EncodeToString(sum[:])

	for !golden {
		// if string(block.PreviousHash)[:chain.Dificulty] == string(hash)[:chain.Dificulty] {
		if "00" == string(hash)[:chain.Dificulty] {
			golden = true
			block.Nonce = int32(nonce)
			block.Hash = string(hash)
			chain.WriteBlock(block)
		} else {
			nonce = nonce + 1
			fmt.Println(nonce, hash)
			// fmt.Println(block.Data + block.PreviousHash + strconv.FormatUint(uint64(block.Index), 10) + strconv.Itoa(nonce))
		}

	}
	return errors.New("error")
}

func main() {

	var chain Chain
	chain.LastHash = "1"
	chain.Dificulty = 2

	for i := 0; i < 5; i++ {

		var b Block
		b.Index = uint64(i)
		b.Data = strconv.Itoa(i)
		b.PreviousHash = chain.LastHash
		chain.PoW(&b)
		chain.LastHash = b.Hash
		// sum := sha256.Sum256([]byte(b.Data + b.PreviousHash + strconv.FormatUint(uint64(b.Nonce), 10)))
		// b.Hash = hex.EncodeToString(sum[:])
		// fmt.Println(b.Hash)
		// fmt.Println(b.Hash[:dificulty])

		// chain.WriteBlock(&b)

	}

	// for _, n := range chain.Blocks {
	// 	fmt.Println(n.Index, n.Data, n.Hash, n.PreviousHash)
	// }

	// fmt.Println(chain)

}
