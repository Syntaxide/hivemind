package node

import (
  "github.com/syntaxide/hivemind/wire"
  "github.com/syntaxide/hivemind/identity"
)

type NodeConfiguration struct {
  SelfAddress string
  Nodes []identity.Identity
  Out wire.WriteWire
}

