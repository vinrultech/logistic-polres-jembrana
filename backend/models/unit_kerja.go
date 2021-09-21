package models

import (
	"database/sql"
	"fmt"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

var tableUnitKerja = "log_unit_kerjas"

type UnitKerja struct {
	ID        int64     `json:"id"`
	Nama      string    `json:"nama" validate:"required"`
	Alamat    string    `json:"alamat" validate:"required"`
	Telepon   string    `json:"telepon" validate:"required"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func (m *Model) CreateUnitKerja(r UnitKerja) error {

	sqlX := `INSERT INTO %s (nama, alamat, telepon) VALUES (?, ?, ?)`

	sqlX = fmt.Sprintf(sqlX, tableUnitKerja)

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Nama, r.Alamat, r.Telepon)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) UpdateUnitKerja(r UnitKerja, id int64) error {

	updatedAt := constants.GetDatetimeNow()

	sqlX := `UPDATE %s SET nama=?, alamat=?, telepon=?, updated_at=? WHERE id=?`

	sqlX = fmt.Sprintf(sqlX, tableUnitKerja)

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Nama, r.Alamat, r.Telepon, updatedAt, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveUnitKerja(id int64) error {

	sqlX := `DELETE FROM %s WHERE id=?`

	sqlX = fmt.Sprintf(sqlX, tableUnitKerja)

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) FetchUnitKerja(id int64) (UnitKerja, error) {

	item := UnitKerja{}

	sqlX := `SELECT id, nama, alamat, telepon, created_at, updated_at FROM %s WHERE id=?`

	sqlX = fmt.Sprintf(sqlX, tableUnitKerja)

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(&item.ID, &item.Nama, &item.Alamat, &item.Telepon, &item.CreatedAt, &item.UpdatedAt)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) GetUnitKerja(lastID int64, limit int) ([]UnitKerja, error) {

	items := []UnitKerja{}

	sqlX := `SELECT id, nama, alamat, telepon, created_at, updated_at FROM %s WHERE id<? ORDER BY id DESC LIMIT ?`

	if lastID == 0 {
		sqlX = `SELECT id, nama, alamat, telepon, created_at, updated_at FROM %s ORDER BY id DESC LIMIT ?`
	}

	sqlX = fmt.Sprintf(sqlX, tableUnitKerja)

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
		item := UnitKerja{}
		err = rows.Scan(&item.ID, &item.Nama, &item.Alamat, &item.Telepon, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (m *Model) AllUnitKerja() ([]UnitKerja, error) {

	items := []UnitKerja{}

	sqlX := `SELECT id, nama, alamat, telepon, created_at, updated_at FROM %s ORDER BY id DESC`

	sqlX = fmt.Sprintf(sqlX, tableUnitKerja)

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

	for rows.Next() {
		item := UnitKerja{}
		err = rows.Scan(&item.ID, &item.Nama, &item.Alamat, &item.Telepon, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}

	return items, nil
}

func (m *Model) SearchUnitKerja(lastID int64, limit int, search string, filter string) ([]UnitKerja, error) {

	items := []UnitKerja{}

	query := `ILIKE '%' || ? || '%'`

	sqlX := `SELECT id, nama, alamat, telepon, created_at, updated_at FROM %s WHERE id<? AND %s %s ORDER BY id DESC LIMIT ?`

	if lastID == 0 {
		sqlX = `SELECT id, nama, alamat, telepon, created_at, updated_at FROM %s WHERE %s %s ORDER BY id DESC LIMIT ?`
	}

	sqlX = fmt.Sprintf(sqlX, tableUnitKerja, filter, query)

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
		item := UnitKerja{}
		err = rows.Scan(&item.ID, &item.Nama, &item.Alamat, &item.Telepon, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}

	return items, nil
}
