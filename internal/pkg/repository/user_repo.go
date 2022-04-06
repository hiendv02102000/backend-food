package repository

import (
	"backend-food/internal/pkg/domain/domain_model/entity"
	"backend-food/pkg/infrastucture/db"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	DB db.Database
}

func (u *UserRepository) FirstUser(condition entity.Users) (entity.Users, error) {
	user := entity.Users{}
	err := u.DB.First(condition, &user)
	return user, err
}
func (u *UserRepository) FindUserList(condition entity.Users) (userList []*entity.Users, err error) {

	result, err := u.DB.Find(condition, &entity.Users{})
	for _, v := range result {
		data, err1 := bson.Marshal(v)
		if err1 != nil {
			err = err1
			userList = nil
			return
		}
		u := entity.Users{}
		err = bson.Unmarshal(data, &u)
		userList = append(userList, &u)
	}
	return
}
func (u *UserRepository) CreateUser(user entity.Users) (entity.Users, error) {
	err := u.DB.Create(&user)
	return user, err
}
func (u *UserRepository) DeleteUser(user entity.Users) error {
	err := u.DB.Delete(&user)
	return err
}
func (u *UserRepository) UpdateUser(oldUser, user entity.Users) error {
	return u.DB.Update(&oldUser, &user)
}
func NewUserRepository(db db.Database) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}
