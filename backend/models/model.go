package models

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

type Model struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *Model {
	m := new(Model)
	m.Db = db
	return m
}

//ToDateSQLDate - konversi string tanggal ke postgresql tanggal
func ToDateSQLDate(t string) string {
	tgl := strings.Split(t, "-")
	tanggal := strings.Join([]string{tgl[2], tgl[1], tgl[0]}, "-")
	return tanggal
}

//ToSQLDateString - konversi dari postgresql ke tanggal string
func ToSQLDateString(t string) string {
	//2019-01-27T00:00:00Z
	z := strings.Split(t, "T")
	tgl := strings.Split(z[0], "-")
	tanggal := strings.Join([]string{tgl[2], tgl[1], tgl[0]}, "-")
	return tanggal
}

func (m *Model) MaxID(table string, sql string, search string, filter string) (int, error) {

	sqlX := `SELECT max(id) as id from ` + table

	if sql != "" {
		sqlX += sql
	}

	id, err := m.GetID(sqlX, search, filter)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return 0, err
	}

	return int(id), nil

}

func (m *Model) MinID(table string, sql string, search string, filter string) (int, error) {

	sqlX := `SELECT min(id) as id from ` + table

	if sql != "" {
		sqlX += sql
	}

	id, err := m.GetID(sqlX, search, filter)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return 0, err
	}

	return int(id), nil

}

func (m *Model) GetID(sqlX string, search string, filter string) (int64, error) {
	var id null.Int

	if search != "" {
		sqlX += db.WhereFilter(filter)
	}

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return id.ValueOrZero(), err
	}

	defer stmt.Close()

	row := stmt.QueryRow()

	if search != "" {
		row = stmt.QueryRow(search)
	}

	err = row.Scan(&id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return id.ValueOrZero(), err
	}

	if id.Valid {
		return id.Int64, nil
	} else {
		return id.ValueOrZero(), nil
	}

}
