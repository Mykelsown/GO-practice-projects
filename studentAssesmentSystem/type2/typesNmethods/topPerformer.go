package type2

func TopPerformer[T Learner](learners []T) T {
	highestScorer := learners[0]
	_, initialHighestScore := learners[0].Progress()
	for _, learner := range learners[1:] {
		if _, score := learner.Progress(); score > initialHighestScore {
			highestScorer = learner
		}
	}
	return highestScorer
}
