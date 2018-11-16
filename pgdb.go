package pgdb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

var User string
var Pass string
var Host string
var Port string
var Name string
var Ping int
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
	DB.Close()
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
	dbLink := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", User, Pass, Host, Port, Name)

	return dbLink
}
