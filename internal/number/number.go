package number

import (
	"bufio"
	"math/rand"
	"sync"
)

func GetRandomNumber(randomaizer *rand.Rand) int {

	var random int

	random = randomaizer.Intn(6) + 1

	if random <= 0 {
		random = GetRandomNumber(randomaizer)
	}

	return random

}

func InputNewNumber(enterednumber *byte, reader *bufio.Reader, writer *bufio.Writer, wg *sync.WaitGroup) {

	var isFirst bool

	defer wg.Done()

	for {

		if !isFirst {
			writer.Write([]byte("\nEnter a random number from 1 to 6: "))
			writer.Flush()
			writer.Reset(writer)
		}

		bytes_, _ := reader.ReadBytes(0b00001010)

		if bytes_[0] == 13 {
			isFirst = true
			continue
		}

		if len(bytes_) > 3 {
			IncorrectNumberEntered(writer)
			continue
		}

		*enterednumber = bytes_[0]

		break

	}

}

func IncorrectNumberEntered(writer *bufio.Writer) {

	writer.Write([]byte("\nYou entered the wrong value! Please try again.\n"))
	writer.Flush()
	writer.Reset(writer)

}
