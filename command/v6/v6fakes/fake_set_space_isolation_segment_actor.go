// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeSetSpaceIsolationSegmentActor struct {
	AssignIsolationSegmentToSpaceByNameAndSpaceStub        func(string, string) (v3action.Warnings, error)
	assignIsolationSegmentToSpaceByNameAndSpaceMutex       sync.RWMutex
	assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	assignIsolationSegmentToSpaceByNameAndSpaceReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	assignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSetSpaceIsolationSegmentActor) AssignIsolationSegmentToSpaceByNameAndSpace(arg1 string, arg2 string) (v3action.Warnings, error) {
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.assignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall[len(fake.assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall)]
	fake.assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall = append(fake.assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("AssignIsolationSegmentToSpaceByNameAndSpace", []interface{}{arg1, arg2})
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.Unlock()
	if fake.AssignIsolationSegmentToSpaceByNameAndSpaceStub != nil {
		return fake.AssignIsolationSegmentToSpaceByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.assignIsolationSegmentToSpaceByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSetSpaceIsolationSegmentActor) AssignIsolationSegmentToSpaceByNameAndSpaceCallCount() int {
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RLock()
	defer fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RUnlock()
	return len(fake.assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall)
}

func (fake *FakeSetSpaceIsolationSegmentActor) AssignIsolationSegmentToSpaceByNameAndSpaceCalls(stub func(string, string) (v3action.Warnings, error)) {
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.Lock()
	defer fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.Unlock()
	fake.AssignIsolationSegmentToSpaceByNameAndSpaceStub = stub
}

func (fake *FakeSetSpaceIsolationSegmentActor) AssignIsolationSegmentToSpaceByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RLock()
	defer fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.assignIsolationSegmentToSpaceByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSetSpaceIsolationSegmentActor) AssignIsolationSegmentToSpaceByNameAndSpaceReturns(result1 v3action.Warnings, result2 error) {
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.Lock()
	defer fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.Unlock()
	fake.AssignIsolationSegmentToSpaceByNameAndSpaceStub = nil
	fake.assignIsolationSegmentToSpaceByNameAndSpaceReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetSpaceIsolationSegmentActor) AssignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.Lock()
	defer fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.Unlock()
	fake.AssignIsolationSegmentToSpaceByNameAndSpaceStub = nil
	if fake.assignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall == nil {
		fake.assignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.assignIsolationSegmentToSpaceByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeSetSpaceIsolationSegmentActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RLock()
	defer fake.assignIsolationSegmentToSpaceByNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSetSpaceIsolationSegmentActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.SetSpaceIsolationSegmentActor = new(FakeSetSpaceIsolationSegmentActor)