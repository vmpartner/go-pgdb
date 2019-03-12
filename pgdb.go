package mqdb

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

func New() (*gorm.DB, error) {
	DB, err := Connect()
	if err != nil {
		return DB, err
	}
	if PingEachMinute > 0 && time.Now().After(lastPing.Add(time.Duration(PingEachMinute)*time.Minute)) {
		lastPing = time.Now()
		err := DB.DB().Ping()
		if err != nil {
			err = DB.Close()
			if err != nil {
				return DB, err
			}
			DB, err = Connect()
			if err != nil {
				return DB, err
			}
			return New()
		}
	}

	return DB, nil
}

func Connect() (*gorm.DB, error) {
	dbLink := GetLInk()
	var err error
	DB, err := gorm.Open("postgres", dbLink)
	DB.LogMode(Debug)
	if err != nil {
		return DB, err
	}
	DB.DB().SetMaxIdleConns(MaxIdleConns)
	DB.DB().SetMaxOpenConns(MaxOpenConns)

	return DB, nil
}

func GetLInk() string {
	dbLink := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", User, Pass, Host, Port, Name)

	return dbLink
}
