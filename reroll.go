package secureRollz

import (
	"fmt"
	"sort"
)

type reRollRoll struct {
	baseRoll
	results []Roll
	total   RollValue
}

func (rrr reRollRoll) Results() []Roll {
	results := make([]Roll, len(rrr.results))
	copy(results, rrr.results)
	return results
}

func (rrr reRollRoll) Total() RollValue {
	return rrr.total
}

type reRollRoller struct {
	original Roller
	reRollN  int
	value    RollValue
	lessThan bool
}

func (rrr reRollRoller) String() string {
	lh := ">"
	if rrr.lessThan {
		lh = "<"
	}
	if rrr.reRollN > 1 {
		return fmt.Sprintf("%sr%d%s%d", rrr.original.String(), rrr.reRollN, lh, rrr.value)
	} else {
		return fmt.Sprintf("%sr%s%d", rrr.original.String(), lh, rrr.value)
	}
}

func (rrr reRollRoller) Roll() Roll {
	sourceRoll := rrr.original.Roll()
	rolls := sourceRoll.Results()
	total := sourceRoll.Total()
	if rrr.lessThan {
		sort.Slice(rolls, func(i, j int) bool {
			return rolls[i].Total() < rolls[j].Total()
		})
		for i := 0; i < rrr.reRollN && rolls[i].Total() <= rrr.value; i++ {
			newRoll := rolls[i].Source().Roll()
			total -= rolls[i].Total()
			rolls[i] = newRoll
			total += newRoll.Total()
		}
	} else {
		sort.Slice(rolls, func(i, j int) bool {
			return rolls[i].Total() > rolls[j].Total()
		})
		for i := 0; i < rrr.reRollN && rolls[i].Total() >= rrr.value; i++ {
			newRoll := rolls[i].Source().Roll()
			total -= rolls[i].Total()
			rolls[i] = newRoll
			total += newRoll.Total()
		}
	}
	return reRollRoll{baseRoll{rrr}, rolls, total}
}

func ReRollRoller(reRollN int, value RollValue, lessThan bool, roller Roller) Roller {
	return reRollRoller{roller, reRollN, value, lessThan}
}
