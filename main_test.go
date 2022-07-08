//
// Copyright 2020 Cody Collier <cody@telnet.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/codycollier/tfs-model-status-probe/tfproto/tfproto"
)

func TestResponseEmpty(t *testing.T) {
	request := &tfproto.GetModelStatusResponse{}
	retval := checkServableResponse(request, 0)
	assert.Equal(t, 11, retval, "Expecting response code for empty")
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
	assert.Equal(t, 12, retval, "Expecting response code for empty")
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
	assert.Equal(t, 30, retval, "Expecting response code for state Unknown")
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
	assert.Equal(t, 31, retval, "Expecting response code for state Start")
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
	assert.Equal(t, 32, retval, "Expecting response code for Loading")
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
	assert.Equal(t, 33, retval, "Expecting response code for state Unloading")
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
	assert.Equal(t, 34, retval, "Expecting response code for state End")
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

func TestResponseStateAvailableOnNoVersion(t *testing.T) {

	// Ensure success when arbitrary version is available (ex: after a rollback)
	request := &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 101,
				State:   tfproto.ModelVersionStatus_END,
			},
			{
				Version: 98,
				State:   tfproto.ModelVersionStatus_END,
			},
			{
				Version: 301,
				State:   tfproto.ModelVersionStatus_AVAILABLE,
			},
			{
				Version: 303,
				State:   tfproto.ModelVersionStatus_END,
			},
		},
	}
	retval := checkServableResponse(request, 0)
	assert.Equal(t, 0, retval)

	// Ensure order is not relevant
	request = &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 101,
				State:   tfproto.ModelVersionStatus_AVAILABLE,
			},
			{
				Version: 301,
				State:   tfproto.ModelVersionStatus_END,
			},
		},
	}
	retval = checkServableResponse(request, 0)
	assert.Equal(t, 0, retval)

	request = &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 301,
				State:   tfproto.ModelVersionStatus_END,
			},
			{
				Version: 101,
				State:   tfproto.ModelVersionStatus_AVAILABLE,
			},
		},
	}
	retval = checkServableResponse(request, 0)
	assert.Equal(t, 0, retval)

	// Ensure still fails if there is no version which is available
	request = &tfproto.GetModelStatusResponse{
		ModelVersionStatus: []*tfproto.ModelVersionStatus{
			{
				Version: 101,
				State:   tfproto.ModelVersionStatus_END,
			},
			{
				Version: 301,
				State:   tfproto.ModelVersionStatus_END,
			},
			{
				Version: 303,
				State:   tfproto.ModelVersionStatus_END,
			},
		},
	}
	retval = checkServableResponse(request, 0)
	assert.Equal(t, 34, retval)

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
	retval := checkServableResponse(request, 101)
	assert.Equal(t, 34, retval)
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
	retval = checkServableResponse(request, 101)
	assert.Equal(t, 34, retval)
	retval = checkServableResponse(request, 301)
	assert.Equal(t, 0, retval)

}
