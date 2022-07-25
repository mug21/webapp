package config

import (
	"fmt"
	"webapp/e-commerce/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDbConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "db",
		Port:     5432,
		User:     "postgres",
		Password: "mypassword",
		DBName:   "postgres",
	}
	return &dbConfig
}

func GetDbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port)
}

func ConnectToDB() {
	dburl := GetDbURL(BuildDbConfig())
	var err error
	// Connect to database
	DB, err = gorm.Open(postgres.Open(dburl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Connected to DB successfullly!!")
	DB.AutoMigrate(&models.User{}, &models.Order{})
	// Check if foreign key constraint present or not
	if !DB.Migrator().HasConstraint(&models.User{}, "fk_users_orders") {
		DB.Migrator().CreateConstraint(&models.User{}, "fk_users_orders")
	}
}
