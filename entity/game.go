package entity

type Game struct {
	ID        uint
	Category  string
	Questions []Question
	Players   []User
}
