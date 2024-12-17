package dbMysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionMysql() (*sql.DB, error) {
	// 读取 MYSQL 环境变量
	dsn := os.Getenv("MYSQL")
	if dsn == "" {
		return nil, fmt.Errorf("MYSQL: 环境变量未设置 示例: root:123456@tcp(127.0.0.1:3306)/")
	}

	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("mysql 连接失败: %v", err)
	}
	defer db.Close()

	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping mysql 失败: %v", err)
	}
	fmt.Println("MySQL 服务器连接成功")

	// 检查数据库是否存在
	database := os.Getenv("MYSQL_DATABASE")
	if database == "" {
		return nil, fmt.Errorf("MYSQL_DATABASE 环境变量未设置")
	}

	var dbName string
	query := "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = ?"
	err = db.QueryRow(query, database).Scan(&dbName)

	if err != nil {
		if err == sql.ErrNoRows {
			// 数据库不存在，创建数据库
			fmt.Printf("数据库 %s 不存在，正在创建...\n", database)
			createDBQuery := fmt.Sprintf("CREATE DATABASE %s", database)
			_, err := db.Exec(createDBQuery)
			if err != nil {
				return nil, fmt.Errorf("创建数据库失败: %v", err)
			}
			fmt.Printf("数据库 %s 创建成功\n", database)
		} else {
			return nil, fmt.Errorf("查询数据库失败: %v", err)
		}
	}

	// 重新连接到指定的数据库
	dsnWithDB := fmt.Sprintf("%s%s", dsn, database) // 正确拼接数据库名
	dbWithDB, err := sql.Open("mysql", dsnWithDB)
	if err != nil {
		return nil, fmt.Errorf("连接数据库 %s 失败: %v", database, err)
	}

	if err := dbWithDB.Ping(); err != nil {
		return nil, fmt.Errorf("连接数据库 %s 后 Ping 失败: %v", database, err)
	}

	fmt.Printf("数据库 %s 连接成功\n", database)
	return dbWithDB, nil
}
