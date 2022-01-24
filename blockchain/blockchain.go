package blockchain

import (
	"fmt"
)

type Block struct {
	Index        uint64
	Nonce        int32
	Data         []byte
	PreviousHash string
	Hash         string
}

type Chain struct {
	Dificulty int
	LastHash  string
	Blocks    []*Block
}

func Test() {
	fmt.Println("blockchain")
}

// func (chain *Chain) WriteBlock(block *Block) error {
// 	chain.Blocks = append(chain.Blocks, block)
// 	return errors.New("error")
// }

// func (chain *Chain) PoW(block *Block) (error, bool){
// 	golden := false
// 	nonce := 0
// 	sum := sha256.Sum256([]byte(block.Data + block.PreviousHash + strconv.FormatUint(uint64(block.Index), 10) + strconv.Itoa(nonce)))
// 	var hash string = hex.EncodeToString(sum[:])

// 	for !golden {
// 		// if string(block.PreviousHash)[:chain.Dificulty] == string(hash)[:chain.Dificulty] {
// 		if "00" == string(hash)[:chain.Dificulty] {
// 			golden = true
// 			block.Nonce = int32(nonce)
// 			block.Hash = hash
// 			chain.WriteBlock(block)
// 		} else {
// 			nonce = nonce + 1
// 			fmt.Println(nonce, hash)
// 			// fmt.Println(block.Data + block.PreviousHash + strconv.FormatUint(uint64(block.Index), 10) + strconv.Itoa(nonce))
// 		}

//
// 	return errors.New("error"), false
// }

// func (chain *Chain) CreateBlock(data []byte) (error, Block) {
// 	var lastBlock = chain.Blocks[len(chain.Blocks) - 1]
//   var b Block
//   b.Data = data
//   b.PreviousHash = lastBlock.Hash
// 	return errors.New("error"), b
// }

// func (chain *Chain) GenerateHash(data []byte, previousHash string, nonce uint32) (error, []byte) {
// 	o := [][]byte{data, []byte(previousHash)}
// 	var n []byte
// 	binary.LittleEndian.PutUint32(n, nonce)
// 	j := bytes.Join(o, n)
// 	sum := sha256.Sum256(j)
// 	return errors.New("error"), sum[:]
// }

// func (chain *Chain) Genesis() error {
// 	var b Block
// 	b.PreviousHash = "0000000000000000000000000000000000000000000000000000000000000000"
// 	b.Index = 0
//   b.Hash = chain.GenerateHash([]byte{}, b.PreviousHash, 0)
// 	chain.WriteBlock(&b)
// 	return errors.New("error")
// }
