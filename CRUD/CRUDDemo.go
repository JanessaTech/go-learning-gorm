package crud

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Testuser struct {
	ID           uint           `gorm:"column:id"`
	Name         string         `gorm:"column:name"`
	Email        string         `gorm:"column:email"`
	Age          uint8          `gorm:"column:age"`
	Birthday     time.Time      `gorm:"column:birthday"`
	MemberNumber sql.NullString `gorm:"column:member_number"`
	ActivatedAt  sql.NullTime   `gorm:"column:activated_at"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
}

func getDB() (*gorm.DB, error) {
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

// CREATE TABLE testusers
// (
//
//	`id`              INT,
//	`name`            VARCHAR(125),
//	`email`           VARCHAR(125),
//	`age`             TINYINT,
//	`birthday`        TIMESTAMP NULL,
//	`member_number`   VARCHAR(125),
//	`activated_at`    TIMESTAMP NULL,
//	`created_at`      TIMESTAMP NULL,
//	`updated_at`      TIMESTAMP NULL
//
// ) ENGINE = InnoDB
//
//	DEFAULT CHARSET = UTF8MB4;
func Create() error {
	db, err := getDB()
	if err != nil {
		return err
	}
	user := Testuser{Name: "JanessaTech", Age: 30, Birthday: time.Now()}
	result := db.Create(&user)
	fmt.Println("user.ID =", user.ID)
	fmt.Println("result.Error =", result.Error)
	fmt.Println("result.RowsAffected =", result.RowsAffected)
	return nil
}

func Main() {
	err := Create()
	if err == nil {
		fmt.Println("Created user successfully")
	}
}
