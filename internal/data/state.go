package data

import (
	"github.com/faiface/pixel"
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
)
