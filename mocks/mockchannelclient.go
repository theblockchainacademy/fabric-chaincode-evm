// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-chaincode-evm/fabproxy"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

type MockChannelClient struct {
	QueryStub        func(request channel.Request, options ...channel.RequestOption) (channel.Response, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		request channel.Request
		options []channel.RequestOption
	}
	queryReturns struct {
		result1 channel.Response
		result2 error
	}
	queryReturnsOnCall map[int]struct {
		result1 channel.Response
		result2 error
	}
	ExecuteStub        func(request channel.Request, options ...channel.RequestOption) (channel.Response, error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		request channel.Request
		options []channel.RequestOption
	}
	executeReturns struct {
		result1 channel.Response
		result2 error
	}
	executeReturnsOnCall map[int]struct {
		result1 channel.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MockChannelClient) Query(request channel.Request, options ...channel.RequestOption) (channel.Response, error) {
	fake.queryMutex.Lock()
	ret, specificReturn := fake.queryReturnsOnCall[len(fake.queryArgsForCall)]
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		request channel.Request
		options []channel.RequestOption
	}{request, options})
	fake.recordInvocation("Query", []interface{}{request, options})
	fake.queryMutex.Unlock()
	if fake.QueryStub != nil {
		return fake.QueryStub(request, options...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.queryReturns.result1, fake.queryReturns.result2
}

func (fake *MockChannelClient) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *MockChannelClient) QueryArgsForCall(i int) (channel.Request, []channel.RequestOption) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return fake.queryArgsForCall[i].request, fake.queryArgsForCall[i].options
}

func (fake *MockChannelClient) QueryReturns(result1 channel.Response, result2 error) {
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *MockChannelClient) QueryReturnsOnCall(i int, result1 channel.Response, result2 error) {
	fake.QueryStub = nil
	if fake.queryReturnsOnCall == nil {
		fake.queryReturnsOnCall = make(map[int]struct {
			result1 channel.Response
			result2 error
		})
	}
	fake.queryReturnsOnCall[i] = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *MockChannelClient) Execute(request channel.Request, options ...channel.RequestOption) (channel.Response, error) {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		request channel.Request
		options []channel.RequestOption
	}{request, options})
	fake.recordInvocation("Execute", []interface{}{request, options})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(request, options...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.executeReturns.result1, fake.executeReturns.result2
}

func (fake *MockChannelClient) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *MockChannelClient) ExecuteArgsForCall(i int) (channel.Request, []channel.RequestOption) {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].request, fake.executeArgsForCall[i].options
}

func (fake *MockChannelClient) ExecuteReturns(result1 channel.Response, result2 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *MockChannelClient) ExecuteReturnsOnCall(i int, result1 channel.Response, result2 error) {
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 channel.Response
			result2 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 channel.Response
		result2 error
	}{result1, result2}
}

func (fake *MockChannelClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MockChannelClient) recordInvocation(key string, args []interface{}) {
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

var _ fabproxy.ChannelClient = new(MockChannelClient)