package grammarmysql

import (
	"fmt"
	"go_study/database/mysql"
)

/*
✅ Query 多行查询结果集
✅ QueryRow 单行查询结果集
✅ Exec 无返回数据SQL
❌ Prepare 预编译SQL
❌ stmt.Exec
❌ stmt.Query
❌ stmt.QueryRow
❌ db.Begin 事务
❌
*/
func MysqlTest() {
	var db = mysql.GetDB()
	// Query: 查询结果集. 返回多行数据
	rows, err := db.Query(`SELECT id FROM t1`)
	if err != nil {
		println("查询错误")
		return
	}
	defer rows.Close() // 需要关闭连接
	// 循环结果集
	for rows.Next() {
		var id int
		// 扫描结果集
		err := rows.Scan(&id)
		if err != nil {
			println("扫描错误")
			return
		}
		println(id)
	}

	// QueryRow: 单行查询
	var id int
	err = db.QueryRow(`SELECT id FROM t1 WHERE id = (SELECT MAX(id) FROM t1)`).Scan(&id)
	if err != nil {
		println("查询错误")
		return
	}
	println("最后一条数据的id为:", id)

	/*
	   Exec: 执行不返回数据的SQL 如(INSERT UPDATE DELETE)等
	   返回 sql.Result. 可以获取受影响的行数
	       LastInsertId() 返回最后受影响的id
	       RowsAffected() 返回受影响的行数
	*/
	result, err := db.Exec(`
    DELETE t1
    FROM t1
    JOIN (
        SELECT MAX(id) AS max_id
        FROM t1
    ) AS temp
    ON t1.id = temp.max_id;`)
	if err != nil {
		println("删除失败")
		return
	}

	// DELETE 不支持 LastInsertId
	r1, _ := result.LastInsertId()
	r2, err := result.RowsAffected()
	if err != nil {
		println("获取行数失败")
		return
	}
	fmt.Printf("最后受影响的id: %d\n", r1)
	fmt.Printf("受影响的行数: %d\n", r2)
}
