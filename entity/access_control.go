package entity

type AccessControl struct {
	ID           uint
	ActorType    ActorType //User or Role
	ActorID      uint
	PermissionID uint
}

type ActorType string

const (
	RoleActorType = "role"
	UserActorType = "user"
)
