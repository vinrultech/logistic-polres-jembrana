package models

import (
	"database/sql"
	"time"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

var tableSuratKeluar = "log_surat_keluars"

type SuratKeluar struct {
	ID           int64     `json:"id"`
	RowID        string    `json:"row_id"`
	NoSurat      string    `json:"no_surat" validate:"required"`
	TanggalSurat string    `json:"tanggal_surat" validate:"required"`
	Tujuan       string    `json:"tujuan" validate:"required"`
	Perihal      string    `json:"perihal" validate:"required"`
	UnitKerja    UnitKerja `json:"unit_kerja"`
	Isi          string    `json:"isi"`
	Files        []File    `json:"files"`
	CreatedAt    null.Time `json:"created_at"`
	UpdatedAt    null.Time `json:"updated_at"`
}

func getSelectSuratKeluar() []string {
	return []string{"id", "row_id", "no_surat", "tanggal_surat", "tujuan", "perihal", "isi", "unit_kerja_id", "created_at", "updated_at"}
}

func getRowSuratKeluar(values []interface{}, m *Model) (SuratKeluar, error) {
	item := SuratKeluar{}
	var unitKerjaID int64
	item.ID = values[0].(int64)
	item.RowID = values[1].(string)
	item.NoSurat = values[2].(string)
	item.TanggalSurat = values[3].(string)
	item.Tujuan = values[4].(string)
	item.Perihal = values[5].(string)
	item.Isi = values[6].(string)
	unitKerjaID = values[7].(int64)
	item.CreatedAt = null.TimeFrom(values[8].(time.Time))
	item.UpdatedAt = null.TimeFrom(values[9].(time.Time))
	if unitKerjaID == 0 {
		item.UnitKerja = GetUnitKerja()
	} else {
		unitKerja, err := m.FetchUnitKerja(int64(unitKerjaID))
		if err != nil {
			return item, err
		}
		item.UnitKerja = unitKerja
	}

	files, err := m.GetFile(item.RowID, "surat_keluar")
	if err != nil {
		return item, err
	}
	item.Files = files
	return item, nil
}

func getRowsSuratKeluar(rows *sql.Rows, m *Model) ([]SuratKeluar, error) {

	items := []SuratKeluar{}

	count := len(getSelectSuratKeluar())

	values, valuesPtrs := getValuePtr(count)

	for rows.Next() {

		err := rows.Scan(valuesPtrs...)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		item, err := getRowSuratKeluar(values, m)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}
	return items, nil
}

func (m *Model) CreateSuratKeluar(r SuratKeluar) error {

	sqlX := db.Insert(tableSuratKeluar, "row_id", "no_surat",
		"tanggal_surat", "tujuan", "perihal", "isi", "unit_kerja_id")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.RowID, r.NoSurat, r.TanggalSurat, r.Tujuan, r.Perihal, r.Isi, r.UnitKerja.ID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) CreateSuratKeluarWithTransaction(r SuratKeluar) error {

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	sqlInsert := db.Insert(tableSuratKeluar, "row_id", "no_surat",
		"tanggal_surat", "tujuan", "perihal", "isi", "unit_kerja_id")

	sqlInsert = m.Db.Rebind(sqlInsert)

	_, err = tx.Exec(sqlInsert, r.RowID, r.NoSurat,
		r.TanggalSurat, r.Tujuan, r.Perihal, r.Isi, r.UnitKerja.ID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	for _, file := range r.Files {
		file.RowID = r.RowID
		file.Tipe = "surat_keluar"

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

func (m *Model) UpdateSuratKeluarWithTransaction(r SuratKeluar, rowID string) error {

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	updatedAt := constants.GetDatetimeNow()

	sqlUpdate := db.Update(tableSuratKeluar, "row_id",
		"no_surat", "tanggal_surat", "tujuan", "perihal", "isi", "updated_at")

	sqlUpdate = m.Db.Rebind(sqlUpdate)

	_, err = tx.Exec(sqlUpdate, r.NoSurat,
		r.TanggalSurat, r.Tujuan, r.Perihal, r.Isi, updatedAt, rowID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	for _, file := range r.Files {
		if file.Status == "new" {
			file.RowID = r.RowID
			file.Tipe = "surat_keluar"

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

func (m *Model) RemoveSuratKeluar(rowID string) error {

	sqlX := db.Delete(tableSuratKeluar, "row_id")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, rowID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveSuratKeluarTransaction(rowID string) ([]File, error) {

	files := []File{}

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return files, err
	}

	sqlFetch := db.Fetch(tableSuratKeluar, "row_id", getSelectSuratKeluar())

	sqlFetch = m.Db.Rebind(sqlFetch)

	stmt, err := tx.Prepare(sqlFetch)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return files, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(rowID)

	count := len(getSelectSuratKeluar())

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

	item, err := getRowSuratKeluar(values, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return files, err
	}

	files = item.Files

	sqlDelete := db.Delete(tableSuratKeluar, "row_id")

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

func (m *Model) FetchSuratKeluar(rowID string) (SuratKeluar, error) {

	item := SuratKeluar{}

	sqlX := db.Fetch(tableSuratKeluar, "row_id", getSelectSuratKeluar())

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(rowID)

	count := len(getSelectSuratKeluar())

	values, valuesPtrs := getValuePtr(count)

	err = row.Scan(valuesPtrs...)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	item, err = getRowSuratKeluar(values, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) GetSuratKeluar(lastID int64, limit int, filters []string, filterValues []interface{}) ([]SuratKeluar, error) {

	items := []SuratKeluar{}

	sqlX := db.QueryPagingJoin(tableSuratKeluar, "id", false, getSelectSuratKeluar(), []db.Join{}, filters)

	if lastID == 0 {
		sqlX = db.QueryPagingJoin(tableSuratKeluar, "id", true, getSelectSuratKeluar(), []db.Join{}, filters)
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

	items, err = getRowsSuratKeluar(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}

func (m *Model) SearchSuratKeluar(lastID int64, limit int, search []string, filters []string, filterValues []interface{}) ([]SuratKeluar, error) {

	items := []SuratKeluar{}

	sqlX := db.QueryPagingJoinSearch(tableSuratKeluar, "id", false, getSelectSuratKeluar(), []db.Join{}, search[0], filters)

	if lastID == 0 {
		sqlX = db.QueryPagingJoinSearch(tableSuratKeluar, "id", true, getSelectSuratKeluar(), []db.Join{}, search[0], filters)
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

	items, err = getRowsSuratKeluar(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}
