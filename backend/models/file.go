package models

import (
	"database/sql"
	"fmt"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

var tableFile = "log_files"

type File struct {
	FileID    string    `json:"file_id" validate:"required"`
	RowID     string    `json:"row_id"`
	Tipe      string    `json:"tipe"`
	Filename  string    `json:"filename" validate:"required"`
	Url       string    `json:"url" validate:"required"`
	Status    string    `json:"status" validate:"required"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func getInsertFile() string {
	return db.Insert(tableFile, "file_id", "filename", "url", "row_id", "tipe")
}

func txInsertFile(query string, tx *sql.Tx, file File) error {

	if _, err := tx.Exec(query, file.FileID, file.Filename, file.Url, file.RowID, file.Tipe); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) TxInsertFiles(tx *sql.Tx, files []File, rowId string, tipe string) error {
	sqlX := getInsertFile()
	sqlX = m.Db.Rebind(sqlX)
	for _, file := range files {
		if file.Status == "new" {
			file.RowID = rowId
			file.Tipe = tipe
			if err := txInsertFile(sqlX, tx, file); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *Model) TxRemoveFiles(tx *sql.Tx, files []File) error {
	sqlX := db.Delete(tableFile, "file_id")
	sqlX = m.Db.Rebind(sqlX)
	for _, file := range files {
		if _, err := tx.Exec(sqlX, file.FileID); err != nil {
			return err
		}
	}

	return nil
}

func (m *Model) CreateFile(r File) error {

	sqlX := `INSERT INTO %s (file_id, filename, url, row_id, tipe) VALUES (?, ?, ?, ?, ?)`

	sqlX = fmt.Sprintf(sqlX, tableFile)

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.FileID, r.Filename, r.Url, r.RowID, r.Tipe)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveFile(fileID string) error {

	sqlX := `DELETE FROM %s WHERE file_id=?`

	sqlX = fmt.Sprintf(sqlX, tableFile)

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, fileID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) GetFile(rowID string, tipe string) ([]File, error) {

	items := []File{}

	sqlX := `SELECT file_id, row_id, tipe, url, filename, created_at, updated_at FROM %s WHERE row_id=? and tipe=?`

	sqlX = fmt.Sprintf(sqlX, tableFile)

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(rowID, tipe)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	for rows.Next() {
		item := File{}

		err = rows.Scan(&item.FileID, &item.RowID, &item.Tipe, &item.Url, &item.Filename, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}
		item.Status = "edit"
		items = append(items, item)
	}

	return items, nil
}
