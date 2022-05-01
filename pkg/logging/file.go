package logging

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Aj002Th/LittlePrince/pkg/file"
	"github.com/Aj002Th/LittlePrince/pkg/setting"
)

// getLogFilePath 获取日志文件保存路径
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.LogSetting.LogSavePath)
}

// getLogFileName 根据时间和项目名称组合出日志名称
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.LogSetting.LogSaveName,
		time.Now().Format(setting.LogSetting.TimeFormat),
		setting.LogSetting.LogFileExt,
	)
}

// getLogFile 生成对应的日志文件
func GetLogFile(prefix string) (*os.File, error) {
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err := file.MustOpen(prefix+fileName, filePath)
	if err != nil {
		log.Fatalf("logging.GetLogFile err: %v", err)
	}
	return F, nil
}
