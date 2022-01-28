package blockchain

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type Block struct {
	Nonce        uint32
	Data         []Transaction
	PreviousHash string
	Hash         string
}

type Chain struct {
	Dificulty int
	LastHash  string
	Blocks    []*Block
}

type Transaction struct {
	Data []byte
}

type TransactionPool struct {
	Transactions []*Transaction
}

func Test() {
	fmt.Println("blockchain")
}

func (chain *Chain) WriteBlock(block *Block) error {
	chain.Blocks = append(chain.Blocks, block)
	return errors.New("error")
}

func (chain *Chain) PoW(block *Block) *Block {
	golden := false
	nonce := 0
	start := time.Now().UnixMilli()

	for !golden {
		hash := chain.GenerateHash(block.Data, block.PreviousHash, uint32(nonce))

		if string(block.PreviousHash)[:chain.Dificulty] == string(hash)[:chain.Dificulty] {
			golden = true
			block.Nonce = uint32(nonce)
			block.Hash = hash
			stop := time.Now().UnixMilli()
			// chain.WriteBlock(block)
			fmt.Println(nonce, hash, stop-start)
		} else {
			nonce++
			// fmt.Println("wrong", nonce, hash)
			// fmt.Println(block.Data + block.PreviousHash + strconv.FormatUint(uint64(block.Index), 10) + strconv.Itoa(nonce))
		}
	}

	return block
}

func (chain *Chain) CreateBlock(data []Transaction) Block {
	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	var b Block
	b.Data = data
	b.PreviousHash = lastBlock.Hash
	return b
}

func (chain *Chain) GenerateHash(data []Transaction, previousHash string, nonce uint32) string {

	// Nonce to byte - ok
	n := make([]byte, 32)
	binary.LittleEndian.PutUint32(n, nonce)

	// Trasaction to byte - ok
	var t []byte
	for _, v := range data {
		t = append(t, v.Data...)
	}

	// Previous hash to byte
	p := []byte(previousHash)

	// j := [][]byte{n, t, p}

	// a := bytes.Join(j, []byte{})

	var a []byte
	a = append(a, n...)
	a = append(a, t...)
	a = append(a, p...)

	sum := sha256.Sum256(a)
	return hex.EncodeToString(sum[:])
}

func (chain *Chain) Genesis() error {
	var b Block
	b.PreviousHash = "0000000000000000000000000000000000000000000000000000000000000000"
	b.Hash = chain.GenerateHash([]Transaction{}, b.PreviousHash, uint32(0))
	chain.WriteBlock(&b)
	return errors.New("error")
}
