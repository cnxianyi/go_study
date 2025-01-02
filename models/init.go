package models

import (
	"fmt"
	"go_study/database/mysql"
	"go_study/database/redis"
	userModels "go_study/models/user"
)

func InitMysql() {
	// 初始化数据库连接
	if err := mysql.InitDB(); err != nil {
		fmt.Printf("数据库初始化失败: %v", err)
	}

	// 初始化表
	if err := userModels.InitUser(); err != nil {
		fmt.Printf("初始化 user 表失败: %v", err)
	}
	return
}

func InitRedis() {
	_, err := redis.ConnectionRedis()
	if err != nil {
		fmt.Printf("数据库初始化失败: %v", err)
		return
	}
}
