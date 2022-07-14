package presenters

import (
	"reflect"
	"testing"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/api/private"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/internal/api/dbapi"
	mock "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/internal/kafka/test/mocks/data_plane"
	"github.com/onsi/gomega"
)

func TestConvertDataPlaneClusterStatus_AvailableStrimziVersions(t *testing.T) {
	tests := []struct {
		name                            string
		inputClusterUpdateStatusRequest func() *private.DataPlaneClusterUpdateStatusRequest
		want                            []api.StrimziVersion
		wantErr                         bool
	}{
		{
			name: "When setting a non empty ordered list of strimzi it is stored as is",
			inputClusterUpdateStatusRequest: func() *private.DataPlaneClusterUpdateStatusRequest {
				request := sampleValidDataPlaneClusterUpdateStatusRequest()
				request.Strimzi = []private.DataPlaneClusterUpdateStatusRequestStrimzi{
					{Version: "v1.0.0-0", Ready: true},
					{Version: "v2.0.0-0", Ready: false},
					{Version: "v3.0.0-0", Ready: true},
				}
				return request
			},
			want: []api.StrimziVersion{
				{Version: "v1.0.0-0", Ready: true},
				{Version: "v2.0.0-0", Ready: false},
				{Version: "v3.0.0-0", Ready: true},
			},
			wantErr: false,
		},
		{
			name: "When setting a non empty unordered list of strimzi that list is stored in semver ascending order from the version attribute",
			inputClusterUpdateStatusRequest: func() *private.DataPlaneClusterUpdateStatusRequest {
				request := sampleValidDataPlaneClusterUpdateStatusRequest()
				request.Strimzi = []private.DataPlaneClusterUpdateStatusRequestStrimzi{
					{Version: "v5.0.0-0", Ready: true},
					{Version: "v2.0.0-0", Ready: false},
					{Version: "v3.0.0-0", Ready: true},
				}
				return request
			},
			want: []api.StrimziVersion{
				{Version: "v2.0.0-0", Ready: false},
				{Version: "v3.0.0-0", Ready: true},
				{Version: "v5.0.0-0", Ready: true},
			},
			wantErr: false,
		},
		{
			name: "When setting an empty list of strimzi that list is stored as the empty list",
			inputClusterUpdateStatusRequest: func() *private.DataPlaneClusterUpdateStatusRequest {
				request := sampleValidDataPlaneClusterUpdateStatusRequest()
				request.Strimzi = []private.DataPlaneClusterUpdateStatusRequestStrimzi{}
				return request
			},
			want:    []api.StrimziVersion{},
			wantErr: false,
		},
		{
			name: "When setting a nil list of strimzi that list is stored as the empty list",
			inputClusterUpdateStatusRequest: func() *private.DataPlaneClusterUpdateStatusRequest {
				request := sampleValidDataPlaneClusterUpdateStatusRequest()
				request.Strimzi = nil
				return request
			},
			want:    []api.StrimziVersion{},
			wantErr: false,
		},
		{
			name: "When strimzi nor strimziVersions are defined an empty list is returned",
			inputClusterUpdateStatusRequest: func() *private.DataPlaneClusterUpdateStatusRequest {
				request := sampleValidDataPlaneClusterUpdateStatusRequest()
				return request
			},
			want:    []api.StrimziVersion{},
			wantErr: false,
		},
		{
			name: "When one of the versions in strimzi does not follow the expected format an error is returned",
			inputClusterUpdateStatusRequest: func() *private.DataPlaneClusterUpdateStatusRequest {
				request := sampleValidDataPlaneClusterUpdateStatusRequest()
				request.Strimzi = []private.DataPlaneClusterUpdateStatusRequestStrimzi{
					{Version: "v1.0.0-0", Ready: true},
					{Version: "v2.0.0", Ready: false},
					{Version: "v3.0.0-0", Ready: true},
				}
				return request
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			inputClusterStatusRequest := tt.inputClusterUpdateStatusRequest()
			res, err := ConvertDataPlaneClusterStatus(*inputClusterStatusRequest)
			gotErr := err != nil
			errResultTestFailed := false
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				errResultTestFailed = true
				t.Errorf("wantErr: %v got: %v", tt.wantErr, gotErr)
			}
			if !errResultTestFailed && !tt.wantErr {
				if !reflect.DeepEqual(res.AvailableStrimziVersions, tt.want) {
					t.Errorf("want: %+v got: %+v", tt.want, res.AvailableStrimziVersions)
				}
			}
		})
	}
}

func TestPresentDataPlaneClusterConfig(t *testing.T) {
	type args struct {
		config *dbapi.DataPlaneClusterConfig
	}

	tests := []struct {
		name string
		args args
		want private.DataplaneClusterAgentConfig
	}{
		{
			name: "should return converted DataplaneClusterAgentConfig",
			args: args{
				config: mock.BuildDataPlaneClusterConfig(func(config *dbapi.DataPlaneClusterConfig) {
					config.DynamicCapacityInfo = map[string]api.DynamicCapacityInfo{
						"key": {
							MaxNodes:       4,
							MaxUnits:       2,
							RemainingUnits: 1,
						},
					}
				}),
			},
			want: mock.BuildDataplaneClusterAgentConfig(func(config private.DataplaneClusterAgentConfig) {
				config.Spec.Capacity["key"] = private.DataplaneClusterAgentConfigSpecCapacity{
					MaxNodes: 4,
				}
			}),
		},
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			g.Expect(PresentDataPlaneClusterConfig(tt.args.config)).To(gomega.Equal(tt.want))
		})
	}
}

func TestGetAvailableStrimziVersions(t *testing.T) {
	type args struct {
		status private.DataPlaneClusterUpdateStatusRequest
	}

	tests := []struct {
		name string
		args args
		want []api.StrimziVersion
	}{
		{
			name: "should return available strimzi versions from DataPlaneClusterUpdateStatusRequest",
			args: args{
				status: mock.BuildDataPlaneClusterUpdateStatusRequest(nil),
			},
			want: *mock.BuildApiStrimziVersions(nil),
		},
	}

	for _, testcase := range tests {
		tt := testcase

		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			convertedConfig, err := getAvailableStrimziVersions(tt.args.status)
			if err == nil {
				g.Expect(convertedConfig).To(gomega.Equal(tt.want))
			}
		})
	}
}

func sampleValidDataPlaneClusterUpdateStatusRequest() *private.DataPlaneClusterUpdateStatusRequest {
	return &private.DataPlaneClusterUpdateStatusRequest{
		Conditions: []private.DataPlaneClusterUpdateStatusRequestConditions{
			{
				Type:   "Ready",
				Status: "True",
			},
		},
	}
}
