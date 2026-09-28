package game

import (
	"bufio"
	"cube/cmd"
	"cube/internal/output"
	"sort"
	"strconv"
)

const (
	NUMBER_THROWS_EASY   uint8 = 6
	NUMBER_THROWS_MIDDLE uint8 = 3
	NUMBER_THROWS_HIGH   uint8 = 1
)

func DifficultyChoice(
	countThrows *uint8,
	writer *bufio.Writer,
	reader *bufio.Reader) {

	var answer uint8
	var isFirst bool

	diff := getDifficultyChoice()

	for {

		if !isFirst {
			outputOptions(diff)
		}

		output.Output("\r")

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

		cmd.GetConsole().Clear()

		output.Output("There is no such level.\n")

		isFirst = false

	}

	switch answer {
	case 0:
		*countThrows = NUMBER_THROWS_EASY
	case 1:
		*countThrows = NUMBER_THROWS_MIDDLE
	case 2:
		*countThrows = NUMBER_THROWS_HIGH
	}

	output.Output("settings succed!")

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

func outputOptions(diff map[uint8]string) {

	output.Output("Choose the difficulty level:\n\t")

	keys := make([]uint8, 0, len(diff))
	for k := range diff {
		keys = append(keys, k)
	}

	sort.Slice(keys,
		func(i, j int) bool {
			return keys[i] < keys[j]
		})

	for _, k := range keys {
		output.Output(strconv.Itoa(int(k)) + "-" + diff[k] + "\n\t")
	}

}
