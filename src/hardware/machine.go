package hardware

import "fmt"

type nybble uint8

// The actual 4RCH virtual machine
type Machine struct {
	Accumulator  nybble
	RAM          [16][16]nybble
	RAMListeners map[byte]RAMListener
	Display      *Screen
	Cart         *Cart
	InsPointer   byte
	isRunning    bool
	ticks        int8
}

func NewMachine(display *Screen) *Machine {
	vm := Machine{
		Accumulator:  0,
		RAM:          [16][16]nybble{},
		RAMListeners: map[byte]RAMListener{},
		Display:      display,
		Cart:         nil,
		InsPointer:   0,
		isRunning:    false,
		ticks:        0,
	}
	vm.AddRAMListener(PERIPHERAL_PAGE, FPG_SCR_VAL, onScreenWritten)
	return &vm
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
		if vm.ticks > 0 {
			vm.Display.DrawBMP(BMP_NO_CART)
		} else {
			vm.Display.Clear()
		}
		vm.ticks++
		return
	}
}

func (vm *Machine) DrawDisplay() {
	vm.Display.DrawVRAM()
}
