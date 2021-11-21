// Package cron 定时任务
package cron

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/axiaoxin-com/investool/datacenter"
	"github.com/axiaoxin-com/investool/services"
	"github.com/axiaoxin-com/logging"
)

// SyncIndustryList 同步行业列表
func SyncIndustryList() {
	ctx := context.Background()
	indlist, err := datacenter.EastMoney.QueryIndustryList(ctx)
	if err != nil {
		logging.Errorf(ctx, "SyncIndustryList QueryIndustryList error:", err)
		promSyncError.WithLabelValues("SyncIndustryList").Inc()
		return
	}
	services.StockIndustryList = indlist

	// 更新文件
	b, err := json.Marshal(indlist)
	if err != nil {
		logging.Errorf(ctx, "SyncIndustryList json marshal error:", err)
		promSyncError.WithLabelValues("SyncIndustryList").Inc()
		return
	}
	if err := ioutil.WriteFile(services.IndustryListFilename, b, 0666); err != nil {
		logging.Errorf(ctx, "SyncIndustryList WriteFile error:", err)
		promSyncError.WithLabelValues("SyncIndustryList").Inc()
		return
	}
}
