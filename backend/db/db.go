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

type Join struct {
	TableJoin   string
	IDTable     string
	IDTableJoin string
}

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

func FetchJoin(table string, field string, params []string, joins []Join) string {

	values := Select(params)

	q := QuerySelect(table, values)

	for _, j := range joins {
		q += LeftJoin(j.TableJoin, j.IDTable, j.IDTableJoin)
	}

	q += Where(field)

	return q
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

func QuerySelect(table string, values string) string {
	return fmt.Sprintf(`SELECT %s FROM %s`, values, table)
}

func LeftJoin(tableJoin string, idTable string, idTableJoin string) string {
	return fmt.Sprintf(" LEFT JOIN %s ON %s=%s ", tableJoin, idTable, idTableJoin)
}

func GetFilter(filters []string, where bool) string {
	q := ""

	if len(filters) > 0 {
		if where {
			q += " WHERE "
		} else {
			q += " AND "
		}
		for i, filter := range filters {
			q += filter
			if i != len(filters)-1 {
				q += " AND "
			}
		}
	}

	return q
}

func QueryPagingJoin(table string, field string, isFirst bool, selects []string, joins []Join, filters []string) string {
	values := Select(selects)

	q := QuerySelect(table, values)

	if len(joins) > 0 {
		for _, j := range joins {
			q += LeftJoin(j.TableJoin, j.IDTable, j.IDTableJoin)
		}
	}

	if isFirst {
		q += GetFilter(filters, true)
		q += OrderByLimit(field, "DESC")
	} else {
		q += WhereLess(field)
		q += GetFilter(filters, false)
		q += OrderByLimit(field, "DESC")
	}

	return q

}

func QueryPagingJoinSearch(table string, field string, isFirst bool, selects []string, joins []Join, search string, filters []string) string {
	values := Select(selects)

	q := QuerySelect(table, values)
	if len(joins) > 0 {
		for _, j := range joins {
			q += LeftJoin(j.TableJoin, j.IDTable, j.IDTableJoin)
		}
	}

	qs := `ILIKE '%' || ? || '%'`

	q = fmt.Sprintf(`%s WHERE %s %s`, q, search, qs)

	if isFirst {
		q += GetFilter(filters, false)
		q += OrderByLimit(field, "DESC")
	} else {
		q += WhereLessAnd(field)
		q += GetFilter(filters, false)
		q += OrderByLimit(field, "DESC")
	}

	return q

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

func And() string {
	return " AND "
}
