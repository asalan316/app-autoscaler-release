// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"database/sql"
	"sync"

	"code.cloudfoundry.org/app-autoscaler/src/autoscaler/db"
	"code.cloudfoundry.org/app-autoscaler/src/autoscaler/models"
)

type FakeSchedulerDB struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	GetActiveSchedulesStub        func() (map[string]*models.ActiveSchedule, error)
	getActiveSchedulesMutex       sync.RWMutex
	getActiveSchedulesArgsForCall []struct {
	}
	getActiveSchedulesReturns struct {
		result1 map[string]*models.ActiveSchedule
		result2 error
	}
	getActiveSchedulesReturnsOnCall map[int]struct {
		result1 map[string]*models.ActiveSchedule
		result2 error
	}
	GetDBStatusStub        func() sql.DBStats
	getDBStatusMutex       sync.RWMutex
	getDBStatusArgsForCall []struct {
	}
	getDBStatusReturns struct {
		result1 sql.DBStats
	}
	getDBStatusReturnsOnCall map[int]struct {
		result1 sql.DBStats
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSchedulerDB) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSchedulerDB) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSchedulerDB) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeSchedulerDB) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSchedulerDB) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSchedulerDB) GetActiveSchedules() (map[string]*models.ActiveSchedule, error) {
	fake.getActiveSchedulesMutex.Lock()
	ret, specificReturn := fake.getActiveSchedulesReturnsOnCall[len(fake.getActiveSchedulesArgsForCall)]
	fake.getActiveSchedulesArgsForCall = append(fake.getActiveSchedulesArgsForCall, struct {
	}{})
	stub := fake.GetActiveSchedulesStub
	fakeReturns := fake.getActiveSchedulesReturns
	fake.recordInvocation("GetActiveSchedules", []interface{}{})
	fake.getActiveSchedulesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSchedulerDB) GetActiveSchedulesCallCount() int {
	fake.getActiveSchedulesMutex.RLock()
	defer fake.getActiveSchedulesMutex.RUnlock()
	return len(fake.getActiveSchedulesArgsForCall)
}

func (fake *FakeSchedulerDB) GetActiveSchedulesCalls(stub func() (map[string]*models.ActiveSchedule, error)) {
	fake.getActiveSchedulesMutex.Lock()
	defer fake.getActiveSchedulesMutex.Unlock()
	fake.GetActiveSchedulesStub = stub
}

func (fake *FakeSchedulerDB) GetActiveSchedulesReturns(result1 map[string]*models.ActiveSchedule, result2 error) {
	fake.getActiveSchedulesMutex.Lock()
	defer fake.getActiveSchedulesMutex.Unlock()
	fake.GetActiveSchedulesStub = nil
	fake.getActiveSchedulesReturns = struct {
		result1 map[string]*models.ActiveSchedule
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) GetActiveSchedulesReturnsOnCall(i int, result1 map[string]*models.ActiveSchedule, result2 error) {
	fake.getActiveSchedulesMutex.Lock()
	defer fake.getActiveSchedulesMutex.Unlock()
	fake.GetActiveSchedulesStub = nil
	if fake.getActiveSchedulesReturnsOnCall == nil {
		fake.getActiveSchedulesReturnsOnCall = make(map[int]struct {
			result1 map[string]*models.ActiveSchedule
			result2 error
		})
	}
	fake.getActiveSchedulesReturnsOnCall[i] = struct {
		result1 map[string]*models.ActiveSchedule
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) GetDBStatus() sql.DBStats {
	fake.getDBStatusMutex.Lock()
	ret, specificReturn := fake.getDBStatusReturnsOnCall[len(fake.getDBStatusArgsForCall)]
	fake.getDBStatusArgsForCall = append(fake.getDBStatusArgsForCall, struct {
	}{})
	stub := fake.GetDBStatusStub
	fakeReturns := fake.getDBStatusReturns
	fake.recordInvocation("GetDBStatus", []interface{}{})
	fake.getDBStatusMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSchedulerDB) GetDBStatusCallCount() int {
	fake.getDBStatusMutex.RLock()
	defer fake.getDBStatusMutex.RUnlock()
	return len(fake.getDBStatusArgsForCall)
}

func (fake *FakeSchedulerDB) GetDBStatusCalls(stub func() sql.DBStats) {
	fake.getDBStatusMutex.Lock()
	defer fake.getDBStatusMutex.Unlock()
	fake.GetDBStatusStub = stub
}

func (fake *FakeSchedulerDB) GetDBStatusReturns(result1 sql.DBStats) {
	fake.getDBStatusMutex.Lock()
	defer fake.getDBStatusMutex.Unlock()
	fake.GetDBStatusStub = nil
	fake.getDBStatusReturns = struct {
		result1 sql.DBStats
	}{result1}
}

func (fake *FakeSchedulerDB) GetDBStatusReturnsOnCall(i int, result1 sql.DBStats) {
	fake.getDBStatusMutex.Lock()
	defer fake.getDBStatusMutex.Unlock()
	fake.GetDBStatusStub = nil
	if fake.getDBStatusReturnsOnCall == nil {
		fake.getDBStatusReturnsOnCall = make(map[int]struct {
			result1 sql.DBStats
		})
	}
	fake.getDBStatusReturnsOnCall[i] = struct {
		result1 sql.DBStats
	}{result1}
}

func (fake *FakeSchedulerDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.getActiveSchedulesMutex.RLock()
	defer fake.getActiveSchedulesMutex.RUnlock()
	fake.getDBStatusMutex.RLock()
	defer fake.getDBStatusMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSchedulerDB) recordInvocation(key string, args []interface{}) {
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

var _ db.SchedulerDB = new(FakeSchedulerDB)