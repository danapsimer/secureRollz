package secureRollz_test

import (
	"github.com/danapsimer/secureRollz"
	"github.com/danapsimer/secureRollz/rolltest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiRoller(t *testing.T) {
	baseRoller := secureRollz.DieRoller(6)
	roller := secureRollz.MultiRoller(10, baseRoller)
	population := rolltest.RollerTest(t, roller, 10, 10, 60, "10d6")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 35, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 5.4, stddev, 0.1)
	}
}

func BenchmarkMultiRoller(b *testing.B) {
	roller := secureRollz.MultiRoller(10,secureRollz.DieRoller(6))
	rolltest.RollerBenchmark(b, roller)
}
