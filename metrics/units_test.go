package metrics_test

import (
	"testing"

	"github.com/JayDalton/simplegoserver/metrics"
	"github.com/stretchr/testify/assert"
)

func TestUnits(t *testing.T) {
	assert.Equal(t, "milliseconds", metrics.Milliseconds)
}
