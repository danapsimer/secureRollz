package secureRollz

import (
	"bytes"
	"fmt"
	"sort"
)

type reRollRoll struct {
	baseRoll
	results  []Roll
	rerolled []Roll
	total    RollValue
}

func (rrr reRollRoll) String() string {
	sw := bytes.NewBuffer(make([]byte,0,256))
	_, err := fmt.Fprintf(sw,"%s{total:%d,results:[",rrr.Source().String(), rrr.total)
	if err != nil {
		sw.WriteString("Error writing string: ")
		sw.WriteString(err.Error())
		return sw.String()
	}
	for i, r := range rrr.results {
		if i > 0 {
			sw.WriteString(",")
		}
		sw.WriteString(r.String())
	}
	sw.WriteString("],rerolled:[")
	for i, r := range rrr.rerolled {
		if i > 0 {
			sw.WriteString(",")
		}
		sw.WriteString(r.String())
	}
	sw.WriteString("]}")
	return sw.String()
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
	rerolled := make([]Roll, 0, rrr.reRollN)
	total := sourceRoll.Total()
	if rrr.lessThan {
		sort.Slice(rolls, func(i, j int) bool {
			return rolls[i].Total() < rolls[j].Total()
		})
		for i := 0; i < rrr.reRollN && rolls[i].Total() <= rrr.value; i++ {
			rerolled = append(rerolled,rolls[i])
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
			rerolled = append(rerolled,rolls[i])
			newRoll := rolls[i].Source().Roll()
			total -= rolls[i].Total()
			rolls[i] = newRoll
			total += newRoll.Total()
		}
	}
	return reRollRoll{baseRoll{rrr}, rolls, rerolled, total}
}

func ReRollRoller(reRollN int, value RollValue, lessThan bool, roller Roller) Roller {
	return reRollRoller{roller, reRollN, value, lessThan}
}
