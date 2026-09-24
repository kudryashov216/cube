package results

import (
	"bufio"
	"cube/internal/paint"
	"strconv"
)

func Output(result *uint8, writer *bufio.Writer) {

	paint.DrawDice(result, writer)

	if string(rune(*result)) == strconv.Itoa(int(*result)) {
		writer.Write([]byte("You win!\n"))
		writer.Flush()
		writer.Reset(writer)
		return
	}

	writer.Write([]byte("You lost!\n"))
	writer.Flush()
	writer.Reset(writer)

	writer.Write([]byte("The correct number: " + strconv.Itoa(int(*result)) + "\n"))
	writer.Flush()
	writer.Reset(writer)

}
