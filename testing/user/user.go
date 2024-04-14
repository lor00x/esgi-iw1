package user

import (
	"esgi-iw1/testing/user/db"
	"esgi-iw1/testing/user/model"
)

func GetUserById(db db.Db, id int) (model.User, error) {
	return db.Query("SELECT * FROM users WHERE id = ?", id)
}
