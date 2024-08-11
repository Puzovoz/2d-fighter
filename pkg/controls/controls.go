package controls

import (
	"github.com/veandco/go-sdl2/sdl"
)

type PlayerAction uint8

const (
	NO_ACTION  PlayerAction = iota
	MOVE_LEFT  PlayerAction = iota
	MOVE_RIGHT PlayerAction = iota
	BUTTON_A   PlayerAction = iota
	BUTTON_B   PlayerAction = iota
	BUTTON_C   PlayerAction = iota
)

type KeyboardControls map[sdl.Scancode]PlayerAction
type GamepadControls map[sdl.GameControllerButton]PlayerAction

func GetDefaultKeyboardControls() KeyboardControls {
	return KeyboardControls{
		sdl.SCANCODE_A: MOVE_LEFT,
		sdl.SCANCODE_D: MOVE_RIGHT,
		sdl.SCANCODE_H: BUTTON_A,
		sdl.SCANCODE_J: BUTTON_B,
		sdl.SCANCODE_K: BUTTON_C,
	}
}

func (kc KeyboardControls) Process(keyboardState []uint8) []PlayerAction {
	actions := make([]PlayerAction, 0, 5)

	for key, action := range kc {
		if pressed := keyboardState[key]; pressed > 0 {
			actions = append(actions, action)
		}
	}

	return actions
}
