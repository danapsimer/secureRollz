package secureRollz

type RollValue int16

type Roll interface {
	// The roller used for the roll
	Source() Roller
	// A representation of the rolls with results.
	Results() []Roll
	// The Total Result
	Total() RollValue
}

type Roller interface {
	// A representation of the source rolls.
	String() string
	// Roll the die
	Roll() Roll
}

type RollerFunc func() Roll

func (rf RollerFunc) Roll() Roll {
	return rf()
}

type baseRoll struct {
	source Roller
}

func (r baseRoll) Source() Roller {
	return r.source
}
