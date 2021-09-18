package models

import (
	"fmt"

	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/constants"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/db"
	"gitlab.com/vinrul.tech/log-polres-jembrana-surat/loggers"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID        int         `json:"id"`
	Nama      string      `json:"nama" validate:"required,gte=3"`
	Username  string      `json:"username" validate:"required,gte=5"`
	Password  string      `json:"password" validate:"required"`
	Role      string      `json:"role" validate:"required"`
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

func (m *Model) CreateUser(r User) error {

	password, _ := constants.HashPassword(r.Password)

	sqlX := `INSERT INTO %s (nama, username, password, role, hp, foto) 
				VALUES (?, ?, ?, ?, ?, ?)`
	sqlX = fmt.Sprintf(sqlX, db.TableUsers)
	sqlX = m.Db.Rebind(sqlX)
	_, err := m.Db.Exec(sqlX, r.Nama, r.Username, password, r.Role, r.Hp, r.Foto)

	if err != nil {
		loggers.Log.Errorln(err.Error())
		return err
	}

	return nil
}

func (m *Model) RemoveUser(id int) error {

	sqlX := `DELETE FROM %s WHERE id=?`
	sqlX = fmt.Sprintf(sqlX, db.TableUsers)
	sqlX = m.Db.Rebind(sqlX)
	_, err := m.Db.Exec(sqlX, id)

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
