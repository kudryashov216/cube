package game

import (
	"bufio"
	results "cube/internal/Results"
	"cube/internal/initialization"
	"cube/internal/number"
	"math/rand"
	"strconv"
	"sync"
)

func StartGame(wg *sync.WaitGroup,
	enteredNumber *byte,
	randomaizer *rand.Rand,
	writer *bufio.Writer,
	reader *bufio.Reader,
	scoreCounter *int64,
	decline *bool,
	countThorws *uint8) {

	wg.Add(1)
	go number.InputNewNumber(enteredNumber, reader, writer, wg)
	wg.Wait()

	number_, _ := strconv.Atoi(string(rune(*enteredNumber)))

	if number_ > 6 || number_ <= 0 {
		number.IncorrectNumberEntered(writer)
		return
	}

	initialization.Initialization(writer)

	randomNumber := uint8(number.GetRandomNumber(randomaizer))

	results.Output(&randomNumber, writer)

}
