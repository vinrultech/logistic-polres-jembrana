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

func Insert(table string, params ...string) string {
	var field string
	var values string

	for i, param := range params {
		field += param
		values += "?"
		if i != len(params)-1 {
			field += ","
			values += ","
		}
	}

	return fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, table, field, values)
}

func Update(table string, field string, params ...string) string {
	var values string

	for i, param := range params {
		values += param + "=?"
		if i != len(params)-1 {
			values += ","
		}
	}

	return fmt.Sprintf(`UPDATE %s SET %s WHERE %s=?`, table, values, field)
}

func Delete(table string, field string) string {
	return fmt.Sprintf(`DELETE FROM %s WHERE %s=?`, table, field)
}

func Select(params []string) string {
	var values string

	for i, param := range params {
		values += param
		if i != len(params)-1 {
			values += ","
		}
	}
	return values
}

func QueryAll(table string, params []string) string {

	values := Select(params)

	return fmt.Sprintf(`SELECT %s FROM %s`, values, table)
}

func Fetch(table string, field string, params []string) string {

	values := Select(params)

	return fmt.Sprintf(`SELECT %s FROM %s %s`, values, table, Where(field))
}

func QueryAllWhere(table string, field string, params []string) string {

	values := Select(params)

	return fmt.Sprintf(`SELECT %s FROM %s %s`, values, table, Where(field))
}

func QueryAllOrderBy(table string, field string, order string, params []string) string {
	values := Select(params)

	return fmt.Sprintf(`SELECT %s FROM %s %s`, values, table, OrderBy(field, order))
}

func QueryAllWhereOrderBy(table string, fieldWhere string, fieldOrder string, order string, params []string) string {
	values := Select(params)

	return fmt.Sprintf(`SELECT %s FROM %s %s %s`, values, table, Where(fieldWhere), OrderBy(fieldOrder, order))
}

func QueryPagingSurat(table string, values string, first bool, field string, params ...string) string {
	q := "SELECT %s FROM %s"
	q = fmt.Sprintf(q, values, table)

	if !first {
		q += WhereLess(field)
	}

	if len(params) > 0 {
		if !first {
			q += WhereBetweenAnd(params[0])
		} else {
			q += WhereBetween(params[0])
		}
	}

	q += OrderByLimit(field, "DESC")

	return q
}

func QueryPagingSuratWithFilter(table string, values string, first bool, field string, filter string, params ...string) string {
	q := "SELECT %s FROM %s WHERE %s"
	q = fmt.Sprintf(q, values, table, filter)

	if !first {
		q += WhereLessAnd(field)
	}

	if len(params) > 0 {
		q += WhereBetweenAnd(params[0])
	}

	q += OrderByLimit(field, "DESC")

	return q
}

func QueryPagingSuratSearch(table string, values string, first bool, field string, search string, params ...string) string {
	s := `ILIKE '%' || ? || '%'`
	q := "SELECT %s FROM %s WHERE %s %s"
	q = fmt.Sprintf(q, values, table, search, s)

	if !first {
		q += WhereLessAnd(field)
	}

	if len(params) > 0 {
		q += WhereBetweenAnd(params[0])
	}

	q += OrderByLimit(field, "DESC")

	return q
}

func QueryPagingSuratSearchWithFilter(table string, values string, first bool, field string, search string, filter string, params ...string) string {

	s := `ILIKE '%' || ? || '%'`
	q := "SELECT %s FROM %s WHERE %s %s AND %s"
	q = fmt.Sprintf(q, values, table, search, s, filter)

	if !first {
		q += WhereLessAnd(field)
	}

	if len(params) > 0 {
		q += WhereBetweenAnd(params[0])
	}

	q += OrderByLimit(field, "DESC")

	return q
}

func QueryPaging(table string, field string, isFirst bool, params []string) string {
	values := Select(params)

	if isFirst {
		return QueryPagingFirst(table, field, values)
	} else {
		return QueryPagingLess(table, field, values)
	}

}

func QueryPagingSearch(table string, field string, isFirst bool, search string, params []string) string {
	values := Select(params)

	if isFirst {
		return QueryPagingFirstSearch(table, field, values, search)
	} else {
		return QueryPagingLessSearch(table, field, values, search)
	}

}

func QueryPagingLess(table string, field string, values string) string {

	return fmt.Sprintf(`SELECT %s FROM %s %s %s`, values, table, WhereLess(field), OrderByLimit(field, "DESC"))
}

func QueryPagingFirst(table string, field string, values string) string {
	return fmt.Sprintf(`SELECT %s FROM %s %s`, values, table, OrderByLimit(field, "DESC"))
}

func QueryPagingLessSearch(table string, field string, values string, search string) string {
	s := `ILIKE '%' || ? || '%'`
	return fmt.Sprintf(`SELECT %s FROM %s %s AND %s %s %s`, values, table, WhereLess(field), search, s, OrderByLimit(field, "DESC"))
}

func QueryPagingFirstSearch(table string, field string, values string, search string) string {
	s := `ILIKE '%' || ? || '%'`
	return fmt.Sprintf(`SELECT %s FROM %s WHERE %s %s %s`, values, table, search, s, OrderByLimit(field, "DESC"))
}

func Filter(field string) string {
	return ` AND ` + field + ` LIKE '%' || ? || '%'`
}

func WhereFilter(field string) string {
	return ` WHERE ` + field + ` LIKE '%' || ? || '%'`
}

func OrderByLimit(field string, order string) string {
	return ` ORDER BY ` + field + ` ` + order + ` LIMIT ?`
}

func OrderBy(field string, order string) string {
	return ` ORDER BY ` + field + ` ` + order + ``
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

func WhereLess(field string) string {
	return ` WHERE ` + field + ` <  ?`
}

func WhereLessAnd(field string) string {
	return ` AND ` + field + ` <  ?`
}

func WhereOr(field string) string {
	return ` OR ` + field + ` =  ?`
}

func WhereIn(field string) string {
	return ` WHERE ` + field + ` IN  (?)`
}

func WhereBetween(field string) string {
	return fmt.Sprintf(" WHERE (%s BETWEEN ? AND ?)", field)
}

func WhereBetweenAnd(field string) string {
	return fmt.Sprintf(" AND (%s BETWEEN ? AND ?)", field)
}
