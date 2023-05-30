package crud

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Customer struct {
	ID    uint   `gorm:"column:id"`
	Name  string `gorm:"column:name;default:JanessaTech"`
	Email string `gorm:"column:email"`
	Age   uint8  `gorm:"column:age;default:10"` // provided default value for age
	//CreditCard   CreditCard
	Birthday     time.Time      `gorm:"column:birthday"`
	MemberNumber sql.NullString `gorm:"column:member_number"`
	ActivatedAt  sql.NullTime   `gorm:"column:activated_at"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
}

type CreditCard struct {
	gorm.Model
	Number string `gorm:"column:number"`
	UserID uint   `gorm:"column:user_id"`
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

// CREATE TABLE customers
// (
//
//		`id`              INT NOT NULL AUTO_INCREMENT,
//		`name`            VARCHAR(125),
//		`email`           VARCHAR(125),
//		`age`             TINYINT,
//		`birthday`        TIMESTAMP NULL,
//		`member_number`   VARCHAR(125),
//		`activated_at`    TIMESTAMP NULL,
//		`created_at`      TIMESTAMP NULL,
//		`updated_at`      TIMESTAMP NULL,
//	 PRIMARY KEY (`id`)
//
// ) ENGINE = InnoDB
//
//	DEFAULT CHARSET = UTF8MB4;
func Insert() error {
	db, err := getDB()
	if err != nil {
		return err
	}
	customer := Customer{Name: "JanessaTech", Age: 30, Birthday: time.Now()}
	result := db.Create(&customer)
	fmt.Println("customer.ID =", customer.ID)
	fmt.Println("result.Error =", result.Error)
	fmt.Println("result.RowsAffected =", result.RowsAffected)
	return nil
}

func MultipleInsert() {
	db, err := getDB()
	if err != nil {
		return
	}
	customers := []Customer{
		{Name: "JanessaTech1", Email: "demo1@gmail.com", Birthday: time.Now()},
		{Name: "JanessaTech2", Email: "demo2@gmail.com", Birthday: time.Now()},
		{Name: "JanessaTech3", Email: "demo3@gmail.com", Birthday: time.Now()}}
	result := db.Create(&customers)
	for _, customer := range customers {
		fmt.Println("customer.ID =", customer.ID)
	}

	fmt.Println("result.Error =", result.Error)
	fmt.Println("result.RowsAffected =", result.RowsAffected)
}

func InsertWithAssociation() {
	db, err := getDB()
	if err != nil {
		return
	}
	custom := Customer{Name: "JanessaTech4",
		Age:      34,
		Birthday: time.Now(),
		//CreditCard: CreditCard{Number: "411111111111"}
	}
	result := db.Create(&custom)
	fmt.Println("result.Error =", result.Error)
	fmt.Println("result.RowsAffected =", result.RowsAffected)

}

func SimpleQuery() {
	db, err := getDB()
	if err != nil {
		return
	}
	customer := Customer{}
	result := db.First(&customer)
	fmt.Println("customer = ", customer)
	fmt.Println("result.RowsAffected =", result.RowsAffected)
	fmt.Println("result.Error", result.Error)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("customer is found")
	} else {
		fmt.Println("customer is not found")
	}
}

func QueryAll() {
	db, err := getDB()
	if err != nil {
		return
	}
	customers := []Customer{}
	result := db.Find(&customers)
	fmt.Println("customers =", customers)
	fmt.Println("result.RowsAffected =", result.RowsAffected)
	fmt.Println("result.Error", result.Error)
}

func QueryByCondition() {
	db, err := getDB()
	if err != nil {
		return
	}

	customer := Customer{}
	result := db.Where("email = ?", "demo3@gmail.com").First(&customer)
	fmt.Println("customer =", customer)
	fmt.Println("result.RowsAffected =", result.RowsAffected)
	fmt.Println("result.Error", result.Error)
}

func Main() {
	// For create
	//Insert()
	//MultipleInsert()
	//InsertWithAssociation() // I will take a look at it later on

	// For query
	//SimpleQuery()
	//QueryAll()
	QueryByCondition()

}
