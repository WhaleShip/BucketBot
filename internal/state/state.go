package state

import (
	"log"
	"sync"
)

var instance *UserState

func InitializeStateMachine() {
	instance = &UserState{
		userData: make(map[int]int),
	}
	log.Println("States initialized")
}

type UserStateInterface interface {
	Set(userID int, value int)
	Get(userID int) (int, bool)
	Delete(userID int)
}

type UserState struct {
	mu       sync.Mutex
	userData map[int]int
}

func SetUserState(userID int, value int) {
	instance.mu.Lock()
	defer instance.mu.Unlock()
	instance.userData[userID] = value
}

func GetUserState(userID int) (int, bool) {
	instance.mu.Lock()
	defer instance.mu.Unlock()
	value, exists := instance.userData[userID]
	return value, exists
}

func Delete(userID int) {
	instance.mu.Lock()
	defer instance.mu.Unlock()
	delete(instance.userData, userID)
}
