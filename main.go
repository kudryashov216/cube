package main

import (
	"bufio"
	"context"
	"cube/cmd"
	"cube/internal/game"
	"cube/internal/keyboardevents"
	"cube/internal/menu"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {

	var (
		gameToStart    bool
		decline        bool
		eneterednumber byte
		countThrows    uint8
		currentCommand uint8       = 0
		scoreCounter   int64       = 100
		keyboardEvent  chan uint16 = make(chan uint16)
		wg             sync.WaitGroup
	)

	randSource := rand.NewSource(time.Now().Unix())
	randomaizer := rand.New(randSource)

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	ctx, cancel := context.WithCancel(context.Background())

	for {

		if !gameToStart {

			menu := menu.GetMenu()

			menu.PaintInitialMenu(currentCommand)

			go keyboardevents.GetKeyEvent(ctx, keyboardEvent)

			for command := range keyboardEvent {

				switch command {

				case 38:
					menu.UpdateCurrentCommand(&currentCommand)
				case 40:
					menu.UpdateCurrentCommand(&currentCommand)
				case 13:
					menu.Execute(cancel, &gameToStart, &keyboardEvent, currentCommand)
				}

				menu.PaintInitialMenu(currentCommand)

			}

			cmd.GetConsole().Clear()
			game.DifficultyChoice(&countThrows, writer, reader)
			cmd.GetConsole().Clear()

		}

		game.StartGame(
			&wg,
			&eneterednumber,
			randomaizer,
			writer,
			reader,
			&scoreCounter,
			&decline,
			&countThrows)

	}

}
