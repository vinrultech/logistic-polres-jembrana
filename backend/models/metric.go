package models

import (
	"database/sql"
	"time"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

var tableMetric = "log_metric"

type Metric struct {
	ID        int64     `json:"id"`
	Nama      string    `json:"nama" validate:"required"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func getSelectMetric() []string {
	return []string{"id", "nama", "created_at", "updated_at"}
}

func getRowMetric(values []interface{}, m *Model) (Metric, error) {
	item := Metric{}
	item.ID = values[0].(int64)
	item.Nama = values[1].(string)
	item.CreatedAt = null.TimeFrom(values[2].(time.Time))
	item.UpdatedAt = null.TimeFrom(values[3].(time.Time))
	return item, nil
}

func getRowsMetric(rows *sql.Rows, m *Model) ([]Metric, error) {

	items := []Metric{}

	count := len(getSelectMetric())

	values, valuesPtrs := getValuePtr(count)

	for rows.Next() {

		err := rows.Scan(valuesPtrs...)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		item, err := getRowMetric(values, m)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}
	return items, nil
}

func (m *Model) CreateMetric(r Metric) error {

	sqlX := db.Insert(tableMetric, "nama")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Nama)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) UpdateMetric(r Metric, id int64) error {

	updatedAt := constants.GetDatetimeNow()

	sqlX := db.Update(tableMetric, "id", "nama", "updated_at")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Nama, updatedAt, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveMetric(id int64) error {

	sqlX := db.Delete(tableMetric, "id")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) FetchMetric(id int64) (Metric, error) {

	item := Metric{}

	sqlX := db.Fetch(tableMetric, "id", getSelectMetric())

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	count := len(getSelectMetric())

	values, valuesPtrs := getValuePtr(count)

	err = row.Scan(valuesPtrs...)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	item, err = getRowMetric(values, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) AllMetric() ([]Metric, error) {

	items := []Metric{}

	sqlX := db.QueryAll(tableMetric, getSelectMetric())

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

	items, err = getRowsMetric(rows, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	return items, nil
}

func (m *Model) GetMetric(lastID int64, limit int) ([]Metric, error) {

	items := []Metric{}

	sqlX := db.QueryPagingJoin(tableMetric, "id", false, getSelectMetric(), []db.Join{}, []string{})

	if lastID == 0 {
		sqlX = db.QueryPagingJoin(tableMetric, "id", true, getSelectMetric(), []db.Join{}, []string{})
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

	items, err = getRowsMetric(rows, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	return items, nil
}

func (m *Model) SearchMetric(lastID int64, limit int, search string, filter string) ([]Metric, error) {

	items := []Metric{}

	sqlX := db.QueryPagingJoinSearch(tableMetric, "id", false, getSelectMetric(), []db.Join{}, filter, []string{})

	if lastID == 0 {
		sqlX = db.QueryPagingJoinSearch(tableMetric, "id", true, getSelectMetric(), []db.Join{}, filter, []string{})
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

	items, err = getRowsMetric(rows, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	return items, nil
}
