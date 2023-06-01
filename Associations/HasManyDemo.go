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

// For Teacher with id being 1, return students whose age < 12
func FindAssociationsForOneMany() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}
	var students []model.Student

	error := db.Model(&model.Teacher{ID: 1}).Where("age < ?", 12).Association("Students").Find(&students)
	fmt.Println(students)

	return error
}
func HasManyDemo() {
	//HasManyGetAll()
	FindAssociationsForOneMany()

}
