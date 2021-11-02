package db

import (
	"fmt"
	_ "github.com/Binaretech/classroom-main/internal/config"
	"log"
	"os"
	"time"

	"github.com/Binaretech/classroom-main/internal/db/model"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func init() {
	for i := 0; i < 5; i++ {
		if err := OpenDatabase(); err != nil {
			logrus.Errorln("Error connecting to database", err)
			time.Sleep(5 * time.Second)
		} else {
			Migrate()
			return
		}
	}

	logrus.Fatal("Could not connect to database")
}

const (
	DESC = "DESC"
	ASC  = "ASC"
)

func OrderByString(value string) string {
	if value == DESC {
		return DESC
	}

	return ASC
}

var db *gorm.DB

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
		&model.Files{},
		&model.Assignment{},
		&model.Participant{},
		&model.Grade{},
	}
}

// Migrate run migrations to update the database
func Migrate() {
	db.AutoMigrate(Models()...)
}

// Drop and recreate the database
func Drop() error {
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
		viper.GetString("DATABASE_HOST"),
		viper.GetString("DATABASE_USER"),
		viper.GetString("DATABASE_PASS"),
		viper.GetString("DATABASE_NAME"),
		viper.GetString("DATABASE_PORT"),
	)
}

// CreateDatabase create a new database based on DATABASE_NAME config
func CreateDatabase() error {
	originalName := viper.GetString("DATABASE_NAME")
	viper.Set("DATABASE_NAME", "postgres")

	conn, err := gorm.Open(postgres.Open(connectionString()), &gorm.Config{})
	if err != nil {
		return err
	}

	if err = conn.Exec("CREATE DATABASE " + originalName).Error; err != nil {
		return err
	}

	viper.Set("DATABASE_NAME", originalName)

	return nil
}

func logLevel() logger.LogLevel {
	if viper.GetBool("debug_db") {
		return logger.Info
	}

	return logger.Error
}

// OpenDatabase opens the connection with database
func OpenDatabase() error {
	var err error
	db, err = gorm.Open(postgres.Open(connectionString()), &gorm.Config{
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
	return err
}

// Create a new record
func Create(value interface{}) *gorm.DB {
	return db.Create(value)
}

// CreateInBatches creates many records in batches
func CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return db.CreateInBatches(value, batchSize)
}

// Model add the given model to the query scope
func Model(model interface{}) *gorm.DB {
	return db.Model(model)
}

// First of the result
func First(model interface{}, conds ...interface{}) *gorm.DB {
	return db.First(model, conds...)
}

// Find results
func Find(model interface{}, conds ...interface{}) *gorm.DB {
	return db.Find(model, conds...)
}

// UpdateOrCreate update the model in case of conflict
func UpdateOrCreate(model interface{}) *gorm.DB {
	return db.Clauses(clause.OnConflict{UpdateAll: true}).Create(model)
}

// Joins join query
func Joins(query string, args ...interface{}) *gorm.DB {
	return db.Joins(query, args...)
}

// Where clause
func Where(query interface{}, args ...interface{}) *gorm.DB {
	return db.Where(query, args...)
}

// Association returns the association query
func Association(association string) *gorm.Association {
	return db.Association(association)
}

// Table add the given table to the query scope
func Table(table string, args ...interface{}) *gorm.DB {
	return db.Table(table, args...)
}

// Paginte paginate the result
func Paginate(pagination int, page int) *gorm.DB {
	return PaginateQuery(db, pagination, page)
}

// PaginateQuery paginate the result with a custom query
func PaginateQuery(query *gorm.DB, pagination, page int) *gorm.DB {
	if pagination == 0 {
		pagination = 10
	}

	if page == 0 {
		page = 1
	}

	return query.Offset((page - 1) * pagination).Limit(pagination)
}

// Offset add a offset to the query
func Offset(offset int) *gorm.DB {
	return db.Offset(offset)
}

// Limit add a limit to the query
func Limit(limit int) *gorm.DB {
	return db.Limit(limit)
}

// Delete delete the given model from the database with the given conditions (optional)
func Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return db.Delete(value, conds)
}

// Query returns the database client instance
func Query() *gorm.DB {
	return db
}

// Preload preload the given relations
func Preload(query string, args ...interface{}) *gorm.DB {
	return db.Preload(query, args...)
}

// Session returns a database session
func Session(config *gorm.Session) *gorm.DB {
	return db.Session(config)
}

// Exists check if the given query has any results
func Exists(query *gorm.DB) bool {
	var count int64
	query.Count(&count)

	return count > 0
}
