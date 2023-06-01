package model

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	ID           uint           `gorm:"column:id;primary_key"`
	Name         string         `gorm:"column:name;default:JanessaTech"`
	Email        string         `gorm:"column:email"`
	Age          uint8          `gorm:"column:age;default:10"` // provided default value for age
	CreditCard   CreditCard     `gorm:"foreignKey:EmpoyeeID"`  // for the case to show has-one
	Birthday     time.Time      `gorm:"column:birthday"`
	MemberNumber sql.NullString `gorm:"column:member_number"`
	ActivatedAt  sql.NullTime   `gorm:"column:activated_at"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
}
type CreditCard struct {
	ID        uint      `gorm:"column:id;primary_key"`
	Number    string    `gorm:"column:number"`
	EmpoyeeID uint      `gorm:"column:empoyee_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Shape struct {
	gorm.Model
	Name      string `gorm:"column:name;default:shape"`
	ChildID   int    `gorm:"column:child_id"`
	ChildType string `gorm:"column:child_type"`
}
type Circle struct {
	gorm.Model
	Name  string `gorm:"column:name;default:circle"`
	Shape Shape  `gorm:"polymorphic:Child;"`
}

type Square struct {
	gorm.Model
	Name  string `gorm:"column:name;default:square"`
	Shape Shape  `gorm:"polymorphic:Child;"`
}

type Teacher struct {
	gorm.Model
	Name     string    `gorm:"column:name"`
	Students []Student `gorm:"foreignKey:TeacherID"`
}

type Student struct {
	gorm.Model
	Name      string `gorm:"column:name"`
	Age       int    `gorm:"column:age"`
	TeacherID uint   `gorm:"column:teacher_id"`
}

func GetDB() (*gorm.DB, error) {
	dsn := "gorm:gorm_demo@tcp(192.168.0.200:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // you could open mysql in this way
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, errors.New("cannot connect to mysql")
	} else {
		fmt.Println("mysql is connected")
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New("failed to return DB()")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
