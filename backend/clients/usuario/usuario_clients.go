package clients

import (
	"Backend/dao"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLClient struct {
	DB *gorm.DB
}

func NewMySQLClient() *MySQLClient {
	dsnFormat := "%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local"
	dsn := fmt.Sprintf(dsnFormat, "root", "root", "127.0.0.1", 3306, "backend")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %w", err))
	}

	for _, table := range []interface{}{
		dao.User{},
		//		dao.Activity{},
		//		dao.TimeSlot{},
		//		dao.Inscription{},
	} {
		if err := db.AutoMigrate(&table); err != nil {
			panic(fmt.Errorf("error migrating table: %w", err))
		}
	}

	return &MySQLClient{
		DB: db,
	}
}

func (c *MySQLClient) GetUserByUsername(username string) (dao.User, error) {
	var userDAO dao.User
	txn := c.DB.First(&userDAO, "username = ?", username)
	if txn.Error != nil {
		return dao.User{}, fmt.Errorf("error getting user: %w", txn.Error)
	}

	return userDAO, nil
}
