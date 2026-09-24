package game

func ScoreCounter(
	score *int64,
	decline *bool) {

	if !*decline {
		*score += 10
	} else {
		*score -= 10
		*decline = false
	}
}
