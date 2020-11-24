package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	r := checkResponse("foo")
	assert.Equal(t, 0, r, "Expecting 0 response code")
}
