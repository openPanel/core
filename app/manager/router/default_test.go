package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_defaultRouteAlgorithm(t *testing.T) {
	setupTestData(t)

	defaultRouteAlgorithm()

	assert.Equal(t, decisions["B"], nodes["B"])
	assert.Equal(t, decisions["C"], nodes["C"])
	assert.Equal(t, decisions["D"], nodes["E"])
	assert.Equal(t, decisions["E"], nodes["E"])
}
