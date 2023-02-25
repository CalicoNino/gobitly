package datastore

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Gobitly struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect"`
	Gobitly  string `json:"gobitly gorm:"unique;not null"`
	Random   bool   `json:"random"`
	Clicked  int    `json:"clicked"`
}

func Setup(ip string, user string, password string, dbname string, port string, sslmode string) {

	dsn := fmt.Sprintf("host-%s user=%s password=%s dbname=%s port=%s sslmode=%s", ip, user, password, dbname, port, sslmode)
	fmt.Println(dsn)
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate((&Gobitly{}))
	if err != nil {
		fmt.Println(err)
	}
}

func GetAllGobitlies() ([]Gobitly, error) {
	var gobitlies []Gobitly

	tx := db.Find((&gobitlies))

	if tx.Error != nil {
		return gobitlies, tx.Error
	}

	return gobitlies, nil
}

func GetGobitly(id uint64) (Gobitly, error) {
	var gobitly Gobitly

	tx := db.Where("id = ?", id).First(&gobitly)
	if tx.Error != nil {
		return Gobitly{}, tx.Error
	}

	return gobitly, nil
}

func CreateGobitly(gobitly Gobitly) error {
	tx := db.Create(&gobitly)
	return tx.Error
}

func UpdateGobitly(gobitly Gobitly) error {
	tx := db.Save(&gobitly)
	return tx.Error
}

func DeleteGobitly(id uint64) error {
	tx := db.Unscoped().Delete(&Gobitly{}, id)
	return tx.Error
}

func FindByGobitly(url string) (Gobitly, error) {
	var gobitly Gobitly
	tx := db.Where("gobitly = ?", url).First(&gobitly)
	return gobitly, tx.Error
}
