package crond

type Task struct{}

func NewTask() *Task {
	return &Task{}
}

//func (t *Task) SendEmailTask() {
//	return
//}

// oss定时更新
//func (t *Task) OssCrond() {
//	result, err := oss.DownloadOssInfo() // 获取oss数据
//	if err != nil {
//		logger.FileLogger.Errorf(fmt.Sprintf("定时刷新oss 失败------ %v", err))
//		return
//	}
//	logger.FileLogger.Info(fmt.Sprintf("定时刷新oss 成功====== %v ", result))
//}
