// This file was generated by counterfeiter
package mocks

import (
	"sync"

	"github.com/cloudfoundry-incubator/lattice/ltc/ssh"
)

type FakeSessionFactory struct {
	NewStub        func(client ssh.Client, width, height int) (ssh.Session, error)
	newMutex       sync.RWMutex
	newArgsForCall []struct {
		client ssh.Client
		width  int
		height int
	}
	newReturns struct {
		result1 ssh.Session
		result2 error
	}
}

func (fake *FakeSessionFactory) New(client ssh.Client, width int, height int) (ssh.Session, error) {
	fake.newMutex.Lock()
	fake.newArgsForCall = append(fake.newArgsForCall, struct {
		client ssh.Client
		width  int
		height int
	}{client, width, height})
	fake.newMutex.Unlock()
	if fake.NewStub != nil {
		return fake.NewStub(client, width, height)
	} else {
		return fake.newReturns.result1, fake.newReturns.result2
	}
}

func (fake *FakeSessionFactory) NewCallCount() int {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return len(fake.newArgsForCall)
}

func (fake *FakeSessionFactory) NewArgsForCall(i int) (ssh.Client, int, int) {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return fake.newArgsForCall[i].client, fake.newArgsForCall[i].width, fake.newArgsForCall[i].height
}

func (fake *FakeSessionFactory) NewReturns(result1 ssh.Session, result2 error) {
	fake.NewStub = nil
	fake.newReturns = struct {
		result1 ssh.Session
		result2 error
	}{result1, result2}
}

var _ ssh.SessionFactory = new(FakeSessionFactory)
