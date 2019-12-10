package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

//DB Insatncia do banco de dados
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig constroi as configurações
func BuildDBConfig() *DBConfig {

	dbHost, _ := getenvStr("DB_HOST")
	dbPort, _ := getenvInt("DB_PORT")
	dbUser, _ := getenvStr("DB_USER")
	dbName, _ := getenvStr("DB_NAME")
	dbPassword, _ := getenvStr("DB_PASSWORD")

	dbConfig := DBConfig{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		DBName:   dbName,
		Password: dbPassword,
	}
	return &dbConfig
}

// DbURL gera a uri de acesso
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		// mysql
		// "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		// postgresql
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

// ErrEnvVarEmpty a variavel não existe ou está vazia
var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func getenvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, ErrEnvVarEmpty
	}
	return v, nil
}

func getenvInt(key string) (int, error) {
	s, err := getenvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func getenvBool(key string) (bool, error) {
	s, err := getenvStr(key)
	if err != nil {
		return false, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}
