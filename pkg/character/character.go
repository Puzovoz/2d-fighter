package character

import (
	"github.com/puzovoz/2d-fighter/pkg/controls"
	"github.com/veandco/go-sdl2/sdl"
)

type Character struct {
	KeyboardControls controls.KeyboardLayout
	Origin           sdl.Point
	State            ActionState
	CurrentMove      *Move
	CurrentMoveFrame uint16
}

type ActionState uint8

const (
	IDLE      ActionState = iota
	ATTACKING ActionState = iota

	// Inactionable
	BLOCKSTUN ActionState = iota
	HITSTUN   ActionState = iota
	KNOCKDOWN ActionState = iota
)

// Process player input and translate it into character activity.
func (c *Character) ProcessActions(actions []controls.PlayerAction) {
	for _, action := range actions {
		switch action {
		case controls.MOVE_LEFT:
			if c.State == IDLE {
				c.Origin.X -= 3
			}
		case controls.MOVE_RIGHT:
			if c.State == IDLE {
				c.Origin.X += 3
			}
		case controls.BUTTON_A:
			move := GetPunch()
			c.StartMove(&move)
		}
	}
}

func (c *Character) StartMove(move *Move) {
	if c.State != IDLE {
		return
	}

	c.State = ATTACKING
	c.CurrentMove = move
	c.CurrentMoveFrame = 1
}

func (c *Character) GetBoundaries() Boundaries {
	var boundaries Boundaries
	switch c.State {
	case IDLE:
		boundaries = Boundaries{
			Hurtbox:   Hitbox{sdl.Rect{X: -10, Y: 0, H: 20, W: 20}},
			Attackbox: nil,
		}
	case ATTACKING:
		boundaries = c.CurrentMove.BoundariesFromFrame(c.CurrentMoveFrame)
	}

	return boundaries.TranslateTo(c.Origin)
}
