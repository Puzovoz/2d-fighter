package render

import "github.com/veandco/go-sdl2/sdl"

// Manager is responsible for visually representing the state of the game.
type Manager interface {
	Init()
	Cleanup()
	Loop(ch chan Context)
}

// Visual representation of a character on the game screen.
type CharacterRenderContext struct {
	Boxes []sdl.Rect
	Color sdl.Color
}

// Context of the game state required to render a visual representation of the game.
type Context struct {
	CharacterRenders []CharacterRenderContext
}
