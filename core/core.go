package core

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/asdine/storm"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port        int
	URL         string
	Compress    bool
	Domain      string
	Dev         bool
	Ssl         bool
	MailDomain  string
	MailPKey    string
	SQLDBName   string
	SQLUser     string
	SQLPassword string
	SQLHost     string
	SQLPort     string
	StripeKey   string
	AppEmail    string
	AppName     string
}

type AppConfig struct {
	AppId      int    `storm:"id,increment" json:"AppId" `
	AppName    string `json:"AppName" storm:"index,unique"`
	DomainName string `json:"DomainName" storm:"index,unique"`
}

var Config ServerConfig
var Secure bool
var DbFile = "csite.db"
var DB = AbsolutePath("dbf/" + DbFile)
var BinPath = BinaryPath()
var SKey = "JPcT1k6SwyqA2JX-oyGjGfOfHzKsN2BQdI4Cr56KG9M="
var StaticUrl = "/static"

func BinaryPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("Error getting file path")
	}
	return dir
}

func CleanPath(path string) string {
	path = strings.Replace(path, "//", "/", -1)
	return path
}

func AbsolutePath(path string) string {
	if flag.Lookup("test.v") == nil {
		return CleanPath(filepath.FromSlash(BinPath + "/" + path))
	} else {
		return CleanPath(filepath.FromSlash(path))
	}
}

func FixPathSlash(path string) string {
	return filepath.FromSlash(path)
}

func SetupConfig() {
	if flag.Lookup("test.v") == nil {
		viper.AddConfigPath(BinPath)
		viper.SetConfigName("config")
	} else {
		log.Println("Testing.......")
		viper.AddConfigPath("./")
		viper.SetConfigName("test-config")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&Config)

	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	//Set env
	if Config.Dev {
		Secure = false

	} else {
		Secure = true
	}

}

func InitDb(dbName string) {
	db, err := storm.Open(DB)
	defer db.Close()
	if err == nil {
		log.Println("App Initilaized")
	}
}

func Start() {
	log.Println("Starting")
	SetupConfig()
	if flag.Lookup("test.v") == nil {
		InitDb(DB)
	}
}

func StartMulti() {
	SetupConfig()
}
