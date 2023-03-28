package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dijkstraRouteAlgorithm(t *testing.T) {
	testLock.Lock()
	defer testLock.Unlock()

	setupTestData()
	dijkstraRouteAlgorithm()

	assert.Equal(t, routerDecision["B"], nodes["B"].Address)
	assert.Equal(t, routerDecision["C"], nodes["B"].Address)
	assert.Equal(t, routerDecision["D"], nodes["B"].Address)
	assert.Equal(t, routerDecision["E"], nodes["E"].Address)
}
