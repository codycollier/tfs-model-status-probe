package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/codycollier/tfs-model-status-probe/tfproto/tfproto"
)

func TestResponseEmpty(t *testing.T) {
	request := &tfproto.GetModelStatusResponse{}
	retval := checkServableResponse(request)
	assert.Equal(t, 8, retval, "Expecting response code for empty")
}

func TestResponseStateUnknown(t *testing.T) {
	request := &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 123,
				State:   tfproto.ModelVersionStatus_UNKNOWN,
			},
		},
	}
	retval := checkServableResponse(request)
	assert.Equal(t, 10, retval, "Expecting response code for state Unknown")
}

func TestResponseStateStart(t *testing.T) {
	request := &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 123,
				State:   tfproto.ModelVersionStatus_START,
			},
		},
	}
	retval := checkServableResponse(request)
	assert.Equal(t, 20, retval, "Expecting response code for state Start")
}

func TestResponseStateLoading(t *testing.T) {
	request := &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 123,
				State:   tfproto.ModelVersionStatus_LOADING,
			},
		},
	}
	retval := checkServableResponse(request)
	assert.Equal(t, 30, retval, "Expecting response code for Loading")
}

func TestResponseStateUnloading(t *testing.T) {
	request := &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 123,
				State:   tfproto.ModelVersionStatus_UNLOADING,
			},
		},
	}
	retval := checkServableResponse(request)
	assert.Equal(t, 40, retval, "Expecting response code for state Unloading")
}

func TestResponseStateEnd(t *testing.T) {
	request := &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 123,
				State:   tfproto.ModelVersionStatus_END,
			},
		},
	}
	retval := checkServableResponse(request)
	assert.Equal(t, 50, retval, "Expecting response code for state End")
}

func TestResponseStateAvailable(t *testing.T) {
	request := &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 123,
				State:   tfproto.ModelVersionStatus_AVAILABLE,
			},
		},
	}
	retval := checkServableResponse(request)
	assert.Equal(t, 0, retval, "Expecting response code for state Available")
}
