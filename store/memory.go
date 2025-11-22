package store

import (
	"burn-secret/models"
	"context"
	"time"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)
var (
	rdb *redis.Client
	ctx = context.Background()
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB:  0,
	})
}

func StoreSecret(secret *models.Secret) error {
	b, err := json.Marshal(*secret)
	if err != nil {
		return err
	}

	err = rdb.Set(ctx, secret.ID, b, time.Duration(secret.ExpiryMinutes) * time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetSecret(id string) (secret *models.Secret, err error) { 
	val, err := rdb.Get(ctx, id).Result()
	if err == redis.Nil{
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	secret = &models.Secret{}
	err = json.Unmarshal([]byte(val), secret)
	if err != nil {
		return secret, err
	}

	if secret.ViewsCount++; secret.ViewsCount >= secret.MaxViews {
		if err = rdb.Del(ctx, secret.ID).Err(); err != nil {
			return secret, err
		}
		return secret, nil
	}

	b, err := json.Marshal(secret)

	if err = rdb.Set(ctx, secret.ID, b, redis.KeepTTL).Err(); err != nil {
		return secret, err
	}

	return secret, nil
}
