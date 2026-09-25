package game

import (
	"bufio"
	"cube/internal/output"
	"strconv"
)

func DifficultyChoice(
	writer *bufio.Writer,
	reader *bufio.Reader) uint8 {

	var answer uint8
	var isFirst bool

	diff := getDifficultyChoice()

	for {

		if !isFirst {

			output.Output("Choose the difficulty level:\n\t")

			for key, value := range diff {
				output.Output(strconv.Itoa(int(key)) + "-" + value + "\n\t")
			}

		}

		writer.WriteByte(0b00001101)
		writer.Flush()
		writer.Reset(writer)

		bytes_, err := reader.ReadBytes(0b00001010)

		if err != nil {
			output.Output("Error: " + err.Error() + "\n\t")
			ExitInGame()
		}

		if bytes_[0] == 0b00001101 {
			isFirst = true
			continue
		}

		number, err := strconv.Atoi(string(bytes_[0]))

		if err != nil {
			output.Output("Error: " + err.Error() + "\n\t")
			ExitInGame()
		}

		answer = uint8(number)

		if findLevel(&answer, diff) {
			break
		}

		output.Output("There is no such level.\n")
		isFirst = false

	}

	return answer

}

func getDifficultyChoice() map[uint8]string {

	difficulty := make(map[uint8]string)

	difficulty[0] = "Easy"
	difficulty[1] = "Middle"
	difficulty[2] = "High"

	return difficulty

}

func findLevel(
	answer *uint8,
	diff map[uint8]string) bool {

	for key := range diff {

		if *answer == key {
			return true
		}

	}

	return false

}
