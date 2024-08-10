package render

import "github.com/veandco/go-sdl2/sdl"

type Manager interface {
	Init()
	Cleanup()
	Render(*Context)
	Loop(ch chan *Context)
}

type CharacterRenderContext struct {
	Box   *sdl.Rect
	Color *sdl.Color
}

type Context struct {
	CharacterRenders []CharacterRenderContext
}
