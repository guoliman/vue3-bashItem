package crond

import (
	"github.com/robfig/cron"
)

func RunTask() {
	c := cron.New()
	// * * * * * * 秒 分 时 日 月 周
	//c.AddFunc("10 30 8 * * *", NewTask().OssCrond)  		// 每天上午8点30分10秒执行
	//c.AddFunc("1 */30 * * * *", NewTask().OssCrond)   	// 每30分钟的第01秒执行
	//c.AddFunc("*/10 * * * * *", NewTask().OssCrond)  		// 每10秒执行

	c.Start()
}
