package models

import (
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

var tableInventoryKeluar = "log_barang_keluars AS ik"

type InventoryKeluar struct {
	ID        int64     `json:"id"`
	RowID     string    `json:"row_id"`
	Tanggal   string    `json:"nama" validate:"required"`
	Jumlah    int64     `json:"jumlah" validate:"required"`
	Barang    Barang    `json:"barang"`
	UnitKerja UnitKerja `json:"unit_kerja"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func getSelectInventoryKeluar() []string {
	return []string{"ik.id", "ik.row_id", "ik.barang_id", "ik.unit_kerja_id", "ik.tanggal",
		"ik.jumlah", "ik.created_at", "ik.updated_at"}
}

func getJoinInventoryKeluar() []db.Join {
	joins := []db.Join{}

	join1 := db.Join{TableJoin: "log_barangs AS b", IDTableJoin: "b.row_id", IDTable: "ik.barang_id"}
	joins = append(joins, join1)

	join2 := db.Join{TableJoin: "log_unit_kerjas AS u", IDTableJoin: "u.id", IDTable: "ik.unit_kerja_id"}
	joins = append(joins, join2)

	return joins
}

func getRowInventoryKeluar(values []interface{}, m *Model) (InventoryKeluar, error) {
	item := InventoryKeluar{}

	item.ID = values[0].(int64)
	item.RowID = values[1].(string)
	barangID := values[2].(string)
	unitKerjaID := values[3].(int64)
	item.Tanggal = values[4].(string)
	item.Jumlah = values[5].(int64)

	item.CreatedAt = null.TimeFrom(values[6].(time.Time))
	item.UpdatedAt = null.TimeFrom(values[7].(time.Time))

	barang, err := m.FetchBarang(barangID)

	if err != nil {
		return item, err
	}

	item.Barang = barang

	if unitKerjaID == 0 {
		item.UnitKerja = GetUnitKerja()
	} else {
		unitKerja, err := m.FetchUnitKerja(unitKerjaID)
		if err != nil {
			return item, err
		}
		item.UnitKerja = unitKerja

	}

	return item, nil
}

func getRowsInventoryKeluar(rows *sql.Rows, m *Model) ([]InventoryKeluar, error) {

	items := []InventoryKeluar{}

	count := len(getSelectInventoryKeluar())

	values, valuesPtrs := getValuePtr(count)

	for rows.Next() {

		err := rows.Scan(valuesPtrs...)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		item, err := getRowInventoryKeluar(values, m)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}
	return items, nil
}

func (m *Model) CreateInventoryKeluar(r InventoryKeluar, b Barang) error {

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	sqlInsert := db.Insert(tableInventoryKeluar, "row_id", "barang_id", "unit_kerja_id",
		"tanggal", "jumlah")

	sqlInsert = m.Db.Rebind(sqlInsert)

	fmt.Println(sqlInsert)

	if _, err = tx.Exec(sqlInsert, r.RowID, b.RowID,
		r.UnitKerja.ID, r.Tanggal, r.Jumlah); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	jumlah := b.Jumlah - r.Jumlah

	if jumlah < 0 {
		jumlah = 0
	}

	updatedAt := constants.GetDatetimeNow()

	sqlUpdate := db.Update(tableBarang, "row_id", "jumlah", "updated_at")

	sqlUpdate = m.Db.Rebind(sqlUpdate)

	if _, err = tx.Exec(sqlUpdate, jumlah, updatedAt, b.RowID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if err = tx.Commit(); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveInventoryKeluar(r InventoryKeluar, b Barang) error {

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	sqlDelete := db.Delete(tableInventoryKeluar, "row_id")

	sqlDelete = m.Db.Rebind(sqlDelete)

	if _, err = tx.Exec(sqlDelete, r.RowID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	jumlah := b.Jumlah + r.Jumlah

	updatedAt := constants.GetDatetimeNow()

	sqlUpdate := db.Update(tableBarang, "row_id", "jumlah", "updated_at")

	sqlUpdate = m.Db.Rebind(sqlUpdate)

	if _, err = tx.Exec(sqlUpdate, jumlah, updatedAt, b.RowID); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	if err = tx.Commit(); err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) FetchInventoryKeluar(rowID string) (InventoryKeluar, error) {

	item := InventoryKeluar{}

	joins := getJoinInventoryKeluar()

	sqlX := db.FetchJoin(tableInventoryKeluar, "ik.row_id", getSelectInventoryKeluar(), joins)

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(rowID)

	count := len(getSelectInventoryKeluar())

	values, valuesPtrs := getValuePtr(count)

	err = row.Scan(valuesPtrs...)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	item, err = getRowInventoryKeluar(values, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) GetInventoryKeluar(lastID int64, limit int, filters []string, filterValues []interface{}) ([]InventoryKeluar, error) {

	items := []InventoryKeluar{}

	joins := getJoinInventoryKeluar()

	sqlX := db.QueryPagingJoin(tableInventoryKeluar, "ik.id", false, getSelectInventoryKeluar(), joins, filters)

	if lastID == 0 {

		sqlX = db.QueryPagingJoin(tableInventoryKeluar, "ik.id", true, getSelectInventoryKeluar(), joins, filters)

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

	items, err = getRowsInventoryKeluar(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}

func (m *Model) SearchInventoryKeluar(lastID int64, limit int, search []string, filters []string, filterValues []interface{}) ([]InventoryKeluar, error) {

	items := []InventoryKeluar{}

	joins := getJoinInventoryKeluar()

	sqlX := db.QueryPagingJoinSearch(tableInventoryKeluar, "ik.id", false, getSelectInventoryKeluar(), joins, search[0], filters)

	if lastID == 0 {

		sqlX = db.QueryPagingJoinSearch(tableInventoryKeluar, "ik.id", true, getSelectInventoryKeluar(), joins, search[0], filters)

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

	items, err = getRowsInventoryKeluar(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}
