package main

import (
	"bufio"
	"cube/internal/game"
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

			writer.Write([]byte("Start the game? -answers:\n\t-y\n\t-n\n"))
			writer.Flush()
			writer.Reset(writer)

			response, _ := reader.ReadByte()

			switch response {
			case 'n':
				writer.Write([]byte("\nYou exit to game!\n"))
				writer.Flush()
				writer.Reset(writer)
				time.Sleep(time.Second * 1)
				game.ExitInGame()
			case 'y':
				writer.Write([]byte("You have \"100\" points."))
				writer.Write([]byte("\nGame beginning!\n"))
				writer.Flush()
				writer.Reset(writer)
				gameToStart = true
			default:
				writer.Write([]byte("\nYou put incorrect response\n"))
				writer.Flush()
				writer.Reset(writer)
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

			writer.Write([]byte("settings succed!"))
			writer.Flush()
			writer.Reset(writer)

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
