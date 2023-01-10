package hardware

import "fmt"

type nybble uint8

// The actual 4RCH virtual machine
type Machine struct {
	Accumulator  nybble
	RAM          [16][16]nybble
	RAMListeners map[byte]RAMListener
	Cart         *Cart
	InsPointer   byte
	peripherals  []Peripheral
	isRunning    bool
	ticks        int8
}

func NewMachine() *Machine {
	vm := Machine{
		Accumulator:  0,
		RAM:          [16][16]nybble{},
		RAMListeners: map[byte]RAMListener{},
		Cart:         nil,
		InsPointer:   0,
		isRunning:    false,
		ticks:        0,
	}
	return &vm
}

func (vm *Machine) PlugIn(prph Peripheral) {
	vm.peripherals = append(vm.peripherals, prph)
	vm.AddRAMListener(prph.GetListener(vm))
}

// Tries to load a cartridge from a file
func (vm *Machine) LoadCartridgeFile(filename string) {
	c, err := LoadCartFromFile(filename)
	if err != nil {
		fmt.Printf("[ERROR]: Failed to load cartridge '%s': %v\n", filename, err)
	} else {
		vm.Cart = c
	}
}

// Tells the VM to perform one action
func (vm *Machine) Tick() {
	if vm.Cart == nil {
		return
	}
}
