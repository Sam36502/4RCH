/*
 *
 *		The PP4 Graphics Card implementation
 *
 *		Roughly based on the NES PPU chip.
 *
 */
package hardware

import (
	"image/color"

	"github.com/Sam36502/Arch40/src/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PP4GraphicsCard struct {
	argBuf      [4]uint8
	tileMem     [16]Tile
	paletteMem  [16]Palette
	frameBuffer [256][256]color.RGBA
	background  color.RGBA
}

type Tile [16]uint8 // 16 sections of 4 pixels (2b each)

type Palette [4]color.RGBA

const EXP_PP4_ID = "pp4"

// Graphics Card Instructions
const (
	EXP_PP4_INS_STTI = 0 // Store Tile Info
	EXP_PP4_INS_STPI = 1 // Store Palette Info
	EXP_PP4_INS_DRTL = 2 // Draw Tile
	EXP_PP4_INS_ALTR = 3 // Alter image
)

// Compile-time implementation check
var _ Expansion = (*PP4GraphicsCard)(nil)

func NewPP4() *PP4GraphicsCard {
	pp4 := &PP4GraphicsCard{}
	return pp4
}

func (crd *PP4GraphicsCard) Tick(instruction uint8) {
	switch instruction {

	case EXP_PP4_INS_STTI:
		tid := crd.argBuf[0]  // nyble 0 is tile ID to set
		sid := crd.argBuf[1]  // nyble 1 is tile section ID
		dat := crd.argBuf[2:] // nyble 2-3 is section data
		crd.tileMem[tid][sid] = argsToTile(dat)

	case EXP_PP4_INS_STPI:
		pid := crd.argBuf[0]  // nyble 0 is palette ID to set
		tbc := crd.argBuf[1:] // nyble 1-3 is 3-bit colour info
		crd.paletteMem[pid] = argsToPalette(tbc)

	case EXP_PP4_INS_DRTL:
		tid := crd.argBuf[0] // nyble 0 is tile ID to set
		pid := crd.argBuf[1] // nyble 1 is palette ID to set
		gx := crd.argBuf[2]  // nyble 2 is the grid x coord
		gy := crd.argBuf[3]  // nyble 3 is the grid y coord

		// Write tile to framebuffer
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				pixel := crd.tileMem[tid].getPixel(x, y)
				if pixel > 0 {
					crd.frameBuffer[int(gy)*8+y][int(gx)*8+x] = crd.paletteMem[pid][pixel]
				}
			}
		}

	case EXP_PP4_INS_ALTR:
		// TODO: More than just set background
		clr := crd.argBuf[0]
		crd.background = color.RGBA{
			R: (clr >> 2) % 2 * 255,
			G: (clr >> 1) % 2 * 255,
			B: clr % 2 * 255,
			A: 255,
		}

	}
}

func (crd *PP4GraphicsCard) SetArguments(args [4]uint8) {
	crd.argBuf = args
}

func (crd *PP4GraphicsCard) DrawScreen() {
	rl.ClearBackground(crd.background)
	for y := 0; y < 256; y++ {
		for x := 0; x < 256; x++ {
			ps := util.GlobalOptions.PixelSize
			rl.DrawRectangle(int32(x)*ps, int32(y)*ps, ps, ps, crd.frameBuffer[y][x])
		}
	}
}

// Takes a tile and internal x/y coords and returns the pixel value
func (t Tile) getPixel(x, y int) uint8 {
	sid := (y*8 + x) / 4
	return (t[sid] >> (x % 4 * 2)) % 4
}

// Takes the tile arguments and inserts the bits in the correct place in the tile
func argsToTile(bits []uint8) uint8 {
	s := uint8(0)
	for x := 0; x < 4; x++ {
		hi := bits[0] >> x % 2
		lo := bits[1] >> x % 2
		pixel := (hi << 1) | lo
		s |= pixel << (x * 2)
	}
	return s
}

// Takes the palette arguments and converts them into TBCs
func argsToPalette(bits []uint8) Palette {
	var p Palette
	for x := 0; x < 4; x++ {
		tbc := color.RGBA{
			R: (bits[0] >> x) % 2 * 255,
			G: (bits[1] >> x) % 2 * 255,
			B: (bits[2] >> x) % 2 * 255,
			A: 255,
		}
		p[x] = tbc
	}
	return p
}
