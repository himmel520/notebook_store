package server

import (
	"context"
	"fmt"
	"strconv"
	"time"

	log "github.com/himmel520/notebook_store/authentication/internal/logger"
	"github.com/himmel520/notebook_store/authentication/internal/models"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedis(config RedisConf) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	ctx := context.Background()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &Redis{rdb: rdb, ctx: ctx}, nil
}

func (r *Redis) SetSessionID(sessionIdID string, registeredUser *models.User) error {
	return r.rdb.Set(
		r.ctx,
		fmt.Sprintf("sessions:%v", sessionIdID),
		strconv.FormatBool(registeredUser.IsAdmin),
		time.Hour*24*30).Err()
}

func (r *Redis) DeleteSessionID(sessionIdID string) error {
	return r.rdb.Del(
		r.ctx,
		fmt.Sprintf("sessions:%v", sessionIdID),
	).Err()
}

func (r *Redis) GetIsAdmin(sessionIdID string) (string, error) {
	role, err := r.rdb.Get(
		r.ctx,
		fmt.Sprintf("sessions:%v", sessionIdID),
	).Result()
	if err != nil {
		log.Logger.Info(err)
		return "", err
	}
	return role, nil
}
