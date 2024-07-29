// Code generated by counterfeiter. DO NOT EDIT.
package flagcustomfakesdirfakes

import (
	"sync"

	packagemodeshim "github.com/jsokel/counterfeiter/v6/fixtures/packagemode/flagcustomfakesdir"
)

type FakePackagemode struct {
	ArgStub        func(int) string
	argMutex       sync.RWMutex
	argArgsForCall []struct {
		arg1 int
	}
	argReturns struct {
		result1 string
	}
	argReturnsOnCall map[int]struct {
		result1 string
	}
	ArgsStub        func() []string
	argsMutex       sync.RWMutex
	argsArgsForCall []struct {
	}
	argsReturns struct {
		result1 []string
	}
	argsReturnsOnCall map[int]struct {
		result1 []string
	}
	BoolStub        func(string, bool, string) *bool
	boolMutex       sync.RWMutex
	boolArgsForCall []struct {
		arg1 string
		arg2 bool
		arg3 string
	}
	boolReturns struct {
		result1 *bool
	}
	boolReturnsOnCall map[int]struct {
		result1 *bool
	}
	BoolVarStub        func(*bool, string, bool, string)
	boolVarMutex       sync.RWMutex
	boolVarArgsForCall []struct {
		arg1 *bool
		arg2 string
		arg3 bool
		arg4 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePackagemode) Arg(arg1 int) string {
	fake.argMutex.Lock()
	ret, specificReturn := fake.argReturnsOnCall[len(fake.argArgsForCall)]
	fake.argArgsForCall = append(fake.argArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.ArgStub
	fakeReturns := fake.argReturns
	fake.recordInvocation("Arg", []interface{}{arg1})
	fake.argMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePackagemode) ArgCallCount() int {
	fake.argMutex.RLock()
	defer fake.argMutex.RUnlock()
	return len(fake.argArgsForCall)
}

func (fake *FakePackagemode) ArgCalls(stub func(int) string) {
	fake.argMutex.Lock()
	defer fake.argMutex.Unlock()
	fake.ArgStub = stub
}

func (fake *FakePackagemode) ArgArgsForCall(i int) int {
	fake.argMutex.RLock()
	defer fake.argMutex.RUnlock()
	argsForCall := fake.argArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePackagemode) ArgReturns(result1 string) {
	fake.argMutex.Lock()
	defer fake.argMutex.Unlock()
	fake.ArgStub = nil
	fake.argReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePackagemode) ArgReturnsOnCall(i int, result1 string) {
	fake.argMutex.Lock()
	defer fake.argMutex.Unlock()
	fake.ArgStub = nil
	if fake.argReturnsOnCall == nil {
		fake.argReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.argReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePackagemode) Args() []string {
	fake.argsMutex.Lock()
	ret, specificReturn := fake.argsReturnsOnCall[len(fake.argsArgsForCall)]
	fake.argsArgsForCall = append(fake.argsArgsForCall, struct {
	}{})
	stub := fake.ArgsStub
	fakeReturns := fake.argsReturns
	fake.recordInvocation("Args", []interface{}{})
	fake.argsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePackagemode) ArgsCallCount() int {
	fake.argsMutex.RLock()
	defer fake.argsMutex.RUnlock()
	return len(fake.argsArgsForCall)
}

func (fake *FakePackagemode) ArgsCalls(stub func() []string) {
	fake.argsMutex.Lock()
	defer fake.argsMutex.Unlock()
	fake.ArgsStub = stub
}

func (fake *FakePackagemode) ArgsReturns(result1 []string) {
	fake.argsMutex.Lock()
	defer fake.argsMutex.Unlock()
	fake.ArgsStub = nil
	fake.argsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakePackagemode) ArgsReturnsOnCall(i int, result1 []string) {
	fake.argsMutex.Lock()
	defer fake.argsMutex.Unlock()
	fake.ArgsStub = nil
	if fake.argsReturnsOnCall == nil {
		fake.argsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.argsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakePackagemode) Bool(arg1 string, arg2 bool, arg3 string) *bool {
	fake.boolMutex.Lock()
	ret, specificReturn := fake.boolReturnsOnCall[len(fake.boolArgsForCall)]
	fake.boolArgsForCall = append(fake.boolArgsForCall, struct {
		arg1 string
		arg2 bool
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.BoolStub
	fakeReturns := fake.boolReturns
	fake.recordInvocation("Bool", []interface{}{arg1, arg2, arg3})
	fake.boolMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePackagemode) BoolCallCount() int {
	fake.boolMutex.RLock()
	defer fake.boolMutex.RUnlock()
	return len(fake.boolArgsForCall)
}

func (fake *FakePackagemode) BoolCalls(stub func(string, bool, string) *bool) {
	fake.boolMutex.Lock()
	defer fake.boolMutex.Unlock()
	fake.BoolStub = stub
}

func (fake *FakePackagemode) BoolArgsForCall(i int) (string, bool, string) {
	fake.boolMutex.RLock()
	defer fake.boolMutex.RUnlock()
	argsForCall := fake.boolArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePackagemode) BoolReturns(result1 *bool) {
	fake.boolMutex.Lock()
	defer fake.boolMutex.Unlock()
	fake.BoolStub = nil
	fake.boolReturns = struct {
		result1 *bool
	}{result1}
}

func (fake *FakePackagemode) BoolReturnsOnCall(i int, result1 *bool) {
	fake.boolMutex.Lock()
	defer fake.boolMutex.Unlock()
	fake.BoolStub = nil
	if fake.boolReturnsOnCall == nil {
		fake.boolReturnsOnCall = make(map[int]struct {
			result1 *bool
		})
	}
	fake.boolReturnsOnCall[i] = struct {
		result1 *bool
	}{result1}
}

func (fake *FakePackagemode) BoolVar(arg1 *bool, arg2 string, arg3 bool, arg4 string) {
	fake.boolVarMutex.Lock()
	fake.boolVarArgsForCall = append(fake.boolVarArgsForCall, struct {
		arg1 *bool
		arg2 string
		arg3 bool
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.BoolVarStub
	fake.recordInvocation("BoolVar", []interface{}{arg1, arg2, arg3, arg4})
	fake.boolVarMutex.Unlock()
	if stub != nil {
		fake.BoolVarStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *FakePackagemode) BoolVarCallCount() int {
	fake.boolVarMutex.RLock()
	defer fake.boolVarMutex.RUnlock()
	return len(fake.boolVarArgsForCall)
}

func (fake *FakePackagemode) BoolVarCalls(stub func(*bool, string, bool, string)) {
	fake.boolVarMutex.Lock()
	defer fake.boolVarMutex.Unlock()
	fake.BoolVarStub = stub
}

func (fake *FakePackagemode) BoolVarArgsForCall(i int) (*bool, string, bool, string) {
	fake.boolVarMutex.RLock()
	defer fake.boolVarMutex.RUnlock()
	argsForCall := fake.boolVarArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakePackagemode) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.argMutex.RLock()
	defer fake.argMutex.RUnlock()
	fake.argsMutex.RLock()
	defer fake.argsMutex.RUnlock()
	fake.boolMutex.RLock()
	defer fake.boolMutex.RUnlock()
	fake.boolVarMutex.RLock()
	defer fake.boolVarMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePackagemode) recordInvocation(key string, args []interface{}) {
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

var _ packagemodeshim.Packagemode = new(FakePackagemode)
