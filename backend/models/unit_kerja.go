package models

import (
	"database/sql"
	"time"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
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

func getSelectUnitKerja() []string {
	return []string{"id", "nama", "alamat", "telepon", "created_at", "updated_at"}
}

func getRowUnitKerja(values []interface{}, m *Model) (UnitKerja, error) {
	item := UnitKerja{}
	item.ID = values[0].(int64)
	item.Nama = values[1].(string)
	item.Alamat = values[2].(string)
	item.Telepon = values[3].(string)
	item.CreatedAt = null.TimeFrom(values[4].(time.Time))
	item.UpdatedAt = null.TimeFrom(values[5].(time.Time))
	return item, nil
}

func getRowsUnitKerja(rows *sql.Rows, m *Model) ([]UnitKerja, error) {

	items := []UnitKerja{}

	count := len(getSelectUnitKerja())

	values, valuesPtrs := getValuePtr(count)

	for rows.Next() {

		err := rows.Scan(valuesPtrs...)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		item, err := getRowUnitKerja(values, m)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}
	return items, nil
}

func (m *Model) CreateUnitKerja(r UnitKerja) error {

	sqlX := db.Insert(tableUnitKerja, "nama", "alamat", "telepon")

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

	sqlX := db.Update(tableUnitKerja, "id", "nama", "alamat", "telepon", "updated_at")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Nama, r.Alamat, r.Telepon, updatedAt, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveUnitKerja(id int64) error {

	sqlX := db.Delete(tableUnitKerja, "id")

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

	sqlX := db.Fetch(tableUnitKerja, "id", getSelectUnitKerja())

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	count := len(getSelectUnitKerja())

	values, valuesPtrs := getValuePtr(count)

	err = row.Scan(valuesPtrs...)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	item, err = getRowUnitKerja(values, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) GetUnitKerja(lastID int64, limit int) ([]UnitKerja, error) {

	items := []UnitKerja{}

	sqlX := db.QueryPagingJoin(tableUnitKerja, "id", false, getSelectUnitKerja(), []db.Join{}, []string{})

	if lastID == 0 {
		sqlX = db.QueryPagingJoin(tableUnitKerja, "id", true, getSelectUnitKerja(), []db.Join{}, []string{})
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

	items, err = getRowsUnitKerja(rows, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	return items, nil
}

func (m *Model) AllUnitKerja() ([]UnitKerja, error) {

	items := []UnitKerja{}

	sqlX := db.QueryAll(tableUnitKerja, getSelectUnitKerja())

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

	items, err = getRowsUnitKerja(rows, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	return items, nil
}

func (m *Model) SearchUnitKerja(lastID int64, limit int, search string, filter string) ([]UnitKerja, error) {

	items := []UnitKerja{}

	sqlX := db.QueryPagingJoinSearch(tableUnitKerja, "id", false, getSelectUnitKerja(), []db.Join{}, filter, []string{})

	if lastID == 0 {
		sqlX = db.QueryPagingJoinSearch(tableUnitKerja, "id", true, getSelectUnitKerja(), []db.Join{}, filter, []string{})
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

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	items, err = getRowsUnitKerja(rows, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	return items, nil
}
