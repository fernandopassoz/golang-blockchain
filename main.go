package main

import "fmt"

type Block struct {
  Data []byte
  PreviousHash []byte
  Hash []byte
}


func main(){

  var chain []Block

  var b Block
  b.Data = []byte("oi")
  chain = append(chain, b)

  fmt.Println("oi")
  fmt.Println(chain)
  for i, n := range chain{
    fmt.Println(i, n)
  }

}
