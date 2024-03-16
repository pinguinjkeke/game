package component

import (
	"game/controls"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/yohamta/donburi"
)

type ControlsHandlerData struct {
	handler *input.Handler
}

func (h *ControlsHandlerData) ActionIsPressed(action controls.Action) bool {
	return h.handler.ActionIsPressed(input.Action(action))
}

func (h *ControlsHandlerData) ActionIsJustPressed(action controls.Action) bool {
	return h.handler.ActionIsJustPressed(input.Action(action))
}

func (h *ControlsHandlerData) ActionIsJustReleased(action controls.Action) bool {
	return h.handler.ActionIsJustReleased(input.Action(action))
}

var ControlsHandler = donburi.NewComponentType[ControlsHandlerData]()
