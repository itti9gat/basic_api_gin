package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"iiujapp.tech/susaday/handle"
	"iiujapp.tech/susaday/repo"
	"iiujapp.tech/susaday/route"
)

var (
	cfg config
)

type config struct {
	Port string `json:"port"`
	SQL  struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		DBname   string `json:"dbname"`
	} `json:"sql"`
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Panic(fmt.Errorf("error config file: %s", err))
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panic(fmt.Errorf("error config file: %s", err))
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	log.Println("connecting to database...")
	db, err := connectSQLServer(cfg.SQL.User, cfg.SQL.Password, cfg.SQL.Host, cfg.SQL.Port, cfg.SQL.DBname)
	if err != nil {
		log.Panicf("unable to connect to SQL server: %s", err)
	}
	defer db.Close()

	rc := repo.NewCategory(db)
	rp := repo.NewProduct(db)

	r := gin.Default()
	h := handle.NewHandler(rc, rp)
	route.NewRoute(r, h)

	log.Printf("Listening and serving HTTP on %s\n", cfg.Port)
	r.Run(cfg.Port)
}

func connectSQLServer(user, password, host, port, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, dbName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
