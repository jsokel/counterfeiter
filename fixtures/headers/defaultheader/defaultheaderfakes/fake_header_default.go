// This is a default header for all the fakes in this package
//

// Code generated by counterfeiter. DO NOT EDIT.
package defaultheaderfakes

import (
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/v6/fixtures/headers/defaultheader"
)

type FakeHeaderDefault struct {
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHeaderDefault) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHeaderDefault) recordInvocation(key string, args []interface{}) {
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

var _ defaultheader.HeaderDefault = new(FakeHeaderDefault)