package secureRollz

import "fmt"

type parenRoll struct {
	baseRoll
	roll Roll
}

func (pr parenRoll) Results() []Roll {
	return pr.roll.Results()
}

func (pr parenRoll) Total() RollValue {
	return pr.roll.Total()
}

func (pr parenRoll) String() string {
	return fmt.Sprintf("(%s){%s}", pr.Source().String(), pr.roll.String())
}

type parenRoller struct {
	inside Roller
}

func (pr parenRoller) String() string {
	return fmt.Sprintf("(%s)", pr.inside.String())
}

func (pr parenRoller) Roll() Roll {
	return parenRoll{baseRoll{pr.inside}, pr.inside.Roll()}
}

func ParenRoller(roller Roller) Roller {
	return parenRoller{roller}
}
