package service

import (
	"strconv"
	"strings"

	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/fishjar/gin-rest-boilerplate/db"
	"github.com/fishjar/gin-rest-boilerplate/model"
)

// SetUserToRedis 保存用户登录信息到redis
func SetUserToRedis(user *model.UserCurrent, issuedAt int64) error {
	// 保存登录信息到redis
	userKey := "user:" + user.UserID
	err := db.Redis.HSet(userKey, map[string]interface{}{
		"aid":   user.AuthID,
		"iss":   strconv.FormatInt(issuedAt, 10),
		"roles": strings.Join(user.Roles, ","),
	}).Err()
	if err != nil {
		return err
	}

	// 设置redis过期时间
	// TODO: 还应判断帐号的过期时间是否小于 config.JWTExpiresIn
	db.Redis.Expire(userKey, config.JWTExpiresIn)

	return nil
}
