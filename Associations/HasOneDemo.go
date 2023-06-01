package associations

import (
	"fmt"
	"time"

	"github.com/hi-supergirl/go-learning-gorm/Associations/model"
)

// Retrieve Employee list with eager loading credit card
func HasOneGetAll() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}

	var employees []model.Employee
	result := db.Model(&model.Employee{}).Preload("CreditCard").Find(&employees)
	fmt.Println(employees)
	fmt.Println("result.Error=", result.Error)
	fmt.Println("result.RowsAffected=", result.RowsAffected)
	return nil
}

func HasOneCreate() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}

	creditCard := model.CreditCard{Number: "22222"}
	employee := model.Employee{Name: "Jianwei1", Email: "jianwei1@gmail.com", CreditCard: creditCard, Birthday: time.Now()}
	result := db.Create(&employee)
	fmt.Println("result.Error=", result.Error)
	fmt.Println("result.RowsAffected=", result.RowsAffected)
	//db.Save(&employee)
	return nil
}

// The case to find creditcard for Employee with id=2
func FindAssociationsForHasOne() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}
	employee := model.Employee{ID: 2}
	var creditCards []model.CreditCard
	result := db.Model(&employee).Association("CreditCard").Find(&creditCards)
	fmt.Println("creditCard=", creditCards)
	fmt.Println("result =", result)

	return nil
}

func PolymorphicDemo() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}
	db.Create(&model.Circle{Name: "circle1", Shape: model.Shape{Name: "shape1"}})
	// insert into circles(`name`) values('circle1') -- new id is 1
	// insert into shapes(`name`, `child_id`,`child_type`) values('shape1', 1, 'circles')
	db.Create(&model.Square{Name: "square1", Shape: model.Shape{Name: "shape2"}})
	// insert into square(`name`) values('square1') -- new id is 2
	// insert into shapes(`name`, `child_id`,`child_type`) values('shape2', 2, 'squares')
	return nil
}

func HasOneDemo() {
	//HasOneGetAll()
	//PolymorphicDemo()
	//HasOneCreate()
	FindAssociationsForHasOne()
}
