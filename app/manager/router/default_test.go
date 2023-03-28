package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_defaultRouteAlgorithm(t *testing.T) {
	testLock.Lock()
	defer testLock.Unlock()
	setupTestData()

	defaultRouteAlgorithm()

	assert.Equal(t, routerDecision["B"], nodes["B"].Address)
	assert.Equal(t, routerDecision["C"], nodes["C"].Address)
	assert.Equal(t, routerDecision["D"], nodes["E"].Address)
	assert.Equal(t, routerDecision["E"], nodes["E"].Address)
}
