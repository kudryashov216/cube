package paint

import (
	"bufio"
)

func DrawDice(result *uint8, writer *bufio.Writer) {

	switch *result {

	case 1:
		writer.Write([]byte("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t|         |\n\t\t\t\t\t\t|    o    |\n\t\t\t\t\t\t|         |\n\t\t\t\t\t\t+---------+\n"))
	case 2:
		writer.Write([]byte("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t| o       |\n\t\t\t\t\t\t|         |\n\t\t\t\t\t\t|       o |\n\t\t\t\t\t\t+---------+\n"))
	case 3:
		writer.Write([]byte("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t| o       |\n\t\t\t\t\t\t|    o    |\n\t\t\t\t\t\t|       o |\n\t\t\t\t\t\t+---------+\n"))
	case 4:
		writer.Write([]byte("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t|         |\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t+---------+\n"))
	case 5:
		writer.Write([]byte("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t|    o    |\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t+---------+\n"))
	case 6:
		writer.Write([]byte("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t+---------+\n"))
	}

	writer.Flush()
	writer.Reset(writer)

}
