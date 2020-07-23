// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric/orderer/common/follower"
)

type BlockPullerCreator struct {
	CreateBlockPullerStub        func() (follower.ChannelPuller, error)
	createBlockPullerMutex       sync.RWMutex
	createBlockPullerArgsForCall []struct {
	}
	createBlockPullerReturns struct {
		result1 follower.ChannelPuller
		result2 error
	}
	createBlockPullerReturnsOnCall map[int]struct {
		result1 follower.ChannelPuller
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BlockPullerCreator) CreateBlockPuller() (follower.ChannelPuller, error) {
	fake.createBlockPullerMutex.Lock()
	ret, specificReturn := fake.createBlockPullerReturnsOnCall[len(fake.createBlockPullerArgsForCall)]
	fake.createBlockPullerArgsForCall = append(fake.createBlockPullerArgsForCall, struct {
	}{})
	fake.recordInvocation("CreateBlockPuller", []interface{}{})
	fake.createBlockPullerMutex.Unlock()
	if fake.CreateBlockPullerStub != nil {
		return fake.CreateBlockPullerStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createBlockPullerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *BlockPullerCreator) CreateBlockPullerCallCount() int {
	fake.createBlockPullerMutex.RLock()
	defer fake.createBlockPullerMutex.RUnlock()
	return len(fake.createBlockPullerArgsForCall)
}

func (fake *BlockPullerCreator) CreateBlockPullerCalls(stub func() (follower.ChannelPuller, error)) {
	fake.createBlockPullerMutex.Lock()
	defer fake.createBlockPullerMutex.Unlock()
	fake.CreateBlockPullerStub = stub
}

func (fake *BlockPullerCreator) CreateBlockPullerReturns(result1 follower.ChannelPuller, result2 error) {
	fake.createBlockPullerMutex.Lock()
	defer fake.createBlockPullerMutex.Unlock()
	fake.CreateBlockPullerStub = nil
	fake.createBlockPullerReturns = struct {
		result1 follower.ChannelPuller
		result2 error
	}{result1, result2}
}

func (fake *BlockPullerCreator) CreateBlockPullerReturnsOnCall(i int, result1 follower.ChannelPuller, result2 error) {
	fake.createBlockPullerMutex.Lock()
	defer fake.createBlockPullerMutex.Unlock()
	fake.CreateBlockPullerStub = nil
	if fake.createBlockPullerReturnsOnCall == nil {
		fake.createBlockPullerReturnsOnCall = make(map[int]struct {
			result1 follower.ChannelPuller
			result2 error
		})
	}
	fake.createBlockPullerReturnsOnCall[i] = struct {
		result1 follower.ChannelPuller
		result2 error
	}{result1, result2}
}

func (fake *BlockPullerCreator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createBlockPullerMutex.RLock()
	defer fake.createBlockPullerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BlockPullerCreator) recordInvocation(key string, args []interface{}) {
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