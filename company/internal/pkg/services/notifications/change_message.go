package notifications

type ChangeType int

const (
	Create ChangeType = iota
	Patch
	Delete
)

type ChangeMessage struct {
	ChangeType ChangeType
	Id         string
}
