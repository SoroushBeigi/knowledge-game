package entity

type Permission struct {
	ID    uint
	Title string
}

// role-action
// Must be inserted into database manually or written in migrations
const (
	UserListPermission   = "user-list"
	UserDeletePermission = "user-delete"
)
