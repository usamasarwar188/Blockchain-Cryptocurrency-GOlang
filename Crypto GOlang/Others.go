// Pass arguments to Satoshi.go (portNumber address of Others , portNumber address of Satoshi)

package main

import (

  a1 "github.com/osamasarwar38/assignment01IBC"
	"fmt"
  "net"
 "encoding/gob"
	"log"
	"os"
//	"strconv"
)

func listen(othersAddress string){
   net.Listen("tcp", ":"+othersAddress)
  fmt.Println("Listening")
}



func printChain(conn net.Conn){
  fmt.Println("Printing the whole chain...")
  var recvdBlock a1.Block
  dec := gob.NewDecoder(conn)
  err := dec.Decode(&recvdBlock)
  if err != nil {
    //handle error
  }

  a1.ListBlocks (&recvdBlock)
}

func connectToAllPeers(conn net.Conn){
  fmt.Println("Connecting to all Peers...")
  var connectedPeers string
  dec := gob.NewDecoder(conn)
  err := dec.Decode(connectedPeers)
  if err != nil {
    //handle error
  }

  fmt.Println(connectedPeers)

}


func main() {
    // Pass arguments to Satoshi.go (portNumber address of Others , portNumber address of Satoshi)
othersAddress := os.Args[1]
satoshiAddress := os.Args[2]
conn, err := net.Dial("tcp", "localhost:"+satoshiAddress)
if err != nil {
  //handle error
}

conn.Write([]byte(""+othersAddress))
ln,err := net.Listen("tcp", ":"+othersAddress)
if err!=nil{
  log.Fatal(err)
  fmt.Println("Error")
}
conn, err2 := ln.Accept()
if err2 != nil {
  log.Println(err2)
  }else{
  //Satoshi sends chain and addresses of connectedPeers

  printChain(conn)
  connectToAllPeers(conn)

}

}
