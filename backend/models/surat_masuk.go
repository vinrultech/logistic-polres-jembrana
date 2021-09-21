package models

import (
	"database/sql"
	"fmt"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
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
	Isi          string    `json:"isi"`
	Files        []File    `json:"files"`
	CreatedAt    null.Time `json:"created_at"`
	UpdatedAt    null.Time `json:"updated_at"`
}

func (m *Model) CreateSuratMasuk(r SuratMasuk) error {

	sqlX := `INSERT INTO %s (row_id, no_surat, tanggal_surat, dari, perihal, isi) VALUES (?, ?, ?, ?, ?, ?)`

	sqlX = fmt.Sprintf(sqlX, tableSuratMasuk)

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.RowID, r.NoSurat, r.TanggalSurat, r.Dari, r.Perihal, r.Isi)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) UpdateSuratMasuk(r SuratMasuk, rowID string) error {

	updatedAt := constants.GetDatetimeNow()

	sqlX := `UPDATE %s SET no_surat=?, tanggal_surat=?, dari=?, perihal=?, isi=?, updated_at=? WHERE row_id=?`

	sqlX = fmt.Sprintf(sqlX, tableSuratMasuk)

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.NoSurat, r.TanggalSurat, r.Dari, r.Perihal, r.Isi, updatedAt, rowID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveSuratMasuk(rowID string) error {

	sqlX := `DELETE FROM %s WHERE row_id=?`

	sqlX = fmt.Sprintf(sqlX, tableSuratMasuk)

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, rowID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) FetchSuratMasuk(rowID string) (SuratMasuk, error) {

	item := SuratMasuk{}

	sqlX := `SELECT id, row_id , no_surat, tanggal_surat, dari, perihal, isi, created_at, updated_at FROM %s WHERE row_id=?`

	sqlX = fmt.Sprintf(sqlX, tableSuratMasuk)

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(rowID)

	err = row.Scan(&item.ID, &item.RowID, &item.NoSurat, &item.TanggalSurat,
		&item.Dari, &item.Perihal, &item.Isi, &item.CreatedAt, &item.UpdatedAt)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	files, err := m.GetFile(rowID, "surat_masuk")

	if err != nil {
		return item, err
	}

	item.Files = files

	return item, nil
}

func (m *Model) GetSuratMasuk(lastID int64, limit int, params ...string) ([]SuratMasuk, error) {

	var startDate string
	var endDate string

	if len(params) > 0 {
		startDate = params[0]
		endDate = params[1]
	}

	items := []SuratMasuk{}

	sqlX := `SELECT id, row_id , no_surat, tanggal_surat, dari, perihal, isi, created_at, updated_at FROM %s WHERE id<? ORDER BY id DESC LIMIT ?`

	if len(params) > 0 {
		sqlX = `SELECT id, row_id , no_surat, tanggal_surat, dari, perihal, isi, created_at, updated_at FROM %s WHERE id<? AND (tanggal_surat BETWEEN ? AND ?) ORDER BY id DESC LIMIT ?`
	}

	if lastID == 0 {
		sqlX = `SELECT id, row_id , no_surat, tanggal_surat, dari, perihal, isi, created_at, updated_at FROM %s ORDER BY id DESC LIMIT ?`

		if len(params) > 0 {
			sqlX = `SELECT id, row_id , no_surat, tanggal_surat, dari, perihal, isi, created_at, updated_at FROM %s WHERE (tanggal_surat BETWEEN ? AND ?) ORDER BY id DESC LIMIT ?`
		}
	}

	sqlX = fmt.Sprintf(sqlX, tableSuratMasuk)

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

	for rows.Next() {
		item := SuratMasuk{}
		err = rows.Scan(&item.ID, &item.RowID, &item.NoSurat, &item.TanggalSurat,
			&item.Dari, &item.Perihal, &item.Isi, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}
		files, err := m.GetFile(item.RowID, "surat_masuk")
		if err != nil {
			return items, err
		}
		item.Files = files
		items = append(items, item)
	}

	return items, nil
}

func (m *Model) SearchSuratMasuk(lastID int64, limit int, search string, filter string, params ...string) ([]SuratMasuk, error) {

	var startDate string
	var endDate string

	if len(params) > 0 {
		startDate = params[0]
		endDate = params[1]
	}

	items := []SuratMasuk{}

	query := `ILIKE '%' || ? || '%'`

	sqlX := `SELECT id, row_id , no_surat, tanggal_surat, dari, perihal, isi, created_at, updated_at FROM %s WHERE id<? AND %s %s ORDER BY id DESC LIMIT ?`

	if len(params) > 0 {
		sqlX = `SELECT id, row_id , no_surat, tanggal_surat, dari, perihal, isi, created_at, updated_at FROM %s WHERE id<? AND %s %s AND (tanggal_surat BETWEEN ? AND ?) ORDER BY id DESC LIMIT ?`
	}

	if lastID == 0 {
		sqlX = `SELECT id, row_id , no_surat, tanggal_surat, dari, perihal, isi, created_at, updated_at FROM %s WHERE %s %s ORDER BY id DESC LIMIT ?`

		if len(params) > 0 {
			sqlX = `SELECT id, row_id , no_surat, tanggal_surat, dari, perihal, isi, created_at, updated_at FROM %s WHERE %s %s AND (tanggal_surat BETWEEN ? AND ?) ORDER BY id DESC LIMIT ?`
		}
	}

	sqlX = fmt.Sprintf(sqlX, tableSuratMasuk, filter, query)

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
		if len(params) > 0 {
			rows, err = stmt.Query(search, startDate, endDate, limit)
		} else {
			rows, err = stmt.Query(search, limit)
		}

	} else {
		if len(params) > 0 {
			rows, err = stmt.Query(lastID, search, startDate, endDate, limit)
		} else {
			rows, err = stmt.Query(lastID, search, limit)
		}

	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	for rows.Next() {
		item := SuratMasuk{}
		err = rows.Scan(&item.ID, &item.RowID, &item.NoSurat, &item.TanggalSurat,
			&item.Dari, &item.Perihal, &item.Isi, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}
		files, err := m.GetFile(item.RowID, "surat_masuk")
		if err != nil {
			return items, err
		}
		item.Files = files
		items = append(items, item)
	}

	return items, nil
}
