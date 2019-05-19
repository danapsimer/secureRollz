package secureRollz_test

import (
	"github.com/stretchr/testify/assert"
	"secureRollz"
	"testing"
)

func TestModifiedRoller(t *testing.T) {
	roller := secureRollz.ModifiedRoller(7, secureRollz.DieRoller(6))
	population := testRoller(t, roller, 1, 8, 13, "d6+7")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 10.5, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 1.7, stddev, 0.1)
	}
}

func TestModifiedRollerNegative(t *testing.T) {
	roller := secureRollz.ModifiedRoller(-3, secureRollz.DieRoller(6))
	population := testRoller(t, roller, 1, -2, 3, "d6-3")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 0.5, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 1.7, stddev, 0.1)
	}
}

func BenchmarkModifiedRoller(b *testing.B) {
	roller := secureRollz.ModifiedRoller(7, secureRollz.DieRoller(6))
	rollerBenchmark(b, roller)
}
