package entity

type Question struct {
	ID              uint
	Question        string
	PossibleAnswers []string
	CorrectAnswer   []string
	Difficulty      string
	Category        string
}
