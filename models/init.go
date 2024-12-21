package models

import (
	"fmt"
	"go_study/database/mysql"
	userModels "go_study/models/user"
)

func InitMysql() error {
	// 初始化数据库连接
	if err := mysql.InitDB(); err != nil {
		return fmt.Errorf("数据库初始化失败: %v", err)
	}

	// 初始化表
	if err := userModels.InitUser(); err != nil {
		return fmt.Errorf("初始化 user 表失败: %v", err)
	}
	return nil
}
