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

	assert.Equal(t, routerDecisions["B"], nodes["B"])
	assert.Equal(t, routerDecisions["C"], nodes["C"])
	assert.Equal(t, routerDecisions["D"], nodes["E"])
	assert.Equal(t, routerDecisions["E"], nodes["E"])
}
