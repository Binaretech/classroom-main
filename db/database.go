package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/Binaretech/classroom-main/config"

	"github.com/Binaretech/classroom-main/db/model"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DESC = "DESC"
	ASC  = "ASC"
)

func Connect() (*gorm.DB, error) {
	for i := 0; i < 5; i++ {
		if db, err := OpenDatabase(); err != nil {
			logrus.Errorln("Error connecting to database", err)
			time.Sleep(5 * time.Second)
		} else {
			Migrate(db)
			return db, nil
		}
	}

	return nil, errors.New("Error connecting to database")
}

func OrderByString(value string) string {
	if value == DESC {
		return DESC
	}

	return ASC
}

// Models returns all the registered models
func Models() []interface{} {
	return []interface{}{
		&model.User{},
		&model.Class{},
		&model.Section{},
		&model.Material{},
		&model.Module{},
		&model.EvaluationDate{},
		&model.Post{},
		&model.File{},
		&model.Assignment{},
		&model.Participant{},
		&model.Grade{},
	}
}

// Migrate run migrations to update the database
func Migrate(db *gorm.DB) {
	db.AutoMigrate(Models()...)
}

// Drop and recreate the database
func Drop(db *gorm.DB) error {
	return db.Exec(`
	DROP SCHEMA public CASCADE;
	CREATE SCHEMA public;

	GRANT ALL ON SCHEMA public TO postgres;
	GRANT ALL ON SCHEMA public TO public;
`).Error
}

// connectionString returns the connection string based on the environment
func connectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
	)
}

// CreateDatabase create a new database based on DATABASE_NAME config
func CreateDatabase() error {
	originalName := viper.GetString("DB_NAME")
	viper.Set("DB_NAME", "postgres")

	conn, err := gorm.Open(postgres.Open(connectionString()), &gorm.Config{})
	if err != nil {
		return err
	}

	if err = conn.Exec("CREATE DATABASE " + originalName).Error; err != nil {
		return err
	}

	viper.Set("DB_NAME", originalName)

	return nil
}

func logLevel() logger.LogLevel {
	if viper.GetBool("debug_db") {
		return logger.Info
	}

	return logger.Error
}

// OpenDatabase opens the connection with database
func OpenDatabase() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(connectionString()), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logLevel(),
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
}

// Joins join query
func PaginatedResponse[T any]() T {
	var resources T
	return resources
}

// Paginte paginate the result
func Paginate(db *gorm.DB, pagination int, page int) *gorm.DB {
	return PaginateQuery(db, pagination, page)
}

// PaginateQuery paginate the result with a custom query
func PaginateQuery(query *gorm.DB, limit, page int) *gorm.DB {
	return query.Offset((page - 1) * limit).Limit(limit)
}

// Exists check if the given query has any results
func Exists(query *gorm.DB) bool {
	var count int64
	query.Count(&count)

	return count > 0
}

func Get[T any](query *gorm.DB) (T, error) {
	var resource T

	if err := query.Find(&resource).Error; err != nil {
		return resource, err
	}

	return resource, nil
}

func GetOne[T any](query *gorm.DB) (T, error) {
	var resource T

	if err := query.First(&resource).Error; err != nil {
		return resource, err
	}

	return resource, nil
}
