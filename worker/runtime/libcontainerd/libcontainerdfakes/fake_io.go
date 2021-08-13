// Code generated by counterfeiter. DO NOT EDIT.
package libcontainerdfakes

import (
	"sync"

	"github.com/containerd/containerd/cio"
)

type FakeIO struct {
	CancelStub        func()
	cancelMutex       sync.RWMutex
	cancelArgsForCall []struct {
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	ConfigStub        func() cio.Config
	configMutex       sync.RWMutex
	configArgsForCall []struct {
	}
	configReturns struct {
		result1 cio.Config
	}
	configReturnsOnCall map[int]struct {
		result1 cio.Config
	}
	WaitStub        func()
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIO) Cancel() {
	fake.cancelMutex.Lock()
	fake.cancelArgsForCall = append(fake.cancelArgsForCall, struct {
	}{})
	stub := fake.CancelStub
	fake.recordInvocation("Cancel", []interface{}{})
	fake.cancelMutex.Unlock()
	if stub != nil {
		fake.CancelStub()
	}
}

func (fake *FakeIO) CancelCallCount() int {
	fake.cancelMutex.RLock()
	defer fake.cancelMutex.RUnlock()
	return len(fake.cancelArgsForCall)
}

func (fake *FakeIO) CancelCalls(stub func()) {
	fake.cancelMutex.Lock()
	defer fake.cancelMutex.Unlock()
	fake.CancelStub = stub
}

func (fake *FakeIO) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIO) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeIO) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeIO) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIO) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIO) Config() cio.Config {
	fake.configMutex.Lock()
	ret, specificReturn := fake.configReturnsOnCall[len(fake.configArgsForCall)]
	fake.configArgsForCall = append(fake.configArgsForCall, struct {
	}{})
	stub := fake.ConfigStub
	fakeReturns := fake.configReturns
	fake.recordInvocation("Config", []interface{}{})
	fake.configMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIO) ConfigCallCount() int {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	return len(fake.configArgsForCall)
}

func (fake *FakeIO) ConfigCalls(stub func() cio.Config) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = stub
}

func (fake *FakeIO) ConfigReturns(result1 cio.Config) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	fake.configReturns = struct {
		result1 cio.Config
	}{result1}
}

func (fake *FakeIO) ConfigReturnsOnCall(i int, result1 cio.Config) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	if fake.configReturnsOnCall == nil {
		fake.configReturnsOnCall = make(map[int]struct {
			result1 cio.Config
		})
	}
	fake.configReturnsOnCall[i] = struct {
		result1 cio.Config
	}{result1}
}

func (fake *FakeIO) Wait() {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
	}{})
	stub := fake.WaitStub
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if stub != nil {
		fake.WaitStub()
	}
}

func (fake *FakeIO) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeIO) WaitCalls(stub func()) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeIO) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cancelMutex.RLock()
	defer fake.cancelMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIO) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ cio.IO = new(FakeIO)