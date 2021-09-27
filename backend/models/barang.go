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

var tableBarang = "log_barangs AS b"

type Barang struct {
	ID        int64     `json:"id"`
	RowID     string    `json:"row_id"`
	Kode      string    `json:"kode" validate:"required"`
	Nama      string    `json:"nama" validate:"required"`
	Jumlah    int64     `json:"jumlah" validate:"required"`
	Kategori  Kategori  `json:"kategori" validate:"required"`
	UnitKerja UnitKerja `json:"unit_kerja"`
	Metric    Metric    `json:"metric"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func getSelectBarang() []string {
	return []string{"b.id", "b.row_id", "b.unit_kerja_id", "b.kategori_id",
		"b.metric_id", "b.kode", "b.nama", "b.jumlah", "k.kode AS kode_kategori", "k.nama AS kategori",
		"m.nama AS metric", "u.nama AS unit_kerja", "u.alamat", "u.telepon", "b.created_at", "b.updated_at"}
}

func getJoinBarang() []db.Join {
	joins := []db.Join{}

	join1 := db.Join{TableJoin: "log_unit_kerjas AS u", IDTableJoin: "u.id", IDTable: "b.unit_kerja_id"}
	joins = append(joins, join1)

	join2 := db.Join{TableJoin: "log_kategoris AS k", IDTableJoin: "k.id", IDTable: "b.kategori_id"}
	joins = append(joins, join2)

	join3 := db.Join{TableJoin: "log_metric AS m", IDTableJoin: "m.id", IDTable: "b.metric_id"}
	joins = append(joins, join3)

	return joins
}

func getRowBarang(values []interface{}, m *Model) (Barang, error) {
	item := Barang{}

	var unitKerja string
	var alamat string
	var telepon string

	item.ID = values[0].(int64)
	item.RowID = values[1].(string)
	unitKerjaID := values[2].(int64)
	kategoriID := values[3].(int64)
	metricID := values[4].(int64)
	item.Kode = values[5].(string)
	item.Nama = values[6].(string)
	item.Jumlah = values[7].(int64)
	kodeKategori := values[8].(string)
	kategori := values[9].(string)
	metric := values[10].(string)

	if unitKerjaID != 0 {
		unitKerja = values[11].(string)
		alamat = values[12].(string)
		telepon = values[13].(string)
	}

	item.CreatedAt = null.TimeFrom(values[14].(time.Time))
	item.UpdatedAt = null.TimeFrom(values[15].(time.Time))

	if unitKerjaID == 0 {
		item.UnitKerja = GetUnitKerja()
	} else {
		unitKerja := UnitKerja{
			ID:      int64(unitKerjaID),
			Nama:    unitKerja,
			Alamat:  alamat,
			Telepon: telepon,
		}
		item.UnitKerja = unitKerja
	}

	item.Kategori = Kategori{
		ID:   int64(kategoriID),
		Kode: kodeKategori,
		Nama: kategori,
	}

	item.Metric = Metric{
		ID:   int64(metricID),
		Nama: metric,
	}

	return item, nil
}

func getRowsBarang(rows *sql.Rows, m *Model) ([]Barang, error) {

	items := []Barang{}

	count := len(getSelectBarang())

	values, valuesPtrs := getValuePtr(count)

	for rows.Next() {

		err := rows.Scan(valuesPtrs...)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		item, err := getRowBarang(values, m)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}
	return items, nil
}

func (m *Model) CreateBarang(r Barang) error {

	sqlX := db.Insert(tableBarang, "row_id", "unit_kerja_id",
		"kategori_id", "metric_id", "kode", "nama", "jumlah")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.RowID, r.UnitKerja.ID, r.Kategori.ID,
		r.Metric.ID, r.Kode, r.Nama, r.Jumlah)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) UpdateBarang(r Barang, row_id string) error {

	updatedAt := constants.GetDatetimeNow()

	sqlX := db.Update(tableBarang, "row_id", "kategori_id", "metric_id", "kode", "nama", "jumlah", "updated_at")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Kategori.ID,
		r.Metric.ID, r.Kode, r.Nama, r.Jumlah, updatedAt, row_id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveBarang(rowID string) error {

	sqlX := db.Delete(tableBarang, "row_id")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, rowID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) FetchBarang(rowID string) (Barang, error) {

	item := Barang{}

	joins := getJoinBarang()

	sqlX := db.FetchJoin(tableBarang, "row_id", getSelectBarang(), joins)

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(rowID)

	count := len(getSelectBarang())

	values, valuesPtrs := getValuePtr(count)

	err = row.Scan(valuesPtrs...)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	item, err = getRowBarang(values, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) AllBarang(filters []string, filterValues []interface{}) ([]Barang, error) {

	items := []Barang{}

	joins := getJoinBarang()

	sqlX := db.QueryAllJoin("log_barangs AS b", "b.id", getSelectBarang(), joins, filters)

	sqlX = m.Db.Rebind(sqlX)

	fmt.Println(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	rows, err := GetQueryRow(stmt, 0, 0, "", "", filterValues)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	items, err = getRowsBarang(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}

func (m *Model) GetBarang(lastID int64, limit int, filters []string, filterValues []interface{}) ([]Barang, error) {

	items := []Barang{}

	joins := getJoinBarang()

	sqlX := db.QueryPagingJoin("log_barangs AS b", "b.id", false, getSelectBarang(), joins, filters)

	if lastID == 0 {

		sqlX = db.QueryPagingJoin("log_barangs AS b", "b.id", true, getSelectBarang(), joins, filters)

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

	items, err = getRowsBarang(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}

func (m *Model) SearchBarang(lastID int64, limit int, search []string, filters []string, filterValues []interface{}) ([]Barang, error) {

	items := []Barang{}

	joins := getJoinBarang()

	sqlX := db.QueryPagingJoinSearch("log_barangs AS b", "b.id", false, getSelectBarang(), joins, search[0], filters)

	if lastID == 0 {

		sqlX = db.QueryPagingJoinSearch("log_barangs AS b", "b.id", true, getSelectBarang(), joins, search[0], filters)

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

	items, err = getRowsBarang(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}
