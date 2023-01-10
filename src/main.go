package main

import (
	"fmt"
	"os"

	hw "github.com/Sam36502/4RCH/src/hardware"
	"github.com/Sam36502/4RCH/src/util"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	OPTIONS_FILE = "options.json"
)

func main() {

	// Load Options from file
	err := util.LoadOptions(OPTIONS_FILE)
	if err != nil {
		fmt.Printf("[ERROR]: Failed to load options file '%s'\nPlease make sure the options file is in the same directory as the Arch40 binary.\n", OPTIONS_FILE)
		return
	}

	// Initialise Console window
	windowSize := 16 * util.GlobalOptions.PixelSize
	rl.InitWindow(windowSize, windowSize, "4RCH")
	if util.GlobalOptions.TargetFPS > 0 {
		rl.SetTargetFPS(util.GlobalOptions.TargetFPS)
	}

	screen := hw.NewScreen(
		util.GlobalOptions.ColourFG,
		util.GlobalOptions.ColourBG,
		util.GlobalOptions.PixelSize,
	)
	vm := hw.NewMachine(screen)

	// Try to load cartridge provided as argument if available
	if len(os.Args) > 1 {
		vm.LoadCartridgeFile(os.Args[1])
	}

	// Main loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Try to load a cartridge from dropped files
		if vm.Cart == nil && rl.IsFileDropped() {
			fs := rl.LoadDroppedFiles()
			vm.LoadCartridgeFile(fs[len(fs)-1])
		}

		vm.Tick()
		vm.DrawDisplay()

		rl.EndDrawing()
	}

	util.LoadOptions(OPTIONS_FILE)
	rl.CloseWindow()
}
