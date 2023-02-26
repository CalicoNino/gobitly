package datastore

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
