package controls

import input "github.com/quasilyte/ebitengine-input"

type Action input.Action

const (
	MoveLeft Action = iota
	MoveRight
	Run
	Jump
)
