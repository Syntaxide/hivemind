package wire

import (
  "github.com/syntaxide/hivemind/identity"
)

type Message interface {
  // String produces a human readable string for debugging
  String() string;
  // Serialize produces the bytes to be sent
  Serialize() []byte;
}

// Naming to match the paper
type RequestVoteRequest struct {}
type RequestVoteResponse struct {}
type AppendEntriesRequest struct {}
type AppendEntriesResponse struct {}

type ReadWire interface {
  OnRead(src identity.Identity, message Message)
}

// Wire is the interface used to communicate with other nodes.
type WriteWire interface {
  // Install registers the node which shall process incoming messages.
  Write(dest identity.Identity, message Message)
}

func Listen(receiver ReadWire) {}
