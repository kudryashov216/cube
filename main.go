package main

import (
	"bufio"
	"cube/internal/game"
	"math/rand"
	"os"
	"sync"
	"time"
)

/*test*/

func main() {

	var gameToStart bool
	var eneterednumber byte
	var wg sync.WaitGroup

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
				exitInGame()
			case 'y':
				writer.Write([]byte("\nGame beginning!\n"))
				writer.Flush()
				writer.Reset(writer)
				gameToStart = true
			default:
				writer.Write([]byte("\nYou put incorrect response\n"))
				writer.Flush()
				writer.Reset(writer)
				exitInGame()
			}

		}

		game.StartGame(&wg, &eneterednumber, randomaizer, writer, reader)

	}

}

func exitInGame() {

	os.Exit(0)

}
