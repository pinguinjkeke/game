package component

import (
	input "github.com/quasilyte/ebitengine-input"
	"github.com/yohamta/donburi"
)

var controlsId uint8 = 0

type ControlsData struct {
	input  input.System
	keymap input.Keymap
}

func (c *ControlsData) NewHandler() *ControlsHandlerData {
	controlsId++

	return &ControlsHandlerData{c.input.NewHandler(controlsId, c.keymap)}
}

func NewControlsData(system input.System, keymap input.Keymap) ControlsData {
	return ControlsData{system, keymap}
}

var Controls = donburi.NewComponentType[ControlsData]()
