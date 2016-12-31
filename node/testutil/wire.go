package testutil

import (
  "fmt"
  "github.com/syntaxide/hivemind/node"
  "github.com/syntaxide/hivemind/wire"
  "github.com/syntaxide/hivemind/identity"
)

type NoopWriteWire struct {
}

var _ wire.WriteWire = (*NoopWriteWire)(nil)

func (n NoopWriteWire) Write(_ identity.Identity, _ wire.Message) {}

// TODO: split testutil
func GenerateNodes(num int, write wire.WriteWire) ([]*node.Node) {
  nodeNames := make([]identity.Identity, num)
  for i := 0; i < num; i++ {
    nodeNames[i] = identity.NetIdentity(fmt.Sprintf("testhostid=%d", i))
  }

  nodes := make([]*node.Node, num)
  for i, name := range(nodeNames) {
    rest := make([]identity.Identity, num-1)
    rest = append(rest, nodeNames[:i]...)
    rest = append(rest, nodeNames[i+1:]...)

    nodes[i] = node.NewNode(node.NodeConfiguration {
      SelfAddress: name,
      Nodes: rest,
      Out: write,
    })
  }
  return nodes
}

// start a goroutine for each node,
// returns a channel that, when written to, shuts down all nodes.
func RunAllNodes(nodes []*node.Node) (chan struct{}) {
  stopChan := make(chan struct{})
  for i := range(nodes) {
    go func(i int) {
      nodes[i].Run()
    }(i)
  }

  go func() {
    select {
      case <- stopChan:
        for i := range(nodes) {
          nodes[i].Stop()
        }
    }
  }()
  return stopChan
}

func StopAllNodes(signal chan struct{}) {
    signal <- struct{}{}
}
