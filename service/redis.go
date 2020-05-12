package service

import (
	"strings"

	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/fishjar/gin-rest-boilerplate/db"
)

// SetUserToRedis 保存用户登录信息到redis
func SetUserToRedis(uid string, aid string, roles []string) error {
	// 保存登录信息到redis
	userKey := "user:" + uid
	err := db.Redis.HSet(userKey, map[string]interface{}{
		"aid":   aid,
		"roles": strings.Join(roles, ","),
	}).Err()
	if err != nil {
		return err
	}

	// 设置redis过期时间
	// TODO: 还应判断帐号的过期时间是否小于 config.JWTExpiresIn
	db.Redis.Expire(userKey, config.JWTExpiresIn)

	return nil
}
