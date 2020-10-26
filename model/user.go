package model

import (
	"github.com/Allifiando/go-gin-boilerplate/model/request"
	responses "github.com/Allifiando/go-gin-boilerplate/model/response"
	Error "github.com/Allifiando/go-gin-boilerplate/pkg/error"
	"github.com/google/uuid"
)

// UserModel ...
type UserModel struct{}

// GetAll ...
func (model *UserModel) GetAll(offset, limit int) (data []responses.UserModel, count int, err error) {
	query := `SELECT * FROM users limit ? offset ?`
	rows, err := GetDB().Query(query, limit, offset)
	if err != nil {
		Error.Error(err)
		return data, count, err
	}
	defer rows.Close()
	for rows.Next() {
		d := responses.UserModel{}
		err = rows.Scan(&d.ID, &d.UUID, &d.Name, &d.Email, &d.Password)
		if err != nil {
			Error.Error(err)
			return data, count, err
		}
		data = append(data, d)
	}
	// Query row count
	query = `SELECT COUNT(*) FROM users`
	err = GetDB().QueryRow(query).Scan(&count)

	return data, count, err
}

// FindByUUID ...
func (model *UserModel) FindByUUID(uuid string) (data responses.UserModel, err error) {
	query := `SELECT uuid, name, email FROM users where uuid = ?`
	err = GetDB().QueryRow(query, uuid).Scan(&data.UUID, &data.Name, &data.Email)
	if err != nil {
		Error.Error(err)
		return data, err
	}
	return data, err
}

// FindByEmail ...
func (model *UserModel) FindByEmail(email string) (data responses.UserModel, err error) {
	query := `SELECT * FROM users where email = ?`
	err = GetDB().QueryRow(query, email).Scan(&data.ID, &data.UUID, &data.Name, &data.Email, &data.Password)

	if err != nil {
		Error.Error(err)
		return data, err
	}
	return data, err
}

// Create ...
func (model *UserModel) Create(body request.Register) (res int64, uid string, err error) {
	tx, err := GetDB().Begin()
	generate := uuid.New().String()
	query, err := tx.Exec(`INSERT INTO users(uuid, name, email, password) 
	values(?,?,?,?)`, generate, body.Name, body.Email, body.Password)

	if err != nil {
		tx.Rollback()
		Error.Error(err)
		return res, generate, err
	}
	res, err = query.LastInsertId()
	if err != nil {
		tx.Rollback()
		Error.Error(err)
		return res, generate, err
	}
	tx.Commit()
	return res, generate, err
}
