package snap

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pmezard/go-difflib/difflib"
)

type SnapshotTester struct {
	config *Config
	buffer bytes.Buffer
}

type testingT interface {
	Helper()
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Name() string
}

// assert SnapshotTester implements io.Writer.
var _ io.Writer = (*SnapshotTester)(nil)

// New returns new SnapshotTester.
func New(options ...Option) *SnapshotTester {
	return &SnapshotTester{
		config: NewConfig(options...),
		buffer: bytes.Buffer{},
	}
}

// Snapshot writes value as snapshot.
func (s *SnapshotTester) Snapshot(value string) error {
	_, err := fmt.Fprintln(&s.buffer, value)
	if err != nil {
		return err
	}
	return nil
}

// SnapshotBytes writes bytes as snapshot.
func (s *SnapshotTester) SnapshotBytes(bs []byte) error {
	if _, err := s.buffer.Write(bs); err != nil {
		return err
	}
	return nil
}

// Write writes bytes.
func (s *SnapshotTester) Write(p []byte) (n int, err error) {
	return s.buffer.Write(p)
}

func (s *SnapshotTester) shouldUpdate() bool {
	return s.config.UpdateSnapshot
}

func (s *SnapshotTester) takeSnapshot() string {
	return s.buffer.String()
}

func (s *SnapshotTester) readSnapshot(snapshotName string) (string, error) {
	f, err := os.Open(filepath.Join(".snapshot", snapshotName))
	if err != nil {
		return "", err
	}
	bs, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (s *SnapshotTester) writeSnapshot(snapshotName string, snapshot string) error {
	path := filepath.Join(".snapshot", snapshotName)
	_ = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.WriteString(f, snapshot)
	return err
}

// Assert reports error if current snapshot differs from older one.
func (s *SnapshotTester) Assert(t testingT) {
	t.Helper()
	newSnap := s.takeSnapshot()

	snap, err := s.readSnapshot(t.Name())
	if os.IsNotExist(err) {
		if err := s.writeSnapshot(t.Name(), newSnap); err != nil {
			t.Error(err)
		} else {
			t.Error("created snapshot")
		}
		return
	}

	diff, foundDiff := s.compare(newSnap, snap)
	if !foundDiff {
		return
	}
	if s.shouldUpdate() {
		err := s.writeSnapshot(t.Name(), newSnap)
		if err != nil {
			t.Error(err)
		} else {
			t.Error("updated snapshot")
		}
	} else {
		t.Errorf("found difference between current snapshot. please update snapshot if this difference is intended\n\n%s", diff)
	}
}

func (s *SnapshotTester) compare(new string, old string) (string, bool) {
	if new == old {
		return "", false
	}
	diff, _ := difflib.GetUnifiedDiffString(
		difflib.UnifiedDiff{
			A:        difflib.SplitLines(old),
			B:        difflib.SplitLines(new),
			FromFile: "Previous",
			FromDate: "",
			ToFile:   "Current",
		})
	return diff, true
}
