// Code generated by counterfeiter. DO NOT EDIT.
package beaconfakes

import (
	"os"
	"sync"

	"github.com/concourse/worker/beacon"
)

type FakeBeaconClient struct {
	RegisterStub        func(signals <-chan os.Signal, ready chan<- struct{}) error
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}
	registerReturns struct {
		result1 error
	}
	registerReturnsOnCall map[int]struct {
		result1 error
	}
	RetireWorkerStub        func(signals <-chan os.Signal, ready chan<- struct{}) error
	retireWorkerMutex       sync.RWMutex
	retireWorkerArgsForCall []struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}
	retireWorkerReturns struct {
		result1 error
	}
	retireWorkerReturnsOnCall map[int]struct {
		result1 error
	}
	MarkandSweepContainersandVolumesStub        func() error
	markandSweepContainersandVolumesMutex       sync.RWMutex
	markandSweepContainersandVolumesArgsForCall []struct{}
	markandSweepContainersandVolumesReturns     struct {
		result1 error
	}
	markandSweepContainersandVolumesReturnsOnCall map[int]struct {
		result1 error
	}
	LandWorkerStub        func(signals <-chan os.Signal, ready chan<- struct{}) error
	landWorkerMutex       sync.RWMutex
	landWorkerArgsForCall []struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}
	landWorkerReturns struct {
		result1 error
	}
	landWorkerReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteWorkerStub        func(signals <-chan os.Signal, ready chan<- struct{}) error
	deleteWorkerMutex       sync.RWMutex
	deleteWorkerArgsForCall []struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}
	deleteWorkerReturns struct {
		result1 error
	}
	deleteWorkerReturnsOnCall map[int]struct {
		result1 error
	}
	DisableKeepAliveStub        func()
	disableKeepAliveMutex       sync.RWMutex
	disableKeepAliveArgsForCall []struct{}
	invocations                 map[string][][]interface{}
	invocationsMutex            sync.RWMutex
}

func (fake *FakeBeaconClient) Register(signals <-chan os.Signal, ready chan<- struct{}) error {
	fake.registerMutex.Lock()
	ret, specificReturn := fake.registerReturnsOnCall[len(fake.registerArgsForCall)]
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}{signals, ready})
	fake.recordInvocation("Register", []interface{}{signals, ready})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		return fake.RegisterStub(signals, ready)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.registerReturns.result1
}

func (fake *FakeBeaconClient) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *FakeBeaconClient) RegisterArgsForCall(i int) (<-chan os.Signal, chan<- struct{}) {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return fake.registerArgsForCall[i].signals, fake.registerArgsForCall[i].ready
}

func (fake *FakeBeaconClient) RegisterReturns(result1 error) {
	fake.RegisterStub = nil
	fake.registerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) RegisterReturnsOnCall(i int, result1 error) {
	fake.RegisterStub = nil
	if fake.registerReturnsOnCall == nil {
		fake.registerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) RetireWorker(signals <-chan os.Signal, ready chan<- struct{}) error {
	fake.retireWorkerMutex.Lock()
	ret, specificReturn := fake.retireWorkerReturnsOnCall[len(fake.retireWorkerArgsForCall)]
	fake.retireWorkerArgsForCall = append(fake.retireWorkerArgsForCall, struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}{signals, ready})
	fake.recordInvocation("RetireWorker", []interface{}{signals, ready})
	fake.retireWorkerMutex.Unlock()
	if fake.RetireWorkerStub != nil {
		return fake.RetireWorkerStub(signals, ready)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.retireWorkerReturns.result1
}

func (fake *FakeBeaconClient) RetireWorkerCallCount() int {
	fake.retireWorkerMutex.RLock()
	defer fake.retireWorkerMutex.RUnlock()
	return len(fake.retireWorkerArgsForCall)
}

func (fake *FakeBeaconClient) RetireWorkerArgsForCall(i int) (<-chan os.Signal, chan<- struct{}) {
	fake.retireWorkerMutex.RLock()
	defer fake.retireWorkerMutex.RUnlock()
	return fake.retireWorkerArgsForCall[i].signals, fake.retireWorkerArgsForCall[i].ready
}

func (fake *FakeBeaconClient) RetireWorkerReturns(result1 error) {
	fake.RetireWorkerStub = nil
	fake.retireWorkerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) RetireWorkerReturnsOnCall(i int, result1 error) {
	fake.RetireWorkerStub = nil
	if fake.retireWorkerReturnsOnCall == nil {
		fake.retireWorkerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.retireWorkerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) MarkandSweepContainersandVolumes() error {
	fake.markandSweepContainersandVolumesMutex.Lock()
	ret, specificReturn := fake.markandSweepContainersandVolumesReturnsOnCall[len(fake.markandSweepContainersandVolumesArgsForCall)]
	fake.markandSweepContainersandVolumesArgsForCall = append(fake.markandSweepContainersandVolumesArgsForCall, struct{}{})
	fake.recordInvocation("MarkandSweepContainersandVolumes", []interface{}{})
	fake.markandSweepContainersandVolumesMutex.Unlock()
	if fake.MarkandSweepContainersandVolumesStub != nil {
		return fake.MarkandSweepContainersandVolumesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.markandSweepContainersandVolumesReturns.result1
}

func (fake *FakeBeaconClient) MarkandSweepContainersandVolumesCallCount() int {
	fake.markandSweepContainersandVolumesMutex.RLock()
	defer fake.markandSweepContainersandVolumesMutex.RUnlock()
	return len(fake.markandSweepContainersandVolumesArgsForCall)
}

func (fake *FakeBeaconClient) MarkandSweepContainersandVolumesReturns(result1 error) {
	fake.MarkandSweepContainersandVolumesStub = nil
	fake.markandSweepContainersandVolumesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) MarkandSweepContainersandVolumesReturnsOnCall(i int, result1 error) {
	fake.MarkandSweepContainersandVolumesStub = nil
	if fake.markandSweepContainersandVolumesReturnsOnCall == nil {
		fake.markandSweepContainersandVolumesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.markandSweepContainersandVolumesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) LandWorker(signals <-chan os.Signal, ready chan<- struct{}) error {
	fake.landWorkerMutex.Lock()
	ret, specificReturn := fake.landWorkerReturnsOnCall[len(fake.landWorkerArgsForCall)]
	fake.landWorkerArgsForCall = append(fake.landWorkerArgsForCall, struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}{signals, ready})
	fake.recordInvocation("LandWorker", []interface{}{signals, ready})
	fake.landWorkerMutex.Unlock()
	if fake.LandWorkerStub != nil {
		return fake.LandWorkerStub(signals, ready)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.landWorkerReturns.result1
}

func (fake *FakeBeaconClient) LandWorkerCallCount() int {
	fake.landWorkerMutex.RLock()
	defer fake.landWorkerMutex.RUnlock()
	return len(fake.landWorkerArgsForCall)
}

func (fake *FakeBeaconClient) LandWorkerArgsForCall(i int) (<-chan os.Signal, chan<- struct{}) {
	fake.landWorkerMutex.RLock()
	defer fake.landWorkerMutex.RUnlock()
	return fake.landWorkerArgsForCall[i].signals, fake.landWorkerArgsForCall[i].ready
}

func (fake *FakeBeaconClient) LandWorkerReturns(result1 error) {
	fake.LandWorkerStub = nil
	fake.landWorkerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) LandWorkerReturnsOnCall(i int, result1 error) {
	fake.LandWorkerStub = nil
	if fake.landWorkerReturnsOnCall == nil {
		fake.landWorkerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.landWorkerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) DeleteWorker(signals <-chan os.Signal, ready chan<- struct{}) error {
	fake.deleteWorkerMutex.Lock()
	ret, specificReturn := fake.deleteWorkerReturnsOnCall[len(fake.deleteWorkerArgsForCall)]
	fake.deleteWorkerArgsForCall = append(fake.deleteWorkerArgsForCall, struct {
		signals <-chan os.Signal
		ready   chan<- struct{}
	}{signals, ready})
	fake.recordInvocation("DeleteWorker", []interface{}{signals, ready})
	fake.deleteWorkerMutex.Unlock()
	if fake.DeleteWorkerStub != nil {
		return fake.DeleteWorkerStub(signals, ready)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteWorkerReturns.result1
}

func (fake *FakeBeaconClient) DeleteWorkerCallCount() int {
	fake.deleteWorkerMutex.RLock()
	defer fake.deleteWorkerMutex.RUnlock()
	return len(fake.deleteWorkerArgsForCall)
}

func (fake *FakeBeaconClient) DeleteWorkerArgsForCall(i int) (<-chan os.Signal, chan<- struct{}) {
	fake.deleteWorkerMutex.RLock()
	defer fake.deleteWorkerMutex.RUnlock()
	return fake.deleteWorkerArgsForCall[i].signals, fake.deleteWorkerArgsForCall[i].ready
}

func (fake *FakeBeaconClient) DeleteWorkerReturns(result1 error) {
	fake.DeleteWorkerStub = nil
	fake.deleteWorkerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) DeleteWorkerReturnsOnCall(i int, result1 error) {
	fake.DeleteWorkerStub = nil
	if fake.deleteWorkerReturnsOnCall == nil {
		fake.deleteWorkerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteWorkerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBeaconClient) DisableKeepAlive() {
	fake.disableKeepAliveMutex.Lock()
	fake.disableKeepAliveArgsForCall = append(fake.disableKeepAliveArgsForCall, struct{}{})
	fake.recordInvocation("DisableKeepAlive", []interface{}{})
	fake.disableKeepAliveMutex.Unlock()
	if fake.DisableKeepAliveStub != nil {
		fake.DisableKeepAliveStub()
	}
}

func (fake *FakeBeaconClient) DisableKeepAliveCallCount() int {
	fake.disableKeepAliveMutex.RLock()
	defer fake.disableKeepAliveMutex.RUnlock()
	return len(fake.disableKeepAliveArgsForCall)
}

func (fake *FakeBeaconClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	fake.retireWorkerMutex.RLock()
	defer fake.retireWorkerMutex.RUnlock()
	fake.markandSweepContainersandVolumesMutex.RLock()
	defer fake.markandSweepContainersandVolumesMutex.RUnlock()
	fake.landWorkerMutex.RLock()
	defer fake.landWorkerMutex.RUnlock()
	fake.deleteWorkerMutex.RLock()
	defer fake.deleteWorkerMutex.RUnlock()
	fake.disableKeepAliveMutex.RLock()
	defer fake.disableKeepAliveMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBeaconClient) recordInvocation(key string, args []interface{}) {
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

var _ beacon.BeaconClient = new(FakeBeaconClient)
