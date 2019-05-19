package secureRollz_test

import (
	"github.com/stretchr/testify/assert"
	"secureRollz"
	"testing"
)

func TestMultiplyRoller(t *testing.T) {
	roller := secureRollz.MultiplyRoller(2, secureRollz.DieRoller(6))
	population := testRoller(t, roller, 1, 2, 12, "d6*2")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 7, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 3.4, stddev, 0.1)
	}
}

func BenchmarkMultiplyRoller(b *testing.B) {
	roller := secureRollz.MultiplyRoller(2, secureRollz.DieRoller(6))
	rollerBenchmark(b, roller)
}
