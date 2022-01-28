package main

import (
	"fmt"
	"shiigo/blockchain"
)

func main() {
	// n := make([]byte, 32)
	// var nonce uint32
	// nonce = 123456
	// blockchain.Test()
	// // binary.LittleEndian.PutUint32(n[0:], nonce)
	// binary.LittleEndian.PutUint32(n, nonce)
	// // binary.Write(n, binary.LittleEndian, nonce)
	// // binary.PutUvarint(n, uint64(nonce))
	// fmt.Println(n)

	var bc blockchain.Chain
	var blk blockchain.Block
	var ts blockchain.Transaction
	var ts2 blockchain.Transaction
	var tss []blockchain.Transaction

	ts.Data = []byte("Hello World!")
	ts2.Data = []byte("Another World!!!")
	tss = append(tss, ts, ts2)

	bc.Dificulty = 6

	blk.Data = tss
	blk.PreviousHash = "0000000000000000000000000000000000000000000000000000000000000000"

	bc.PoW(&blk)

	fmt.Println(blk)
}
