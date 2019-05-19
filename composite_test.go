package secureRollz_test

import (
	"github.com/stretchr/testify/assert"
	"secureRollz"
	"testing"
)

func TestCompositeRoller(t *testing.T) {
	roller := secureRollz.CompositeRoller([]secureRollz.Roller{
		secureRollz.DieRoller(6),
		secureRollz.ModifiedRoller(7, secureRollz.DieRoller(6)),
		secureRollz.MultiRoller(10,  secureRollz.DieRoller(6)),
		secureRollz.MultiplyRoller(2, secureRollz.DieRoller(6)),
	})
	population := testRoller(t, roller, 4, 21, 91, "d6+d6+7+10d6+d6*2")
	mean, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 56, mean, 1.0)
	}
	stddev, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, 6.8, stddev, 0.1)
	}
}

func BenchmarkCompositeRoller(b *testing.B) {
	roller := secureRollz.CompositeRoller([]secureRollz.Roller{
		secureRollz.DieRoller(6),
		secureRollz.ModifiedRoller(7, secureRollz.DieRoller(6)),
		secureRollz.MultiRoller(10,  secureRollz.DieRoller(6)),
		secureRollz.MultiplyRoller(2, secureRollz.DieRoller(6)),
	})
	rollerBenchmark(b, roller);
}
