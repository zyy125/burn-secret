package store

import (
	"burn-secret/models"
	"sync"
	"time"
)

var (
	Store = make(map[string]*models.Secret)

	Mutex sync.RWMutex
)

func StoreSecret(secret *models.Secret) {
	Mutex.Lock()
	Store[secret.ID] = secret
	Mutex.Unlock()
}

func GetSecret(id string) (secret *models.Secret, exist bool) { 
	Mutex.Lock()
	if Store[id] == nil {
		Mutex.Unlock()
		return nil, false
	}

	secret = Store[id]
	Mutex.Unlock()
	secret.ViewsCount++
	duration := time.Duration(secret.ExpiryMinutes) * time.Minute

	if time.Now().After(secret.CreatedAt.Add(duration)) || secret.ViewsCount > secret.MaxViews{
		Mutex.Lock()
		delete(Store, id)
		Mutex.Unlock()
		return nil, false
	}

	return secret, true
}

func CleanTask() {
	for {
		time.Sleep(time.Minute * 1)
		Mutex.Lock()
		for id, secret := range Store{
			duration := time.Duration(secret.ExpiryMinutes) * time.Minute
			if time.Now().After(secret.CreatedAt.Add(duration)) || secret.ViewsCount > secret.MaxViews{
				delete(Store, id)
			}
		}	
		Mutex.Unlock()
	}
}
