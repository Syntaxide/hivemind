package node_test

import (
    "testing"

    "github.com/syntaxide/hivemind/node/testutil"
)


func TestSilentWithoutFailure(t *testing.T) {
  // In this test, all other nodes receive messages immediately.
  wire := testutil.NoopWriteWire {}
  nodes := testutil.GenerateNodes(2, &wire)
  stopper := testutil.RunAllNodes(nodes)
  testutil.StopAllNodes(stopper)
}
