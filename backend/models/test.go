package models

import (
	"database/sql"
	"strconv"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

type Test struct {
	ID        int64     `json:"id"`
	Nama      string    `json:"nama" validate:"required"`
	Alamat    string    `json:"alamat" validate:"required"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func (m *Model) CreateTest(r Test) error {

	createdAt := constants.GetDatetimeNow()

	sqlX := `INSERT INTO tests (nama, alamat, created_at, updated_at) VALUES (?, ?, ?, ?)`

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Nama, r.Alamat, createdAt, createdAt)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) PaginatorTest(params ...string) ([]Test, error) {
	results := []Test{}
	action := params[0]
	no, _ := strconv.Atoi(params[1])
	limit, _ := strconv.Atoi(params[2])

	search := ""
	filter := ""
	if len(params) >= 4 {
		search = params[3]
		filter = params[4]
	}

	sqlX := `SELECT t.id, t.nama, t.alamat, t.created_at, t.updated_at FROM tests AS t`

	if action == "next" {
		sqlX += db.WhereOrderBy("t.id", "<")

		if search != "" {
			sqlX += db.Filter(filter)
		}
		sqlX += db.OrderBy("t.id", "DESC")

	} else {
		if search != "" {
			sqlX += db.WhereFilter(filter)
		}
		sqlX += db.OrderBy("u.id", "DESC")
	}

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return results, err
	}

	defer stmt.Close()

	rows, err := new(sql.Rows), nil

	if action == "next" {
		rows, err = stmt.Query(no, limit)
		if search != "" {
			rows, err = stmt.Query(no, search, limit)
		}
	} else {
		rows, err = stmt.Query(limit)
		if search != "" {
			rows, err = stmt.Query(search, limit)
		}
	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return results, err
	}

	for rows.Next() {
		r := Test{}

		err = rows.Scan(&r.ID, &r.Nama, &r.Alamat, &r.CreatedAt, &r.UpdatedAt)
		results = append(results, r)
	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return results, err
	}

	return results, nil

}
