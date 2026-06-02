package entity

type Permission struct {
	ID    uint
	Title string
}

// role-action
const (
	UserListPermission   = "user-list"
	UserDeletePermission = "user-delete"
)
