package secureRollz

import (
	"bytes"
	"fmt"
)

type multiRoll struct {
	baseRoll
	results []Roll
	total   RollValue
}

func (mr multiRoll) String() string {
	sw := bytes.NewBuffer(make([]byte,0,256))
	sw.WriteString(mr.Source().String())
	_, err := fmt.Fprintf(sw,"%s{total:%d,results:[", mr.Source().String(), mr.total)
	if err != nil {
		sw.WriteString("Error while creating string: ")
		sw.WriteString(err.Error())
		return sw.String()
	}
	for i, r := range mr.results {
		if i > 0 {
			sw.WriteString(",")
		}
		sw.WriteString(r.String())
	}
	sw.WriteString("]}")
	return sw.String()
}

func (mr multiRoll) Results() []Roll {
	cp := make([]Roll, len(mr.results))
	copy(cp, mr.results)
	return cp
}

func (mr multiRoll) Total() RollValue {
	return mr.total
}

type multiRoller struct {
	n      int
	roller Roller
}

func (mr multiRoller) String() string {
	return fmt.Sprintf("%d%s", mr.n, mr.roller.String())
}

func (mr multiRoller) Roll() Roll {
	roll := multiRoll{baseRoll: baseRoll{mr}, results: make([]Roll, 0, mr.n), total: 0}
	for i := 0; i < mr.n; i++ {
		die := mr.roller.Roll()
		roll.results = append(roll.results, die)
		roll.total += die.Total()
	}
	return roll
}

func MultiRoller(n int, roller Roller) Roller {
	return multiRoller{n, roller}
}
