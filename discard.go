package secureRollz

import (
	"bytes"
	"fmt"
	"sort"
)

type discardRoll struct {
	baseRoll
	results []Roll
	discarded []Roll
	total   RollValue
}

func (dr discardRoll) String() string {
	sw := bytes.NewBuffer(make([]byte,0,256))
	_, err := fmt.Fprintf(sw,"%s{total:%d,results:[", dr.Source().String(), dr.total)
	if err != nil {
		sw.WriteString("Error creating string: ")
		sw.WriteString(err.Error())
		return sw.String()
	}
	for i, r := range dr.results {
		if i > 0 {
			sw.WriteString(",")
		}
		sw.WriteString(r.String())
	}
	sw.WriteString("],discarded:[")
	for i, r := range dr.discarded {
		if i > 0 {
			sw.WriteString(",")
		}
		sw.WriteString(r.String())
	}
	sw.WriteString("]}")
	return sw.String()
}

func (dr discardRoll) Results() []Roll {
	r := make([]Roll, len(dr.results))
	copy(r, dr.results)
	return r
}

func (dr discardRoll) Total() RollValue {
	return dr.total
}

type discardRoller struct {
	n        int
	lowest   bool
	original Roller
}

func (dr discardRoller) String() string {
	direction := ">"
	if dr.lowest {
		direction = ""
	}
	return fmt.Sprintf("%sD%s%d", dr.original.String(), direction, dr.n)
}

func (dr discardRoller) Roll() Roll {
	roll := dr.original.Roll()
	results := roll.Results()
	sort.Slice(results, func(i, j int) bool {
		if dr.lowest {
			return results[i].Total() < results[j].Total()
		}
		return results[i].Total() > results[j].Total()
	})
	newResults := results[dr.n:]
	discarded := results[:dr.n]
	total := RollValue(0)
	for _, nr := range newResults {
		total += nr.Total()
	}
	return discardRoll{baseRoll{dr}, newResults, discarded, total}
}

func DiscardRoller(n int, lowest bool, roller Roller) Roller {
	return discardRoller{n, lowest, roller}
}
