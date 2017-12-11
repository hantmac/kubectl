// Code generated by counterfeiter. DO NOT EDIT.
package testfakes

import (
	"sync"

	"k8s.io/kubectl/pkg/framework/test"
)

type FakeFixtureProcess struct {
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub          func()
	stopMutex         sync.RWMutex
	stopArgsForCall   []struct{}
	GetURLStub        func() string
	getURLMutex       sync.RWMutex
	getURLArgsForCall []struct{}
	getURLReturns     struct {
		result1 string
	}
	getURLReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFixtureProcess) Start() error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startReturns.result1
}

func (fake *FakeFixtureProcess) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeFixtureProcess) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFixtureProcess) StartReturnsOnCall(i int, result1 error) {
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFixtureProcess) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		fake.StopStub()
	}
}

func (fake *FakeFixtureProcess) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeFixtureProcess) GetURL() string {
	fake.getURLMutex.Lock()
	ret, specificReturn := fake.getURLReturnsOnCall[len(fake.getURLArgsForCall)]
	fake.getURLArgsForCall = append(fake.getURLArgsForCall, struct{}{})
	fake.recordInvocation("GetURL", []interface{}{})
	fake.getURLMutex.Unlock()
	if fake.GetURLStub != nil {
		return fake.GetURLStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getURLReturns.result1
}

func (fake *FakeFixtureProcess) GetURLCallCount() int {
	fake.getURLMutex.RLock()
	defer fake.getURLMutex.RUnlock()
	return len(fake.getURLArgsForCall)
}

func (fake *FakeFixtureProcess) GetURLReturns(result1 string) {
	fake.GetURLStub = nil
	fake.getURLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFixtureProcess) GetURLReturnsOnCall(i int, result1 string) {
	fake.GetURLStub = nil
	if fake.getURLReturnsOnCall == nil {
		fake.getURLReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getURLReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeFixtureProcess) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.getURLMutex.RLock()
	defer fake.getURLMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFixtureProcess) recordInvocation(key string, args []interface{}) {
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

var _ test.FixtureProcess = new(FakeFixtureProcess)
