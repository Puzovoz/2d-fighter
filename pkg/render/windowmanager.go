package render

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type WindowManager struct {
	window      *sdl.Window
	surface     *sdl.Surface
	refreshRate int32
}

func (wm *WindowManager) Init() {
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	displayIndex, err := window.GetDisplayIndex()
	if err != nil {
		panic(err)
	}

	displayMode, err := sdl.GetCurrentDisplayMode(displayIndex)
	if err != nil {
		panic(err)
	}

	wm.window = window
	wm.surface = surface
	wm.refreshRate = displayMode.RefreshRate
}

func (wm *WindowManager) Cleanup() {
	wm.window.Destroy()
}

// A loop meant to be used as a goroutine to update the game screen created with windowmanager.Init().
// Takes a channel argument that accepts context needed to render the next frame.
func (wm *WindowManager) Loop(ch chan *Context) {
	var refreshRate = uint64(wm.refreshRate)
	var desiredDelta = float32(1000) / float32(refreshRate) // Time between frames for preferred refresh rate.

	var lastFrameTime uint64 = 0
	for ctx := range ch {
		// Render the game state as soon as it is received to reduce the time between player input and visual feedback.
		wm.Render(ctx)

		// Try to achieve the target time between frames by delaying the next frame.
		now := sdl.GetTicks64()
		timeBetweenFrames := now - lastFrameTime
		if float32(timeBetweenFrames) < desiredDelta {
			time.Sleep(time.Duration(uint64(desiredDelta)-timeBetweenFrames) * time.Millisecond)
		}
		lastFrameTime = now
	}
}

func (wm *WindowManager) Render(ctx *Context) {
	wm.surface.FillRect(nil, 0)

	for _, characterRender := range ctx.CharacterRenders {
		if len(characterRender.Boxes) == 0 {
			continue
		}

		color := characterRender.Color
		pixel := sdl.MapRGBA(wm.surface.Format, color.R, color.G, color.B, color.A)
		wm.surface.FillRects(characterRender.Boxes, pixel)
	}

	wm.window.UpdateSurface()
}
