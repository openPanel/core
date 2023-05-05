package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dijkstraRouteAlgorithm(t *testing.T) {
	setupTestData(t)

	dijkstraRouteAlgorithm()

	assert.Equal(t, routerDecisions["B"], nodes["B"])
	assert.Equal(t, routerDecisions["C"], nodes["B"])
	assert.Equal(t, routerDecisions["D"], nodes["B"])
	assert.Equal(t, routerDecisions["E"], nodes["E"])
}
