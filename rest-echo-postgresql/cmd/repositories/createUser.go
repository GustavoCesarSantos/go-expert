package repositories

import (
	"github.com/GustavoCesarSantos/go-expert/rest-echo-postgresql/cmd/models"
	"github.com/GustavoCesarSantos/go-expert/rest-echo-postgresql/cmd/storage"
)

func CreateUser(user models.IUser) (models.IUser, error) {
	db := storage.GetDB()
	sqlStatement := `
		INSERT INTO users (
			name,
			email,
			password
		)
		VALUES (
			$1,
			$2,
			$3
		)
		RETURNING id
	`
	var id int
	err := db.QueryRow(sqlStatement, user.GetName(), user.GetEmail(), user.GetPassword()).Scan(&id)
	if err != nil {
		return nil, err
	}
	user.SetID(id)
	return user, nil
}