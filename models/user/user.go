package user

import (
	"fmt"
	"log"

	"go_study/database/mysql"
)

// User 表的结构体定义
type User struct {
	ID    int
	Name  string
	Email string
}

// InitTable 确保 user 表存在
func InitUser() error {
	// 获取数据库连接
	db := mysql.GetDB()

	// 创建 user 表（如果不存在）
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS user (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL
		);`
	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("创建 user 表失败: %v", err)
	}

	return nil
}

// AddUser 添加新用户
func AddUser(name, email, password string) (int64, error) {
	// 获取数据库连接
	db := mysql.GetDB()

	// 插入数据
	insertQuery := "INSERT INTO user (name, email , password) VALUES (?, ? , ?)"
	result, err := db.Exec(insertQuery, name, email, password)
	if err != nil {
		return 0, fmt.Errorf("插入用户失败: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("获取插入用户的 ID 失败: %v", err)
	}

	log.Println("用户添加成功")
	return id, nil
}

// GetAllUsers 获取所有用户
func GetAllUsers() ([]User, error) {
	// 获取数据库连接
	db, err := mysql.ConnectionMysql()
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %v", err)
	}
	defer db.Close()

	// 查询所有用户
	query := "SELECT id, name, email FROM user"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %v", err)
	}
	defer rows.Close()

	// 解析结果
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("解析用户失败: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}
