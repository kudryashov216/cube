package menu

import (
	"context"
	"cube/internal/game"
	"cube/internal/output"
	"time"
)

type Menu_ struct {
	Commands map[uint8]string
}

func GetMenu() Menu_ {

	commands := make(map[uint8]string)

	commands[0] = ">> Start game\n   Exit in game\n"
	commands[1] = "   Start game\n>> Exit in game\n"

	var newMenu Menu_

	newMenu.Commands = commands

	return newMenu

}

func (m Menu_) Execute(cancel context.CancelFunc, gameToStart *bool, ch *chan uint16, currentCommand uint8) {

	defer cancel()
	defer close(*ch)

	switch currentCommand {

	case 0:
		output.Output("\nYou have \"100\" points.\n\nGame beginning!\n")
		*gameToStart = true
	case 1:
		output.Output("\nYou exit to game!\n")
		time.Sleep(time.Second * 1)
		game.ExitInGame()
	}

}
