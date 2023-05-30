package getstarted

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// sudo docker run -p 3306:3306 --name mysql
// -v /home/jane/docker/mysql/data:/data
// -e MYSQL_ROOT_PASSWORD=123456
// --restart=always
// -itd mysql:5.7.19
//
// sudo docker exec -it mysql /bin/bash
//
// mysql -uroot -p123456
//
// show databases;
//
// CREATE DATABASE gorm_demo;
// CREATE USER 'gorm'@'%' IDENTIFIED BY 'gorm_demo';
// SELECT USER, host from mysql.user;
// GRANT ALL on *.* TO 'gorm'@'%';
// GRANT super on *.* to 'gorm'@'%';
// GRANT show view on *.* to 'gorm'@'%';
//
// mysql -ugorm -pgorm_demo gorm_demo
//
// show tables
func FirstGormDemo() {
	dsn := "gorm:gorm_demo@tcp(192.168.0.200:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // you could open mysql in this way
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("cannot connect to mysql")
	} else {
		fmt.Println("mysql is connected")
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("failed to return DB()")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println(db)
}
