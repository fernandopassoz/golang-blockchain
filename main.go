package main

import (
	"encoding/binary"
	"fmt"
	"shiigo/blockchain"
)

func main() {
	var n []byte
	var nonce uint32
	blockchain.Test()
	binary.LittleEndian.PutUint32(n[0:], nonce)
	fmt.Println(n)
}
