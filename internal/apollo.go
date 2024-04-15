// Package internal
/***********************************************************************************************************************
* ProjectName:  consumerOrders
* FileName:     apollo.go
* Description:  TODO
* Author:       ckf10000
* CreateDate:   2024/04/15 04:00:31
* Copyright ©2011-2024. Hunan xyz Company limited. All rights reserved.
* *********************************************************************************************************************/
package internal

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/ckf10000/gologger/v3/log"
)

func GetApolloCache(log *log.FileLogger) agcache.CacheInterface {
	c := &config.AppConfig{
		AppID:          "org-system-order-consumer",
		Cluster:        "PRO",
		IP:             "http://192.168.3.232:8080",
		NamespaceName:  "application",
		IsBackupConfig: true,
		Secret:         "8c64c50f8ea0452db1b00cc0e8f2c9a1",
	}

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	log.Info("初始化Apollo配置成功.")
	cache := client.GetConfigCache(c.NamespaceName)
	return cache
}
