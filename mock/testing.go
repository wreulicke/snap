package mock

type TestingT struct {
	TestName string
	OnError  func(args ...interface{})
}

func (*TestingT) Helper() {}

func (m *TestingT) Name() string {
	return m.TestName
}

func (m *TestingT) Error(args ...interface{}) {
	if m.OnError != nil {
		m.OnError(args...)
	}
}

func (m *TestingT) Errorf(format string, args ...interface{}) {
	if m.OnError != nil {
		m.OnError(append([]interface{}{format}, args...))
	}
}
