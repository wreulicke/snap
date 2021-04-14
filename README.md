# snap

Snap is snapshot testing library.

## Usage

You can update snapshot via `go test ./... -update`.

### Examples

```go
func TestSnapshot(t *testing.T) {
	s := New()
	_ = s.Snapshot("xxx")
	_ = s.Snapshot("xxx")
	_ = s.Snapshot("xxxy")
	s.Assert(t)
}

func TestSnapshotBytes(t *testing.T) {
	s := New()
	_ = s.SnapshotBytes([]byte("吉野家"))
	s.Assert(t)
}

func TestCanUseAsIoWriter(t *testing.T) {
	s := New()
	b := bytes.NewBufferString("test")
	_, _ = io.Copy(s, b)
	s.Assert(t)
}
```