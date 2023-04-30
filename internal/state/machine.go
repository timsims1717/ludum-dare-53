package state

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/ludum-dare-53/internal/constants"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/internal/systems"
	"timsims1717/ludum-dare-53/pkg/debug"
	"timsims1717/ludum-dare-53/pkg/img"
	"timsims1717/ludum-dare-53/pkg/options"
	"timsims1717/ludum-dare-53/pkg/state"
	"timsims1717/ludum-dare-53/pkg/timing"
	"timsims1717/ludum-dare-53/pkg/viewport"
)

const (
	MenuStateKey = "menu_state"
	GameStateKey = "game_state"
)

var (
	MenuState = &menuState{}
	GameState = &gameState{}

	States = map[string]*state.AbstractState{
		MenuStateKey: state.New(MenuState, false),
		GameStateKey: state.New(GameState, false),
	}

	switchState = true
	currState   = "unknown"
	nextState   = GameStateKey
	loading     = false
	loadingDone = false
	done        = make(chan struct{})
)

func Update(win *pixelgl.Window) {
	timing.Update()
	debug.Clear()
	options.WindowUpdate(win)
	updateState()
	if loading {
		select {
		case <-done:
			loading = false
			loadingDone = true
			currState = nextState
		default:
			//LoadingState.Update(win)
		}
	} else {
		debugInput.Update(win, viewport.MainCamera.Mat)

		if debugInput.Get("fullscreen").JustPressed() {
			options.FullScreen = !options.FullScreen
		}
		if debugInput.Get("debugText").JustPressed() {
			debug.Text = !debug.Text
		}
		if debugInput.Get("debugClear").JustPressed() {
			systems.FailCondition = false
			systems.ClearBoard()
			systems.ClearFactory()
		}

		if cState, ok := States[currState]; ok {
			cState.Update(win)
		}
	}
	viewport.MainCamera.Update()
	myecs.UpdateManager()
}

func Draw(win *pixelgl.Window) {
	img.Clear()
	cState, ok1 := States[currState]
	nState, ok2 := States[nextState]
	if !ok2 {
		panic(fmt.Sprintf("state %s doesn't exist", nextState))
	}
	if loading && nState.ShowLoad || !ok1 {
		//LoadingState.Draw(win)
	} else {
		win.Clear(constants.BlackColor)
		cState.Draw(win)
	}
	debug.Draw(win)
	win.Update()
}

func updateState() {
	if !loading && (currState != nextState || switchState) {
		// uninitialize
		img.FullClear()
		if cState, ok := States[currState]; ok {
			go cState.Unload()
		}
		// initialize
		if nState, ok := States[nextState]; ok {
			go nState.Load(done)
			loading = true
			loadingDone = false
		}
		switchState = false
	}
}

func SwitchState(s string) {
	if !switchState {
		switchState = true
		nextState = s
	}
}
