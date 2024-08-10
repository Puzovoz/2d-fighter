package character

import (
	"github.com/puzovoz/2d-fighter/pkg/controls"

	"github.com/veandco/go-sdl2/sdl"
)

type Character struct {
	KeyboardControls controls.KeyboardControls
	Hitbox           *sdl.Rect
}

func (c *Character) ProcessActions(actions []controls.PlayerAction) {
	for _, action := range actions {
		switch action {
		case controls.MOVE_LEFT:
			c.Hitbox.X -= 3
		case controls.MOVE_RIGHT:
			c.Hitbox.X += 3
		}
	}
}
