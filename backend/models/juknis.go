package models

import (
	"database/sql"
	"time"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

var tableJuknis = "log_jukniss"

type Juknis struct {
	ID           int64     `json:"id"`
	RowID        string    `json:"row_id"`
	NoSurat      string    `json:"no_surat" validate:"required"`
	TanggalSurat string    `json:"tanggal_surat" validate:"required"`
	Perihal      string    `json:"perihal" validate:"required"`
	UnitKerja    UnitKerja `json:"unit_kerja"`
	Isi          string    `json:"isi"`
	Files        []File    `json:"files"`
	CreatedAt    null.Time `json:"created_at"`
	UpdatedAt    null.Time `json:"updated_at"`
}

func getSelectJuknis() []string {
	return []string{"id", "row_id", "no_surat", "tanggal_surat",
		"perihal", "isi", "unit_kerja_id", "created_at", "updated_at"}
}

func getInsertJuknis() string {
	return db.Insert(tableJuknis, "row_id", "no_surat",
		"tanggal_surat", "perihal", "isi", "unit_kerja_id")
}

func getUpdateJuknis() string {
	return db.Update(tableJuknis, "row_id",
		"no_surat", "tanggal_surat", "perihal", "isi", "updated_at")
}

func getRowJuknis(values []interface{}, m *Model) (Juknis, error) {
	item := Juknis{}
	var unitKerjaID int64
	item.ID = values[0].(int64)
	item.RowID = values[1].(string)
	item.NoSurat = values[2].(string)
	item.TanggalSurat = values[3].(string)
	item.Perihal = values[4].(string)
	item.Isi = values[5].(string)
	unitKerjaID = values[6].(int64)
	item.CreatedAt = null.TimeFrom(values[7].(time.Time))
	item.UpdatedAt = null.TimeFrom(values[8].(time.Time))
	if unitKerjaID == 0 {
		item.UnitKerja = GetUnitKerja()
	} else {
		unitKerja, err := m.FetchUnitKerja(int64(unitKerjaID))
		if err != nil {
			return item, err
		}
		item.UnitKerja = unitKerja
	}

	files, err := m.GetFile(item.RowID, "juknis")
	if err != nil {
		return item, err
	}
	item.Files = files
	return item, nil
}

func getRowsJuknis(rows *sql.Rows, m *Model) ([]Juknis, error) {

	items := []Juknis{}

	count := len(getSelectJuknis())

	values, valuesPtrs := getValuePtr(count)

	for rows.Next() {

		err := rows.Scan(valuesPtrs...)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		item, err := getRowJuknis(values, m)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}
	return items, nil
}

func (m *Model) CreateJuknis(r Juknis) error {

	sqlX := getInsertJuknis()

	sqlX = m.Db.Rebind(sqlX)

	if _, err := m.Db.Exec(sqlX, r.RowID, r.NoSurat, r.TanggalSurat, r.Perihal, r.Isi, r.UnitKerja.ID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) CreateJuknisWithTransaction(r Juknis) error {

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	sqlX := getInsertJuknis()

	sqlX = m.Db.Rebind(sqlX)

	if _, err = tx.Exec(sqlX, r.RowID, r.NoSurat,
		r.TanggalSurat, r.Perihal, r.Isi, r.UnitKerja.ID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if err := m.TxInsertFiles(tx, r.Files, r.RowID, "juknis"); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if err = tx.Commit(); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) UpdateJuknisWithTransaction(r Juknis, rowID string) error {

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	updatedAt := constants.GetDatetimeNow()

	sqlX := getUpdateJuknis()

	sqlX = m.Db.Rebind(sqlX)

	if _, err = tx.Exec(sqlX, r.NoSurat,
		r.TanggalSurat, r.Perihal, r.Isi, updatedAt, rowID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if err := m.TxInsertFiles(tx, r.Files, r.RowID, "juknis"); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if err = tx.Commit(); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveJuknis(rowID string) error {

	sqlX := db.Delete(tableJuknis, "row_id")

	sqlX = m.Db.Rebind(sqlX)

	if _, err := m.Db.Exec(sqlX, rowID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveJuknisTransaction(rowID string, files []File) error {

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	sqlX := db.Delete(tableJuknis, "row_id")

	sqlX = m.Db.Rebind(sqlX)

	if _, err = tx.Exec(sqlX, rowID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if err := m.TxRemoveFiles(tx, files); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if tx.Commit(); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) FetchJuknis(rowID string) (Juknis, error) {

	item := Juknis{}

	sqlX := db.Fetch(tableJuknis, "row_id", getSelectJuknis())

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(rowID)

	count := len(getSelectJuknis())

	values, valuesPtrs := getValuePtr(count)

	err = row.Scan(valuesPtrs...)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	item, err = getRowJuknis(values, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) GetJuknis(lastID int64, limit int, filters []string, filterValues []interface{}) ([]Juknis, error) {

	items := []Juknis{}

	sqlX := db.QueryPagingJoin(tableJuknis, "id", false, getSelectJuknis(), []db.Join{}, filters)

	if lastID == 0 {
		sqlX = db.QueryPagingJoin(tableJuknis, "id", true, getSelectJuknis(), []db.Join{}, filters)
	}

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	rows, err := GetQueryRow(stmt, lastID, limit, "", "", filterValues)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	items, err = getRowsJuknis(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}

func (m *Model) SearchJuknis(lastID int64, limit int, search []string, filters []string, filterValues []interface{}) ([]Juknis, error) {

	items := []Juknis{}

	sqlX := db.QueryPagingJoinSearch(tableJuknis, "id", false, getSelectJuknis(), []db.Join{}, search[0], filters)

	if lastID == 0 {
		sqlX = db.QueryPagingJoinSearch(tableJuknis, "id", true, getSelectJuknis(), []db.Join{}, search[0], filters)
	}

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	rows, err := GetQueryRow(stmt, lastID, limit, search[1], search[0], filterValues)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	items, err = getRowsJuknis(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}
