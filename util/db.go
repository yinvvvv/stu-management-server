package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	Db, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}
}

func QueryAndParse(Db *sql.DB, queryStr string) map[string]string {
	rows, err := Db.Query(queryStr)
	defer rows.Close()
	if err != nil {
		log.Printf("查询出错,SQL语句:%s\n错误详情:%s\n", queryStr, err.Error())
		return nil
	}
	cols, _ := rows.Columns()
	if len(cols) > 0 {
		buff := make([]interface{}, len(cols))
		data := make([][]byte, len(cols))
		dataKv := make(map[string]string, len(cols))
		for i, _ := range buff {
			buff[i] = &data[i]
		}

		for rows.Next() {
			rows.Scan(buff...)
		}

		for k, col := range data {
			dataKv[cols[k]] = string(col)
		}
		return dataKv
	} else {
		return nil
	}
}

func QueryAndParseRows(Db *sql.DB, queryStr string) []map[string]string {
	rows, err := Db.Query(queryStr)
	defer rows.Close()
	if err != nil {
		fmt.Printf("查询出错:\nSQL:\n%s, 错误详情:%s\n", queryStr, err.Error())
		return nil
	}
	cols, _ := rows.Columns()
	if len(cols) > 0 {
		var ret []map[string]string
		for rows.Next() {
			buff := make([]interface{}, len(cols))
			data := make([][]byte, len(cols))
			for i, _ := range buff {
				buff[i] = &data[i]
			}
			rows.Scan(buff...)
			dataKv := make(map[string]string, len(cols))
			for k, col := range data {
				fmt.Printf("%30s:\t%s\n", cols[k], col)
				dataKv[cols[k]] = string(col)
			}
			ret = append(ret, dataKv)
		}
		return ret
	} else {
		return nil
	}
}

func Data2Json(anyData interface{}) string {
	JsonByte, err := json.Marshal(anyData)
	if err != nil {
		log.Printf("数据序列化为json出错:\n%s\n", err.Error())
	}
	return string(JsonByte)
}
