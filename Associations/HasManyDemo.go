package associations

import (
	"fmt"
	"time"

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

func HasManyCreate() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}

	teacher := model.Teacher{Name: "teacher2"}
	students := []model.Student{{Name: "student6", Age: 21},
		{Name: "student7", Age: 20},
		{Name: "student8", Age: 24}}
	teacher.Students = students
	result := db.Create(&teacher)
	fmt.Println("result.Error=", result.Error)
	fmt.Println("result.RowsAffected=", result.RowsAffected)
	return nil
}

func HasManyUpdate() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}
	var teachers []model.Teacher
	result := db.Model(&model.Teacher{}).Where("name = ?", "teacher2").Preload("Students").Find(&teachers)
	fmt.Println("result.Error=", result.Error)
	fmt.Println("result.RowsAffected=", result.RowsAffected)
	teachers[0].Students[0].Age = 30

	/** it doesn't work
	updates := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&teachers)
	fmt.Println("result.Error=", updates.Error)
	fmt.Println("result.RowsAffected=", updates.RowsAffected)**/
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
	cnt := db.Model(&model.Teacher{ID: 1}).Association("Students").Count()
	fmt.Println("cnt=", cnt)

	return error
}

// how to add association data
func AppendAssociations() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}
	//newStudent := &model.Student{Name: "student9", Age: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	newStudents := []model.Student{{Name: "student10", Age: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "student11", Age: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}}

	error := db.Model(&model.Teacher{ID: 1}).Association("Students").Append(newStudents)

	return error
}

// the replace doesn't actually delete data to be replaced, only emptify teacher_id column for the matched rows
func ReplaceAssociations() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}
	newStudents := []model.Student{{Name: "student12", Age: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "student13", Age: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}}
	error := db.Model(&model.Teacher{ID: 1}).Association("Students").Replace(newStudents)
	return error

}

func DeleteAssociations() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}
	error := db.Model(&model.Teacher{ID: 1}).Association("Students").Delete(model.Student{ID: 17}) // specify the id of student to be deleted
	// only the reference is deleted
	// | id | name      | age  | teacher_id | created_at          | updated_at          | deleted_at |
	// | 17 | student12 |    9 |       NULL | 2023-06-01 15:00:47 | 2023-06-01 15:00:47 | NULL       |
	return error
}

func DeleteAssociationsInReality() error {
	db, err := model.GetDB()
	if err != nil {
		return err
	}
	error := db.Model(&model.Teacher{ID: 1}).Association("Students").Unscoped().Clear()
	return error
}

func HasManyDemo() {
	//HasManyGetAll()
	//HasManyCreate()
	//HasManyUpdate()
	//FindAssociationsForOneMany()
	//AppendAssociations()
	//ReplaceAssociations()
	//DeleteAssociations()
	DeleteAssociationsInReality()
}
