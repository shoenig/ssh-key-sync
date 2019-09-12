package netapi

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"gophers.dev/cmds/ssh-key-sync/internal/ssh"
)

// ClientMock implements Client
type ClientMock struct {
	t minimock.Tester

	funcGetKeys          func(user string) (ka1 []ssh.Key, err error)
	inspectFuncGetKeys   func(user string)
	afterGetKeysCounter  uint64
	beforeGetKeysCounter uint64
	GetKeysMock          mClientMockGetKeys
}

// NewClientMock returns a mock for Client
func NewClientMock(t minimock.Tester) *ClientMock {
	m := &ClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetKeysMock = mClientMockGetKeys{mock: m}
	m.GetKeysMock.callArgs = []*ClientMockGetKeysParams{}

	return m
}

type mClientMockGetKeys struct {
	mock               *ClientMock
	defaultExpectation *ClientMockGetKeysExpectation
	expectations       []*ClientMockGetKeysExpectation

	callArgs []*ClientMockGetKeysParams
	mutex    sync.RWMutex
}

// ClientMockGetKeysExpectation specifies expectation struct of the Client.GetKeys
type ClientMockGetKeysExpectation struct {
	mock    *ClientMock
	params  *ClientMockGetKeysParams
	results *ClientMockGetKeysResults
	Counter uint64
}

// ClientMockGetKeysParams contains parameters of the Client.GetKeys
type ClientMockGetKeysParams struct {
	user string
}

// ClientMockGetKeysResults contains results of the Client.GetKeys
type ClientMockGetKeysResults struct {
	ka1 []ssh.Key
	err error
}

// Expect sets up expected params for Client.GetKeys
func (mmGetKeys *mClientMockGetKeys) Expect(user string) *mClientMockGetKeys {
	if mmGetKeys.mock.funcGetKeys != nil {
		mmGetKeys.mock.t.Fatalf("ClientMock.GetKeys mock is already set by Set")
	}

	if mmGetKeys.defaultExpectation == nil {
		mmGetKeys.defaultExpectation = &ClientMockGetKeysExpectation{}
	}

	mmGetKeys.defaultExpectation.params = &ClientMockGetKeysParams{user}
	for _, e := range mmGetKeys.expectations {
		if minimock.Equal(e.params, mmGetKeys.defaultExpectation.params) {
			mmGetKeys.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetKeys.defaultExpectation.params)
		}
	}

	return mmGetKeys
}

// Inspect accepts an inspector function that has same arguments as the Client.GetKeys
func (mmGetKeys *mClientMockGetKeys) Inspect(f func(user string)) *mClientMockGetKeys {
	if mmGetKeys.mock.inspectFuncGetKeys != nil {
		mmGetKeys.mock.t.Fatalf("Inspect function is already set for ClientMock.GetKeys")
	}

	mmGetKeys.mock.inspectFuncGetKeys = f

	return mmGetKeys
}

// Return sets up results that will be returned by Client.GetKeys
func (mmGetKeys *mClientMockGetKeys) Return(ka1 []ssh.Key, err error) *ClientMock {
	if mmGetKeys.mock.funcGetKeys != nil {
		mmGetKeys.mock.t.Fatalf("ClientMock.GetKeys mock is already set by Set")
	}

	if mmGetKeys.defaultExpectation == nil {
		mmGetKeys.defaultExpectation = &ClientMockGetKeysExpectation{mock: mmGetKeys.mock}
	}
	mmGetKeys.defaultExpectation.results = &ClientMockGetKeysResults{ka1, err}
	return mmGetKeys.mock
}

//Set uses given function f to mock the Client.GetKeys method
func (mmGetKeys *mClientMockGetKeys) Set(f func(user string) (ka1 []ssh.Key, err error)) *ClientMock {
	if mmGetKeys.defaultExpectation != nil {
		mmGetKeys.mock.t.Fatalf("Default expectation is already set for the Client.GetKeys method")
	}

	if len(mmGetKeys.expectations) > 0 {
		mmGetKeys.mock.t.Fatalf("Some expectations are already set for the Client.GetKeys method")
	}

	mmGetKeys.mock.funcGetKeys = f
	return mmGetKeys.mock
}

// When sets expectation for the Client.GetKeys which will trigger the result defined by the following
// Then helper
func (mmGetKeys *mClientMockGetKeys) When(user string) *ClientMockGetKeysExpectation {
	if mmGetKeys.mock.funcGetKeys != nil {
		mmGetKeys.mock.t.Fatalf("ClientMock.GetKeys mock is already set by Set")
	}

	expectation := &ClientMockGetKeysExpectation{
		mock:   mmGetKeys.mock,
		params: &ClientMockGetKeysParams{user},
	}
	mmGetKeys.expectations = append(mmGetKeys.expectations, expectation)
	return expectation
}

// Then sets up Client.GetKeys return parameters for the expectation previously defined by the When method
func (e *ClientMockGetKeysExpectation) Then(ka1 []ssh.Key, err error) *ClientMock {
	e.results = &ClientMockGetKeysResults{ka1, err}
	return e.mock
}

// GetKeys implements Client
func (mmGetKeys *ClientMock) GetKeys(user string) (ka1 []ssh.Key, err error) {
	mm_atomic.AddUint64(&mmGetKeys.beforeGetKeysCounter, 1)
	defer mm_atomic.AddUint64(&mmGetKeys.afterGetKeysCounter, 1)

	if mmGetKeys.inspectFuncGetKeys != nil {
		mmGetKeys.inspectFuncGetKeys(user)
	}

	mm_params := &ClientMockGetKeysParams{user}

	// Record call args
	mmGetKeys.GetKeysMock.mutex.Lock()
	mmGetKeys.GetKeysMock.callArgs = append(mmGetKeys.GetKeysMock.callArgs, mm_params)
	mmGetKeys.GetKeysMock.mutex.Unlock()

	for _, e := range mmGetKeys.GetKeysMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ka1, e.results.err
		}
	}

	if mmGetKeys.GetKeysMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetKeys.GetKeysMock.defaultExpectation.Counter, 1)
		mm_want := mmGetKeys.GetKeysMock.defaultExpectation.params
		mm_got := ClientMockGetKeysParams{user}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetKeys.t.Errorf("ClientMock.GetKeys got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetKeys.GetKeysMock.defaultExpectation.results
		if mm_results == nil {
			mmGetKeys.t.Fatal("No results are set for the ClientMock.GetKeys")
		}
		return (*mm_results).ka1, (*mm_results).err
	}
	if mmGetKeys.funcGetKeys != nil {
		return mmGetKeys.funcGetKeys(user)
	}
	mmGetKeys.t.Fatalf("Unexpected call to ClientMock.GetKeys. %v", user)
	return
}

// GetKeysAfterCounter returns a count of finished ClientMock.GetKeys invocations
func (mmGetKeys *ClientMock) GetKeysAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetKeys.afterGetKeysCounter)
}

// GetKeysBeforeCounter returns a count of ClientMock.GetKeys invocations
func (mmGetKeys *ClientMock) GetKeysBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetKeys.beforeGetKeysCounter)
}

// Calls returns a list of arguments used in each call to ClientMock.GetKeys.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetKeys *mClientMockGetKeys) Calls() []*ClientMockGetKeysParams {
	mmGetKeys.mutex.RLock()

	argCopy := make([]*ClientMockGetKeysParams, len(mmGetKeys.callArgs))
	copy(argCopy, mmGetKeys.callArgs)

	mmGetKeys.mutex.RUnlock()

	return argCopy
}

// MinimockGetKeysDone returns true if the count of the GetKeys invocations corresponds
// the number of defined expectations
func (m *ClientMock) MinimockGetKeysDone() bool {
	for _, e := range m.GetKeysMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetKeysMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetKeysCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetKeys != nil && mm_atomic.LoadUint64(&m.afterGetKeysCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetKeysInspect logs each unmet expectation
func (m *ClientMock) MinimockGetKeysInspect() {
	for _, e := range m.GetKeysMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ClientMock.GetKeys with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetKeysMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetKeysCounter) < 1 {
		if m.GetKeysMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ClientMock.GetKeys")
		} else {
			m.t.Errorf("Expected call to ClientMock.GetKeys with params: %#v", *m.GetKeysMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetKeys != nil && mm_atomic.LoadUint64(&m.afterGetKeysCounter) < 1 {
		m.t.Error("Expected call to ClientMock.GetKeys")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetKeysInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ClientMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetKeysDone()
}