package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	cfg *ini.File

	RunMode string

	App      = &AppSetting{}
	Log      = &LogSetting{}
	Server   = &ServerSetting{}
	Database = &DatabaseSetting{}
	Redis    = &RedisSetting{}
)

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	RunMode = cfg.Section("").Key("RUN_MODE").MustString("debug")
	mapTo("app", App)
	mapTo("log", Log)
	mapTo("server", Server)
	mapTo("database", Database)
	mapTo("redis", Redis)

	// 特殊的转换
	Server.ReadTimeout = Server.ReadTimeout * time.Second
	Server.WriteTimeout = Server.WriteTimeout * time.Second
}

// 将每一个 section 映射到结构体里
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

type AppSetting struct {
	SessionKey  string `ini:"SESSION_KEY"`
	SessionName string `ini:"SESSION_NAME"`
}

type LogSetting struct {
	SavePath   string `ini:"SAVE_PATH"`
	FilePrefix string `ini:"FILE_PREFIX"`
	FileExt    string `ini:"FILE_EXT"`
	TimeFormat string `ini:"TIME_FORMAT"`
}

type ServerSetting struct {
	HttpPort     string        `ini:"HTTP_PORT"`
	ReadTimeout  time.Duration `ini:"READ_TIMEOUT"`
	WriteTimeout time.Duration `ini:"WRITE_TIMEOUT"`
}

type DatabaseSetting struct {
	Type     string `ini:"TYPE"`
	User     string `ini:"USER"`
	Password string `ini:"PASSWORD"`
	Host     string `ini:"HOST"`
	Name     string `ini:"NAME"`
}

type RedisSetting struct {
	Host     string `ini:"HOST"`
	DB       int    `ini:"DB"`
	Password string `ini:"PASSWORD"`
	PoolSize int    `ini:"POOL_SIZE"`
}
