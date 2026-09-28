package menu

import (
	"cube/cmd"
	"cube/internal/output"
)

func (m Menu_) PaintInitialMenu(currentCommand uint8) {

	cmd := cmd.GetConsole()
	cmd.Clear()

	switch currentCommand {
	case 0:
		output.Output(m.Commands[currentCommand])
	case 1:
		output.Output(m.Commands[currentCommand])
	}

}
