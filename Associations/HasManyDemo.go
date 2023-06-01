package associations

import (
	"fmt"

	"github.com/hi-supergirl/go-learning-gorm/Associations/model"
)

func HasManyGetAll() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}
	var teachers []model.Teacher
	result := db.Model(&model.Teacher{}).Preload("Students").Find(&teachers)
	fmt.Println("result.Error=", result.Error)
	fmt.Println("result.RowsAffected=", result.RowsAffected)
	return nil
}

func HasManyDemo() {
	HasManyGetAll()

}
