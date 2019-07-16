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

var Debug bool
var PingEachMinute int
var MaxIdleConns int
var MaxOpenConns int

var lastPing time.Time

var DB *gorm.DB

func New() (*gorm.DB, error) {
	if DB == nil {
		Connect()
		return New()
	}
	if PingEachMinute > 0 && time.Now().After(lastPing.Add(time.Duration(PingEachMinute)*time.Minute)) {
		lastPing = time.Now()
		err := DB.DB().Ping()
		if err != nil {
			DB.Close()
			Connect()
			return New()
		}
	}

	return DB, nil
}

func Connect() {
	dbLink := GetLInk()
	var err error
	DB, err = gorm.Open("postgres", dbLink)
	if err != nil {
		panic(err)
	}
	DB.LogMode(Debug)
	DB.DB().SetMaxIdleConns(MaxIdleConns)
	DB.DB().SetMaxOpenConns(MaxOpenConns)
}

func GetLInk() string {
	dbLink := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", User, Pass, Host, Port, Name)

	return dbLink
}
