package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/codycollier/tfs-model-status-probe/tfproto/tfproto"
)

func TestResponseEmpty(t *testing.T) {
	request := &tfproto.GetModelStatusResponse{}
	retval := checkServableResponse(request, 0)
	assert.Equal(t, 303, retval, "Expecting response code for empty")
}

func TestResponseMissingVersion(t *testing.T) {
	request := &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 100,
				State:   tfproto.ModelVersionStatus_UNKNOWN,
			},
		},
	}
	retval := checkServableResponse(request, 300)
	assert.Equal(t, 304, retval, "Expecting response code for empty")
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
	retval := checkServableResponse(request, 0)
	assert.Equal(t, 310, retval, "Expecting response code for state Unknown")
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
	retval := checkServableResponse(request, 0)
	assert.Equal(t, 320, retval, "Expecting response code for state Start")
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
	retval := checkServableResponse(request, 0)
	assert.Equal(t, 330, retval, "Expecting response code for Loading")
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
	retval := checkServableResponse(request, 0)
	assert.Equal(t, 340, retval, "Expecting response code for state Unloading")
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
	retval := checkServableResponse(request, 0)
	assert.Equal(t, 350, retval, "Expecting response code for state End")
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
	retval := checkServableResponse(request, 0)
	assert.Equal(t, 0, retval, "Expecting response code for state Available")
}

func TestResponseStateAvailableOnSpecificVersion(t *testing.T) {

	request := &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 101,
				State:   tfproto.ModelVersionStatus_END,
			},
			{
				Version: 301,
				State:   tfproto.ModelVersionStatus_AVAILABLE,
			},
		},
	}
	retval := checkServableResponse(request, 0)
	assert.Equal(t, 350, retval)
	retval = checkServableResponse(request, 101)
	assert.Equal(t, 350, retval)
	retval = checkServableResponse(request, 301)
	assert.Equal(t, 0, retval)

	request = &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 301,
				State:   tfproto.ModelVersionStatus_AVAILABLE,
			},
			{
				Version: 101,
				State:   tfproto.ModelVersionStatus_END,
			},
		},
	}
	retval = checkServableResponse(request, 0)
	assert.Equal(t, 0, retval)
	retval = checkServableResponse(request, 101)
	assert.Equal(t, 350, retval)
	retval = checkServableResponse(request, 301)
	assert.Equal(t, 0, retval)

}
