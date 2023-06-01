package associations

import (
	"fmt"

	"github.com/hi-supergirl/go-learning-gorm/Associations/model"
	"gorm.io/gorm"
)

type Human struct {
	gorm.Model
	Name      string     `gorm:"column:name"`
	Languages []Language `gorm:"many2many:human_languages;"`
}
type Language struct {
	gorm.Model
	Name string `gorm:"column:name"`
}

func GetAll() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}
	//db.AutoMigrate(&Human{})
	var humen []Human
	result := db.Model(&Human{}).Preload("Languages").Find(&humen)
	fmt.Println("result.Error=", result.Error)
	fmt.Println("result.RowsAffected=", result.RowsAffected)
	return nil
}

func ManyToManyDemo() {
	GetAll()
}
