  // Pass arguments to Satoshi.go (portNumber address of Satoshi , Number of Nodes)
package main

import (

  a1 "github.com/osamasarwar38/assignment01IBC"
	"fmt"
  "net"
 "encoding/gob"
	"log"
	"os"
"strings"
	"strconv"
)
var chainHead *a1.Block
var numberOfNodes string
var connectedPeers = ""




func handleConnection(c net.Conn, nconnectedNodes int) {
    log.Println("A client has connected", c.RemoteAddr())
    buf := make([]byte, 4)

    n, err := c.Read(buf)
    if (err != nil || n == 0) {
      c.Close()
    }

  //  connectedPeers = append(connectedPeers,string(buf[0:n]))
    connectedPeers= connectedPeers+string(buf[0:n])+","

    chainHead = a1.InsertBlock("Others: Node"+ strconv.Itoa(nconnectedNodes), chainHead)

}



func serveAllNodes(nconnectedNodes int) {

    fmt.Println("Now serving all nodes")
    for i:=0;i<nconnectedNodes;i++{

        s := strings.Split(connectedPeers, ",")
        fmt.Println(s)
        conn, err := net.Dial("tcp", "localhost:"+s[i])
        if err != nil {
        //handle error
        fmt.Println("Error")
          continue;
        }

        gobEncoder := gob.NewEncoder(conn)
        err2 := gobEncoder.Encode(chainHead)
        if err2 != nil {
            log.Println(err2)
          }

        gobEncoder2 := gob.NewEncoder(conn)
        err3 := gobEncoder2.Encode(connectedPeers)
          if err3 != nil {
              log.Println(err3)
            }
      }

}





func main(){

  // Pass arguments to Satoshi.go (portNumber address of Satoshi , Number of Nodes)
  satoshiAddress := os.Args[1]
  numberOfNodes := os.Args[2]
  nNodes, err := strconv.Atoi(numberOfNodes)


  ln,err := net.Listen("tcp", ":"+satoshiAddress)
  if err!=nil{
    log.Fatal(err)
    fmt.Println("Error")
  }
  chainHead = a1.InsertBlock("Satoshi: 100YoloCoins->Satoshi", nil)

  var nconnectedNodes=0

    for {
      fmt.Println("Listening for Others...")
        conn, err := ln.Accept()
        if err != nil {
          log.Println(err)
          continue
          }else{

      nconnectedNodes++
      go handleConnection(conn,nconnectedNodes)
      fmt.Println("ConnectedPeers")
      fmt.Println(connectedPeers)
      if (nconnectedNodes>=nNodes){
        serveAllNodes(nconnectedNodes)
        break
      }
      }
    }
  }
