package entity

type Question struct {
	ID              uint
	Question        string
	PossibleAnswers []PossibleAnswer
	CorrectAnswerID uint
	Difficulty      QuestionDifficulty
	CategoryID      uint
}

type PossibleAnswer struct {
	ID      uint
	Content string
	Choice  PossibleAnswerChoice
}

type PossibleAnswerChoice uint8

const (
	PossibleAnswerA PossibleAnswerChoice = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)

func (pac PossibleAnswerChoice) IsValid() bool {
	if pac >= PossibleAnswerA && pac <= PossibleAnswerD {

		return true
	}

	return false
}

type QuestionDifficulty uint8

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)

func (qd QuestionDifficulty) IsValid() bool {
	if qd >= QuestionDifficultyEasy && qd <= QuestionDifficultyHard {

		return true
	}

	return false
}
