package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunc(t *testing.T) {
	assert.Equal(t, "Hello World!", GetMessage())
}
