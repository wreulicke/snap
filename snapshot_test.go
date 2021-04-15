package snap

import (
	"bytes"
	"io"
	"testing"

	"github.com/wreulicke/snap/mock"
)

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

func TestAssert(t *testing.T) {
	s := New()
	_ = s.Snapshot("test")
	var called bool
	mockT := &mock.MockTestingT{onError: func(args ...interface{}) { called = true }}
	s.Assert(mockT)
	if !called {
		t.Error("not called onError")
	}
}
