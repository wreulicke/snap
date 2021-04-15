package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wreulicke/snap"
	"github.com/wreulicke/snap/mock"
)

func TestSnapshotDir(t *testing.T) {
	os.RemoveAll(".snapshot")
	defer os.RemoveAll(".snapshot")
	s := snap.New()
	s.Snapshot("test")
	var called bool
	s.Assert(&mock.TestingT{OnError: func(args ...interface{}) {
		called = true
	}})
	if !called {
		t.Fatal("not called")
	}
	_, err := os.Stat(filepath.Join(".snapshot", t.Name()))
	if os.IsNotExist(err) {
		t.Error(err)
	}
}
