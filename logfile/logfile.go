package logfile
// Assumptions about Logfile behavior:
// Writes can fail. If this happens, append may return error, or not return at
// all. Other readers should be able to read all data besides this block.

// LogfileHeader does not necessarily reflect the in file
// structure, but the data that can be retrieved from it.
type LogfileHeader struct {
  lastOffset int
}

type LogObject {
  termNumber int;
  index int;
  data []byte;
}

type Logfile interface {
  ReadHeader(header interface{}) error;
  Seek(offset int) error;
  Read(offset int, object *LogObject) error;
  Append(data []byte) error;
  // Consider removing. A flaky client may not remember to close,
  // but data should still exist.
  Close() error;
}

// The following objects implement Logfile as an in-memory list.
// This is not fault tolerant, and should only be used for testing /
// development.

var _ (*LogFile) = (nil)(MemoryLogObject);

type MemoryLogfile struct {
  buffer []Objects
}
