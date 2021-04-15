package mock

type MockTestingT struct {
	TestName string
	OnError  func(args ...interface{})
}

func (*MockTestingT) Helper() {}

func (m *MockTestingT) Name() string {
	return m.TestName
}

func (m *MockTestingT) Error(args ...interface{}) {
	if m.OnError != nil {
		m.OnError(args...)
	}
}

func (m *MockTestingT) Errorf(format string, args ...interface{}) {
	if m.OnError != nil {
		m.OnError(append([]interface{}{format}, args...))
	}
}
