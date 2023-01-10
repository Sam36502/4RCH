package hardware

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Screen struct {
	vram  [16]uint16
	fg    rl.Color
	bg    rl.Color
	scale int32
}

func NewScreen(fg, bg rl.Color, scale int32) *Screen {
	return &Screen{
		vram:  [16]uint16{},
		fg:    fg,
		bg:    bg,
		scale: scale,
	}
}

func (s *Screen) Clear() {
	s.vram = [16]uint16{}
	rl.ClearBackground(s.bg)
}

func (s *Screen) Invert() {
	for i, row := range s.vram {
		s.vram[i] = ^row
	}
}

func (s *Screen) DrawBMP(bmp [16]uint16) {
	rl.ClearBackground(s.bg)
	var x, y int32
	for y = 0; y < 16; y++ {
		for x = 0; x < 16; x++ {
			if (bmp[y]>>(15-x))%2 == 1 {
				ps := s.scale
				rl.DrawRectangle(
					x*ps, y*ps,
					ps, ps,
					s.fg,
				)
			}
		}
	}
}

func (s *Screen) DrawVRAM() {
	s.DrawBMP(s.vram)
}

func (s *Screen) Tick(vm *Machine) {

	// Check if top opt bits are used
	opt := vm.RAM[PERIPHERAL_PAGE][FPG_SND_OPT]
	if (opt>>3)%2 == 1 {
		s.Clear()
		opt &= 0b0111
	}
	if (opt>>2)%2 == 1 {
		s.Invert()
		opt &= 0b1011
	}
	vm.RAM[PERIPHERAL_PAGE][FPG_SCR_OPT] = opt

	// Update state of the screen_value address
	x := vm.RAM[PERIPHERAL_PAGE][FPG_SCR_X]
	y := vm.RAM[PERIPHERAL_PAGE][FPG_SCR_Y]
	var val nybble = 0
	switch opt % 4 {

	case 0b00:
		val = nybble(s.vram[y] >> (8 - 1 - x) % 2)

	case 0b01:
		val = nybble(s.vram[y] >> (8 - 4 - x) % 16)

	case 0b10:
		for i := nybble(0); i < 4; i++ {
			val |= nybble(s.vram[y+i]>>(8-1-x)%2) << (4 - i)
		}

	case 0b11:
		val |= nybble(s.vram[y+0]>>(8-2-x)%4) << 2
		val |= nybble(s.vram[y+1]>>(8-2-x)%4) << 0
	}
	vm.RAM[PERIPHERAL_PAGE][FPG_SCR_VAL] = val

}

func (s *Screen) GetListener(vm *Machine) ([]byte, RAMListener) {
	return []byte{(PERIPHERAL_PAGE << 4) | FPG_SCR_VAL},
		func(val nybble) {

			x := vm.RAM[PERIPHERAL_PAGE][FPG_SCR_X]
			y := vm.RAM[PERIPHERAL_PAGE][FPG_SCR_Y]
			opt := vm.RAM[PERIPHERAL_PAGE][FPG_SCR_OPT]
			switch opt % 4 {

			case 0b00:
				s.vram[y] ^= uint16(val<<(8-1-x)) % 2

			case 0b01:
				s.vram[y] ^= uint16(val<<(8-4-x)) % 16

			case 0b10:
				for i := nybble(0); i < 4; i++ {
					s.vram[y+i] ^= (uint16(val<<(8-1-x)) % 2) << (4 - i)
				}

			case 0b11:
				s.vram[y+0] ^= (uint16(val<<(8-2-x)) % 4) << 2
				s.vram[y+1] ^= (uint16(val<<(8-2-x)) % 4) << 0
			}

		}
}

// func ErrorPopup(msg string) {
// // Draw Box
// red := color.RGBA{255, 64, 64, 255}
// darkRed := color.RGBA{200, 32, 32, 255}
// width := 300
// height := 150
// x := (util.g_options.PixelSize * 16 / 2) - width/2
// y := (util.g_options.PixelSize * 16 / 2) - height/2
// rec := rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(width), Height: float32(height)}
// rl.DrawRectangleRec(rec, red)
// rl.DrawRectangleLinesEx(rec, 5, darkRed)

// // Draw Text
// rl.DrawText(msg, int32(x+25), int32(y+25), 20, darkRed)
// rl.EndDrawing()
// time.Sleep(3 * time.Second)
// }
