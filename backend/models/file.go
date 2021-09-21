package models

import (
	"fmt"

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
