package mocks

import (
	"context"
	mm_habit "habits/api/pkg/habit"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

type HabitCreatorMock struct {
	t minimock.Tester

	funcAdd          func(ctx context.Context, habit mm_habit.Habit) (err error)
	inspectFuncAdd   func(ctx context.Context, habit mm_habit.Habit)
	afterAddCounter  uint64
	beforeAddCounter uint64
	AddMock          mHabitCreatorMockAdd
}

// NewHabitCreatorMock returns a mock for habit.habitCreator
func NewHabitCreatorMock(t minimock.Tester) *HabitCreatorMock {
	m := &HabitCreatorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddMock = mHabitCreatorMockAdd{mock: m}
	m.AddMock.callArgs = []*HabitCreatorMockAddParams{}

	return m
}

type mHabitCreatorMockAdd struct {
	mock               *HabitCreatorMock
	defaultExpectation *HabitCreatorMockAddExpectation
	expectations       []*HabitCreatorMockAddExpectation

	callArgs []*HabitCreatorMockAddParams
	mutex    sync.RWMutex
}

// HabitCreatorMockAddExpectation specifies expectation struct of the habitCreator.Add
type HabitCreatorMockAddExpectation struct {
	mock    *HabitCreatorMock
	params  *HabitCreatorMockAddParams
	results *HabitCreatorMockAddResults
	Counter uint64
}

// HabitCreatorMockAddParams contains parameters of the habitCreator.Add
type HabitCreatorMockAddParams struct {
	ctx   context.Context
	habit mm_habit.Habit
}

// HabitCreatorMockAddResults contains results of the habitCreator.Add
type HabitCreatorMockAddResults struct {
	err error
}

// Expect sets up expected params for habitCreator.Add
func (mmAdd *mHabitCreatorMockAdd) Expect(ctx context.Context, habit mm_habit.Habit) *mHabitCreatorMockAdd {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("HabitCreatorMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &HabitCreatorMockAddExpectation{}
	}

	mmAdd.defaultExpectation.params = &HabitCreatorMockAddParams{ctx, habit}
	for _, e := range mmAdd.expectations {
		if minimock.Equal(e.params, mmAdd.defaultExpectation.params) {
			mmAdd.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAdd.defaultExpectation.params)
		}
	}

	return mmAdd
}

// Inspect accepts an inspector function that has same arguments as the habitCreator.Add
func (mmAdd *mHabitCreatorMockAdd) Inspect(f func(ctx context.Context, habit mm_habit.Habit)) *mHabitCreatorMockAdd {
	if mmAdd.mock.inspectFuncAdd != nil {
		mmAdd.mock.t.Fatalf("Inspect function is already set for HabitCreatorMock.Add")
	}

	mmAdd.mock.inspectFuncAdd = f

	return mmAdd
}

// Return sets up results that will be returned by habitCreator.Add
func (mmAdd *mHabitCreatorMockAdd) Return(err error) *HabitCreatorMock {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("HabitCreatorMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &HabitCreatorMockAddExpectation{mock: mmAdd.mock}
	}
	mmAdd.defaultExpectation.results = &HabitCreatorMockAddResults{err}
	return mmAdd.mock
}

// Set uses given function f to mock the habitCreator.Add method
func (mmAdd *mHabitCreatorMockAdd) Set(f func(ctx context.Context, habit mm_habit.Habit) (err error)) *HabitCreatorMock {
	if mmAdd.defaultExpectation != nil {
		mmAdd.mock.t.Fatalf("Default expectation is already set for the habitCreator.Add method")
	}

	if len(mmAdd.expectations) > 0 {
		mmAdd.mock.t.Fatalf("Some expectations are already set for the habitCreator.Add method")
	}

	mmAdd.mock.funcAdd = f
	return mmAdd.mock
}

// When sets expectation for the habitCreator.Add which will trigger the result defined by the following
// Then helper
func (mmAdd *mHabitCreatorMockAdd) When(ctx context.Context, habit mm_habit.Habit) *HabitCreatorMockAddExpectation {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("HabitCreatorMock.Add mock is already set by Set")
	}

	expectation := &HabitCreatorMockAddExpectation{
		mock:   mmAdd.mock,
		params: &HabitCreatorMockAddParams{ctx, habit},
	}
	mmAdd.expectations = append(mmAdd.expectations, expectation)
	return expectation
}

// Then sets up habitCreator.Add return parameters for the expectation previously defined by the When method
func (e *HabitCreatorMockAddExpectation) Then(err error) *HabitCreatorMock {
	e.results = &HabitCreatorMockAddResults{err}
	return e.mock
}

// Add implements habit.habitCreator
func (mmAdd *HabitCreatorMock) Add(ctx context.Context, habit mm_habit.Habit) (err error) {
	mm_atomic.AddUint64(&mmAdd.beforeAddCounter, 1)
	defer mm_atomic.AddUint64(&mmAdd.afterAddCounter, 1)

	if mmAdd.inspectFuncAdd != nil {
		mmAdd.inspectFuncAdd(ctx, habit)
	}

	mm_params := &HabitCreatorMockAddParams{ctx, habit}

	// Record call args
	mmAdd.AddMock.mutex.Lock()
	mmAdd.AddMock.callArgs = append(mmAdd.AddMock.callArgs, mm_params)
	mmAdd.AddMock.mutex.Unlock()

	for _, e := range mmAdd.AddMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAdd.AddMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAdd.AddMock.defaultExpectation.Counter, 1)
		mm_want := mmAdd.AddMock.defaultExpectation.params
		mm_got := HabitCreatorMockAddParams{ctx, habit}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAdd.t.Errorf("HabitCreatorMock.Add got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAdd.AddMock.defaultExpectation.results
		if mm_results == nil {
			mmAdd.t.Fatal("No results are set for the HabitCreatorMock.Add")
		}
		return (*mm_results).err
	}
	if mmAdd.funcAdd != nil {
		return mmAdd.funcAdd(ctx, habit)
	}
	mmAdd.t.Fatalf("Unexpected call to HabitCreatorMock.Add. %v %v", ctx, habit)
	return
}

// AddAfterCounter returns a count of finished HabitCreatorMock.Add invocations
func (mmAdd *HabitCreatorMock) AddAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.afterAddCounter)
}

// AddBeforeCounter returns a count of HabitCreatorMock.Add invocations
func (mmAdd *HabitCreatorMock) AddBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.beforeAddCounter)
}

// Calls returns a list of arguments used in each call to HabitCreatorMock.Add.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAdd *mHabitCreatorMockAdd) Calls() []*HabitCreatorMockAddParams {
	mmAdd.mutex.RLock()

	argCopy := make([]*HabitCreatorMockAddParams, len(mmAdd.callArgs))
	copy(argCopy, mmAdd.callArgs)

	mmAdd.mutex.RUnlock()

	return argCopy
}

// MinimockAddDone returns true if the count of the Add invocations corresponds
// the number of defined expectations
func (m *HabitCreatorMock) MinimockAddDone() bool {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddInspect logs each unmet expectation
func (m *HabitCreatorMock) MinimockAddInspect() {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to HabitCreatorMock.Add with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		if m.AddMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to HabitCreatorMock.Add")
		} else {
			m.t.Errorf("Expected call to HabitCreatorMock.Add with params: %#v", *m.AddMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && mm_atomic.LoadUint64(&m.afterAddCounter) < 1 {
		m.t.Error("Expected call to HabitCreatorMock.Add")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *HabitCreatorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *HabitCreatorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *HabitCreatorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddDone()
}