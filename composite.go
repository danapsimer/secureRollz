package secureRollz

import (
	"bytes"
)

type compositeRoll struct {
	baseRoll
	results []Roll
	total RollValue
}

func (cr compositeRoll) Results() []Roll {
	cp := make([]Roll,len(cr.results))
	copy(cp, cr.results)
	return cp
}

func (cr compositeRoll) Total() RollValue {
	return cr.total
}

type compositeRoller struct {
	components []Roller
}

func (cr compositeRoller) String() string {
	buffer := bytes.NewBuffer(make([]byte,0,100))
	for idx, result := range cr.components {
		if idx > 0 {
			buffer.WriteString("+")
		}
		buffer.WriteString(result.String())
	}
	return buffer.String()
}

func (crer compositeRoller) Roll() Roll {
	rolls := make([]Roll, len(crer.components))
	total := RollValue(0)
	for idx, component := range crer.components {
		rolls[idx] = component.Roll()
		total += rolls[idx].Total()
	}
	return compositeRoll{baseRoll{crer}, rolls, total}
}

func CompositeRoller(components []Roller) Roller {
	return compositeRoller{components}
}
