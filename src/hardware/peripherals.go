package hardware

type Peripheral interface {
	Tick(vm *Machine)
	GetListener(vm *Machine) ([]byte, RAMListener)
}

type RAMListener func(val nybble)

const (
	PERIPHERAL_PAGE = 0xF
)

// F-page Addresses
const (
	FPG_P1_DPAD = 0x0 // Player 1 Direction-Pad
	FPG_P1_BTNS = 0x1 // Player 1 Buttons
	FPG_P2_DPAD = 0x2 // Player 2 Direction-Pad
	FPG_P2_BTNS = 0x3 // Player 2 Buttons

	FPG_SCR_X   = 0x4 // Screen X Coord
	FPG_SCR_Y   = 0x5 // Screen Y Coord
	FPG_SCR_VAL = 0x6 // Screen Pixel Value
	FPG_SCR_OPT = 0x7 // Screen Options

	FPG_SND_VOL = 0x8 // Sound Card Volume
	FPG_SND_PTC = 0x9 // Sound Card Pitch
	FPG_SND_OPT = 0xA // Sound Card reserved

	FPG_RAND = 0xB // Pseudo-Random Number

	FPG_DSK_H   = 0xC // High-nyble of disk address   \
	FPG_DSK_M   = 0xD // Middle-nyble of disk address  } 12-bit Address
	FPG_DSK_L   = 0xE // Low-nyble of disk address    /
	FPG_DSK_VAL = 0xF // Value of the selected disk nyble
)

func (vm *Machine) AddRAMListener(trigAddrs []byte, lnr RAMListener) {
	for _, addr := range trigAddrs {
		vm.RAMListeners[addr] = lnr
	}
}
