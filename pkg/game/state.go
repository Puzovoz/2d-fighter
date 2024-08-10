package game

import (
	"github.com/puzovoz/2d-fighter/pkg/character"
	"github.com/puzovoz/2d-fighter/pkg/controls"
	"github.com/puzovoz/2d-fighter/pkg/render"

	"github.com/veandco/go-sdl2/sdl"
)

type state struct {
	characters    []*character.Character
	keyboardState []uint8
	running       bool
}

type tickState struct {
	characterActions actionsByCharacter
}

type playerActions []controls.PlayerAction
type actionsByCharacter map[*character.Character]playerActions

func NewState() *state {
	return &state{
		characters: []*character.Character{
			{
				KeyboardControls: controls.GetDefaultKeyboardControls(),
				Hitbox:           &sdl.Rect{X: 100, Y: 100, W: 20, H: 20},
			},
		},
		keyboardState: sdl.GetKeyboardState(),
		running:       true,
	}
}

func ProcessEvents(gs *state) *tickState {
	var ts = &tickState{characterActions: make(actionsByCharacter)}

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

func UpdateGameState(gs *state, ts *tickState) {
	for _, chara := range gs.characters {
		chara.ProcessActions(ts.characterActions[chara])
	}
}

func GenerateRenderContext(gs *state) *render.Context {
	characterRenders := make([]render.CharacterRenderContext, 0)

	for _, chara := range gs.characters {
		characterRenders = append(characterRenders, render.CharacterRenderContext{
			Box:   chara.Hitbox,
			Color: &sdl.Color{R: 255, G: 0, B: 0, A: 255},
		})
	}

	return &render.Context{
		CharacterRenders: characterRenders,
	}
}
