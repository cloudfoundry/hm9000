// This file was generated by counterfeiter
package fakestore

import (
	"sync"
	"time"

	"github.com/cloudfoundry/hm9000/models"
	"github.com/cloudfoundry/hm9000/store"
)

type FakeStore struct {
	BumpDesiredFreshnessStub        func(timestamp time.Time) error
	bumpDesiredFreshnessMutex       sync.RWMutex
	bumpDesiredFreshnessArgsForCall []struct {
		timestamp time.Time
	}
	bumpDesiredFreshnessReturns struct {
		result1 error
	}
	BumpActualFreshnessStub        func(timestamp time.Time) error
	bumpActualFreshnessMutex       sync.RWMutex
	bumpActualFreshnessArgsForCall []struct {
		timestamp time.Time
	}
	bumpActualFreshnessReturns struct {
		result1 error
	}
	RevokeActualFreshnessStub        func() error
	revokeActualFreshnessMutex       sync.RWMutex
	revokeActualFreshnessArgsForCall []struct{}
	revokeActualFreshnessReturns     struct {
		result1 error
	}
	IsDesiredStateFreshStub        func() (bool, error)
	isDesiredStateFreshMutex       sync.RWMutex
	isDesiredStateFreshArgsForCall []struct{}
	isDesiredStateFreshReturns     struct {
		result1 bool
		result2 error
	}
	IsActualStateFreshStub        func(time.Time) (bool, error)
	isActualStateFreshMutex       sync.RWMutex
	isActualStateFreshArgsForCall []struct {
		arg1 time.Time
	}
	isActualStateFreshReturns struct {
		result1 bool
		result2 error
	}
	VerifyFreshnessStub        func(time.Time) error
	verifyFreshnessMutex       sync.RWMutex
	verifyFreshnessArgsForCall []struct {
		arg1 time.Time
	}
	verifyFreshnessReturns struct {
		result1 error
	}
	AppKeyStub        func(appGuid string, appVersion string) string
	appKeyMutex       sync.RWMutex
	appKeyArgsForCall []struct {
		appGuid    string
		appVersion string
	}
	appKeyReturns struct {
		result1 string
	}
	GetAppsStub        func() (map[string]*models.App, error)
	getAppsMutex       sync.RWMutex
	getAppsArgsForCall []struct{}
	getAppsReturns     struct {
		result1 map[string]*models.App
		result2 error
	}
	GetAppStub        func(appGuid string, appVersion string) (*models.App, error)
	getAppMutex       sync.RWMutex
	getAppArgsForCall []struct {
		appGuid    string
		appVersion string
	}
	getAppReturns struct {
		result1 *models.App
		result2 error
	}
	SyncDesiredStateStub        func(desiredStates ...models.DesiredAppState) error
	syncDesiredStateMutex       sync.RWMutex
	syncDesiredStateArgsForCall []struct {
		desiredStates []models.DesiredAppState
	}
	syncDesiredStateReturns struct {
		result1 error
	}
	GetDesiredStateStub        func() (map[string]models.DesiredAppState, error)
	getDesiredStateMutex       sync.RWMutex
	getDesiredStateArgsForCall []struct{}
	getDesiredStateReturns     struct {
		result1 map[string]models.DesiredAppState
		result2 error
	}
	SyncHeartbeatsStub        func(heartbeat ...models.Heartbeat) error
	syncHeartbeatsMutex       sync.RWMutex
	syncHeartbeatsArgsForCall []struct {
		heartbeat []models.Heartbeat
	}
	syncHeartbeatsReturns struct {
		result1 error
	}
	GetInstanceHeartbeatsStub        func() (results []models.InstanceHeartbeat, err error)
	getInstanceHeartbeatsMutex       sync.RWMutex
	getInstanceHeartbeatsArgsForCall []struct{}
	getInstanceHeartbeatsReturns     struct {
		result1 []models.InstanceHeartbeat
		result2 error
	}
	GetInstanceHeartbeatsForAppStub        func(appGuid string, appVersion string) (results []models.InstanceHeartbeat, err error)
	getInstanceHeartbeatsForAppMutex       sync.RWMutex
	getInstanceHeartbeatsForAppArgsForCall []struct {
		appGuid    string
		appVersion string
	}
	getInstanceHeartbeatsForAppReturns struct {
		result1 []models.InstanceHeartbeat
		result2 error
	}
	SaveCrashCountsStub        func(crashCounts ...models.CrashCount) error
	saveCrashCountsMutex       sync.RWMutex
	saveCrashCountsArgsForCall []struct {
		crashCounts []models.CrashCount
	}
	saveCrashCountsReturns struct {
		result1 error
	}
	SavePendingStartMessagesStub        func(startMessages ...models.PendingStartMessage) error
	savePendingStartMessagesMutex       sync.RWMutex
	savePendingStartMessagesArgsForCall []struct {
		startMessages []models.PendingStartMessage
	}
	savePendingStartMessagesReturns struct {
		result1 error
	}
	GetPendingStartMessagesStub        func() (map[string]models.PendingStartMessage, error)
	getPendingStartMessagesMutex       sync.RWMutex
	getPendingStartMessagesArgsForCall []struct{}
	getPendingStartMessagesReturns     struct {
		result1 map[string]models.PendingStartMessage
		result2 error
	}
	DeletePendingStartMessagesStub        func(startMessages ...models.PendingStartMessage) error
	deletePendingStartMessagesMutex       sync.RWMutex
	deletePendingStartMessagesArgsForCall []struct {
		startMessages []models.PendingStartMessage
	}
	deletePendingStartMessagesReturns struct {
		result1 error
	}
	SavePendingStopMessagesStub        func(stopMessages ...models.PendingStopMessage) error
	savePendingStopMessagesMutex       sync.RWMutex
	savePendingStopMessagesArgsForCall []struct {
		stopMessages []models.PendingStopMessage
	}
	savePendingStopMessagesReturns struct {
		result1 error
	}
	GetPendingStopMessagesStub        func() (map[string]models.PendingStopMessage, error)
	getPendingStopMessagesMutex       sync.RWMutex
	getPendingStopMessagesArgsForCall []struct{}
	getPendingStopMessagesReturns     struct {
		result1 map[string]models.PendingStopMessage
		result2 error
	}
	DeletePendingStopMessagesStub        func(stopMessages ...models.PendingStopMessage) error
	deletePendingStopMessagesMutex       sync.RWMutex
	deletePendingStopMessagesArgsForCall []struct {
		stopMessages []models.PendingStopMessage
	}
	deletePendingStopMessagesReturns struct {
		result1 error
	}
	SaveMetricStub        func(metric string, value float64) error
	saveMetricMutex       sync.RWMutex
	saveMetricArgsForCall []struct {
		metric string
		value  float64
	}
	saveMetricReturns struct {
		result1 error
	}
	GetMetricStub        func(metric string) (float64, error)
	getMetricMutex       sync.RWMutex
	getMetricArgsForCall []struct {
		metric string
	}
	getMetricReturns struct {
		result1 float64
		result2 error
	}
	CompactStub        func() error
	compactMutex       sync.RWMutex
	compactArgsForCall []struct{}
	compactReturns     struct {
		result1 error
	}
}

func (fake *FakeStore) BumpDesiredFreshness(timestamp time.Time) error {
	fake.bumpDesiredFreshnessMutex.Lock()
	fake.bumpDesiredFreshnessArgsForCall = append(fake.bumpDesiredFreshnessArgsForCall, struct {
		timestamp time.Time
	}{timestamp})
	fake.bumpDesiredFreshnessMutex.Unlock()
	if fake.BumpDesiredFreshnessStub != nil {
		return fake.BumpDesiredFreshnessStub(timestamp)
	} else {
		return fake.bumpDesiredFreshnessReturns.result1
	}
}

func (fake *FakeStore) BumpDesiredFreshnessCallCount() int {
	fake.bumpDesiredFreshnessMutex.RLock()
	defer fake.bumpDesiredFreshnessMutex.RUnlock()
	return len(fake.bumpDesiredFreshnessArgsForCall)
}

func (fake *FakeStore) BumpDesiredFreshnessArgsForCall(i int) time.Time {
	fake.bumpDesiredFreshnessMutex.RLock()
	defer fake.bumpDesiredFreshnessMutex.RUnlock()
	return fake.bumpDesiredFreshnessArgsForCall[i].timestamp
}

func (fake *FakeStore) BumpDesiredFreshnessReturns(result1 error) {
	fake.BumpDesiredFreshnessStub = nil
	fake.bumpDesiredFreshnessReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) BumpActualFreshness(timestamp time.Time) error {
	fake.bumpActualFreshnessMutex.Lock()
	fake.bumpActualFreshnessArgsForCall = append(fake.bumpActualFreshnessArgsForCall, struct {
		timestamp time.Time
	}{timestamp})
	fake.bumpActualFreshnessMutex.Unlock()
	if fake.BumpActualFreshnessStub != nil {
		return fake.BumpActualFreshnessStub(timestamp)
	} else {
		return fake.bumpActualFreshnessReturns.result1
	}
}

func (fake *FakeStore) BumpActualFreshnessCallCount() int {
	fake.bumpActualFreshnessMutex.RLock()
	defer fake.bumpActualFreshnessMutex.RUnlock()
	return len(fake.bumpActualFreshnessArgsForCall)
}

func (fake *FakeStore) BumpActualFreshnessArgsForCall(i int) time.Time {
	fake.bumpActualFreshnessMutex.RLock()
	defer fake.bumpActualFreshnessMutex.RUnlock()
	return fake.bumpActualFreshnessArgsForCall[i].timestamp
}

func (fake *FakeStore) BumpActualFreshnessReturns(result1 error) {
	fake.BumpActualFreshnessStub = nil
	fake.bumpActualFreshnessReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) RevokeActualFreshness() error {
	fake.revokeActualFreshnessMutex.Lock()
	fake.revokeActualFreshnessArgsForCall = append(fake.revokeActualFreshnessArgsForCall, struct{}{})
	fake.revokeActualFreshnessMutex.Unlock()
	if fake.RevokeActualFreshnessStub != nil {
		return fake.RevokeActualFreshnessStub()
	} else {
		return fake.revokeActualFreshnessReturns.result1
	}
}

func (fake *FakeStore) RevokeActualFreshnessCallCount() int {
	fake.revokeActualFreshnessMutex.RLock()
	defer fake.revokeActualFreshnessMutex.RUnlock()
	return len(fake.revokeActualFreshnessArgsForCall)
}

func (fake *FakeStore) RevokeActualFreshnessReturns(result1 error) {
	fake.RevokeActualFreshnessStub = nil
	fake.revokeActualFreshnessReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) IsDesiredStateFresh() (bool, error) {
	fake.isDesiredStateFreshMutex.Lock()
	fake.isDesiredStateFreshArgsForCall = append(fake.isDesiredStateFreshArgsForCall, struct{}{})
	fake.isDesiredStateFreshMutex.Unlock()
	if fake.IsDesiredStateFreshStub != nil {
		return fake.IsDesiredStateFreshStub()
	} else {
		return fake.isDesiredStateFreshReturns.result1, fake.isDesiredStateFreshReturns.result2
	}
}

func (fake *FakeStore) IsDesiredStateFreshCallCount() int {
	fake.isDesiredStateFreshMutex.RLock()
	defer fake.isDesiredStateFreshMutex.RUnlock()
	return len(fake.isDesiredStateFreshArgsForCall)
}

func (fake *FakeStore) IsDesiredStateFreshReturns(result1 bool, result2 error) {
	fake.IsDesiredStateFreshStub = nil
	fake.isDesiredStateFreshReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) IsActualStateFresh(arg1 time.Time) (bool, error) {
	fake.isActualStateFreshMutex.Lock()
	fake.isActualStateFreshArgsForCall = append(fake.isActualStateFreshArgsForCall, struct {
		arg1 time.Time
	}{arg1})
	fake.isActualStateFreshMutex.Unlock()
	if fake.IsActualStateFreshStub != nil {
		return fake.IsActualStateFreshStub(arg1)
	} else {
		return fake.isActualStateFreshReturns.result1, fake.isActualStateFreshReturns.result2
	}
}

func (fake *FakeStore) IsActualStateFreshCallCount() int {
	fake.isActualStateFreshMutex.RLock()
	defer fake.isActualStateFreshMutex.RUnlock()
	return len(fake.isActualStateFreshArgsForCall)
}

func (fake *FakeStore) IsActualStateFreshArgsForCall(i int) time.Time {
	fake.isActualStateFreshMutex.RLock()
	defer fake.isActualStateFreshMutex.RUnlock()
	return fake.isActualStateFreshArgsForCall[i].arg1
}

func (fake *FakeStore) IsActualStateFreshReturns(result1 bool, result2 error) {
	fake.IsActualStateFreshStub = nil
	fake.isActualStateFreshReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) VerifyFreshness(arg1 time.Time) error {
	fake.verifyFreshnessMutex.Lock()
	fake.verifyFreshnessArgsForCall = append(fake.verifyFreshnessArgsForCall, struct {
		arg1 time.Time
	}{arg1})
	fake.verifyFreshnessMutex.Unlock()
	if fake.VerifyFreshnessStub != nil {
		return fake.VerifyFreshnessStub(arg1)
	} else {
		return fake.verifyFreshnessReturns.result1
	}
}

func (fake *FakeStore) VerifyFreshnessCallCount() int {
	fake.verifyFreshnessMutex.RLock()
	defer fake.verifyFreshnessMutex.RUnlock()
	return len(fake.verifyFreshnessArgsForCall)
}

func (fake *FakeStore) VerifyFreshnessArgsForCall(i int) time.Time {
	fake.verifyFreshnessMutex.RLock()
	defer fake.verifyFreshnessMutex.RUnlock()
	return fake.verifyFreshnessArgsForCall[i].arg1
}

func (fake *FakeStore) VerifyFreshnessReturns(result1 error) {
	fake.VerifyFreshnessStub = nil
	fake.verifyFreshnessReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) AppKey(appGuid string, appVersion string) string {
	fake.appKeyMutex.Lock()
	fake.appKeyArgsForCall = append(fake.appKeyArgsForCall, struct {
		appGuid    string
		appVersion string
	}{appGuid, appVersion})
	fake.appKeyMutex.Unlock()
	if fake.AppKeyStub != nil {
		return fake.AppKeyStub(appGuid, appVersion)
	} else {
		return fake.appKeyReturns.result1
	}
}

func (fake *FakeStore) AppKeyCallCount() int {
	fake.appKeyMutex.RLock()
	defer fake.appKeyMutex.RUnlock()
	return len(fake.appKeyArgsForCall)
}

func (fake *FakeStore) AppKeyArgsForCall(i int) (string, string) {
	fake.appKeyMutex.RLock()
	defer fake.appKeyMutex.RUnlock()
	return fake.appKeyArgsForCall[i].appGuid, fake.appKeyArgsForCall[i].appVersion
}

func (fake *FakeStore) AppKeyReturns(result1 string) {
	fake.AppKeyStub = nil
	fake.appKeyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStore) GetApps() (map[string]*models.App, error) {
	fake.getAppsMutex.Lock()
	fake.getAppsArgsForCall = append(fake.getAppsArgsForCall, struct{}{})
	fake.getAppsMutex.Unlock()
	if fake.GetAppsStub != nil {
		return fake.GetAppsStub()
	} else {
		return fake.getAppsReturns.result1, fake.getAppsReturns.result2
	}
}

func (fake *FakeStore) GetAppsCallCount() int {
	fake.getAppsMutex.RLock()
	defer fake.getAppsMutex.RUnlock()
	return len(fake.getAppsArgsForCall)
}

func (fake *FakeStore) GetAppsReturns(result1 map[string]*models.App, result2 error) {
	fake.GetAppsStub = nil
	fake.getAppsReturns = struct {
		result1 map[string]*models.App
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetApp(appGuid string, appVersion string) (*models.App, error) {
	fake.getAppMutex.Lock()
	fake.getAppArgsForCall = append(fake.getAppArgsForCall, struct {
		appGuid    string
		appVersion string
	}{appGuid, appVersion})
	fake.getAppMutex.Unlock()
	if fake.GetAppStub != nil {
		return fake.GetAppStub(appGuid, appVersion)
	} else {
		return fake.getAppReturns.result1, fake.getAppReturns.result2
	}
}

func (fake *FakeStore) GetAppCallCount() int {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return len(fake.getAppArgsForCall)
}

func (fake *FakeStore) GetAppArgsForCall(i int) (string, string) {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return fake.getAppArgsForCall[i].appGuid, fake.getAppArgsForCall[i].appVersion
}

func (fake *FakeStore) GetAppReturns(result1 *models.App, result2 error) {
	fake.GetAppStub = nil
	fake.getAppReturns = struct {
		result1 *models.App
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) SyncDesiredState(desiredStates ...models.DesiredAppState) error {
	fake.syncDesiredStateMutex.Lock()
	fake.syncDesiredStateArgsForCall = append(fake.syncDesiredStateArgsForCall, struct {
		desiredStates []models.DesiredAppState
	}{desiredStates})
	fake.syncDesiredStateMutex.Unlock()
	if fake.SyncDesiredStateStub != nil {
		return fake.SyncDesiredStateStub(desiredStates...)
	} else {
		return fake.syncDesiredStateReturns.result1
	}
}

func (fake *FakeStore) SyncDesiredStateCallCount() int {
	fake.syncDesiredStateMutex.RLock()
	defer fake.syncDesiredStateMutex.RUnlock()
	return len(fake.syncDesiredStateArgsForCall)
}

func (fake *FakeStore) SyncDesiredStateArgsForCall(i int) []models.DesiredAppState {
	fake.syncDesiredStateMutex.RLock()
	defer fake.syncDesiredStateMutex.RUnlock()
	return fake.syncDesiredStateArgsForCall[i].desiredStates
}

func (fake *FakeStore) SyncDesiredStateReturns(result1 error) {
	fake.SyncDesiredStateStub = nil
	fake.syncDesiredStateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) GetDesiredState() (map[string]models.DesiredAppState, error) {
	fake.getDesiredStateMutex.Lock()
	fake.getDesiredStateArgsForCall = append(fake.getDesiredStateArgsForCall, struct{}{})
	fake.getDesiredStateMutex.Unlock()
	if fake.GetDesiredStateStub != nil {
		return fake.GetDesiredStateStub()
	} else {
		return fake.getDesiredStateReturns.result1, fake.getDesiredStateReturns.result2
	}
}

func (fake *FakeStore) GetDesiredStateCallCount() int {
	fake.getDesiredStateMutex.RLock()
	defer fake.getDesiredStateMutex.RUnlock()
	return len(fake.getDesiredStateArgsForCall)
}

func (fake *FakeStore) GetDesiredStateReturns(result1 map[string]models.DesiredAppState, result2 error) {
	fake.GetDesiredStateStub = nil
	fake.getDesiredStateReturns = struct {
		result1 map[string]models.DesiredAppState
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) SyncHeartbeats(heartbeat ...models.Heartbeat) error {
	fake.syncHeartbeatsMutex.Lock()
	fake.syncHeartbeatsArgsForCall = append(fake.syncHeartbeatsArgsForCall, struct {
		heartbeat []models.Heartbeat
	}{heartbeat})
	fake.syncHeartbeatsMutex.Unlock()
	if fake.SyncHeartbeatsStub != nil {
		return fake.SyncHeartbeatsStub(heartbeat...)
	} else {
		return fake.syncHeartbeatsReturns.result1
	}
}

func (fake *FakeStore) SyncHeartbeatsCallCount() int {
	fake.syncHeartbeatsMutex.RLock()
	defer fake.syncHeartbeatsMutex.RUnlock()
	return len(fake.syncHeartbeatsArgsForCall)
}

func (fake *FakeStore) SyncHeartbeatsArgsForCall(i int) []models.Heartbeat {
	fake.syncHeartbeatsMutex.RLock()
	defer fake.syncHeartbeatsMutex.RUnlock()
	return fake.syncHeartbeatsArgsForCall[i].heartbeat
}

func (fake *FakeStore) SyncHeartbeatsReturns(result1 error) {
	fake.SyncHeartbeatsStub = nil
	fake.syncHeartbeatsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) GetInstanceHeartbeats() (results []models.InstanceHeartbeat, err error) {
	fake.getInstanceHeartbeatsMutex.Lock()
	fake.getInstanceHeartbeatsArgsForCall = append(fake.getInstanceHeartbeatsArgsForCall, struct{}{})
	fake.getInstanceHeartbeatsMutex.Unlock()
	if fake.GetInstanceHeartbeatsStub != nil {
		return fake.GetInstanceHeartbeatsStub()
	} else {
		return fake.getInstanceHeartbeatsReturns.result1, fake.getInstanceHeartbeatsReturns.result2
	}
}

func (fake *FakeStore) GetInstanceHeartbeatsCallCount() int {
	fake.getInstanceHeartbeatsMutex.RLock()
	defer fake.getInstanceHeartbeatsMutex.RUnlock()
	return len(fake.getInstanceHeartbeatsArgsForCall)
}

func (fake *FakeStore) GetInstanceHeartbeatsReturns(result1 []models.InstanceHeartbeat, result2 error) {
	fake.GetInstanceHeartbeatsStub = nil
	fake.getInstanceHeartbeatsReturns = struct {
		result1 []models.InstanceHeartbeat
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) GetInstanceHeartbeatsForApp(appGuid string, appVersion string) (results []models.InstanceHeartbeat, err error) {
	fake.getInstanceHeartbeatsForAppMutex.Lock()
	fake.getInstanceHeartbeatsForAppArgsForCall = append(fake.getInstanceHeartbeatsForAppArgsForCall, struct {
		appGuid    string
		appVersion string
	}{appGuid, appVersion})
	fake.getInstanceHeartbeatsForAppMutex.Unlock()
	if fake.GetInstanceHeartbeatsForAppStub != nil {
		return fake.GetInstanceHeartbeatsForAppStub(appGuid, appVersion)
	} else {
		return fake.getInstanceHeartbeatsForAppReturns.result1, fake.getInstanceHeartbeatsForAppReturns.result2
	}
}

func (fake *FakeStore) GetInstanceHeartbeatsForAppCallCount() int {
	fake.getInstanceHeartbeatsForAppMutex.RLock()
	defer fake.getInstanceHeartbeatsForAppMutex.RUnlock()
	return len(fake.getInstanceHeartbeatsForAppArgsForCall)
}

func (fake *FakeStore) GetInstanceHeartbeatsForAppArgsForCall(i int) (string, string) {
	fake.getInstanceHeartbeatsForAppMutex.RLock()
	defer fake.getInstanceHeartbeatsForAppMutex.RUnlock()
	return fake.getInstanceHeartbeatsForAppArgsForCall[i].appGuid, fake.getInstanceHeartbeatsForAppArgsForCall[i].appVersion
}

func (fake *FakeStore) GetInstanceHeartbeatsForAppReturns(result1 []models.InstanceHeartbeat, result2 error) {
	fake.GetInstanceHeartbeatsForAppStub = nil
	fake.getInstanceHeartbeatsForAppReturns = struct {
		result1 []models.InstanceHeartbeat
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) SaveCrashCounts(crashCounts ...models.CrashCount) error {
	fake.saveCrashCountsMutex.Lock()
	fake.saveCrashCountsArgsForCall = append(fake.saveCrashCountsArgsForCall, struct {
		crashCounts []models.CrashCount
	}{crashCounts})
	fake.saveCrashCountsMutex.Unlock()
	if fake.SaveCrashCountsStub != nil {
		return fake.SaveCrashCountsStub(crashCounts...)
	} else {
		return fake.saveCrashCountsReturns.result1
	}
}

func (fake *FakeStore) SaveCrashCountsCallCount() int {
	fake.saveCrashCountsMutex.RLock()
	defer fake.saveCrashCountsMutex.RUnlock()
	return len(fake.saveCrashCountsArgsForCall)
}

func (fake *FakeStore) SaveCrashCountsArgsForCall(i int) []models.CrashCount {
	fake.saveCrashCountsMutex.RLock()
	defer fake.saveCrashCountsMutex.RUnlock()
	return fake.saveCrashCountsArgsForCall[i].crashCounts
}

func (fake *FakeStore) SaveCrashCountsReturns(result1 error) {
	fake.SaveCrashCountsStub = nil
	fake.saveCrashCountsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) SavePendingStartMessages(startMessages ...models.PendingStartMessage) error {
	fake.savePendingStartMessagesMutex.Lock()
	fake.savePendingStartMessagesArgsForCall = append(fake.savePendingStartMessagesArgsForCall, struct {
		startMessages []models.PendingStartMessage
	}{startMessages})
	fake.savePendingStartMessagesMutex.Unlock()
	if fake.SavePendingStartMessagesStub != nil {
		return fake.SavePendingStartMessagesStub(startMessages...)
	} else {
		return fake.savePendingStartMessagesReturns.result1
	}
}

func (fake *FakeStore) SavePendingStartMessagesCallCount() int {
	fake.savePendingStartMessagesMutex.RLock()
	defer fake.savePendingStartMessagesMutex.RUnlock()
	return len(fake.savePendingStartMessagesArgsForCall)
}

func (fake *FakeStore) SavePendingStartMessagesArgsForCall(i int) []models.PendingStartMessage {
	fake.savePendingStartMessagesMutex.RLock()
	defer fake.savePendingStartMessagesMutex.RUnlock()
	return fake.savePendingStartMessagesArgsForCall[i].startMessages
}

func (fake *FakeStore) SavePendingStartMessagesReturns(result1 error) {
	fake.SavePendingStartMessagesStub = nil
	fake.savePendingStartMessagesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) GetPendingStartMessages() (map[string]models.PendingStartMessage, error) {
	fake.getPendingStartMessagesMutex.Lock()
	fake.getPendingStartMessagesArgsForCall = append(fake.getPendingStartMessagesArgsForCall, struct{}{})
	fake.getPendingStartMessagesMutex.Unlock()
	if fake.GetPendingStartMessagesStub != nil {
		return fake.GetPendingStartMessagesStub()
	} else {
		return fake.getPendingStartMessagesReturns.result1, fake.getPendingStartMessagesReturns.result2
	}
}

func (fake *FakeStore) GetPendingStartMessagesCallCount() int {
	fake.getPendingStartMessagesMutex.RLock()
	defer fake.getPendingStartMessagesMutex.RUnlock()
	return len(fake.getPendingStartMessagesArgsForCall)
}

func (fake *FakeStore) GetPendingStartMessagesReturns(result1 map[string]models.PendingStartMessage, result2 error) {
	fake.GetPendingStartMessagesStub = nil
	fake.getPendingStartMessagesReturns = struct {
		result1 map[string]models.PendingStartMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) DeletePendingStartMessages(startMessages ...models.PendingStartMessage) error {
	fake.deletePendingStartMessagesMutex.Lock()
	fake.deletePendingStartMessagesArgsForCall = append(fake.deletePendingStartMessagesArgsForCall, struct {
		startMessages []models.PendingStartMessage
	}{startMessages})
	fake.deletePendingStartMessagesMutex.Unlock()
	if fake.DeletePendingStartMessagesStub != nil {
		return fake.DeletePendingStartMessagesStub(startMessages...)
	} else {
		return fake.deletePendingStartMessagesReturns.result1
	}
}

func (fake *FakeStore) DeletePendingStartMessagesCallCount() int {
	fake.deletePendingStartMessagesMutex.RLock()
	defer fake.deletePendingStartMessagesMutex.RUnlock()
	return len(fake.deletePendingStartMessagesArgsForCall)
}

func (fake *FakeStore) DeletePendingStartMessagesArgsForCall(i int) []models.PendingStartMessage {
	fake.deletePendingStartMessagesMutex.RLock()
	defer fake.deletePendingStartMessagesMutex.RUnlock()
	return fake.deletePendingStartMessagesArgsForCall[i].startMessages
}

func (fake *FakeStore) DeletePendingStartMessagesReturns(result1 error) {
	fake.DeletePendingStartMessagesStub = nil
	fake.deletePendingStartMessagesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) SavePendingStopMessages(stopMessages ...models.PendingStopMessage) error {
	fake.savePendingStopMessagesMutex.Lock()
	fake.savePendingStopMessagesArgsForCall = append(fake.savePendingStopMessagesArgsForCall, struct {
		stopMessages []models.PendingStopMessage
	}{stopMessages})
	fake.savePendingStopMessagesMutex.Unlock()
	if fake.SavePendingStopMessagesStub != nil {
		return fake.SavePendingStopMessagesStub(stopMessages...)
	} else {
		return fake.savePendingStopMessagesReturns.result1
	}
}

func (fake *FakeStore) SavePendingStopMessagesCallCount() int {
	fake.savePendingStopMessagesMutex.RLock()
	defer fake.savePendingStopMessagesMutex.RUnlock()
	return len(fake.savePendingStopMessagesArgsForCall)
}

func (fake *FakeStore) SavePendingStopMessagesArgsForCall(i int) []models.PendingStopMessage {
	fake.savePendingStopMessagesMutex.RLock()
	defer fake.savePendingStopMessagesMutex.RUnlock()
	return fake.savePendingStopMessagesArgsForCall[i].stopMessages
}

func (fake *FakeStore) SavePendingStopMessagesReturns(result1 error) {
	fake.SavePendingStopMessagesStub = nil
	fake.savePendingStopMessagesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) GetPendingStopMessages() (map[string]models.PendingStopMessage, error) {
	fake.getPendingStopMessagesMutex.Lock()
	fake.getPendingStopMessagesArgsForCall = append(fake.getPendingStopMessagesArgsForCall, struct{}{})
	fake.getPendingStopMessagesMutex.Unlock()
	if fake.GetPendingStopMessagesStub != nil {
		return fake.GetPendingStopMessagesStub()
	} else {
		return fake.getPendingStopMessagesReturns.result1, fake.getPendingStopMessagesReturns.result2
	}
}

func (fake *FakeStore) GetPendingStopMessagesCallCount() int {
	fake.getPendingStopMessagesMutex.RLock()
	defer fake.getPendingStopMessagesMutex.RUnlock()
	return len(fake.getPendingStopMessagesArgsForCall)
}

func (fake *FakeStore) GetPendingStopMessagesReturns(result1 map[string]models.PendingStopMessage, result2 error) {
	fake.GetPendingStopMessagesStub = nil
	fake.getPendingStopMessagesReturns = struct {
		result1 map[string]models.PendingStopMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) DeletePendingStopMessages(stopMessages ...models.PendingStopMessage) error {
	fake.deletePendingStopMessagesMutex.Lock()
	fake.deletePendingStopMessagesArgsForCall = append(fake.deletePendingStopMessagesArgsForCall, struct {
		stopMessages []models.PendingStopMessage
	}{stopMessages})
	fake.deletePendingStopMessagesMutex.Unlock()
	if fake.DeletePendingStopMessagesStub != nil {
		return fake.DeletePendingStopMessagesStub(stopMessages...)
	} else {
		return fake.deletePendingStopMessagesReturns.result1
	}
}

func (fake *FakeStore) DeletePendingStopMessagesCallCount() int {
	fake.deletePendingStopMessagesMutex.RLock()
	defer fake.deletePendingStopMessagesMutex.RUnlock()
	return len(fake.deletePendingStopMessagesArgsForCall)
}

func (fake *FakeStore) DeletePendingStopMessagesArgsForCall(i int) []models.PendingStopMessage {
	fake.deletePendingStopMessagesMutex.RLock()
	defer fake.deletePendingStopMessagesMutex.RUnlock()
	return fake.deletePendingStopMessagesArgsForCall[i].stopMessages
}

func (fake *FakeStore) DeletePendingStopMessagesReturns(result1 error) {
	fake.DeletePendingStopMessagesStub = nil
	fake.deletePendingStopMessagesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) SaveMetric(metric string, value float64) error {
	fake.saveMetricMutex.Lock()
	fake.saveMetricArgsForCall = append(fake.saveMetricArgsForCall, struct {
		metric string
		value  float64
	}{metric, value})
	fake.saveMetricMutex.Unlock()
	if fake.SaveMetricStub != nil {
		return fake.SaveMetricStub(metric, value)
	} else {
		return fake.saveMetricReturns.result1
	}
}

func (fake *FakeStore) SaveMetricCallCount() int {
	fake.saveMetricMutex.RLock()
	defer fake.saveMetricMutex.RUnlock()
	return len(fake.saveMetricArgsForCall)
}

func (fake *FakeStore) SaveMetricArgsForCall(i int) (string, float64) {
	fake.saveMetricMutex.RLock()
	defer fake.saveMetricMutex.RUnlock()
	return fake.saveMetricArgsForCall[i].metric, fake.saveMetricArgsForCall[i].value
}

func (fake *FakeStore) SaveMetricReturns(result1 error) {
	fake.SaveMetricStub = nil
	fake.saveMetricReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) GetMetric(metric string) (float64, error) {
	fake.getMetricMutex.Lock()
	fake.getMetricArgsForCall = append(fake.getMetricArgsForCall, struct {
		metric string
	}{metric})
	fake.getMetricMutex.Unlock()
	if fake.GetMetricStub != nil {
		return fake.GetMetricStub(metric)
	} else {
		return fake.getMetricReturns.result1, fake.getMetricReturns.result2
	}
}

func (fake *FakeStore) GetMetricCallCount() int {
	fake.getMetricMutex.RLock()
	defer fake.getMetricMutex.RUnlock()
	return len(fake.getMetricArgsForCall)
}

func (fake *FakeStore) GetMetricArgsForCall(i int) string {
	fake.getMetricMutex.RLock()
	defer fake.getMetricMutex.RUnlock()
	return fake.getMetricArgsForCall[i].metric
}

func (fake *FakeStore) GetMetricReturns(result1 float64, result2 error) {
	fake.GetMetricStub = nil
	fake.getMetricReturns = struct {
		result1 float64
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) Compact() error {
	fake.compactMutex.Lock()
	fake.compactArgsForCall = append(fake.compactArgsForCall, struct{}{})
	fake.compactMutex.Unlock()
	if fake.CompactStub != nil {
		return fake.CompactStub()
	} else {
		return fake.compactReturns.result1
	}
}

func (fake *FakeStore) CompactCallCount() int {
	fake.compactMutex.RLock()
	defer fake.compactMutex.RUnlock()
	return len(fake.compactArgsForCall)
}

func (fake *FakeStore) CompactReturns(result1 error) {
	fake.CompactStub = nil
	fake.compactReturns = struct {
		result1 error
	}{result1}
}

var _ store.Store = new(FakeStore)
