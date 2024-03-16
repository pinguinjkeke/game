package factory

import (
	"game/archetype"
	"game/component"
	"game/controls"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateControls(ecs *ecs.ECS) *donburi.Entry {
	controlsEntry := archetype.Controls.Spawn(ecs)
	controlsInput := input.System{}
	controlsInput.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	keymap := input.Keymap{
		input.Action(controls.MoveLeft):  {input.KeyGamepadLeft, input.KeyLeft},
		input.Action(controls.MoveRight): {input.KeyGamepadRight, input.KeyRight},
		input.Action(controls.Run):       {input.KeyGamepadR2, input.KeyShift},
		input.Action(controls.Jump):      {input.KeyGamepadX, input.KeySpace},
	}

	component.Controls.SetValue(controlsEntry, component.NewControlsData(controlsInput, keymap))

	return controlsEntry
}
