package secureRollz_test

import (
	"github.com/danapsimer/secureRollz"
	"github.com/danapsimer/secureRollz/rolltest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDieRoller(t *testing.T) {
	roller := secureRollz.DieRoller(6)
	population := rolltest.RollerTest(t, roller, 0, 1, 6, "d6")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 3.5, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 1.7, stddev, 0.1)
	}
}

func BenchmarkDieRoller(b *testing.B) {
	roller := secureRollz.DieRoller(6)
	rolltest.RollerBenchmark(b, roller)
}

