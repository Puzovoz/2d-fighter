package render

import "github.com/veandco/go-sdl2/sdl"

type Manager interface {
	Init()
	Cleanup()
	Render(*Context)
	Loop(ch chan *Context)
}

type CharacterRenderContext struct {
	Boxes []sdl.Rect
	Color sdl.Color
}

type Context struct {
	CharacterRenders []CharacterRenderContext
}
