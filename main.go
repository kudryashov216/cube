package main

import (
	"bufio"
	"cube/internal/game"
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
		scoreCounter   int64 = 100
		countThrows    uint8
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

	for {

		if !gameToStart {

			output.Output("Start the game? -answers:\n\t-y\n\t-n\n")

			response, _ := reader.ReadByte()

			switch response {
			case 'n':
				output.Output("\nYou exit to game!\n")
				time.Sleep(time.Second * 1)
				game.ExitInGame()
			case 'y':
				output.Output("\nYou have \"100\" points.\n\nGame beginning!\n")
				gameToStart = true
			default:
				output.Output("\nYou put incorrect response\n")
				game.ExitInGame()
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
