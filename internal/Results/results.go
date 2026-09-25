package results

import (
	"bufio"
	"cube/internal/output"
	"cube/internal/paint"
	"strconv"
)

func Output(enteredNumber *byte, result *uint8, writer *bufio.Writer) {

	paint.DrawDice(result)

	if string(rune(*enteredNumber)) == strconv.Itoa(int(*result)) {
		output.Output("You win!\n")
		return
	}

	output.Output("You lost!\n")
	output.Output("The correct number: " + strconv.Itoa(int(*result)) + "\n")

}
