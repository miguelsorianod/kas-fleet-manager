// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package services

import (
	"context"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/api/dbapi"
	serviceError "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"sync"
)

// Ensure, that DataPlaneKafkaServiceMock does implement DataPlaneKafkaService.
// If this is not the case, regenerate this file with moq.
var _ DataPlaneKafkaService = &DataPlaneKafkaServiceMock{}

// DataPlaneKafkaServiceMock is a mock implementation of DataPlaneKafkaService.
//
// 	func TestSomethingThatUsesDataPlaneKafkaService(t *testing.T) {
//
// 		// make and configure a mocked DataPlaneKafkaService
// 		mockedDataPlaneKafkaService := &DataPlaneKafkaServiceMock{
// 			UpdateDataPlaneKafkaServiceFunc: func(ctx context.Context, clusterId string, status []*dbapi.DataPlaneKafkaStatus) *serviceError.ServiceError {
// 				panic("mock out the UpdateDataPlaneKafkaService method")
// 			},
// 		}
//
// 		// use mockedDataPlaneKafkaService in code that requires DataPlaneKafkaService
// 		// and then make assertions.
//
// 	}
type DataPlaneKafkaServiceMock struct {
	// UpdateDataPlaneKafkaServiceFunc mocks the UpdateDataPlaneKafkaService method.
	UpdateDataPlaneKafkaServiceFunc func(ctx context.Context, clusterId string, status []*dbapi.DataPlaneKafkaStatus) *serviceError.ServiceError

	// calls tracks calls to the methods.
	calls struct {
		// UpdateDataPlaneKafkaService holds details about calls to the UpdateDataPlaneKafkaService method.
		UpdateDataPlaneKafkaService []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ClusterId is the clusterId argument value.
			ClusterId string
			// Status is the status argument value.
			Status []*dbapi.DataPlaneKafkaStatus
		}
	}
	lockUpdateDataPlaneKafkaService sync.RWMutex
}

// UpdateDataPlaneKafkaService calls UpdateDataPlaneKafkaServiceFunc.
func (mock *DataPlaneKafkaServiceMock) UpdateDataPlaneKafkaService(ctx context.Context, clusterId string, status []*dbapi.DataPlaneKafkaStatus) *serviceError.ServiceError {
	if mock.UpdateDataPlaneKafkaServiceFunc == nil {
		panic("DataPlaneKafkaServiceMock.UpdateDataPlaneKafkaServiceFunc: method is nil but DataPlaneKafkaService.UpdateDataPlaneKafkaService was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		ClusterId string
		Status    []*dbapi.DataPlaneKafkaStatus
	}{
		Ctx:       ctx,
		ClusterId: clusterId,
		Status:    status,
	}
	mock.lockUpdateDataPlaneKafkaService.Lock()
	mock.calls.UpdateDataPlaneKafkaService = append(mock.calls.UpdateDataPlaneKafkaService, callInfo)
	mock.lockUpdateDataPlaneKafkaService.Unlock()
	return mock.UpdateDataPlaneKafkaServiceFunc(ctx, clusterId, status)
}

// UpdateDataPlaneKafkaServiceCalls gets all the calls that were made to UpdateDataPlaneKafkaService.
// Check the length with:
//     len(mockedDataPlaneKafkaService.UpdateDataPlaneKafkaServiceCalls())
func (mock *DataPlaneKafkaServiceMock) UpdateDataPlaneKafkaServiceCalls() []struct {
	Ctx       context.Context
	ClusterId string
	Status    []*dbapi.DataPlaneKafkaStatus
} {
	var calls []struct {
		Ctx       context.Context
		ClusterId string
		Status    []*dbapi.DataPlaneKafkaStatus
	}
	mock.lockUpdateDataPlaneKafkaService.RLock()
	calls = mock.calls.UpdateDataPlaneKafkaService
	mock.lockUpdateDataPlaneKafkaService.RUnlock()
	return calls
}
