package pgdb

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	Host            string
	Port            int
	Name            string
	User            string
	Pass            string
	SslMode         string
	TimeZone        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func New(c Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s TimeZone=%s", c.User, c.Pass, c.Name, c.Host, c.Port, c.SslMode, c.TimeZone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	tx, err := db.DB()
	if err != nil {
		return
	}
	tx.SetMaxIdleConns(c.MaxIdleConns)
	tx.SetMaxOpenConns(c.MaxOpenConns)
	tx.SetConnMaxLifetime(c.ConnMaxLifetime)

	return
}

//
//var User string
//var Pass string
//var Host string
//var Port string
//var Name string
//
//var Debug bool
//var PingEachMinute int
//var MaxIdleConns int
//var MaxOpenConns int
//
//var lastPing time.Time
//
//var DB *gorm.DB
//
//func init() {
//	Debug = false
//	PingEachMinute = 10
//	MaxIdleConns = 10
//	MaxOpenConns = 100
//	lastPing = time.Now()
//}
//
//// Get instance
//func New() (*gorm.DB, error) {
//	if DB == nil {
//		var err error
//		DB, err = Connect()
//		if err != nil {
//			return DB, err
//		}
//		return New()
//	}
//	if PingEachMinute > 0 && time.Now().After(lastPing.Add(time.Duration(PingEachMinute)*time.Minute)) {
//		lastPing = time.Now()
//		err := DB.DB().Ping()
//		if err != nil {
//			DB.Close()
//			DB, err = Connect()
//			if err != nil {
//				return DB, err
//			}
//			return New()
//		}
//	}
//
//	return DB, nil
//}
//
//// Connect to DB
//func Connect() (*gorm.DB, error) {
//	DB, err := gorm.Open("postgres", GetLInk())
//	if err != nil {
//		return DB, err
//	}
//	DB.LogMode(Debug)
//	DB.DB().SetMaxIdleConns(MaxIdleConns)
//	DB.DB().SetMaxOpenConns(MaxOpenConns)
//
//	return DB, err
//}
//
//// Format link
//func GetLInk() string {
//	dbLink := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", User, Pass, Host, Port, Name)
//
//	return dbLink
//}
