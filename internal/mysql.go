// Package internal
/***********************************************************************************************************************
* ProjectName:  consumerOrders
* FileName:     mysql.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 03:18:54
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package internal

import (
	"database/sql"

	"github.com/ckf10000/gologger/v3/log"
)

func ConnectMysql(mysqlURI string, log *log.FileLogger) (*sql.DB, error) {
	// 连接 MySQL
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		log.Error("Failed to connect to MySQL: %v", err)
		return nil, err
	} else {
		log.Info("连接Mysql数据库：%s 正常.", mysqlURI)
	}
	return db, nil
}
