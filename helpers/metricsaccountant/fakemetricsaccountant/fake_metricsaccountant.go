// This file was generated by counterfeiter
package fakemetricsaccountant

import (
	"sync"

	"github.com/cloudfoundry/hm9000/helpers/metricsaccountant"
	"github.com/cloudfoundry/hm9000/models"
)

type FakeMetricsAccountant struct {
	TrackReceivedHeartbeatsStub        func(metric int) error
	trackReceivedHeartbeatsMutex       sync.RWMutex
	trackReceivedHeartbeatsArgsForCall []struct {
		metric int
	}
	trackReceivedHeartbeatsReturns struct {
		result1 error
	}
	TrackSavedHeartbeatsStub        func(metric int) error
	trackSavedHeartbeatsMutex       sync.RWMutex
	trackSavedHeartbeatsArgsForCall []struct {
		metric int
	}
	trackSavedHeartbeatsReturns struct {
		result1 error
	}
	IncrementSentMessageMetricsStub        func(starts []models.PendingStartMessage, stops []models.PendingStopMessage) error
	incrementSentMessageMetricsMutex       sync.RWMutex
	incrementSentMessageMetricsArgsForCall []struct {
		starts []models.PendingStartMessage
		stops  []models.PendingStopMessage
	}
	incrementSentMessageMetricsReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMetricsAccountant) TrackReceivedHeartbeats(metric int) error {
	fake.trackReceivedHeartbeatsMutex.Lock()
	fake.trackReceivedHeartbeatsArgsForCall = append(fake.trackReceivedHeartbeatsArgsForCall, struct {
		metric int
	}{metric})
	fake.recordInvocation("TrackReceivedHeartbeats", []interface{}{metric})
	fake.trackReceivedHeartbeatsMutex.Unlock()
	if fake.TrackReceivedHeartbeatsStub != nil {
		return fake.TrackReceivedHeartbeatsStub(metric)
	} else {
		return fake.trackReceivedHeartbeatsReturns.result1
	}
}

func (fake *FakeMetricsAccountant) TrackReceivedHeartbeatsCallCount() int {
	fake.trackReceivedHeartbeatsMutex.RLock()
	defer fake.trackReceivedHeartbeatsMutex.RUnlock()
	return len(fake.trackReceivedHeartbeatsArgsForCall)
}

func (fake *FakeMetricsAccountant) TrackReceivedHeartbeatsArgsForCall(i int) int {
	fake.trackReceivedHeartbeatsMutex.RLock()
	defer fake.trackReceivedHeartbeatsMutex.RUnlock()
	return fake.trackReceivedHeartbeatsArgsForCall[i].metric
}

func (fake *FakeMetricsAccountant) TrackReceivedHeartbeatsReturns(result1 error) {
	fake.TrackReceivedHeartbeatsStub = nil
	fake.trackReceivedHeartbeatsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMetricsAccountant) TrackSavedHeartbeats(metric int) error {
	fake.trackSavedHeartbeatsMutex.Lock()
	fake.trackSavedHeartbeatsArgsForCall = append(fake.trackSavedHeartbeatsArgsForCall, struct {
		metric int
	}{metric})
	fake.recordInvocation("TrackSavedHeartbeats", []interface{}{metric})
	fake.trackSavedHeartbeatsMutex.Unlock()
	if fake.TrackSavedHeartbeatsStub != nil {
		return fake.TrackSavedHeartbeatsStub(metric)
	} else {
		return fake.trackSavedHeartbeatsReturns.result1
	}
}

func (fake *FakeMetricsAccountant) TrackSavedHeartbeatsCallCount() int {
	fake.trackSavedHeartbeatsMutex.RLock()
	defer fake.trackSavedHeartbeatsMutex.RUnlock()
	return len(fake.trackSavedHeartbeatsArgsForCall)
}

func (fake *FakeMetricsAccountant) TrackSavedHeartbeatsArgsForCall(i int) int {
	fake.trackSavedHeartbeatsMutex.RLock()
	defer fake.trackSavedHeartbeatsMutex.RUnlock()
	return fake.trackSavedHeartbeatsArgsForCall[i].metric
}

func (fake *FakeMetricsAccountant) TrackSavedHeartbeatsReturns(result1 error) {
	fake.TrackSavedHeartbeatsStub = nil
	fake.trackSavedHeartbeatsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMetricsAccountant) IncrementSentMessageMetrics(starts []models.PendingStartMessage, stops []models.PendingStopMessage) error {
	var startsCopy []models.PendingStartMessage
	if starts != nil {
		startsCopy = make([]models.PendingStartMessage, len(starts))
		copy(startsCopy, starts)
	}
	var stopsCopy []models.PendingStopMessage
	if stops != nil {
		stopsCopy = make([]models.PendingStopMessage, len(stops))
		copy(stopsCopy, stops)
	}
	fake.incrementSentMessageMetricsMutex.Lock()
	fake.incrementSentMessageMetricsArgsForCall = append(fake.incrementSentMessageMetricsArgsForCall, struct {
		starts []models.PendingStartMessage
		stops  []models.PendingStopMessage
	}{startsCopy, stopsCopy})
	fake.recordInvocation("IncrementSentMessageMetrics", []interface{}{startsCopy, stopsCopy})
	fake.incrementSentMessageMetricsMutex.Unlock()
	if fake.IncrementSentMessageMetricsStub != nil {
		return fake.IncrementSentMessageMetricsStub(starts, stops)
	} else {
		return fake.incrementSentMessageMetricsReturns.result1
	}
}

func (fake *FakeMetricsAccountant) IncrementSentMessageMetricsCallCount() int {
	fake.incrementSentMessageMetricsMutex.RLock()
	defer fake.incrementSentMessageMetricsMutex.RUnlock()
	return len(fake.incrementSentMessageMetricsArgsForCall)
}

func (fake *FakeMetricsAccountant) IncrementSentMessageMetricsArgsForCall(i int) ([]models.PendingStartMessage, []models.PendingStopMessage) {
	fake.incrementSentMessageMetricsMutex.RLock()
	defer fake.incrementSentMessageMetricsMutex.RUnlock()
	return fake.incrementSentMessageMetricsArgsForCall[i].starts, fake.incrementSentMessageMetricsArgsForCall[i].stops
}

func (fake *FakeMetricsAccountant) IncrementSentMessageMetricsReturns(result1 error) {
	fake.IncrementSentMessageMetricsStub = nil
	fake.incrementSentMessageMetricsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMetricsAccountant) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.trackReceivedHeartbeatsMutex.RLock()
	defer fake.trackReceivedHeartbeatsMutex.RUnlock()
	fake.trackSavedHeartbeatsMutex.RLock()
	defer fake.trackSavedHeartbeatsMutex.RUnlock()
	fake.incrementSentMessageMetricsMutex.RLock()
	defer fake.incrementSentMessageMetricsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMetricsAccountant) recordInvocation(key string, args []interface{}) {
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

var _ metricsaccountant.MetricsAccountant = new(FakeMetricsAccountant)
