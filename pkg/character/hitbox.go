package character

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Hitbox is an interactable area primarily used to calculate collisions with other entities.
type Hitbox []sdl.Rect

// Boundaries is a collection of different types of hitboxes of an entity.
type Boundaries struct {
	Attackbox Hitbox
	Hurtbox   Hitbox
}

// Move all boxes of a hitbox in the direction of a vector moveTo.
func (hb Hitbox) TranslateTo(moveTo sdl.Point) Hitbox {
	newHitbox := make(Hitbox, len(hb))

	for _, box := range hb {
		newHitbox = append(newHitbox, sdl.Rect{
			X: box.X + moveTo.X,
			Y: box.Y + moveTo.Y,
			H: box.H,
			W: box.W,
		})
	}

	return newHitbox
}

// Horizontally flip all boxes of a hitbox over a point.
func (hb Hitbox) Flip(over sdl.Point) Hitbox {
	newHitbox := make(Hitbox, len(hb))

	for _, box := range hb {
		// Distance from left border of the box to the target point.
		distanceToPoint := over.X - box.X
		// Move twice the distance to the point to get to the other side
		// and subtract the width to "flip" the rectangle along its left border to make it face the point.
		distanceToMove := distanceToPoint*2 - box.W

		newHitbox = append(newHitbox, sdl.Rect{
			X: box.X + distanceToMove,
			Y: box.Y,
			H: box.H,
			W: box.W,
		})
	}

	return newHitbox
}

// Move all hitboxes of an entity in the direction of a vector moveTo.
func (b Boundaries) TranslateTo(moveTo sdl.Point) Boundaries {
	return Boundaries{
		Attackbox: b.Attackbox.TranslateTo(moveTo),
		Hurtbox:   b.Hurtbox.TranslateTo(moveTo),
	}
}

// Horizontally flip all hitboxes of an entity over a point.
func (b Boundaries) Flip(over sdl.Point) Boundaries {
	return Boundaries{
		Attackbox: b.Attackbox.Flip(over),
		Hurtbox:   b.Hurtbox.Flip(over),
	}
}
