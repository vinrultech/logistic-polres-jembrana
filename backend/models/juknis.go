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

func getValuesJuknis() string {
	return db.Select(getSelectJuknis())
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

	sqlX := db.Insert(tableJuknis, "row_id", "no_surat",
		"tanggal_surat", "perihal", "isi", "unit_kerja_id")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.RowID, r.NoSurat, r.TanggalSurat, r.Perihal, r.Isi, r.UnitKerja.ID)

	if err != nil {
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

	sqlInsert := db.Insert(tableJuknis, "row_id", "no_surat",
		"tanggal_surat", "perihal", "isi", "unit_kerja_id")

	sqlInsert = m.Db.Rebind(sqlInsert)

	_, err = tx.Exec(sqlInsert, r.RowID, r.NoSurat,
		r.TanggalSurat, r.Perihal, r.Isi, r.UnitKerja.ID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	for _, file := range r.Files {
		file.RowID = r.RowID
		file.Tipe = "juknis"

		sqlFileInsert := db.Insert(tableFile, "file_id", "filename", "url", "row_id", "tipe")

		sqlFileInsert = m.Db.Rebind(sqlFileInsert)

		_, err := tx.Exec(sqlFileInsert, file.FileID, file.Filename, file.Url, file.RowID, file.Tipe)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return err
		}

	}

	err = tx.Commit()

	if err != nil {
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

	sqlUpdate := db.Update(tableJuknis, "row_id",
		"no_surat", "tanggal_surat", "perihal", "isi", "updated_at")

	sqlUpdate = m.Db.Rebind(sqlUpdate)

	_, err = tx.Exec(sqlUpdate, r.NoSurat,
		r.TanggalSurat, r.Perihal, r.Isi, updatedAt, rowID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	for _, file := range r.Files {
		if file.Status == "new" {
			file.RowID = r.RowID
			file.Tipe = "juknis"

			sqlFileInsert := db.Insert(tableFile, "file_id", "filename", "url", "row_id", "tipe")

			sqlFileInsert = m.Db.Rebind(sqlFileInsert)

			_, err := tx.Exec(sqlFileInsert, file.FileID, file.Filename, file.Url, file.RowID, file.Tipe)

			if err != nil {
				loggers.Log.Errorln(err.Error())
				return err
			}
		}
	}

	err = tx.Commit()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveJuknis(rowID string) error {

	sqlX := db.Delete(tableJuknis, "row_id")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, rowID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveJuknisTransaction(rowID string) ([]File, error) {

	files := []File{}

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return files, err
	}

	sqlFetch := db.Fetch(tableJuknis, "row_id", getSelectJuknis())

	sqlFetch = m.Db.Rebind(sqlFetch)

	stmt, err := tx.Prepare(sqlFetch)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return files, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(rowID)

	count := len(getSelectJuknis())

	values := make([]interface{}, count)
	valuesPtrs := make([]interface{}, count)

	for i := 0; i < count; i++ {
		valuesPtrs[i] = &values[i]
	}

	err = row.Scan(valuesPtrs...)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return files, err
	}

	item, err := getRowJuknis(values, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return files, err
	}

	files = item.Files

	sqlDelete := db.Delete(tableJuknis, "row_id")

	sqlDelete = m.Db.Rebind(sqlDelete)

	_, err = tx.Exec(sqlDelete, rowID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return files, err
	}

	for _, file := range files {
		sqlFileDelete := db.Delete(tableFile, "file_id")

		sqlFileDelete = m.Db.Rebind(sqlFileDelete)

		_, err := tx.Exec(sqlFileDelete, file.FileID)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return files, err
		}
	}

	tx.Commit()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return files, err
	}

	return files, nil
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

func (m *Model) GetJuknis(lastID int64, limit int, params ...string) ([]Juknis, error) {

	items := []Juknis{}

	var startDate string
	var endDate string

	if len(params) > 0 {
		startDate = params[0]
		endDate = params[1]
	}

	sqlX := db.QueryPagingSurat(tableJuknis, getValuesJuknis(), false, "id")

	if len(params) > 0 {
		sqlX = db.QueryPagingSurat(tableJuknis, getValuesJuknis(), false, "id", "tanggal_surat")

	}

	if lastID == 0 {

		sqlX = db.QueryPagingSurat(tableJuknis, getValuesJuknis(), true, "id")
		if len(params) > 0 {

			sqlX = db.QueryPagingSurat(tableJuknis, getValuesJuknis(), true, "id", "tanggal_surat")
		}
	}

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	var rows *sql.Rows

	if lastID == 0 {

		if len(params) > 0 {
			rows, err = stmt.Query(startDate, endDate, limit)
		} else {
			rows, err = stmt.Query(limit)
		}
	} else {

		if len(params) > 0 {
			rows, err = stmt.Query(lastID, startDate, endDate, limit)
		} else {
			rows, err = stmt.Query(lastID, limit)
		}
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

func (m *Model) GetJuknisWithFilter(lastID int64, limit int, unitKerjaID int64, params ...string) ([]Juknis, error) {

	items := []Juknis{}

	var startDate string
	var endDate string

	if len(params) > 0 {
		startDate = params[0]
		endDate = params[1]
	}

	sqlX := db.QueryPagingSuratWithFilter(tableJuknis, getValuesJuknis(), false, "id", "unit_kerja_id=?")

	if len(params) > 0 {
		sqlX = db.QueryPagingSuratWithFilter(tableJuknis, getValuesJuknis(), false, "id", "unit_kerja_id=?", "tanggal_surat")

	}

	if lastID == 0 {

		sqlX = db.QueryPagingSuratWithFilter(tableJuknis, getValuesJuknis(), true, "id", "unit_kerja_id=?")
		if len(params) > 0 {

			sqlX = db.QueryPagingSuratWithFilter(tableJuknis, getValuesJuknis(), true, "id", "unit_kerja_id=?", "tanggal_surat")
		}
	}

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	var rows *sql.Rows

	if lastID == 0 {

		if len(params) > 0 {
			rows, err = stmt.Query(unitKerjaID, startDate, endDate, limit)
		} else {
			rows, err = stmt.Query(unitKerjaID, limit)
		}
	} else {

		if len(params) > 0 {
			rows, err = stmt.Query(unitKerjaID, lastID, startDate, endDate, limit)
		} else {
			rows, err = stmt.Query(unitKerjaID, lastID, limit)
		}
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

func (m *Model) SearchJuknis(lastID int64, limit int, search string, filter string, params ...string) ([]Juknis, error) {

	items := []Juknis{}

	var startDate string
	var endDate string

	if len(params) > 0 {
		startDate = params[0]
		endDate = params[1]
	}

	sqlX := db.QueryPagingSuratSearch(tableJuknis, getValuesJuknis(), false, "id", filter)

	if len(params) > 0 {
		sqlX = db.QueryPagingSuratSearch(tableJuknis, getValuesJuknis(), false, "id", filter, "tanggal_surat")
	}

	if lastID == 0 {
		sqlX = db.QueryPagingSuratSearch(tableJuknis, getValuesJuknis(), true, "id", filter)

		if len(params) > 0 {
			sqlX = db.QueryPagingSuratSearch(tableJuknis, getValuesJuknis(), true, "id", filter, "tanggal_surat")
		}
	}

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	var rows *sql.Rows

	if lastID == 0 {
		if len(params) > 0 {
			rows, err = stmt.Query(search, startDate, endDate, limit)
		} else {
			rows, err = stmt.Query(search, limit)
		}

	} else {
		if len(params) > 0 {
			rows, err = stmt.Query(search, lastID, startDate, endDate, limit)
		} else {
			rows, err = stmt.Query(search, lastID, limit)
		}

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

func (m *Model) SearchJuknisWithFilter(lastID int64, limit int, search string, filter string, unitKerjaID int64, params ...string) ([]Juknis, error) {

	items := []Juknis{}

	var startDate string
	var endDate string

	if len(params) > 0 {
		startDate = params[0]
		endDate = params[1]
	}

	sqlX := db.QueryPagingSuratSearchWithFilter(tableJuknis, getValuesJuknis(), false, "id", filter, "unit_kerja_id=?")

	if len(params) > 0 {
		sqlX = db.QueryPagingSuratSearchWithFilter(tableJuknis, getValuesJuknis(), false, "id", filter, "unit_kerja_id=?", "tanggal_surat")
	}

	if lastID == 0 {
		sqlX = db.QueryPagingSuratSearchWithFilter(tableJuknis, getValuesJuknis(), true, "id", filter, "unit_kerja_id=?")

		if len(params) > 0 {
			sqlX = db.QueryPagingSuratSearchWithFilter(tableJuknis, getValuesJuknis(), true, "id", filter, "unit_kerja_id=?", "tanggal_surat")
		}
	}

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	var rows *sql.Rows

	if lastID == 0 {
		if len(params) > 0 {
			rows, err = stmt.Query(search, unitKerjaID, startDate, endDate, limit)
		} else {
			rows, err = stmt.Query(search, unitKerjaID, limit)
		}

	} else {
		if len(params) > 0 {
			rows, err = stmt.Query(search, unitKerjaID, lastID, startDate, endDate, limit)
		} else {
			rows, err = stmt.Query(search, unitKerjaID, lastID, limit)
		}

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
