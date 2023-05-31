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

func SimpleUpdate() {
	db, err := getDB()
	if err != nil {
		return
	}

	customer := Customer{}
	db.First(&customer)
	fmt.Println("Before updated. customer:", customer)
	customer.Name = "JanessaTechUpdated"
	result := db.Save(&customer)
	fmt.Println("After updated. customer =", customer)
	fmt.Println("result.RowsAffected =", result.RowsAffected)
	fmt.Println("result.Error", result.Error)

}

func BatchUpdates() {
	db, err := getDB()
	if err != nil {
		return
	}

	result := db.Model(&Customer{}).Where("id IN ?", []int{11, 12}).Updates(Customer{Name: "demo"})
	fmt.Println("result.RowsAffected =", result.RowsAffected)
	fmt.Println("result.Error", result.Error)
}

func EnableGlobalUpdates() {
	db, err := getDB()
	if err != nil {
		return
	}
	result := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&Customer{}).Updates(Customer{Name: "demo2"})
	fmt.Println("result.RowsAffected =", result.RowsAffected)
	fmt.Println("result.Error", result.Error)
}

func RawSqlForQuery() {
	db, err := getDB()
	if err != nil {
		return
	}
	var customer Customer
	result := db.Raw("select * from customers where id = ?", 10).Scan(&customer)
	fmt.Println("customer=", customer)
	fmt.Println("result.RowsAffected =", result.RowsAffected)
	fmt.Println("result.Error", result.Error)
}

func RawSqlForUpdates() {
	db, err := getDB()
	if err != nil {
		return
	}
	result := db.Exec("update customers set name = ? where id = ?", "demo3", 10)
	fmt.Println("result.RowsAffected =", result.RowsAffected)
	fmt.Println("result.Error", result.Error)
}

func DryRun() {
	db, err := getDB()
	if err != nil {
		return
	}

	var customer Customer
	stm := db.Session(&gorm.Session{DryRun: true}).First(&customer).Statement
	fmt.Println(stm.SQL.String())
	fmt.Println(stm.Vars)
}

func ToSQL() {
	db, err := getDB()
	if err != nil {
		return
	}
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&Customer{}).Where("id IN ?", []int{11, 12}).Updates(Customer{Name: "demo"})
	})
	fmt.Println(sql)
}

// Quick summary abour CRUD:
//
// db *gorm.DB
//
// db.Create(...)
//
// db.First(..)
//
// db.Find(..)
//
// db.Where(..)
//
// db.Save(...)
//
// db.Model(..).Where(..).Updates(..)
//
// db.Session(&gorm.Session{..}).Model(..)
//
// db.Table(..) // eg: db.Table("users as u")
//
// db.Raw(..)  //RawSql
//
// db.Exec(..) //RawSql
//
// db.Session(&gorm.Session{DryRun: true})  //for dry run
//
// print sql
//
// stm := db.Session(&gorm.Session{DryRun: true}).First(&customer).Statement
//
// fmt.Println(stm.SQL.String())
//
// similar to dry run
//
//	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
//			return tx.Model(&Customer{}).Where("id IN ?", []int{11, 12}).Updates(Customer{Name: "demo"})
//		})
//
// fmt.Println(sql)
func Main() {
	// For create
	//Insert()
	//MultipleInsert()
	//InsertWithAssociation() // I will take a look at it later on

	// For query
	//SimpleQuery()
	//QueryAll()
	//QueryByCondition()

	//For update
	//SimpleUpdate()
	//BatchUpdates()
	//EnableGlobalUpdates()

	//raw sql
	//RawSqlForQuery()
	//RawSqlForUpdates()

	//DryRun()
	ToSQL()

}
