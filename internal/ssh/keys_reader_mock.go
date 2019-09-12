package ssh

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// KeysReaderMock implements KeysReader
type KeysReaderMock struct {
	t minimock.Tester

	funcReadKeys          func(filename string) (ka1 []Key, err error)
	inspectFuncReadKeys   func(filename string)
	afterReadKeysCounter  uint64
	beforeReadKeysCounter uint64
	ReadKeysMock          mKeysReaderMockReadKeys
}

// NewKeysReaderMock returns a mock for KeysReader
func NewKeysReaderMock(t minimock.Tester) *KeysReaderMock {
	m := &KeysReaderMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ReadKeysMock = mKeysReaderMockReadKeys{mock: m}
	m.ReadKeysMock.callArgs = []*KeysReaderMockReadKeysParams{}

	return m
}

type mKeysReaderMockReadKeys struct {
	mock               *KeysReaderMock
	defaultExpectation *KeysReaderMockReadKeysExpectation
	expectations       []*KeysReaderMockReadKeysExpectation

	callArgs []*KeysReaderMockReadKeysParams
	mutex    sync.RWMutex
}

// KeysReaderMockReadKeysExpectation specifies expectation struct of the KeysReader.ReadKeys
type KeysReaderMockReadKeysExpectation struct {
	mock    *KeysReaderMock
	params  *KeysReaderMockReadKeysParams
	results *KeysReaderMockReadKeysResults
	Counter uint64
}

// KeysReaderMockReadKeysParams contains parameters of the KeysReader.ReadKeys
type KeysReaderMockReadKeysParams struct {
	filename string
}

// KeysReaderMockReadKeysResults contains results of the KeysReader.ReadKeys
type KeysReaderMockReadKeysResults struct {
	ka1 []Key
	err error
}

// Expect sets up expected params for KeysReader.ReadKeys
func (mmReadKeys *mKeysReaderMockReadKeys) Expect(filename string) *mKeysReaderMockReadKeys {
	if mmReadKeys.mock.funcReadKeys != nil {
		mmReadKeys.mock.t.Fatalf("KeysReaderMock.ReadKeys mock is already set by Set")
	}

	if mmReadKeys.defaultExpectation == nil {
		mmReadKeys.defaultExpectation = &KeysReaderMockReadKeysExpectation{}
	}

	mmReadKeys.defaultExpectation.params = &KeysReaderMockReadKeysParams{filename}
	for _, e := range mmReadKeys.expectations {
		if minimock.Equal(e.params, mmReadKeys.defaultExpectation.params) {
			mmReadKeys.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmReadKeys.defaultExpectation.params)
		}
	}

	return mmReadKeys
}

// Inspect accepts an inspector function that has same arguments as the KeysReader.ReadKeys
func (mmReadKeys *mKeysReaderMockReadKeys) Inspect(f func(filename string)) *mKeysReaderMockReadKeys {
	if mmReadKeys.mock.inspectFuncReadKeys != nil {
		mmReadKeys.mock.t.Fatalf("Inspect function is already set for KeysReaderMock.ReadKeys")
	}

	mmReadKeys.mock.inspectFuncReadKeys = f

	return mmReadKeys
}

// Return sets up results that will be returned by KeysReader.ReadKeys
func (mmReadKeys *mKeysReaderMockReadKeys) Return(ka1 []Key, err error) *KeysReaderMock {
	if mmReadKeys.mock.funcReadKeys != nil {
		mmReadKeys.mock.t.Fatalf("KeysReaderMock.ReadKeys mock is already set by Set")
	}

	if mmReadKeys.defaultExpectation == nil {
		mmReadKeys.defaultExpectation = &KeysReaderMockReadKeysExpectation{mock: mmReadKeys.mock}
	}
	mmReadKeys.defaultExpectation.results = &KeysReaderMockReadKeysResults{ka1, err}
	return mmReadKeys.mock
}

//Set uses given function f to mock the KeysReader.ReadKeys method
func (mmReadKeys *mKeysReaderMockReadKeys) Set(f func(filename string) (ka1 []Key, err error)) *KeysReaderMock {
	if mmReadKeys.defaultExpectation != nil {
		mmReadKeys.mock.t.Fatalf("Default expectation is already set for the KeysReader.ReadKeys method")
	}

	if len(mmReadKeys.expectations) > 0 {
		mmReadKeys.mock.t.Fatalf("Some expectations are already set for the KeysReader.ReadKeys method")
	}

	mmReadKeys.mock.funcReadKeys = f
	return mmReadKeys.mock
}

// When sets expectation for the KeysReader.ReadKeys which will trigger the result defined by the following
// Then helper
func (mmReadKeys *mKeysReaderMockReadKeys) When(filename string) *KeysReaderMockReadKeysExpectation {
	if mmReadKeys.mock.funcReadKeys != nil {
		mmReadKeys.mock.t.Fatalf("KeysReaderMock.ReadKeys mock is already set by Set")
	}

	expectation := &KeysReaderMockReadKeysExpectation{
		mock:   mmReadKeys.mock,
		params: &KeysReaderMockReadKeysParams{filename},
	}
	mmReadKeys.expectations = append(mmReadKeys.expectations, expectation)
	return expectation
}

// Then sets up KeysReader.ReadKeys return parameters for the expectation previously defined by the When method
func (e *KeysReaderMockReadKeysExpectation) Then(ka1 []Key, err error) *KeysReaderMock {
	e.results = &KeysReaderMockReadKeysResults{ka1, err}
	return e.mock
}

// ReadKeys implements KeysReader
func (mmReadKeys *KeysReaderMock) ReadKeys(filename string) (ka1 []Key, err error) {
	mm_atomic.AddUint64(&mmReadKeys.beforeReadKeysCounter, 1)
	defer mm_atomic.AddUint64(&mmReadKeys.afterReadKeysCounter, 1)

	if mmReadKeys.inspectFuncReadKeys != nil {
		mmReadKeys.inspectFuncReadKeys(filename)
	}

	mm_params := &KeysReaderMockReadKeysParams{filename}

	// Record call args
	mmReadKeys.ReadKeysMock.mutex.Lock()
	mmReadKeys.ReadKeysMock.callArgs = append(mmReadKeys.ReadKeysMock.callArgs, mm_params)
	mmReadKeys.ReadKeysMock.mutex.Unlock()

	for _, e := range mmReadKeys.ReadKeysMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ka1, e.results.err
		}
	}

	if mmReadKeys.ReadKeysMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmReadKeys.ReadKeysMock.defaultExpectation.Counter, 1)
		mm_want := mmReadKeys.ReadKeysMock.defaultExpectation.params
		mm_got := KeysReaderMockReadKeysParams{filename}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmReadKeys.t.Errorf("KeysReaderMock.ReadKeys got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmReadKeys.ReadKeysMock.defaultExpectation.results
		if mm_results == nil {
			mmReadKeys.t.Fatal("No results are set for the KeysReaderMock.ReadKeys")
		}
		return (*mm_results).ka1, (*mm_results).err
	}
	if mmReadKeys.funcReadKeys != nil {
		return mmReadKeys.funcReadKeys(filename)
	}
	mmReadKeys.t.Fatalf("Unexpected call to KeysReaderMock.ReadKeys. %v", filename)
	return
}

// ReadKeysAfterCounter returns a count of finished KeysReaderMock.ReadKeys invocations
func (mmReadKeys *KeysReaderMock) ReadKeysAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReadKeys.afterReadKeysCounter)
}

// ReadKeysBeforeCounter returns a count of KeysReaderMock.ReadKeys invocations
func (mmReadKeys *KeysReaderMock) ReadKeysBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReadKeys.beforeReadKeysCounter)
}

// Calls returns a list of arguments used in each call to KeysReaderMock.ReadKeys.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmReadKeys *mKeysReaderMockReadKeys) Calls() []*KeysReaderMockReadKeysParams {
	mmReadKeys.mutex.RLock()

	argCopy := make([]*KeysReaderMockReadKeysParams, len(mmReadKeys.callArgs))
	copy(argCopy, mmReadKeys.callArgs)

	mmReadKeys.mutex.RUnlock()

	return argCopy
}

// MinimockReadKeysDone returns true if the count of the ReadKeys invocations corresponds
// the number of defined expectations
func (m *KeysReaderMock) MinimockReadKeysDone() bool {
	for _, e := range m.ReadKeysMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReadKeysMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReadKeysCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReadKeys != nil && mm_atomic.LoadUint64(&m.afterReadKeysCounter) < 1 {
		return false
	}
	return true
}

// MinimockReadKeysInspect logs each unmet expectation
func (m *KeysReaderMock) MinimockReadKeysInspect() {
	for _, e := range m.ReadKeysMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to KeysReaderMock.ReadKeys with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReadKeysMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReadKeysCounter) < 1 {
		if m.ReadKeysMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to KeysReaderMock.ReadKeys")
		} else {
			m.t.Errorf("Expected call to KeysReaderMock.ReadKeys with params: %#v", *m.ReadKeysMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReadKeys != nil && mm_atomic.LoadUint64(&m.afterReadKeysCounter) < 1 {
		m.t.Error("Expected call to KeysReaderMock.ReadKeys")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *KeysReaderMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockReadKeysInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *KeysReaderMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *KeysReaderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockReadKeysDone()
}