package models

import (
	"database/sql"
	"time"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

var tableAsetTetap = "log_aset_tetaps"

type AsetTetap struct {
	ID        int64     `json:"id"`
	RowID     string    `json:"row_id"`
	Kode      string    `json:"kode" validate:"required"`
	Nama      string    `json:"nama" validate:"required"`
	Jumlah    int64     `json:"jumlah" validate:"required"`
	Status    int       `json:"status" validate:"required"`
	Kategori  Kategori  `json:"kategori" validate:"required"`
	UnitKerja UnitKerja `json:"unit_kerja"`
	Metric    Metric    `json:"metric" validate:"required"`
	CreatedAt null.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func getSelectAsetTetap() []string {
	return []string{"a.id", "a.row_id", "a.unit_kerja_id", "a.kategori_id",
		"a.metric_id", "a.kode", "a.nama", "a.jumlah", "a.status", "k.kode AS kode_kategori", "k.nama AS kategori",
		"m.nama AS metric", "u.nama AS unit_kerja", "u.alamat", "u.telepon", "a.created_at", "a.updated_at"}
}

func getJoinAsetTetap() []db.Join {
	joins := []db.Join{}

	join1 := db.Join{TableJoin: "log_unit_kerjas AS u", IDTableJoin: "u.id", IDTable: "a.unit_kerja_id"}
	joins = append(joins, join1)

	join2 := db.Join{TableJoin: "log_kategoris AS k", IDTableJoin: "k.id", IDTable: "a.kategori_id"}
	joins = append(joins, join2)

	join3 := db.Join{TableJoin: "log_metric AS m", IDTableJoin: "m.id", IDTable: "a.metric_id"}
	joins = append(joins, join3)

	return joins
}

func getRowAsetTetap(values []interface{}, m *Model) (AsetTetap, error) {
	item := AsetTetap{}

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
	item.Status = int(values[8].(int64))
	kodeKategori := values[9].(string)
	kategori := values[10].(string)
	metric := values[11].(string)

	if unitKerjaID != 0 {
		unitKerja = values[12].(string)
		alamat = values[13].(string)
		telepon = values[14].(string)
	}

	item.CreatedAt = null.TimeFrom(values[15].(time.Time))
	item.UpdatedAt = null.TimeFrom(values[16].(time.Time))

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

func getRowsAsetTetap(rows *sql.Rows, m *Model) ([]AsetTetap, error) {

	items := []AsetTetap{}

	count := len(getSelectAsetTetap())

	values, valuesPtrs := getValuePtr(count)

	for rows.Next() {

		err := rows.Scan(valuesPtrs...)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		item, err := getRowAsetTetap(values, m)

		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		items = append(items, item)
	}
	return items, nil
}

func (m *Model) CreateAsetTetap(r AsetTetap) error {

	sqlX := db.Insert(tableAsetTetap, "row_id", "unit_kerja_id",
		"kategori_id", "metric_id", "kode", "nama", "jumlah", "status")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.RowID, r.UnitKerja.ID, r.Kategori.ID,
		r.Metric.ID, r.Kode, r.Nama, r.Jumlah, r.Status)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) UpdateAsetTetap(r AsetTetap, row_id string) error {

	updatedAt := constants.GetDatetimeNow()

	sqlX := db.Update(tableAsetTetap, "row_id", "kategori_id", "metric_id", "kode", "nama",
		"jumlah", "status", "updated_at")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Kategori.ID,
		r.Metric.ID, r.Kode, r.Nama, r.Jumlah, r.Status, updatedAt, row_id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveAsetTetap(rowID string) error {

	sqlX := db.Delete(tableAsetTetap, "row_id")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, rowID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) FetchAsetTetap(rowID string) (AsetTetap, error) {

	item := AsetTetap{}

	joins := getJoinAsetTetap()

	sqlX := db.FetchJoin(tableAsetTetap, "row_id", getSelectAsetTetap(), joins)

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(rowID)

	count := len(getSelectAsetTetap())

	values, valuesPtrs := getValuePtr(count)

	err = row.Scan(valuesPtrs...)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	item, err = getRowAsetTetap(values, m)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) GetAsetTetap(lastID int64, limit int, filters []string, filterValues []interface{}) ([]AsetTetap, error) {

	items := []AsetTetap{}

	joins := getJoinAsetTetap()

	sqlX := db.QueryPagingJoin("log_aset_tetaps AS a", "a.id", false, getSelectAsetTetap(), joins, filters)

	if lastID == 0 {

		sqlX = db.QueryPagingJoin("log_aset_tetaps AS a", "a.id", true, getSelectAsetTetap(), joins, filters)

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

	items, err = getRowsAsetTetap(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}

func (m *Model) SearchAsetTetap(lastID int64, limit int, search []string, filters []string, filterValues []interface{}) ([]AsetTetap, error) {

	items := []AsetTetap{}

	joins := getJoinAsetTetap()

	sqlX := db.QueryPagingJoinSearch("log_aset_tetaps AS a", "a.id", false, getSelectAsetTetap(), joins, search[0], filters)

	if lastID == 0 {

		sqlX = db.QueryPagingJoinSearch("log_aset_tetaps AS a", "a.id", true, getSelectAsetTetap(), joins, search[0], filters)

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

	items, err = getRowsAsetTetap(rows, m)

	if err != nil {
		return items, err
	}

	return items, nil
}
