package character

import (
	"sort"

	"github.com/veandco/go-sdl2/sdl"
)

type MoveState uint8

const (
	STARTUP  MoveState = iota // Has not yet produced an Attackbox, counterable
	ACTIVE   MoveState = iota // Is currently producing Attackboxes, counterable
	RECOVERY MoveState = iota // Finished producing Attackboxes, punishable
)

type Move struct {
	totalFrameCount uint16

	boundariesProgression      []Boundaries // Snapshots of the Move's hitboxes that progress over its duration.
	boundariesFrameProgression []uint16     // Sorted slice of frame numbers at which the move progresses.
}

func GetPunch() Move {
	punch := Move{
		totalFrameCount: 11,

		boundariesProgression: []Boundaries{
			{Attackbox: nil, Hurtbox: Hitbox{sdl.Rect{X: -10, Y: 0, W: 15, H: 20}}},
			{Attackbox: Hitbox{sdl.Rect{X: 10, Y: 10, W: 8, H: 3}}, Hurtbox: Hitbox{sdl.Rect{X: -10, Y: 0, W: 20, H: 20}}},
			{Attackbox: nil, Hurtbox: Hitbox{sdl.Rect{X: -5, Y: 0, W: 20, H: 20}}},
		},
		boundariesFrameProgression: []uint16{4, 7},
	}

	return punch
}

// Given a frame number, calculate Character's Boundaries from the hitbox progression of the move.
func (m Move) BoundariesFromFrame(frame uint16) Boundaries {
	progressionLength := len(m.boundariesFrameProgression)

	// Given a sorted slice of frame numbers, at which Character's Boundaries should transition to the next stage,
	// find the index at which character's
	currentIndex := sort.Search(progressionLength, func(i int) bool {
		return frame <= m.boundariesFrameProgression[i]
	})

	return m.boundariesProgression[currentIndex]
}

// Calculate the total frame count of a move, during which a character regularly would not be able to act until it ends.
func (m Move) TotalFrameCount() uint16 {
	return m.totalFrameCount
}
