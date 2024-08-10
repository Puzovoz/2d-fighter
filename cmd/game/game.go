package main

import (
	"runtime"

	"github.com/puzovoz/2d-fighter/pkg/game"
	"github.com/puzovoz/2d-fighter/pkg/render"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/sys/windows"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	if runtime.GOOS == "windows" {
		// Regulate clock resolution for higher precision on time.Sleep() within loops
		windows.TimeBeginPeriod(1)
		defer windows.TimeEndPeriod(1)
	}

	var manager = &render.WindowManager{}
	manager.Init()
	defer manager.Cleanup()

	var gameLoop = game.NewLoop(manager)
	gameLoop.Run()
}
