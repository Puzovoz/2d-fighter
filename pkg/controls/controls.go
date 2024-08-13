package controls

import (
	"github.com/veandco/go-sdl2/sdl"
)

type PlayerAction uint8

// Set of all possible player inputs.
const (
	MOVE_LEFT  PlayerAction = iota
	MOVE_RIGHT PlayerAction = iota
	BUTTON_A   PlayerAction = iota
	BUTTON_B   PlayerAction = iota
	BUTTON_C   PlayerAction = iota
)

// Control layouts for different input devices.
type KeyboardLayout map[sdl.Scancode]PlayerAction
type GamepadLayout map[sdl.GameControllerButton]PlayerAction

// Generate a control layout for keyboard mapped to default keys.
func GetDefaultKeyboardControls() KeyboardLayout {
	return KeyboardLayout{
		sdl.SCANCODE_A: MOVE_LEFT,
		sdl.SCANCODE_D: MOVE_RIGHT,
		sdl.SCANCODE_H: BUTTON_A,
		sdl.SCANCODE_J: BUTTON_B,
		sdl.SCANCODE_K: BUTTON_C,
	}
}

// Process keyboard input and translate it to player input.
func (kc KeyboardLayout) Process(keyboardState []uint8) []PlayerAction {
	actions := make([]PlayerAction, 0)

	for key, action := range kc {
		if pressed := keyboardState[key]; pressed > 0 {
			actions = append(actions, action)
		}
	}

	return actions
}
