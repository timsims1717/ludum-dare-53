package state

import (
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/ludum-dare-53/internal/systems"
	"timsims1717/ludum-dare-53/pkg/debug"
	"timsims1717/ludum-dare-53/pkg/state"
)

type menuState struct {
	*state.AbstractState
}

func (s *menuState) Unload() {

}

func (s *menuState) Load(done chan struct{}) {
	done <- struct{}{}
}

func (s *menuState) Update(win *pixelgl.Window) {
	debug.AddText("Menu State")
	systems.ParentSystem()
	systems.ObjectSystem()
}

func (s *menuState) Draw(win *pixelgl.Window) {

}

func (s *menuState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}
