package models

import (
	"database/sql"
	"fmt"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
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

func (m *Model) CreateKategori(r Kategori) error {

	sqlX := `INSERT INTO %s (kode, nama) VALUES (?, ?)`

	sqlX = fmt.Sprintf(sqlX, tableKategori)

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

	sqlX := `UPDATE %s SET kode=?, nama=?, updated_at=? WHERE id=?`

	sqlX = fmt.Sprintf(sqlX, tableKategori)

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Kode, r.Nama, updatedAt, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveKategori(id int64) error {

	sqlX := `DELETE FROM %s WHERE id=?`

	sqlX = fmt.Sprintf(sqlX, tableKategori)

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

	sqlX := `SELECT id, kode, nama, created_at, updated_at FROM %s WHERE id=?`

	sqlX = fmt.Sprintf(sqlX, tableKategori)

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(&item.ID, &item.Kode, &item.Nama, &item.CreatedAt, &item.UpdatedAt)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) GetKategori(lastID int64, limit int) ([]Kategori, error) {

	items := []Kategori{}

	sqlX := `SELECT id, kode, nama, created_at, updated_at FROM %s WHERE id<? ORDER BY id DESC LIMIT ?`

	if lastID == 0 {
		sqlX = `SELECT id, kode, nama, created_at, updated_at FROM %s ORDER BY id DESC LIMIT ?`
	}

	sqlX = fmt.Sprintf(sqlX, tableKategori)

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	var rows *sql.Rows

	if lastID == 0 {
		rows, err = stmt.Query(limit)
	} else {
		rows, err = stmt.Query(lastID, limit)
	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	for rows.Next() {
		item := Kategori{}
		err = rows.Scan(&item.ID, &item.Kode, &item.Nama, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (m *Model) SearchKategori(lastID int64, limit int, search string, filter string) ([]Kategori, error) {

	items := []Kategori{}

	query := `ILIKE '%' || ? || '%'`

	sqlX := `SELECT id, kode, nama, created_at, updated_at FROM %s WHERE id<? AND %s %s ORDER BY id DESC LIMIT ?`

	if lastID == 0 {
		sqlX = `SELECT id, kode, nama, created_at, updated_at FROM %s WHERE %s %s ORDER BY id DESC LIMIT ?`
	}

	sqlX = fmt.Sprintf(sqlX, tableKategori, filter, query)

	fmt.Println(sqlX)

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	var rows *sql.Rows

	if lastID == 0 {
		rows, err = stmt.Query(search, limit)
	} else {
		rows, err = stmt.Query(lastID, search, limit)
	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	for rows.Next() {
		item := Kategori{}
		err = rows.Scan(&item.ID, &item.Kode, &item.Nama, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (m *Model) CheckKategoriKode(kode string) bool {

	sqlX := `SELECT kode FROM %s WHERE kode=?`
	sqlX = fmt.Sprintf(sqlX, tableKategori)
	sqlX = m.Db.Rebind(sqlX)

	return m.RowExists(sqlX, kode)

}
