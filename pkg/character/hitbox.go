package character

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Hitbox []sdl.Rect

type Boundaries struct {
	Attackbox Hitbox
	Hurtbox   Hitbox
}

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

func (b *Boundaries) TranslateTo(moveTo sdl.Point) Boundaries {
	return Boundaries{
		Attackbox: b.Attackbox.TranslateTo(moveTo),
		Hurtbox:   b.Hurtbox.TranslateTo(moveTo),
	}
}
