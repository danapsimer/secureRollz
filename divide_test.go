package secureRollz_test

import (
	"github.com/stretchr/testify/assert"
	"secureRollz"
	"testing"
)

func TestDivideRoller(t *testing.T) {
	roller := secureRollz.DivideRoller(2, secureRollz.DieRoller(20))
	population := testRoller(t, roller, 1, 0, 10, "d20/2")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 5, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 3, stddev, 0.1)
	}
}

func TestDivideRoller2(t *testing.T) {
	roller := secureRollz.DivideRoller(3, secureRollz.MultiRoller(3, secureRollz.DieRoller(6)))
	population := testRoller(t, roller, 1, 1, 6, "3d6/3")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 3, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 1, stddev, 0.1)
	}
}

func BenchmarkDivideRoller(b *testing.B) {
	roller := secureRollz.DivideRoller(2, secureRollz.DieRoller(20))
	rollerBenchmark(b, roller)
}
