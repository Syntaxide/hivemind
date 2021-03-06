package main

import (
  "fmt"
  "github.com/syntaxide/hivemind/identity"
  "github.com/syntaxide/hivemind/node"
  "github.com/syntaxide/hivemind/wire"
)

type NetWire struct {}
func (NetWire) Write(dest identity.Identity, message wire.Message) {}
var _ wire.WriteWire = (*NetWire)(nil)

func main() {
  fmt.Println("main running")
  configuration := node.NodeConfiguration {
    SelfAddress: NetIdentity("localhost:1233"),
    Nodes: []identity.Identity {
         NetIdentity("localhost:1234"),
         NetIdentity("localhost:1235"),
         NetIdentity("localhost:1236"),
       },
    Out: &NetWire{},
  }
  node := node.NewNode(configuration)
  node.Run()

}
