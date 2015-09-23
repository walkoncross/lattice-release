// This file was generated by counterfeiter
package mocks

import (
	"sync"

	"github.com/cloudfoundry-incubator/lattice/ltc/ssh"
)

type FakeSession struct {
	KeepAliveStub        func() (stopChan chan<- struct{})
	keepAliveMutex       sync.RWMutex
	keepAliveArgsForCall []struct{}
	keepAliveReturns struct {
		result1 chan<- struct{}
	}
	ResizeStub        func(width, height int) error
	resizeMutex       sync.RWMutex
	resizeArgsForCall []struct {
		width  int
		height int
	}
	resizeReturns struct {
		result1 error
	}
	ShellStub        func() error
	shellMutex       sync.RWMutex
	shellArgsForCall []struct{}
	shellReturns struct {
		result1 error
	}
	RunStub        func(string) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 string
	}
	runReturns struct {
		result1 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns struct {
		result1 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns struct {
		result1 error
	}
}

func (fake *FakeSession) KeepAlive() (stopChan chan<- struct{}) {
	fake.keepAliveMutex.Lock()
	fake.keepAliveArgsForCall = append(fake.keepAliveArgsForCall, struct{}{})
	fake.keepAliveMutex.Unlock()
	if fake.KeepAliveStub != nil {
		return fake.KeepAliveStub()
	} else {
		return fake.keepAliveReturns.result1
	}
}

func (fake *FakeSession) KeepAliveCallCount() int {
	fake.keepAliveMutex.RLock()
	defer fake.keepAliveMutex.RUnlock()
	return len(fake.keepAliveArgsForCall)
}

func (fake *FakeSession) KeepAliveReturns(result1 chan<- struct{}) {
	fake.KeepAliveStub = nil
	fake.keepAliveReturns = struct {
		result1 chan<- struct{}
	}{result1}
}

func (fake *FakeSession) Resize(width int, height int) error {
	fake.resizeMutex.Lock()
	fake.resizeArgsForCall = append(fake.resizeArgsForCall, struct {
		width  int
		height int
	}{width, height})
	fake.resizeMutex.Unlock()
	if fake.ResizeStub != nil {
		return fake.ResizeStub(width, height)
	} else {
		return fake.resizeReturns.result1
	}
}

func (fake *FakeSession) ResizeCallCount() int {
	fake.resizeMutex.RLock()
	defer fake.resizeMutex.RUnlock()
	return len(fake.resizeArgsForCall)
}

func (fake *FakeSession) ResizeArgsForCall(i int) (int, int) {
	fake.resizeMutex.RLock()
	defer fake.resizeMutex.RUnlock()
	return fake.resizeArgsForCall[i].width, fake.resizeArgsForCall[i].height
}

func (fake *FakeSession) ResizeReturns(result1 error) {
	fake.ResizeStub = nil
	fake.resizeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSession) Shell() error {
	fake.shellMutex.Lock()
	fake.shellArgsForCall = append(fake.shellArgsForCall, struct{}{})
	fake.shellMutex.Unlock()
	if fake.ShellStub != nil {
		return fake.ShellStub()
	} else {
		return fake.shellReturns.result1
	}
}

func (fake *FakeSession) ShellCallCount() int {
	fake.shellMutex.RLock()
	defer fake.shellMutex.RUnlock()
	return len(fake.shellArgsForCall)
}

func (fake *FakeSession) ShellReturns(result1 error) {
	fake.ShellStub = nil
	fake.shellReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSession) Run(arg1 string) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1)
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeSession) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeSession) RunArgsForCall(i int) string {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1
}

func (fake *FakeSession) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSession) Wait() error {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	} else {
		return fake.waitReturns.result1
	}
}

func (fake *FakeSession) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeSession) WaitReturns(result1 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSession) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeSession) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSession) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

var _ ssh.Session = new(FakeSession)
