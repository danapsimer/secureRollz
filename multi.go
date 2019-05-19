package secureRollz

import "fmt"

type multiRoll struct {
	baseRoll
	results []Roll
	total   RollValue
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
