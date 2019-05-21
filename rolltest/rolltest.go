package rolltest

import (
	"bytes"
	"fmt"
	"github.com/danapsimer/secureRollz"
	"github.com/montanaflynn/stats"
	"github.com/stretchr/testify/assert"
	"io"
	"math"
	"testing"
)

func RollerTest(t *testing.T, roller secureRollz.Roller, numResults int, min, max secureRollz.RollValue, source string) stats.Float64Data {
	return rollerTest(t, roller, numResults, min, max, source, true)
}

func RollerTestStats(t *testing.T, roller secureRollz.Roller, numResults int, min, max secureRollz.RollValue, mean, stddev float64, source string, printHisto bool) {
	population := rollerTest(t, roller, numResults, min, max, source, printHisto)
	m, err := population.Mean()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, mean, m, 1.0)
	}
	sd, err := population.StandardDeviation()
	if (assert.NoError(t, err)) {
		assert.InDelta(t, stddev, sd, 0.1)
	}
}

func rollerTest(t *testing.T, roller secureRollz.Roller, numResults int, min, max secureRollz.RollValue, source string, printHisto bool) stats.Float64Data {
	assert.Equal(t, source, roller.String())
	samples := int(max-min+1)*1000
	population := stats.Float64Data(make([]float64,samples))
	histo := make([]int, max-min+1)
	if assert.NotNil(t, roller) {
		for i := 0; i < samples; i++ {
			roll := roller.Roll()
			if assert.NotNil(t, roll) {
				results := roll.Results()
				assert.Equalf(t, numResults, len(results),"expected %d results but found %d: %+v",numResults, len(results), results)
				total := roll.Total()
				if assert.Conditionf(t, func() (bool) {
					return min <= total && total <= max
				}, "Roller returned a number out of range: expected a number between %d, and %d inclusive but got %d", min, max, total) {
					population[i] = float64(total)
					histo[total-min] += 1
				}
			}
		}
	}
	if printHisto {
		printHistogram(t, min, max, histo, 80)
	}
	return population
}

var fractionalBlocks = []rune(" \u258F\u258E\u258D\u258C\u258B\u258A\u2589\u2588")
func printHistogram(t *testing.T, minV, maxV secureRollz.RollValue, histo []int, maxWidth int) {
	w := bytes.NewBuffer(make([]byte, 0, 512))
	var max int
	for _, h := range histo {
		if h > max {
			max = h
		}
	}
	scale := float64(maxWidth) / float64(max)
	for i := range histo {
		io.WriteString(w, fmt.Sprintf(" %-8d: ", minV + secureRollz.RollValue(i)))
		blocksCalc := scale * float64(histo[i])
		blocks := int(math.Floor(blocksCalc))
		blocksFractionValue := int(math.Ceil((blocksCalc - float64(blocks)) / 0.125))
		for i = 0; i < blocks; i++ {
			io.WriteString(w, "\u2588")
		}
		io.WriteString(w, string(fractionalBlocks[blocksFractionValue]))
		io.WriteString(w, "\n")
	}
	t.Logf("\n%s", w.String())
}

var result secureRollz.Roll

func RollerBenchmark(b *testing.B, roller secureRollz.Roller) {
	var roll secureRollz.Roll
	for n := 0; n < b.N; n++ {
		roll = roller.Roll()
	}
	result = roll
}
