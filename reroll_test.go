package secureRollz_test

import (
	"github.com/stretchr/testify/assert"
	"secureRollz"
	"testing"
)

func TestReRollRollerLower(t *testing.T) {
	roller := secureRollz.ReRollRoller(2, 3, true, secureRollz.MultiRoller(5,secureRollz.DieRoller(10)))
	population := testRoller(t, roller, 5, 5, 50, "5d10r2<3")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 32, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 5.8, stddev, 0.5)
	}
}

func TestReRollRollerHigher(t *testing.T) {
	roller := secureRollz.ReRollRoller(2, 7, false, secureRollz.MultiRoller(5,secureRollz.DieRoller(10)))
	population := testRoller(t, roller, 5, 5, 50, "5d10r2>7")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 22, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 5.8, stddev, 0.5)
	}
}

func BenchmarkReRollRoller(b *testing.B) {
	roller := secureRollz.ReRollRoller(2, 3, true, secureRollz.MultiRoller(5,secureRollz.DieRoller(10)))
	rollerBenchmark(b, roller)
}
