package entity

type WaitingMember struct {
	UserID    uint
	Timestamp int64
	Category  Category
}

type MatchedUsers struct {
	Category Category
	UserIDs  []uint
}
