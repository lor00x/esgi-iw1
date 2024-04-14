package user

import (
	"esgi-iw1/testing/user/db/mock"
	"esgi-iw1/testing/user/model"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetUserById(t *testing.T) {
	id := 18
	db := mock.NewMockDb(gomock.NewController(t))
	db.EXPECT().
		Query("SELECT * FROM users WHERE id = ?", 18).
		Return(model.User{Id: 18}, nil)
	GetUserById(db, id)
}
