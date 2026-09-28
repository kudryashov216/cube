package main

import (
	"bufio"
	"context"
	"cube/internal/game"
	"cube/internal/keyboardevents"
	"cube/internal/menu"
	"cube/internal/output"
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

	const (
		NUMBER_THROWS_EASY   uint8 = 6
		NUMBER_THROWS_MIDDLE uint8 = 3
		NUMBER_THROWS_HIGH   uint8 = 1
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

			result := game.DifficultyChoice(writer, reader)

			switch result {
			case 0:
				countThrows = NUMBER_THROWS_EASY
			case 1:
				countThrows = NUMBER_THROWS_MIDDLE
			case 2:
				countThrows = NUMBER_THROWS_HIGH
			}

			output.Output("settings succed!")

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
