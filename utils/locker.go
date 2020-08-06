package utils

import (
	"github.com/bsm/redislock"
	"github.com/fishjar/gin-rest-boilerplate/db"
)

// Locker redis锁
var Locker *redislock.Client = redislock.New(db.Redis)
