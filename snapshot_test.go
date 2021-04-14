package snap

import (
	"bytes"
	"io"
	"testing"
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

type mockTestingT struct {
	name    string
	onError func(args ...interface{})
}

func (*mockTestingT) Helper() {}

func (m *mockTestingT) Name() string {
	return m.name
}

func (m *mockTestingT) Error(args ...interface{}) {
	if m.onError != nil {
		m.onError(args...)
	}
}

func (m *mockTestingT) Errorf(format string, args ...interface{}) {
	if m.onError != nil {
		m.onError(append([]interface{}{format}, args...))
	}
}

func TestAssert(t *testing.T) {
	s := New()
	_ = s.Snapshot("test")
	var called bool
	mockT := &mockTestingT{onError: func(args ...interface{}) { called = true }}
	s.Assert(mockT)
	if !called {
		t.Error("not called onError")
	}
}
