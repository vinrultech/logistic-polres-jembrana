package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	TableUsers = "log_users"
)

//InitDb -> fungsi untuk koneksi ke database
func InitDb(config *viper.Viper) *sqlx.DB {

	port := config.GetString("db_port")
	user := config.GetString("db_user")
	name := config.GetString("db_name")
	password := config.GetString("db_password")
	host := config.GetString("db_host")

	conn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, name, host, port)

	db, err := sqlx.Open("postgres", conn)

	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(2 * time.Minute)
	db.SetMaxOpenConns(95)

	if err != nil {
		panic(err)
		//loggers.Log.Errorln(err.Error())
	}

	if db == nil {
		panic("db nil")
		//loggers.Log.Errorln("Database null")
	}

	return db
}

func Filter(field string) string {
	return ` AND ` + field + ` LIKE '%' || ? || '%'`
}

func WhereFilter(field string) string {
	return ` WHERE ` + field + ` LIKE '%' || ? || '%'`
}

func OrderBy(field string, order string) string {
	return ` ORDER BY ` + field + ` ` + order + ` LIMIT ?`
}

func WhereOrderBy(field string, order string) string {
	return ` WHERE ` + field + ` ` + order + ` ?`
}

func WhereOrderAnd(field string, order string) string {
	return ` AND ` + field + ` ` + order + ` ?`
}

func WhereAnd(field string) string {
	return ` AND ` + field + ` =  ?`
}

func Where(field string) string {
	return ` WHERE ` + field + ` =  ?`
}

func WhereOr(field string) string {
	return ` OR ` + field + ` =  ?`
}

func WhereIn(field string) string {
	return ` WHERE ` + field + ` IN  (?)`
}
