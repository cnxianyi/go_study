package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB   // 全局数据库连接池变量，初始值为 nil
	once sync.Once // 全局 sync.Once，确保数据库连接只会初始化一次
)

func ConnectionMysql() (*sql.DB, error) {
	// 读取 MYSQL 环境变量
	dsn := os.Getenv("MYSQL")
	if dsn == "" {
		return nil, fmt.Errorf("MYSQL: 环境变量未设置")
	}
	// 检查是否包含 /
	if !strings.Contains(dsn, "/") {
		return nil, fmt.Errorf("MYSQL: 环境变量格式错误")
	}

	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		println("MySQL 连接失败")
		return nil, fmt.Errorf("MySQL 连接失败: %v", err)
	}

	// 测试连接
	if err := db.Ping(); err != nil {
		// mysqlErr 是具体错误, ok是mysqlErr是否属于MySQLError类型, 即断言是否成功的布尔值
		// 即 ok: 断言成功,是MySQLError错误. mysqlErr: 具体的错误 mysqlErr.Number: 错误码
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1049 {
			println("MySQL数据库不存在 , 尝试创建数据库")

			db.Close()

			// 重连MySQL
			db, err = sql.Open("mysql", strings.Split(dsn, "/")[0]+"/")
			if err != nil {
				println("MySQL 连接失败")
				return nil, fmt.Errorf("MySQL 连接失败: %v", err)
			}

			// 创建
			_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", strings.Split(dsn, "/")[1]))
			if err != nil {
				return nil, fmt.Errorf("创建数据库失败: %v", err)
			}

			fmt.Printf("数据库 %s 创建成功", strings.Split(dsn, "/")[1])

			db.Close()

			// 再次重连
			db, err := sql.Open("mysql", dsn)
			if err != nil {
				println("MySQL 连接失败")
				return nil, fmt.Errorf("MySQL 连接失败: %v", err)
			}

			println("MySQL 连接成功")

			return db, nil

		}
		return nil, fmt.Errorf("ping MySQL 失败: %v", err)
	}
	fmt.Println("MySQL 连接成功")

	return db, nil
}

// InitDB 初始化数据库连接
func InitDB() error {
	var err error
	once.Do(func() { // 保证 ConnectionMysql 只会被调用一次
		db, err = ConnectionMysql()
	})
	return err
}

// GetDB 获取数据库连接实例
func GetDB() *sql.DB {
	return db // 返回全局数据库连接池实例
}
