package pgdb

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
	"strings"
	"time"
)

var User string
var Pass string
var Host string
var Port string
var Name string
var Ping int
var ConfigPath string
var Debug bool

var DB *gorm.DB
var lastPing time.Time

func New() *gorm.DB {

	if Ping <= 0 {
		Ping = 1
	}

	if DB == nil {
		Connect()
		return New()
	}
	if time.Now().After(lastPing.Add(time.Duration(Ping) * time.Minute)) {
		lastPing = time.Now()
		err := DB.DB().Ping()
		if err != nil {
			DB.Close()
			Connect()
			return New()
		}
	}

	return DB
}

func Close() {
	err := DB.Close()
	if err != nil {
		panic(err.Error())
	}
}

func Connect() {
	dbLink := GetLInk()
	var err error
	DB, err = gorm.Open("postgres", dbLink)
	DB.LogMode(Debug)
	if err != nil {
		panic(err.Error())
	}
	DB.DB().SetMaxIdleConns(100)
	DB.DB().SetMaxOpenConns(1000)
}

func GetLInk() string {

	if ConfigPath != "" {
		for k := 0; k < 5; k ++ {
			path := strings.Repeat("../", k) + ConfigPath
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				ConfigPath = path
				break
			}
		}
		cfg, err := config.NewConfig("ini", ConfigPath)
		if err != nil {
			panic(err.Error())
		}
		User = cfg.String("SqlUser")
		Pass = cfg.String("SqlPass")
		Host = cfg.String("SqlHost")
		Port = cfg.String("SqlPort")
		Name = cfg.String("SqlName")
	}

	dbLink := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", User, Pass, Host, Port, Name)

	return dbLink
}
