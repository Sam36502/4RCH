package hardware

type Expansion interface {
	Tick(instruction uint8)
	SetArguments(args [4]uint8)
}

var AvailableExpansions = map[string]Expansion{
	// None, yet
}
