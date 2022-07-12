package logging

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Aj002Th/LittlePrince/pkg/setting"
)

// getLogFilePath 获取日志文件保存路径
func getLogFilePath() string {
	return setting.Log.SavePath
}

// getLogFileName 根据时间和项目名称组合出日志名称
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.Log.FilePrefix,
		time.Now().Format(setting.Log.TimeFormat),
		setting.Log.FileExt,
	)
}

// getLogFile 生成对应的日志文件
func GetLogFile(prefix string) (*os.File, error) {
	filePath := getLogFilePath()
	fileName := getLogFileName()

	fp, err := os.OpenFile(filePath+prefix+fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalf("logging.GetLogFile err: %v", err)
	}
	return fp, nil
}
