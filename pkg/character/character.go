package character

import (
	"github.com/puzovoz/2d-fighter/pkg/controls"
	"github.com/veandco/go-sdl2/sdl"
)

type Character struct {
	KeyboardControls controls.KeyboardControls
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

func (c *Character) ProcessActions(actions []controls.PlayerAction) {
	for _, action := range actions {
		switch action {
		case controls.MOVE_LEFT:
			c.Origin.X -= 3
		case controls.MOVE_RIGHT:
			c.Origin.X += 3
		case controls.BUTTON_A:
			c.StartMove(GetPunch())
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
			Hurtbox:   Hitbox{sdl.Rect{X: 0, Y: 0, H: 20, W: 20}},
			Attackbox: nil,
		}
	case ATTACKING:
		boundaries = c.CurrentMove.BoundariesFromFrame(c.CurrentMoveFrame)
	}

	return boundaries.TranslateTo(c.Origin)
}
