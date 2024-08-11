package character

import (
	"sort"

	"github.com/veandco/go-sdl2/sdl"
)

type MoveState uint8

const (
	STARTUP  MoveState = iota
	ACTIVE   MoveState = iota
	RECOVERY MoveState = iota
)

type Move struct {
	StartupFrameCount  uint16
	ActiveFrameCount   uint16
	RecoveryFrameCount uint16

	BoundariesProgression      []Boundaries
	BoundariesFrameProgression []uint16
}

func GetPunch() *Move {
	punch := &Move{
		StartupFrameCount:  20,
		ActiveFrameCount:   20,
		RecoveryFrameCount: 20,

		BoundariesProgression: []Boundaries{
			{Attackbox: nil, Hurtbox: Hitbox{sdl.Rect{X: 0, Y: 0, W: 15, H: 20}}},
			{Attackbox: Hitbox{sdl.Rect{X: 20, Y: 10, W: 8, H: 3}}, Hurtbox: Hitbox{sdl.Rect{X: 0, Y: 0, W: 20, H: 20}}},
			{Attackbox: nil, Hurtbox: Hitbox{sdl.Rect{X: 5, Y: 0, W: 20, H: 20}}},
		},
		BoundariesFrameProgression: []uint16{21, 41},
	}

	return punch
}

func (m *Move) BoundariesFromFrame(frame uint16) Boundaries {
	progressionLength := len(m.BoundariesFrameProgression)
	currentIndex := sort.Search(progressionLength, func(i int) bool {
		return frame <= m.BoundariesFrameProgression[i]
	})

	return m.BoundariesProgression[currentIndex]
}
