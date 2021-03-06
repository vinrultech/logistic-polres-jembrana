package models

import (
	"database/sql"
	"time"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

var tableSuratMasuk = "log_surat_masuks"

type SuratMasuk struct {
	ID           int64     `json:"id"`
	RowID        string    `json:"row_id"`
	NoSurat      string    `json:"no_surat" validate:"required"`
	TanggalSurat string    `json:"tanggal_surat" validate:"required"`
	Dari         string    `json:"dari" validate:"required"`
	Perihal      string    `json:"perihal" validate:"required"`
	UnitKerja    UnitKerja `json:"unit_kerja"`
	Isi          string    `json:"isi"`
	Files        []File    `json:"files"`
	CreatedAt    null.Time `json:"created_at"`
	UpdatedAt    null.Time `json:"updated_at"`
}

func getSelectSuratMasuk() []string {
	return []string{"id", "row_id", "no_surat", "tanggal_surat", "dari", "perihal", "isi", "unit_kerja_id", "created_at", "updated_at"}
}

func getInsertSuratMasuk() string {
	return db.Insert(tableSuratMasuk, "row_id", "no_surat",
		"tanggal_surat", "dari", "perihal", "isi", "unit_kerja_id")
}

func getUpdateSuratMasuk() string {
	return db.Update(tableSuratMasuk, "row_id",
		"no_surat", "tanggal_surat", "dari", "perihal", "isi", "updated_at")
}

func getRowSuratMasuk(values []interface{}, m *Model) (SuratMasuk, error) {
	item := SuratMasuk{}
	var unitKerjaID int64
	item.ID = values[0].(int64)
	item.RowID = values[1].(string)
	item.NoSurat = values[2].(string)
	item.TanggalSurat = values[3].(string)
	item.Dari = values[4].(string)
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

	files, err := m.GetFile(item.RowID, "surat_masuk")
	if err != nil {
		return item, err
	}
	item.Files = files
	return item, nil
}

func getRowsSuratMasuk(rows *sql.Rows, m *Model) ([]SuratMasuk, error) {

	items := []SuratMasuk{}

	count := len(getSelectSuratMasuk())

	values, valuesPtrs := getValuePtr(count)

	for rows.Next() {

		err := rows.Scan(valuesPtrs...)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		item, err := getRowSuratMasuk(values, m)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}
	return items, nil
}

func (m *Model) CreateSuratMasuk(r SuratMasuk) error {

	sqlX := getInsertSuratMasuk()

	sqlX = m.Db.Rebind(sqlX)

	if _, err := m.Db.Exec(sqlX, r.RowID, r.NoSurat, r.TanggalSurat,
		r.Dari, r.Perihal, r.Isi, r.UnitKerja.ID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) CreateSuratMasukWithTransaction(r SuratMasuk) error {

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	sqlX := getInsertSuratMasuk()

	sqlX = m.Db.Rebind(sqlX)

	if _, err = tx.Exec(sqlX, r.RowID, r.NoSurat,
		r.TanggalSurat, r.Dari, r.Perihal, r.Isi, r.UnitKerja.ID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if err := m.TxInsertFiles(tx, r.Files, r.RowID, "surat_masuk"); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if err = tx.Commit(); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) UpdateSuratMasukWithTransaction(r SuratMasuk, rowID string) error {

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	updatedAt := constants.GetDatetimeNow()

	sqlX := getUpdateSuratMasuk()

	sqlX = m.Db.Rebind(sqlX)

	if _, err = tx.Exec(sqlX, r.NoSurat,
		r.TanggalSurat, r.Dari, r.Perihal, r.Isi, updatedAt, rowID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if err := m.TxInsertFiles(tx, r.Files, r.RowID, "surat_masuk"); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if err = tx.Commit(); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveSuratMasuk(rowID string) error {

	sqlX := db.Delete(tableSuratMasuk, "row_id")

	sqlX = m.Db.Rebind(sqlX)

	if _, err := m.Db.Exec(sqlX, rowID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveSuratMasukTransaction(rowID string, files []File) error {

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	sqlX := db.Delete(tableSuratMasuk, "row_id")

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

func (m *Model) FetchSuratMasuk(rowID string) (SuratMasuk, error) {

	item := SuratMasuk{}

	sqlX := db.Fetch(tableSuratMasuk, "row_id", getSelectSuratMasuk())

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(rowID)

	count := len(getSelectSuratMasuk())

	values, valuesPtrs := getValuePtr(count)

	err = row.Scan(valuesPtrs...)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	item, err = getRowSuratMasuk(values, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) GetSuratMasuk(lastID int64, limit int, filters []string, filterValues []interface{}) ([]SuratMasuk, error) {

	items := []SuratMasuk{}

	sqlX := db.QueryPagingJoin(tableSuratMasuk, "id", false, getSelectSuratMasuk(), []db.Join{}, filters)

	if lastID == 0 {
		sqlX = db.QueryPagingJoin(tableSuratMasuk, "id", true, getSelectSuratMasuk(), []db.Join{}, filters)
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

	items, err = getRowsSuratMasuk(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}

func (m *Model) SearchSuratMasuk(lastID int64, limit int, search []string, filters []string, filterValues []interface{}) ([]SuratMasuk, error) {

	items := []SuratMasuk{}

	sqlX := db.QueryPagingJoinSearch(tableSuratMasuk, "id", false, getSelectSuratMasuk(), []db.Join{}, search[0], filters)

	if lastID == 0 {
		sqlX = db.QueryPagingJoinSearch(tableSuratMasuk, "id", true, getSelectSuratMasuk(), []db.Join{}, search[0], filters)
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

	items, err = getRowsSuratMasuk(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}
