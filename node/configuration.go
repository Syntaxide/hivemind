package node

import (
  "github.com/syntaxide/hivemind/wire"
  "github.com/syntaxide/hivemind/identity"
)

type NodeConfiguration struct {
  SelfAddress identity.Identity
  Nodes []identity.Identity
  Out wire.WriteWire
}

