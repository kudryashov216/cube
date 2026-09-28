package menu

import "context"

type Menu interface {
	PaintInitialMenu(currentCommand uint8)
	UpdateCurrentCommand(currentCommand *uint8)
	Execute(cancel context.CancelFunc, gameToStart *bool, currentCommand uint8)
}
