package models

import (
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

var tableKategori = "log_kategoris"

type Kategori struct {
	ID        int64     `json:"id"`
	Kode      string    `json:"kode" validate:"required"`
	Nama      string    `json:"nama" validate:"required"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func getSelectKategori() []string {
	return []string{"id", "kode", "nama", "created_at", "updated_at"}
}

func getRowKategori(values []interface{}, m *Model) (Kategori, error) {
	item := Kategori{}
	item.ID = values[0].(int64)
	item.Kode = values[1].(string)
	item.Nama = values[2].(string)
	item.CreatedAt = null.TimeFrom(values[3].(time.Time))
	item.UpdatedAt = null.TimeFrom(values[4].(time.Time))
	return item, nil
}

func getRowsKategori(rows *sql.Rows, m *Model) ([]Kategori, error) {

	items := []Kategori{}

	count := len(getSelectKategori())

	values, valuesPtrs := getValuePtr(count)

	for rows.Next() {

		err := rows.Scan(valuesPtrs...)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		item, err := getRowKategori(values, m)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}
	return items, nil
}

func (m *Model) CreateKategori(r Kategori) error {

	sqlX := db.Insert(tableKategori, "kode", "nama")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Kode, r.Nama)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) UpdateKategori(r Kategori, id int64) error {

	updatedAt := constants.GetDatetimeNow()

	sqlX := db.Update(tableKategori, "id", "kode", "nama", "updated_at")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Kode, r.Nama, updatedAt, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveKategori(id int64) error {

	sqlX := db.Delete(tableKategori, "id")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) FetchKategori(id int64) (Kategori, error) {

	item := Kategori{}

	sqlX := db.Fetch(tableKategori, "id", getSelectKategori())

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	count := len(getSelectKategori())

	values, valuesPtrs := getValuePtr(count)

	err = row.Scan(valuesPtrs...)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	item, err = getRowKategori(values, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) AllKategori() ([]Kategori, error) {

	items := []Kategori{}

	sqlX := db.QueryAll(tableKategori, getSelectKategori())

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	items, err = getRowsKategori(rows, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	return items, nil
}

func (m *Model) GetKategori(lastID int64, limit int) ([]Kategori, error) {

	items := []Kategori{}

	sqlX := db.QueryPagingJoin(tableKategori, "id", false, getSelectKategori(), []db.Join{}, []string{})

	if lastID == 0 {
		sqlX = db.QueryPagingJoin(tableKategori, "id", true, getSelectKategori(), []db.Join{}, []string{})
	}
	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	rows, err := GetQueryRow(stmt, lastID, limit, "", "", nil)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	items, err = getRowsKategori(rows, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	return items, nil
}

func (m *Model) SearchKategori(lastID int64, limit int, search string, filter string) ([]Kategori, error) {

	items := []Kategori{}

	sqlX := db.QueryPagingJoinSearch(tableKategori, "id", false, getSelectKategori(), []db.Join{}, filter, []string{})

	if lastID == 0 {
		sqlX = db.QueryPagingJoinSearch(tableKategori, "id", true, getSelectKategori(), []db.Join{}, filter, []string{})
	}

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	rows, err := GetQueryRow(stmt, lastID, limit, search, filter, nil)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	items, err = getRowsKategori(rows, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	return items, nil
}

func (m *Model) CheckKategoriKode(kode string) bool {

	sqlX := `SELECT kode FROM %s WHERE kode=?`
	sqlX = fmt.Sprintf(sqlX, tableKategori)
	sqlX = m.Db.Rebind(sqlX)

	return m.RowExists(sqlX, kode)

}
