// Code generated by counterfeiter. DO NOT EDIT.
package parsleyfakes

import (
	"sync"

	"github.com/opsidian/parsley/data"
	"github.com/opsidian/parsley/parsley"
)

type FakeParser struct {
	ParseStub        func(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet)
	parseMutex       sync.RWMutex
	parseArgsForCall []struct {
		h          parsley.History
		leftRecCtx data.IntMap
		r          parsley.Reader
		pos        parsley.Pos
	}
	parseReturns struct {
		result1 parsley.Node
		result2 parsley.Error
		result3 data.IntSet
	}
	parseReturnsOnCall map[int]struct {
		result1 parsley.Node
		result2 parsley.Error
		result3 data.IntSet
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeParser) Parse(h parsley.History, leftRecCtx data.IntMap, r parsley.Reader, pos parsley.Pos) (parsley.Node, parsley.Error, data.IntSet) {
	fake.parseMutex.Lock()
	ret, specificReturn := fake.parseReturnsOnCall[len(fake.parseArgsForCall)]
	fake.parseArgsForCall = append(fake.parseArgsForCall, struct {
		h          parsley.History
		leftRecCtx data.IntMap
		r          parsley.Reader
		pos        parsley.Pos
	}{h, leftRecCtx, r, pos})
	fake.recordInvocation("Parse", []interface{}{h, leftRecCtx, r, pos})
	fake.parseMutex.Unlock()
	if fake.ParseStub != nil {
		return fake.ParseStub(h, leftRecCtx, r, pos)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.parseReturns.result1, fake.parseReturns.result2, fake.parseReturns.result3
}

func (fake *FakeParser) ParseCallCount() int {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return len(fake.parseArgsForCall)
}

func (fake *FakeParser) ParseArgsForCall(i int) (parsley.History, data.IntMap, parsley.Reader, parsley.Pos) {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return fake.parseArgsForCall[i].h, fake.parseArgsForCall[i].leftRecCtx, fake.parseArgsForCall[i].r, fake.parseArgsForCall[i].pos
}

func (fake *FakeParser) ParseReturns(result1 parsley.Node, result2 parsley.Error, result3 data.IntSet) {
	fake.ParseStub = nil
	fake.parseReturns = struct {
		result1 parsley.Node
		result2 parsley.Error
		result3 data.IntSet
	}{result1, result2, result3}
}

func (fake *FakeParser) ParseReturnsOnCall(i int, result1 parsley.Node, result2 parsley.Error, result3 data.IntSet) {
	fake.ParseStub = nil
	if fake.parseReturnsOnCall == nil {
		fake.parseReturnsOnCall = make(map[int]struct {
			result1 parsley.Node
			result2 parsley.Error
			result3 data.IntSet
		})
	}
	fake.parseReturnsOnCall[i] = struct {
		result1 parsley.Node
		result2 parsley.Error
		result3 data.IntSet
	}{result1, result2, result3}
}

func (fake *FakeParser) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeParser) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeParser) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeParser) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeParser) recordInvocation(key string, args []interface{}) {
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

var _ parsley.Parser = new(FakeParser)
