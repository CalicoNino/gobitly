package datastore

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

type Gobitly struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect"`
	Gobitly  string `json:"gobitly gorm:"unique;not null"`
	Random   bool   `json:"random"`
	Clicked  int    `json:"clicked"`
}

func Init(ip string, user string, password string, dbname string, port string, sslmode string) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", ip, user, password, dbname, port, sslmode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	autoMigration()
}

func Close() {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func autoMigration() {
	err = db.AutoMigrate((&Gobitly{}))
	if err != nil {
		fmt.Println(err)
	}
}
