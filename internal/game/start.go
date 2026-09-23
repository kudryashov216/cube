package game

import (
	"bufio"
	"cube/internal/initialization"
	"cube/internal/number"
	"math/rand"
	"strconv"
	"sync"
)

func StartGame(wg *sync.WaitGroup, enteredNumber *byte, randomaizer *rand.Rand, writer *bufio.Writer, reader *bufio.Reader) {

	wg.Add(1)
	go number.InputNewNumber(enteredNumber, reader, writer, wg)
	wg.Wait()

	if number_, _ := strconv.Atoi(string(rune(*enteredNumber))); number_ > 6 || number_ <= 0 {
		number.IncorrectNumberEntered(writer)
		return
	}

	initialization.Initialization(writer)

	randomNumber := number.GetRandomNumber(randomaizer)

	if string(rune(*enteredNumber)) == strconv.Itoa(randomNumber) {
		writer.Write([]byte("You win!\n"))
		writer.Flush()
		writer.Reset(writer)
		return
	}

	writer.Write([]byte("You lost!\n"))
	writer.Flush()
	writer.Reset(writer)

	writer.Write([]byte("The correct number: " + strconv.Itoa(randomNumber) + "\n"))
	writer.Flush()
	writer.Reset(writer)

}
