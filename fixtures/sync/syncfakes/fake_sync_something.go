// Code generated by counterfeiter. DO NOT EDIT.
package syncfakes

import (
	"sync"

	synca "github.com/jsokel/counterfeiter/v6/fixtures/sync"
)

type FakeSyncSomething struct {
	DoASliceStub        func([]byte)
	doASliceMutex       sync.RWMutex
	doASliceArgsForCall []struct {
		arg1 []byte
	}
	DoAnArrayStub        func([4]byte)
	doAnArrayMutex       sync.RWMutex
	doAnArrayArgsForCall []struct {
		arg1 [4]byte
	}
	DoNothingStub        func()
	doNothingMutex       sync.RWMutex
	doNothingArgsForCall []struct {
	}
	DoThingsStub        func(string, uint64) (int, error)
	doThingsMutex       sync.RWMutex
	doThingsArgsForCall []struct {
		arg1 string
		arg2 uint64
	}
	doThingsReturns struct {
		result1 int
		result2 error
	}
	doThingsReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSyncSomething) DoASlice(arg1 []byte) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.doASliceMutex.Lock()
	fake.doASliceArgsForCall = append(fake.doASliceArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.DoASliceStub
	fake.recordInvocation("DoASlice", []interface{}{arg1Copy})
	fake.doASliceMutex.Unlock()
	if stub != nil {
		fake.DoASliceStub(arg1)
	}
}

func (fake *FakeSyncSomething) DoASliceCallCount() int {
	fake.doASliceMutex.RLock()
	defer fake.doASliceMutex.RUnlock()
	return len(fake.doASliceArgsForCall)
}

func (fake *FakeSyncSomething) DoASliceCalls(stub func([]byte)) {
	fake.doASliceMutex.Lock()
	defer fake.doASliceMutex.Unlock()
	fake.DoASliceStub = stub
}

func (fake *FakeSyncSomething) DoASliceArgsForCall(i int) []byte {
	fake.doASliceMutex.RLock()
	defer fake.doASliceMutex.RUnlock()
	argsForCall := fake.doASliceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSyncSomething) DoAnArray(arg1 [4]byte) {
	fake.doAnArrayMutex.Lock()
	fake.doAnArrayArgsForCall = append(fake.doAnArrayArgsForCall, struct {
		arg1 [4]byte
	}{arg1})
	stub := fake.DoAnArrayStub
	fake.recordInvocation("DoAnArray", []interface{}{arg1})
	fake.doAnArrayMutex.Unlock()
	if stub != nil {
		fake.DoAnArrayStub(arg1)
	}
}

func (fake *FakeSyncSomething) DoAnArrayCallCount() int {
	fake.doAnArrayMutex.RLock()
	defer fake.doAnArrayMutex.RUnlock()
	return len(fake.doAnArrayArgsForCall)
}

func (fake *FakeSyncSomething) DoAnArrayCalls(stub func([4]byte)) {
	fake.doAnArrayMutex.Lock()
	defer fake.doAnArrayMutex.Unlock()
	fake.DoAnArrayStub = stub
}

func (fake *FakeSyncSomething) DoAnArrayArgsForCall(i int) [4]byte {
	fake.doAnArrayMutex.RLock()
	defer fake.doAnArrayMutex.RUnlock()
	argsForCall := fake.doAnArrayArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSyncSomething) DoNothing() {
	fake.doNothingMutex.Lock()
	fake.doNothingArgsForCall = append(fake.doNothingArgsForCall, struct {
	}{})
	stub := fake.DoNothingStub
	fake.recordInvocation("DoNothing", []interface{}{})
	fake.doNothingMutex.Unlock()
	if stub != nil {
		fake.DoNothingStub()
	}
}

func (fake *FakeSyncSomething) DoNothingCallCount() int {
	fake.doNothingMutex.RLock()
	defer fake.doNothingMutex.RUnlock()
	return len(fake.doNothingArgsForCall)
}

func (fake *FakeSyncSomething) DoNothingCalls(stub func()) {
	fake.doNothingMutex.Lock()
	defer fake.doNothingMutex.Unlock()
	fake.DoNothingStub = stub
}

func (fake *FakeSyncSomething) DoThings(arg1 string, arg2 uint64) (int, error) {
	fake.doThingsMutex.Lock()
	ret, specificReturn := fake.doThingsReturnsOnCall[len(fake.doThingsArgsForCall)]
	fake.doThingsArgsForCall = append(fake.doThingsArgsForCall, struct {
		arg1 string
		arg2 uint64
	}{arg1, arg2})
	stub := fake.DoThingsStub
	fakeReturns := fake.doThingsReturns
	fake.recordInvocation("DoThings", []interface{}{arg1, arg2})
	fake.doThingsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSyncSomething) DoThingsCallCount() int {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return len(fake.doThingsArgsForCall)
}

func (fake *FakeSyncSomething) DoThingsCalls(stub func(string, uint64) (int, error)) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.DoThingsStub = stub
}

func (fake *FakeSyncSomething) DoThingsArgsForCall(i int) (string, uint64) {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	argsForCall := fake.doThingsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSyncSomething) DoThingsReturns(result1 int, result2 error) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.DoThingsStub = nil
	fake.doThingsReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeSyncSomething) DoThingsReturnsOnCall(i int, result1 int, result2 error) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.DoThingsStub = nil
	if fake.doThingsReturnsOnCall == nil {
		fake.doThingsReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.doThingsReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeSyncSomething) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doASliceMutex.RLock()
	defer fake.doASliceMutex.RUnlock()
	fake.doAnArrayMutex.RLock()
	defer fake.doAnArrayMutex.RUnlock()
	fake.doNothingMutex.RLock()
	defer fake.doNothingMutex.RUnlock()
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSyncSomething) recordInvocation(key string, args []interface{}) {
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

var _ synca.SyncSomething = new(FakeSyncSomething)
