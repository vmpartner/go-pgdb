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

func init() {
	Debug = false
	PingEachMinute = 10
	MaxIdleConns = 10
	MaxOpenConns = 100
	lastPing = time.Now()
}

// Get instance
func New() (*gorm.DB, error) {
	if DB == nil {
		var err error
		DB, err = Connect()
		if err != nil {
			return DB, err
		}
		return New()
	}
	if PingEachMinute > 0 && time.Now().After(lastPing.Add(time.Duration(PingEachMinute)*time.Minute)) {
		lastPing = time.Now()
		err := DB.DB().Ping()
		if err != nil {
			DB.Close()
			DB, err = Connect()
			if err != nil {
				return DB, err
			}
			return New()
		}
	}

	return DB, nil
}

// Connect to DB
func Connect() (*gorm.DB, error) {
	DB, err := gorm.Open("postgres", GetLInk())
	if err != nil {
		return DB, err
	}
	DB.LogMode(Debug)
	DB.DB().SetMaxIdleConns(MaxIdleConns)
	DB.DB().SetMaxOpenConns(MaxOpenConns)

	return DB, err
}

// Format link
func GetLInk() string {
	dbLink := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", User, Pass, Host, Port, Name)

	return dbLink
}
