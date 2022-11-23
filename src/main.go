package main

import (
	"fmt"

	"github.com/Sam36502/Arch40/src/hardware"
	"github.com/Sam36502/Arch40/src/util"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	OPTIONS_FILE = "options.json"
)

func main() {
	err := util.LoadOptions(OPTIONS_FILE)
	if err != nil {
		fmt.Printf("[ERROR]: Failed to load options file '%s'\nPlease make sure the options file is in the same directory as the Arch40 binary.\n", OPTIONS_FILE)
		return
	}

	windowSize := 128 * util.GlobalOptions.PixelSize
	rl.InitWindow(windowSize, windowSize, "Arch40")

	//// DEBUG ////
	pp4 := hardware.NewPP4()
	// Clear Background
	pp4.SetArguments([4]uint8{
		0b0111,
		0b0000,
		0b0000,
		0b0000,
	})
	pp4.Tick(hardware.EXP_PP4_INS_ALTR)
	// Load Test Palette
	pp4.SetArguments([4]uint8{
		0b0000,
		0b0010,
		0b0100,
		0b1000,
	})
	pp4.Tick(hardware.EXP_PP4_INS_STPI)
	// Load Test Graphic
	pp4.SetArguments([4]uint8{
		0b0000,
		0b0000,
		0b1000,
		0b0001,
	})
	pp4.Tick(hardware.EXP_PP4_INS_STTI)
	pp4.SetArguments([4]uint8{
		0b0000,
		0b0100,
		0b1001,
		0b1001,
	})
	pp4.Tick(hardware.EXP_PP4_INS_STTI)
	pp4.SetArguments([4]uint8{
		0b0000,
		0b0110,
		0b1111,
		0b1111,
	})
	pp4.Tick(hardware.EXP_PP4_INS_STTI)
	pp4.SetArguments([4]uint8{
		0b0000,
		0b0000,
		0b0010,
		0b0010,
	})
	pp4.Tick(hardware.EXP_PP4_INS_DRTL)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		pp4.DrawScreen()

		rl.EndDrawing()
	}

	util.LoadOptions(OPTIONS_FILE)
	rl.CloseWindow()
}
