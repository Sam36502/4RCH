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
