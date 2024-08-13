package game

import (
	"github.com/puzovoz/2d-fighter/pkg/render"

	"github.com/veandco/go-sdl2/sdl"
)

const LOOP_TICKS_PER_SECOND = 60
const DESIRED_DELTA_MS = 1000 / LOOP_TICKS_PER_SECOND

type loop struct {
	gameState     matchState
	renderManager render.Manager
	renderChannel chan render.Context
}

func NewLoop(rm render.Manager) loop {
	return loop{
		gameState:     NewMatch(),
		renderManager: rm,
		renderChannel: make(chan render.Context),
	}
}

func (l *loop) Run() {
	go l.renderManager.Loop(l.renderChannel)
	defer close(l.renderChannel)

	for l.gameState.running {
		// Update the SDL events.
		sdl.PumpEvents()

		// Read the new events and modify the game state accordingly.
		tickState := ProcessEvents(&l.gameState)
		UpdateGameState(l.gameState, tickState)

		// Update the game screen with new information.
		renderContext := GenerateRenderContext(l.gameState)
		sendToRender(renderContext, l.renderChannel)

		sdl.Delay(DESIRED_DELTA_MS)
	}
}

func sendToRender(ctx render.Context, renderChannel chan render.Context) {
	// Try to send the game state to render but not block the loop
	select {
	case renderChannel <- ctx:
	default:
	}
}
