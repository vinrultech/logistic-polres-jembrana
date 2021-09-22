package models

import (
	"database/sql"
	"fmt"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

var tableUser = "log_users"

type User struct {
	ID        int         `json:"id"`
	Nama      string      `json:"nama" validate:"required,gte=3"`
	Username  string      `json:"username" validate:"required,gte=5"`
	Password  string      `json:"password" validate:"required"`
	Role      string      `json:"role" validate:"required"`
	UnitKerja UnitKerja   `json:"unit_kerja" validate:"required"`
	Hp        null.String `json:"hp"`
	Foto      null.String `json:"foto"`
	CreatedAt null.Time   `json:"created_at"`
	UpdatedAt null.Time   `json:"updated_at"`
}

func (m *Model) GetUser(username string) (User, error) {
	sqlX := `SELECT 
				id, username, password, role, nama, hp, foto,
				created_at, updated_at 
			 FROM %s 
			 WHERE username=?`
	sqlX = fmt.Sprintf(sqlX, db.TableUsers)

	sqlX = m.Db.Rebind(sqlX)
	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return User{}, err
	}

	defer stmt.Close()

	r := User{}

	row := stmt.QueryRow(username)

	err = row.Scan(&r.ID, &r.Username, &r.Password,
		&r.Role, &r.Nama, &r.Foto,
		&r.Hp, &r.CreatedAt, &r.UpdatedAt)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return User{}, err
	}

	return r, nil

}

func (m *Model) ChangePassword(Password string, id int) error {

	password, _ := constants.HashPassword(Password)
	updatedAt := constants.GetDatetimeNow()

	sqlX := `UPDATE %s SET password=?, updated_at=? WHERE id=?`
	sqlX = fmt.Sprintf(sqlX, db.TableUsers)

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, password, updatedAt, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil

}

func (m *Model) ChangeAccount(Nama string, Hp string, id int) error {

	updatedAt := constants.GetDatetimeNow()

	sqlX := `UPDATE %s SET nama=?, hp=?, updated_at=? WHERE id=?`
	sqlX = fmt.Sprintf(sqlX, db.TableUsers)

	sqlX = m.Db.Rebind(sqlX)
	_, err := m.Db.Exec(sqlX, Nama, Hp, updatedAt, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil

}

func (m *Model) ChangeImage(Foto string, id int) error {

	sqlX := `UPDATE %s SET foto=? WHERE id=?`
	sqlX = fmt.Sprintf(sqlX, db.TableUsers)

	sqlX = m.Db.Rebind(sqlX)
	_, err := m.Db.Exec(sqlX, Foto, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil

}

func (m *Model) FetchAccount(id int) (User, error) {

	r := User{}

	sqlX := `SELECT id, nama, hp, foto FROM %s WHERE id=?`
	sqlX = fmt.Sprintf(sqlX, db.TableUsers)

	sqlX = m.Db.Rebind(sqlX)
	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return r, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(&r.ID, &r.Nama, &r.Hp, &r.Foto)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return r, err
	}

	return r, nil
}

func (m *Model) CreateUser(r User) (int64, error) {

	var lastInsertedID int64

	password, _ := constants.HashPassword(r.Password)

	sqlX := `INSERT INTO %s (nama, username, password, role, hp, foto) 
				VALUES (?, ?, ?, ?, ?, ?) RETURNING id`
	sqlX = fmt.Sprintf(sqlX, db.TableUsers)
	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return lastInsertedID, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(r.Nama, r.Username, password, r.Role, r.Hp, r.Foto).Scan(&lastInsertedID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return lastInsertedID, err
	}

	return lastInsertedID, nil
}

func (m *Model) CreateUserUnitKerja(userID int64, unitKerjaID int64) error {

	sqlX := `INSERT INTO %s (user_id, unit_kerja_id) 
				VALUES (?, ?)`
	sqlX = fmt.Sprintf(sqlX, "log_user_unit_kerja")
	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, userID, unitKerjaID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) CreateUserWithUnitKerja(r User) error {

	password, _ := constants.HashPassword(r.Password)

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	sqlX := `INSERT INTO %s (nama, username, password, role, hp, foto) 
				VALUES (?, ?, ?, ?, ?, ?) RETURNING id`

	sqlX = fmt.Sprintf(sqlX, db.TableUsers)
	sqlX = m.Db.Rebind(sqlX)

	stmt, err := tx.Prepare(sqlX)

	if err != nil {

		return err
	}

	defer stmt.Close()

	var lastInsertedID int

	err = stmt.QueryRow(r.Nama, r.Username, password, r.Role, r.Hp, r.Foto).Scan(&lastInsertedID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	sqlUserUnitKerja := `INSERT INTO %s (user_id, unit_kerja_id) 
				VALUES (?, ?)`
	sqlUserUnitKerja = fmt.Sprintf(sqlUserUnitKerja, "log_user_unit_kerja")
	sqlUserUnitKerja = m.Db.Rebind(sqlUserUnitKerja)

	_, err = tx.Exec(sqlUserUnitKerja, lastInsertedID, r.UnitKerja.ID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	tx.Commit()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveUser(id int64) error {

	tx, err := m.Db.Begin()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	sqlUserDelete := `DELETE FROM %s WHERE id=?`
	sqlUserDelete = fmt.Sprintf(sqlUserDelete, db.TableUsers)
	sqlUserDelete = m.Db.Rebind(sqlUserDelete)
	_, err = tx.Exec(sqlUserDelete, id)
	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	sqlUserUnitKerjaDelete := `DELETE FROM %s WHERE user_id=?`
	sqlUserUnitKerjaDelete = fmt.Sprintf(sqlUserUnitKerjaDelete, "log_user_unit_kerja")
	sqlUserUnitKerjaDelete = m.Db.Rebind(sqlUserUnitKerjaDelete)
	_, err = tx.Exec(sqlUserUnitKerjaDelete, id)
	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	err = tx.Commit()

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil

}

func (m *Model) CheckUser(username string) bool {

	r := User{}

	sqlX := `SELECT username FROM %s WHERE username=?`
	sqlX = fmt.Sprintf(sqlX, db.TableUsers)
	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		return false
	}

	defer stmt.Close()

	row := stmt.QueryRow(username)

	err = row.Scan(&r.Username)

	if err != nil {
		return false
	}

	return true

}

func (m *Model) UpdateUser(r User, id int64) error {

	updatedAt := constants.GetDatetimeNow()

	sqlX := `UPDATE %s SET nama=?, hp=?, updated_at=? WHERE id=?`

	sqlX = fmt.Sprintf(sqlX, tableUser)

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, r.Nama, r.Hp, updatedAt, id)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) UpdateUserUnitKerja(userID int64, unitKerjaID int64) error {

	sqlX := `UPDATE %s SET unit_kerja_id=? WHERE user_id=?`

	sqlX = fmt.Sprintf(sqlX, "log_user_unit_kerja")

	sqlX = m.Db.Rebind(sqlX)

	_, err := m.Db.Exec(sqlX, unitKerjaID, userID)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) FetchUserUnitKerja(id int64) (UnitKerja, error) {

	item := UnitKerja{}

	sqlX := `SELECT unit_kerja_id, nama, alamat, telepon, created_at, updated_at FROM log_user_unit_kerja
		INNER JOIN log_unit_kerjas ON log_unit_kerjas.id=log_user_unit_kerja.unit_kerja_id
		WHERE user_id=?`

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(&item.ID, &item.Nama, &item.Alamat, &item.Telepon, &item.CreatedAt, &item.UpdatedAt)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return item, err
	}

	return item, nil
}

func (m *Model) GetsUser(lastID int64, limit int) ([]User, error) {

	items := []User{}

	sqlX := `SELECT id, nama, username, role, hp, foto, created_at, updated_at FROM %s 
			WHERE id<? ORDER BY id DESC LIMIT ?`

	if lastID == 0 {
		sqlX = `SELECT id, nama, username, role, hp, foto, created_at, updated_at FROM %s 
		ORDER BY id DESC LIMIT ?`
	}

	sqlX = fmt.Sprintf(sqlX, tableUser)

	sqlX = m.Db.Rebind(sqlX)

	stmt, err := m.Db.Preparex(sqlX)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	defer stmt.Close()

	var rows *sql.Rows

	if lastID == 0 {
		rows, err = stmt.Query(limit)
	} else {
		rows, err = stmt.Query(lastID, limit)
	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	for rows.Next() {
		item := User{}

		err = rows.Scan(&item.ID, &item.Nama, &item.Username, &item.Role, &item.Hp, &item.Foto,
			&item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}

		unitKerja, err := m.FetchUserUnitKerja(int64(item.ID))
		if err == sql.ErrNoRows {
			if item.Role == "superuser" {
				item.UnitKerja = GetUnitKerja()
			}
		} else {
			item.UnitKerja = unitKerja
		}

		items = append(items, item)
	}

	return items, nil
}

func (m *Model) SearchUser(lastID int64, limit int, search string, filter string) ([]User, error) {

	items := []User{}

	query := `ILIKE '%' || ? || '%'`

	sqlX := `SELECT id, nama, username, role, hp, foto, created_at, updated_at FROM %s 
			WHERE id<? AND %s %s ORDER BY id DESC LIMIT ?`

	if lastID == 0 {
		sqlX = `SELECT id, nama, username, role, hp, foto, created_at, updated_at FROM %s 
				WHERE %s %s ORDER BY id DESC LIMIT ?`
	}

	sqlX = fmt.Sprintf(sqlX, tableUser, filter, query)

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
		rows, err = stmt.Query(search, limit)
	} else {
		rows, err = stmt.Query(lastID, search, limit)
	}

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return items, err
	}

	for rows.Next() {
		item := User{}
		err = rows.Scan(&item.ID, &item.Nama, &item.Username, &item.Role, &item.Hp, &item.Foto,
			&item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			loggers.Log.Errorln(err.Error())
			return items, err
		}
		unitKerja, err := m.FetchUserUnitKerja(int64(item.ID))
		if err == sql.ErrNoRows {
			if item.Role == "superuser" {
				item.UnitKerja = GetUnitKerja()
			}
		} else {
			item.UnitKerja = unitKerja
		}

		items = append(items, item)
	}

	return items, nil
}
