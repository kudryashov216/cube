package paint

import (
	"cube/internal/output"
)

func DrawDice(result *uint8) {

	switch *result {

	case 1:
		output.Output("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t|         |\n\t\t\t\t\t\t|    o    |\n\t\t\t\t\t\t|         |\n\t\t\t\t\t\t+---------+\n")
	case 2:
		output.Output("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t| o       |\n\t\t\t\t\t\t|         |\n\t\t\t\t\t\t|       o |\n\t\t\t\t\t\t+---------+\n")
	case 3:
		output.Output("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t| o       |\n\t\t\t\t\t\t|    o    |\n\t\t\t\t\t\t|       o |\n\t\t\t\t\t\t+---------+\n")
	case 4:
		output.Output("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t|         |\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t+---------+\n")
	case 5:
		output.Output("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t|    o    |\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t+---------+\n")
	case 6:
		output.Output("\t\t\t\t\t\t+---------+\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t| o     o |\n\t\t\t\t\t\t+---------+\n")
	}

}
