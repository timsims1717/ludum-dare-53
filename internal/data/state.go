package data

import (
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/object"
	"timsims1717/ludum-dare-53/pkg/typeface"
	"timsims1717/ludum-dare-53/pkg/viewport"
)

var (
	TetrisViewport  *viewport.ViewPort
	FactoryViewport *viewport.ViewPort
	StickyViewport  *viewport.ViewPort

	SBLabels *typeface.Text
	SBScores *typeface.Text

	StickyText *typeface.Text

	Paused     bool
	PauseMenu  bool
	StickyOpen bool
	TinyNote   *img.Sprite
)

type StickyMsg struct {
	Message string
	Offset  pixel.Vec
}

var (
	PauseMsg = &StickyMsg{
		Message: "Paused",
		Offset:  pixel.V(-6., 293.),
	}
	CreditsMsg = &StickyMsg{
		Message: "Credits\n\nEverything by Ben and Tim Sims\nGet back to work.",
		Offset:  pixel.V(40., 115.),
	}
)

var (
	RestartButton *Button
	PauseButton   *Button

	RestartButSprs []*img.Sprite
	PauseButSprs   []*img.Sprite
)

type Button struct {
	Click   func()
	Sprites *img.Sprite
	Object  *object.Object
}
