package menu

func (m Menu_) UpdateCurrentCommand(currentCommand *uint8) {

	switch *currentCommand {
	case 0:
		*currentCommand++
	case 1:
		*currentCommand--
	}
}
