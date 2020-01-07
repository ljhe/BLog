package db

import (
	"Blog/util"
	"database/sql"
	"fmt"
	"strings"
)

// 插入数据
func (conn Connector) Insert(dbname string, args map[string]interface{}) (int64, error) {
	keys, value, err := handleArgs(args, 0)
	util.CheckErr(err)

	stmt, err := conn.Db.Prepare("INSERT " + dbname + " SET " + keys + "")
	util.CheckErr(err)

	// value... 将切片打散 传入 接收任意多个参数的函数
	res, err := stmt.Exec(value...)
	util.CheckErr(err)

	id, err := res.LastInsertId()
	util.CheckErr(err)
	return id, nil
}

// 更新数据
func (conn Connector) Update(dbname string, args map[string]interface{}, where map[string]interface{}) (int64, error) {
	var value []interface{}
	argsKeys, argsValues, err := handleArgs(args, 0)
	util.CheckErr(err)

	whereKeys, whereValues, err := handleArgs(where, 0)
	util.CheckErr(err)

	value = append(argsValues, whereValues...)

	stmt, err := conn.Db.Prepare("UPDATE " + dbname + " SET " + argsKeys + " WHERE " + whereKeys + "")
	util.CheckErr(err)

	res, err := stmt.Exec(value...)
	util.CheckErr(err)

	// 返回影响的行数
	affect, err := res.RowsAffected()
	util.CheckErr(err)
	return affect, err
}

// 查询数据
func (conn Connector) Select(dbname string, field string, where map[string]interface{}, orderBy string) ([]map[string]interface{}, error) {
	if field == "" {
		field = "*"
	}

	if orderBy != "" {
		orderBy = " ORDER BY " + orderBy + ""
	}

	var rows *sql.Rows
	var err error
	if where == nil {
		rows, err = conn.Db.Query("SELECT " + field + " FROM " + dbname + orderBy + "")
	} else {
		keys, value, err := handleArgs(where, 1)
		util.CheckErr(err)
		rows, err = conn.Db.Query("SELECT " + field + " FROM " + dbname + " WHERE " + keys + orderBy + "", value...)
	}
	util.CheckErr(err)

	// 返回列名 如果 rows 已经关闭会返回错误
	columns, err := rows.Columns()
	util.CheckErr(err)

	columnsLength := len(columns)
	// 临时存储每行数据
	cache := make([]interface{}, columnsLength)
	for k, _ := range cache {
		var a interface{}
		cache[k] = &a
	}

	var result []map[string]interface{}
	for rows.Next() {
		err := rows.Scan(cache...)
		util.CheckErr(err)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{})
		}
		result = append(result, item)
	}
	return result, nil
}

// 删除数据
func (conn Connector) Delete(dbname string, where map[string]interface{}) (int64, error) {
	keys, value, err := handleArgs(where, 1)
	util.CheckErr(err)

	stmt, err := conn.Db.Prepare("DELETE FROM " + dbname + " where " + keys + "")
	util.CheckErr(err)

	res, err := stmt.Exec(value...)
	util.CheckErr(err)

	affect, err := res.RowsAffected()
	util.CheckErr(err)
	return affect, nil
}

// 处理传过来的参数
func handleArgs (args map[string]interface{}, keyType int) (keys string, value []interface{}, err error) {
	for k, v := range args {
		if v != "" && v != nil {
			if keyType == 0 {
				keys = keys + k + "?,"
			} else {
				keys = keys + k + "? AND "
			}
			value = append(value, v)
		} else {
			return "", nil, fmt.Errorf("handleArgs is error")
		}
	}
	if keyType == 0 {
		keys = strings.TrimRight(keys, ",")
	} else {
		keys = strings.TrimRight(keys, "AND ")
	}
	return keys, value, nil
}
