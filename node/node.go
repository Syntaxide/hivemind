package node

import (
  "github.com/syntaxide/hivemind/wire"
  "github.com/syntaxide/hivemind/identity"
)

type NodeMode int
const (
    Follower NodeMode = iota
    Leader
    Candidate
)

type NodeState struct {
  Mode NodeMode;
}

type Node struct {
  config NodeConfiguration
  state NodeState
}
var _ wire.ReadWire = (*Node)(nil)

func NewNode(config NodeConfiguration) *Node {
  return &Node {
    config: config,
    state: NodeState {
      Mode: Follower,
    },
  }
}

func (n *Node) OnRead(src identity.Identity, message wire.Message) {}

func (n *Node) Run() {
  wire.Listen(n)
}

func (n *Node) Stop() {}


