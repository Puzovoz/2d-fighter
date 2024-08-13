package game

import (
	"github.com/puzovoz/2d-fighter/pkg/character"
	"github.com/puzovoz/2d-fighter/pkg/controls"
	"github.com/puzovoz/2d-fighter/pkg/render"

	"github.com/veandco/go-sdl2/sdl"
)

type matchState struct {
	characters    []*character.Character
	keyboardState []uint8
	running       bool
}

type tickState struct {
	characterActions actionsByCharacter
}

type playerActions []controls.PlayerAction
type actionsByCharacter map[*character.Character]playerActions

func NewMatch() matchState {
	return matchState{
		characters: []*character.Character{
			{
				KeyboardControls: controls.GetDefaultKeyboardControls(),
				Origin:           sdl.Point{X: 100, Y: 100},
			},
		},
		keyboardState: sdl.GetKeyboardState(),
		running:       true,
	}
}

// Read SDL events and generate the state of the current game tick.
func ProcessEvents(gs *matchState) tickState {
	var ts = tickState{characterActions: make(actionsByCharacter)}

	for _, chara := range gs.characters {
		processedActions := chara.KeyboardControls.Process(gs.keyboardState)
		ts.characterActions[chara] = processedActions
	}

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			gs.running = false
		}
	}

	return ts
}

// Based on the state of the current tick, update the state of the game accordingly.
func UpdateGameState(gs matchState, ts tickState) {
	for _, chara := range gs.characters {
		if chara.CurrentMove != nil {
			chara.CurrentMoveFrame += 1

			currentMove := chara.CurrentMove
			if chara.CurrentMoveFrame > currentMove.TotalFrameCount() {
				chara.CurrentMove = nil
				chara.State = character.IDLE
			}
		}

		chara.ProcessActions(ts.characterActions[chara])
	}
}

// Produce visual state info for render.Manager to update the game screen with.
func GenerateRenderContext(gs matchState) render.Context {
	characterRenders := make([]render.CharacterRenderContext, 0)

	for _, chara := range gs.characters {
		boundaries := chara.GetBoundaries()
		if boundaries.Attackbox != nil {
			characterRenders = append(characterRenders, render.CharacterRenderContext{
				Boxes: boundaries.Attackbox,
				Color: sdl.Color{R: 255, G: 0, B: 0, A: 100},
			})
		}
		if boundaries.Hurtbox != nil {
			characterRenders = append(characterRenders, render.CharacterRenderContext{
				Boxes: boundaries.Hurtbox,
				Color: sdl.Color{R: 0, G: 0, B: 255, A: 100},
			})
		}
	}

	return render.Context{
		CharacterRenders: characterRenders,
	}
}
