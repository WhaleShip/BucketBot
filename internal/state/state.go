package state

import "sync"

const (
	NewNoteState = iota + 1
)

var UserStates sync.Map

func SetState(userID int, state int) {
	UserStates.Store(userID, state)
}

func GetState(userID int) (int, bool) {
	val, ok := UserStates.Load(userID)
	return val.(int), ok
}
