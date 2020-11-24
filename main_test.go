package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/codycollier/tfs-model-status-probe/tfproto/tfproto"
)

func TestEmptyResponse(t *testing.T) {
	r := checkServableResponse(&tfproto.GetModelStatusResponse{})
	assert.Equal(t, 0, r, "Expecting 0 response code")
}
