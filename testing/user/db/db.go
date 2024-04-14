package db

import "esgi-iw1/testing/user/model"

type Db interface {
	Query(string, int) (model.User, error)
	Exec(string, int) error
}
