package game

import (
	"bufio"
	"cube/internal/initialization"
	"cube/internal/number"
	"cube/internal/paint"
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

	paint.DrawDice(&randomNumber, writer)

	if string(rune(*enteredNumber)) == strconv.Itoa(int(randomNumber)) {
		writer.Write([]byte("You win!\n"))
		writer.Flush()
		writer.Reset(writer)
		return
	}

	writer.Write([]byte("You lost!\n"))
	writer.Flush()
	writer.Reset(writer)

	writer.Write([]byte("The correct number: " + strconv.Itoa(int(randomNumber)) + "\n"))
	writer.Flush()
	writer.Reset(writer)

}
