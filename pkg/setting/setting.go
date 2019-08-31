package setting

import (
	"github.com/go-ini/ini"
	"log"
	"strings"
	"time"
)

var cfg *ini.File

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	RedisHost        string
	RedisPassword    string
	RedisMaxidle     int
	RedisMaxActive   int
	RedisIdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
	ViewUrl   string

	UploadUrl  string
	UploadPath string

	LogoMobileUrl string
	LogoPcUrl     string
	LogoWaterUrl  string

	SitePre        string
	SiteUrl        string
	SiteName       string
	SiteBrief      string
	SiteSeoWord    string
	TimeZone       string
	Lang           string
	RunLevel       int
	RunLevelReason string

	CookieDomain string
	CookiePath   string

	RuntimeRootPath string
	LogPath         string
	LogSavePath     string
	LogSaveName     string
	TimeFormat      string
	LogFileExt      string
}

var ServerSetting = &Server{}

type Image struct {
	ImageSavePath   string
	ImageMaxSize    int
	ImageAlloweXts  string
	ImageAllowExts  []string
	RuntimeRootPath string
}

var ImageSetting = &Image{}

type Smtp struct {
	EmailUser string
	EmailPass string
	EmailHost string
	EmailPort string
}

var SmtpSetting = &Smtp{}

func init() {
	var err error
	//读取配置
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("无法解析 'conf/config.ini':%v ", err)
	}

	//加载服务
	loadServer()

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

	RedisSetting.RedisIdleTimeout = RedisSetting.RedisIdleTimeout * time.Second
	ImageSetting.ImageAllowExts = strings.Split(ImageSetting.ImageAlloweXts, ",")

}

//加载服务
func loadServer() {
	mapToSection("server", ServerSetting)
	mapToSection("database", DatabaseSetting)
	mapToSection("redis", RedisSetting)
	mapToSection("image", ImageSetting)
	mapToSection("smtp", SmtpSetting)
}

//映射服务
func mapToSection(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("无法获取'server': %v", err)
	}
}
